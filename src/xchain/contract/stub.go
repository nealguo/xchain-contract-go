package contract

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"xchain-contract-go/src/xchain/config"
	"xchain-contract-go/src/xchain/model"
	pb "xchain-contract-go/src/xchain/proto"
	"xchain-contract-go/src/xchain/reply"
	"xchain-contract-go/src/xchain/vo"
)

type Stub struct {
	NodeId  string
	Method  string
	Channel string
	Callers []*pb.ContractID
	Tx      *model.Transaction
	Trace   *model.ContractInvokeTrace
}

func NewStub(nodeId, method, channel string, callers []*pb.ContractID) *Stub {
	// trace
	trace := &model.ContractInvokeTrace{
		Invoke:           method,
		ContractIdentity: GetContractID().GetName(),
		ContractVersion:  GetContractID().GetVersion(),
		SubTraceList:     nil,
	}
	// tx
	tx := &model.Transaction{
		Invoke:              method,
		ContractIdentity:    GetContractID().GetName(),
		ContractVersion:     GetContractID().GetVersion(),
		ChannelName:         channel,
		ContractInvokeTrace: trace,
	}
	return &Stub{
		NodeId:  nodeId,
		Method:  method,
		Channel: channel,
		Callers: callers,
		Tx:      tx,
		Trace:   trace,
	}
}

// invoke contract
func (stub *Stub) Exec(payload string) (*pb.RpcReply, error) {
	// prepare
	m := stub.Method
	method, err := library.Lookup(m)
	if err != nil {
		log.Printf("Dose not exist method:%s, err:%v\n", m, err)
		return reply.RpcReplyError("Dose not exist method " + m), err
	}
	params := &vo.ContractVo{
		Method: m,
		Data:   payload,
	}
	// invoke
	resp, err := method.(func(*Stub, *vo.ContractVo) (*vo.Response, error))(stub, params)
	if err != nil {
		log.Printf("Wrong when invoke method:%s, err:%v\n", m, err)
		return reply.RpcReplyError("Wrong when invoke " + m), err
	}
	if resp == nil {
		log.Printf("Response is empty when invoke method:%s, err:%v\n", m, err)
		return reply.RpcReplyError("Response is empty when invoke " + m), nil
	}
	// result
	stub.prepareTxReads()
	result := model.ResponseModel{Value: resp.Payload, Transaction: stub.Tx}
	bytes, err := json.Marshal(result)
	if err != nil {
		log.Printf("Wrong when serializing to json, invoke method:%s, err:%v\n", m, err)
		return reply.RpcReplyError("Wrong when serializing to json, invoke " + m), err
	}
	return reply.RpcReply(resp.Code, resp.Message, string(bytes)), err
}

// contract may invoke directly
func (stub *Stub) GetState(key string) string {
	if key == "" {
		log.Println("Wrong when getting state, key is empty")
		return ""
	}
	payload, err := GetState(stub.Channel, stub.NodeId, key)
	if err != nil {
		log.Printf("Wrong when getting state, key:%v, err:%v\n", key, err)
		return ""
	}
	if payload == "" {
		return ""
	}
	var stateModel model.StateModel
	err = json.Unmarshal([]byte(payload), &stateModel)
	if err != nil {
		log.Printf("Wrong when deserializing json, payload:%s, err:%v\n", payload, err)
		return ""
	}
	return stateModel.Value
}

// contract may invoke directly
func (stub *Stub) PutState(key, value string) {
	stub.modifyWrites(strings.TrimSpace(key), strings.TrimSpace(value), false)
}

// contract may invoke directly
func (stub *Stub) DeleteState(key string) {
	stub.modifyWrites(strings.TrimSpace(key), "", true)
}

// contract may invoke indirectly
func (stub *Stub) InvokeOtherContract(identity, version, method, params string) (*vo.Response, error) {
	// callers
	if stub.Callers == nil {
		stub.Callers = make([]*pb.ContractID, 0)
	}
	stub.Callers = append(stub.Callers, GetContractID())
	// trace
	if stub.Trace == nil {
		trace := &model.ContractInvokeTrace{
			Invoke:           stub.Method,
			ContractIdentity: GetContractID().GetName(),
			ContractVersion:  GetContractID().GetVersion(),
			SubTraceList:     nil,
		}
		stub.Trace = trace
	}
	// invoke other contract
	other := &pb.ContractID{
		Name:         identity,
		Version:      version,
		Type:         "2", // 用户合约
		LanguageType: "1", // Golang
	}
	return stub.handleOtherContractResult(other, method, params)
}

func (stub *Stub) handleOtherContractResult(other *pb.ContractID, method, params string) (*vo.Response, error) {
	// invoke other contract
	payload, err := InvokeContract(stub.Channel, stub.NodeId, method, params, other, stub.Callers)
	if err != nil {
		return nil, err
	}
	// check result
	var resp *vo.Response
	if payload == "" {
		msg := fmt.Sprintf("result is empty when invoking contract:%s-%s, please avoid calling in loop", other.GetName(), other.GetVersion())
		resp = &vo.Response{
			Code:    config.Success,
			Message: msg,
			Payload: "",
		}
		return resp, nil
	}
	var responseModel model.ResponseModel
	err = json.Unmarshal([]byte(payload), &responseModel)
	if err != nil {
		msg := fmt.Sprintf("wrong when deserializing json")
		resp = &vo.Response{
			Code:    config.Error,
			Message: msg,
			Payload: payload,
		}
		return resp, err
	}
	tx := responseModel.Transaction
	if tx == nil {
		msg := fmt.Sprintf("transaction is empty")
		resp = &vo.Response{
			Code:    config.Error,
			Message: msg,
			Payload: payload,
		}
		return resp, err
	}
	// merge writes and reads
	if tx.Writes != nil && len(tx.Writes) > 0 {
		stub.Tx.Writes = append(stub.Tx.Writes, tx.Writes...)
		stub.Tx.ContractInvokeTrace.SubTraceList = append(stub.Tx.ContractInvokeTrace.SubTraceList, tx.ContractInvokeTrace)
	}
	if tx.Reads != nil && len(tx.Reads) > 0 {
		stub.Tx.Reads = append(stub.Tx.Reads, tx.Reads...)
	}
	// return
	resp = &vo.Response{
		Code:    config.Success,
		Message: "success",
		Payload: responseModel.Value,
	}
	return resp, nil
}

func (stub *Stub) prepareTxReads() {
	// reads
	writes := stub.Tx.Writes
	if writes == nil {
		return
	}
	reads := stub.Tx.Reads
	if reads == nil {
		reads = make([]*model.RwSetRead, 0, len(writes))
	}
	for _, write := range writes {
		read := stub.getTxRead(write.Key)
		if read == nil {
			continue
		}
		reads = append(reads, read)
	}
	stub.Tx.Reads = reads
}

func (stub *Stub) modifyWrites(key, value string, delete bool) {
	if key == "" {
		return
	}
	writes := stub.Tx.Writes
	if writes == nil {
		writes = make([]*model.RwSetWrite, 0)
	}
	self := GetContractID()
	for _, write := range writes {
		if write.Key == key && write.Collection == self.GetName() {
			write.Delete = delete
			write.Value = value
			return
		}
	}
	write := &model.RwSetWrite{
		Key:             key,
		Value:           value,
		Delete:          delete,
		AppId:           stub.NodeId,
		Collection:      self.GetName(),
		ContractVersion: self.GetVersion(),
	}
	stub.Tx.Writes = append(stub.Tx.Writes, write)
}

func (stub *Stub) getTxRead(key string) *model.RwSetRead {
	if key == "" {
		log.Println("Wrong when getting tx read, key is empty")
		return nil
	}
	payload, err := GetState(stub.Channel, stub.NodeId, key)
	if err != nil {
		log.Printf("Wrong when getting tx read, key:%v, err:%v\n", key, err)
		return nil
	}
	if payload == "" {
		return nil
	}
	var stateModel model.StateModel
	err = json.Unmarshal([]byte(payload), &stateModel)
	if err != nil {
		log.Printf("Wrong when deserializing json, payload:%s, err:%v\n", payload, err)
		return nil
	}
	return &model.RwSetRead{
		Key:             key,
		Version:         stateModel.Version,
		Collection:      GetContractID().GetName(),
		ContractVersion: GetContractID().GetVersion(),
	}
}
