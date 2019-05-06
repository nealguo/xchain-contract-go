package reply

import (
	"encoding/json"
	"xchain-contract-go/src/xchain/config"
	"xchain-contract-go/src/xchain/model"
	pb "xchain-contract-go/src/xchain/proto/contract"
	"xchain-contract-go/src/xchain/stub"
)

func HandleReply(req *pb.PeerContractRequest, reply *pb.PeerContractReply, stub *stub.ContractStub) *pb.PeerContractReply {
	// 参数检查
	if req == nil || reply == nil || stub == nil {
		return reply
	}
	code := reply.GetCode()
	if code != int32(config.Success) {
		return reply
	}

	// 准备读集合和交易哈希
	stub.PrepareTxReads()
	stub.PrepareTxHash()

	// 返回结果
	resp := model.ResponseModel{Value: reply.GetPayload(), Transaction: stub.Tx}
	bytes, err := json.Marshal(resp)
	if err != nil {
		return PeerContractReplyError("Wrong when serializing to json", req.GetPeerIp())
	}
	return PeerContractReplySuccess(string(bytes), req.GetPeerIp())
}

func PeerContractReply(status int, msg, payload, peerIp string) *pb.PeerContractReply {
	var original = false
	peerNode := config.GetPeerNode()
	if peerNode.Host == peerIp {
		original = true
	}
	return &pb.PeerContractReply{Code: int32(status), Message: msg, Payload: payload, Is_Original: original}
}

func PeerContractReplySuccess(payload, peerIp string) *pb.PeerContractReply {
	return PeerContractReply(config.Success, "success", payload, peerIp)
}

func PeerContractReplyError(msg, peerIp string) *pb.PeerContractReply {
	return PeerContractReply(config.Error, msg, "payload", peerIp)
}
