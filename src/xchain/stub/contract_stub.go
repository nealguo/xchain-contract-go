package stub

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"xchain-contract-go/src/xchain/client"
	"xchain-contract-go/src/xchain/config"
	"xchain-contract-go/src/xchain/model"
	"xchain-contract-go/src/xchain/util"
)

type ContractStub struct {
	NodeId string
	Tx     *model.Transaction
}

func (stub *ContractStub) InitTx(nodeId, method, contract, channel string) {
	stub.NodeId = nodeId
	stub.Tx.Invoke = method
	stub.Tx.Contract = contract
	stub.Tx.ChannelName = channel
}

func (stub *ContractStub) PrepareTxReads() {
	writes := stub.Tx.Writes
	if writes == nil {
		return
	}

	reads := make([]*model.RwSetRead, len(writes))
	for _, write := range writes {
		read := stub.getTxRead(write.Key)
		if read == nil {
			continue
		}
		reads = append(reads, read)
	}
	stub.Tx.Reads = reads
}

func (stub *ContractStub) PrepareTxHash() {
	reads, err := json.Marshal(stub.Tx.Reads)
	if err != nil {
		return
	}
	writes, err := json.Marshal(stub.Tx.Writes)
	if err != nil {
		return
	}
	info := fmt.Sprintf("%d:%s:%s:%s:%s", stub.Tx.TxId, stub.Tx.Invoke, stub.Tx.Contract, reads, writes)
	stub.Tx.Hash = util.SHA256(info)
}

func (stub *ContractStub) PutState(key, value string) {
	stub.modifyWrite(strings.TrimSpace(key), strings.TrimSpace(value), false)
}

func (stub *ContractStub) DeleteState(key, value string) {
	stub.modifyWrite(strings.TrimSpace(key), "", true)
}

func (stub *ContractStub) modifyWrite(key, value string, delete bool) {
	if key == "" {
		return
	}
	writes := stub.Tx.Writes
	if writes == nil {
		writes = make([]*model.RwSetWrite, 0)
	}

	for _, write := range writes {
		if write.Key == key {
			write.Delete = delete
			write.Value = value
			return
		}
	}

	write := &model.RwSetWrite{
		Key:        key,
		Value:      value,
		Delete:     delete,
		Collection: config.GetContractID().GetName(),
	}
	stub.Tx.Writes = append(stub.Tx.Writes, write)
}

func (stub *ContractStub) GetState(key string) string {
	if key == "" {
		log.Printf("Wrong when getting state, key is empty")
		return ""
	}
	payload, err := client.GetState(stub.Tx.ChannelName, stub.NodeId, key)
	if err != nil {
		log.Printf("Wrong when getting state, key:%v, err:%v", key, err)
		return ""
	}
	var stateModel model.StateModel
	err = json.Unmarshal([]byte(payload), &stateModel)
	if err != nil {
		log.Printf("Wrong when deserializing json, payload:%s, err:%v", payload, err)
		return ""
	}
	return stateModel.Value
}

func (stub *ContractStub) getTxRead(key string) *model.RwSetRead {
	if key == "" {
		log.Printf("Wrong when getting tx read, key is empty")
		return nil
	}
	payload, err := client.GetState(stub.Tx.ChannelName, stub.NodeId, key)
	if err != nil {
		log.Printf("Wrong when getting tx read, key:%v, err:%v", key, err)
		return nil
	}
	var stateModel model.StateModel
	err = json.Unmarshal([]byte(payload), &stateModel)
	if err != nil {
		log.Printf("Wrong when deserializing json, payload:%s, err:%v", payload, err)
		return nil
	}
	return &model.RwSetRead{
		Key:        key,
		Version:    stateModel.Version,
		Collection: config.GetContractID().GetName(),
	}
}
