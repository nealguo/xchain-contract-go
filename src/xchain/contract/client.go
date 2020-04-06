package contract

import (
	"context"
	"log"
	pb "xchain-contract-go/src/xchain/proto"
)

var peerClient pb.ContractPeerServiceClient

func SetPeerClient(client pb.ContractPeerServiceClient) {
	peerClient = client
}

func GetState(channel, nodeId, key string) (string, error) {
	req := &pb.ContractPeerRequest{
		ContractId:  GetContractID(),
		Key:         key,
		AppId:       nodeId,
		ChannelName: channel,
	}
	reply, err := peerClient.GetState(context.Background(), req)
	if err != nil {
		log.Printf("Wrong when getting state, req:%v, err:%v\n", req, err)
		return "", err
	}
	return reply.GetPayload(), nil
}

func InvokeContract(channel, nodeId, method, params string, other *pb.ContractID, callers []*pb.ContractID) (string, error) {
	req := &pb.InvokeContractRequest{
		ContractId:  other,
		Method:      method,
		Payload:     params,
		ChannelName: channel,
		AppId:       nodeId,
		Caller:      callers,
	}
	reply, err := peerClient.InvokeContract(context.Background(), req)
	if err != nil {
		log.Printf("Wrong when invoking contract, req:%v, err:%v\n", req, err)
		return "", err
	}
	return reply.GetPayload(), nil
}
