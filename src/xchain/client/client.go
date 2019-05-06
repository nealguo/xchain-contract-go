package client

import (
	"context"
	"log"
	"xchain-contract-go/src/xchain/config"
	pb "xchain-contract-go/src/xchain/proto/peer"
)

var peerClient pb.ContractPeerServiceClient

func SetPeerClient(client pb.ContractPeerServiceClient) {
	peerClient = client
}

func GetState(channel, nodeId, key string) (string, error) {
	req := &pb.ContractPeerRequest{
		ContractId:  config.GetContractID(),
		Key:         key,
		NodeId:      nodeId,
		ChannelName: channel,
	}
	reply, err := peerClient.GetState(context.Background(), req)
	if err != nil {
		log.Printf("Wrong when getting state, req:%v, err:%v", req, err)
		return "", err
	}
	return reply.GetPayload(), nil
}
