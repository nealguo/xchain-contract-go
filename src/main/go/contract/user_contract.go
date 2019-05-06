package contract

import (
	"context"
	"errors"
	"log"
	"time"
	"xchain-contract-go/src/xchain/config"
	"xchain-contract-go/src/xchain/model"
	pb "xchain-contract-go/src/xchain/proto/contract"
	"xchain-contract-go/src/xchain/reply"
	"xchain-contract-go/src/xchain/stub"
	"xchain-contract-go/src/xchain/util"
)

type UserContract struct {
}

var appId string

// 初始化合约的ID
func Init() {
	// TODO 用户需要关注的部分
	appId = "3606597710829453336"
}

// 实现合约的对外接口
func (u *UserContract) Invoke(ctx context.Context, req *pb.PeerContractRequest) (*pb.PeerContractReply, error) {
	// 参数检查
	if ctx == nil || req == nil {
		log.Printf("empty params when calling Invoke of Contract")
		return nil, errors.New("empty params when calling Invoke of Contract")
	}

	// 每次请求就创建一个stub，对应一个tx
	stub := stub.ContractStub{Tx: &model.Transaction{TxId: util.NewId()}}
	stub.InitTx(appId, req.GetMethod(), config.GetContractID().GetName(), req.GetChannelName())

	// TODO 用户需要关注的部分
	var rep *pb.PeerContractReply
	var err error
	method := req.GetMethod()
	switch method {
	case "initAccount":
		rep, err = initAccount(req, &stub)
	case "queryAccount":
		rep, err = queryAccount(req, &stub)
	case "querySingle":
		rep, err = querySingle(req, &stub)
	case "transfer":
		rep, err = transfer(req, &stub)
	case "enableOrDisable":
		rep, err = enableOrDisable(req, &stub)
	default:
		return reply.PeerContractReplyError("Unsupported method", req.GetPeerIp()), nil
	}

	// 处理返回
	if err != nil {
		log.Printf("Wrong when calling %s of Contract:%v", method, err)
		return reply.PeerContractReplyError(rep.Message, req.GetPeerIp()), err
	}
	return reply.HandleReply(req, rep, &stub), nil
}

// 创建账户
func initAccount(req *pb.PeerContractRequest, stub *stub.ContractStub) (*pb.PeerContractReply, error) {
	// 检查参数
	payload := req.GetPayload()
	if payload == "" {
		log.Printf("empty payload when calling " + req.GetMethod())
		return nil, errors.New("empty payload when calling " + req.GetMethod())
	}
	params := util.JsonToMap(payload)

	// 检查账户
	key := util.ToString(params["key"])
	state := stub.GetState(key)
	if state != "" {
		log.Printf("账户已存在, key:%s", key)
		return reply.PeerContractReplyError("账户已存在", req.GetPeerIp()), errors.New("none-empty state when calling " + req.GetMethod())
	}

	// 创建帐户
	balance := params["balance"]
	now := time.Now().Format("2006-01-02 15:04:05")
	value := make(map[string]interface{})
	value["balance"] = balance
	value["createTime"] = now
	value["enable"] = true
	value["lastOperate"] = now
	stub.PutState(key, util.ToJson(value))

	// 返回结果
	log.Printf("账户初始化成功, key:%s", key)
	return reply.PeerContractReplySuccess("账户初始化成功", req.GetPeerIp()), nil
}

// 查询多个账户
func queryAccount(req *pb.PeerContractRequest, stub *stub.ContractStub) (*pb.PeerContractReply, error) {
	// 检查参数
	payload := req.GetPayload()
	if payload == "" {
		log.Printf("empty payload when calling " + req.GetMethod())
		return nil, errors.New("empty payload when calling " + req.GetMethod())
	}
	params := util.JsonToMap(payload)

	// 查询多个账户
	keys := util.ToSlice(params["keys"])
	value := make(map[string]interface{}, len(keys))
	var k string
	for _, key := range keys {
		k = util.ToString(key)
		value[k] = stub.GetState(k)
	}

	// 返回结果
	result := util.ToJson(value)
	log.Printf("查询多个账户成功, keys:%s, result:%s", keys, result)
	return reply.PeerContractReplySuccess(result, req.GetPeerIp()), nil
}

// 查询单个账户
func querySingle(req *pb.PeerContractRequest, stub *stub.ContractStub) (*pb.PeerContractReply, error) {
	// 检查参数
	payload := req.GetPayload()
	if payload == "" {
		log.Printf("empty payload when calling " + req.GetMethod())
		return nil, errors.New("empty payload when calling " + req.GetMethod())
	}
	params := util.JsonToMap(payload)

	// 查询单个账户
	key := util.ToString(params["key"])
	state := stub.GetState(key)

	// 检查参数
	log.Printf("查询单个账户成功, key:%s, state:%s", key, state)
	return reply.PeerContractReplySuccess(state, req.GetPeerIp()), nil
}

// 转账
func transfer(req *pb.PeerContractRequest, stub *stub.ContractStub) (*pb.PeerContractReply, error) {
	// 检查参数
	payload := req.GetPayload()
	if payload == "" {
		log.Printf("empty payload when calling " + req.GetMethod())
		return nil, errors.New("empty payload when calling " + req.GetMethod())
	}
	params := util.JsonToMap(payload)

	// 检查A账户
	keyA := util.ToString(params["keyA"])
	stateA := stub.GetState(keyA)
	if stateA == "" {
		log.Printf("A账户不存在, keyA:%s", keyA)
		return reply.PeerContractReplyError("A账户不存在", req.GetPeerIp()), errors.New("empty stateA when calling " + req.GetMethod())
	}
	valueA := util.JsonToMap(stateA)
	enable := util.ToBool(valueA["enable"])
	if !enable {
		log.Printf("A账户已禁用, keyA:%s", keyA)
		return reply.PeerContractReplyError("A账户已禁用", req.GetPeerIp()), errors.New("disabled stateA when calling " + req.GetMethod())
	}
	balanceA := util.ToFloat64(valueA["balance"])
	amount := util.ToFloat64(params["amount"])
	if balanceA < amount {
		log.Printf("A账户余额不足, keyA:%s", keyA)
		return reply.PeerContractReplyError("A账户余额不足", req.GetPeerIp()), errors.New("none-enough balanceA when calling " + req.GetMethod())
	}

	// 检查B账户
	keyB := util.ToString(params["keyB"])
	stateB := stub.GetState(keyB)
	if stateB == "" {
		log.Printf("B账户不存在, keyB:%s", keyB)
		return reply.PeerContractReplyError("B账户不存在", req.GetPeerIp()), errors.New("empty stateB when calling " + req.GetMethod())
	}

	// 转账
	now := time.Now().Format("2006-01-02 15:04:05")
	valueA["balance"] = balanceA - amount
	valueA["lastOperate"] = now
	stub.PutState(keyA, util.ToJson(valueA))
	valueB := util.JsonToMap(stateB)
	balanceB := util.ToFloat64(valueB["balance"])
	valueB["balance"] = balanceB + amount
	valueB["lastOperate"] = now
	stub.PutState(keyB, util.ToJson(valueB))

	// 返回结果
	log.Printf("转账成功, keyA:%s, stateA:%s, keyB:%s, stateB:%s", keyA, stateA, keyB, stateB)
	return reply.PeerContractReplySuccess("转账成功", req.GetPeerIp()), nil
}

// 启用或禁用
func enableOrDisable(req *pb.PeerContractRequest, stub *stub.ContractStub) (*pb.PeerContractReply, error) {
	// 检查参数
	payload := req.GetPayload()
	if payload == "" {
		log.Printf("empty payload when calling " + req.GetMethod())
		return nil, errors.New("empty payload when calling " + req.GetMethod())
	}
	params := util.JsonToMap(payload)

	// 检查账户
	key := util.ToString(params["key"])
	state := stub.GetState(key)
	if state == "" {
		log.Printf("账户不存在, key:%s", key)
		return reply.PeerContractReplyError("账户不存在", req.GetPeerIp()), errors.New("empty state when calling " + req.GetMethod())
	}

	// 启用或禁用账户
	value := util.JsonToMap(state)
	enable := util.ToBool(value["enable"])
	value["enable"] = !enable
	value["lastOperate"] = time.Now().Format("2006-01-02 15:04:05")
	stub.PutState(key, util.ToJson(value))

	// 返回结果
	var msg string
	if enable {
		msg = "禁用成功"
	} else {
		msg = "启用成功"
	}
	log.Printf("msg:%s, key:%s", msg, key)
	return reply.PeerContractReplySuccess(msg, req.GetPeerIp()), nil
}
