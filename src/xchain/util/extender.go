package util

import (
	pb "xchain-contract-go/src/xchain/proto"
)

func Contains(slice []*pb.ContractID, value *pb.ContractID) bool {
	if slice == nil || len(slice) == 0 {
		return false
	}
	if value == nil {
		return false
	}
	for _, item := range slice {
		if item == nil {
			continue
		}
		if item.String() == value.String() {
			return true
		}
	}
	return false
}
