package reply

import (
	"xchain-contract-go/src/xchain/config"
	pb "xchain-contract-go/src/xchain/proto"
	"xchain-contract-go/src/xchain/vo"
)

func ContractResponse(status int, msg, payload string) *vo.Response {
	return &vo.Response{
		Code:    status,
		Message: msg,
		Payload: payload,
	}
}

func ContractResponseSuccess(payload string) *vo.Response {
	return ContractResponse(config.Success, "success", payload)
}

func ContractResponseError(msg string) *vo.Response {
	return ContractResponse(config.Error, msg, "payload")
}

func RpcReply(status int, msg, payload string) *pb.RpcReply {
	return &pb.RpcReply{
		Code:            int32(status),
		Message:         msg,
		Payload:         payload,
		TransactionHash: "",
		File:            nil}
}

func RpcReplySuccess(payload string) *pb.RpcReply {
	return RpcReply(config.Success, "success", payload)
}

func RpcReplyError(msg string) *pb.RpcReply {
	return RpcReply(config.Error, msg, "payload")
}
