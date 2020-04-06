package contract

import (
	"context"
	"log"
	pb "xchain-contract-go/src/xchain/proto"
	"xchain-contract-go/src/xchain/reply"
	"xchain-contract-go/src/xchain/util"
)

// 合约作为服务提供的接口
type Service interface {
	Invoke(context.Context, *pb.PeerContractRequest) (*pb.RpcReply, error)
	InitContract(context.Context, *pb.PeerContractRequest) (*pb.RpcReply, error)
}

// 合约服务
type PeerContractService struct {
}

func (service *PeerContractService) Invoke(c context.Context, req *pb.PeerContractRequest) (*pb.RpcReply, error) {
	// check param
	if req == nil {
		msg := "Request is nil, failed to invoke contract"
		log.Println(msg)
		return reply.RpcReplyError(msg), nil
	}
	CheckPeerHost(req.GetPeerIp())
	// check invocation
	callers := req.GetCaller()
	if callers != nil && util.Contains(callers, GetContractID()) {
		msg := "Call in loop, failed to invoke contract"
		log.Println(msg)
		return reply.RpcReplyError(msg), nil
	}
	// exec
	nodeId := req.GetAppId()
	method := req.GetMethod()
	channel := req.GetChannelName()
	stub := NewStub(nodeId, method, channel, callers)
	log.Println("succeeded to invoke method:Invoke")
	return stub.Exec(req.GetPayload())
}

func (service *PeerContractService) InitContract(c context.Context, req *pb.PeerContractRequest) (*pb.RpcReply, error) {
	// check
	if req == nil {
		msg := "Request is nil, failed to init contract"
		log.Println(msg)
		return reply.RpcReplyError(msg), nil
	}
	CheckPeerHost(req.GetPeerIp())
	go Build(req.GetPayload())
	log.Println("succeeded to invoke method:InitContract")
	return reply.RpcReplySuccess(""), nil
}
