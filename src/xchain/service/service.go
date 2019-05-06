package service

import (
	"context"
	pb "xchain-contract-go/src/xchain/proto/contract"
)

// 合约作为服务提供的接口，需要用户来实现
type Contract interface {
	Invoke(context.Context, *pb.PeerContractRequest) (*pb.PeerContractReply, error)
}
