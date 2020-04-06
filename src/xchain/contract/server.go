package contract

// 加载用户自定义的合约
var peerContractService Service = &PeerContractService{}

func GetService() Service {
	return peerContractService
}
