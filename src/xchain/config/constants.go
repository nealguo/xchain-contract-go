package config

// Docker中的环境变量
const (
	ContractPath    = "CONTRACT_PATH"
	ContractName    = "CONTRACT_NAME"
	ContractVersion = "CONTRACT_VERSION"
	ContractPort    = "CONTRACT_PORT"
	PeerAddress     = "PEER_ADDRESS"
	PeerPort        = "PEER_PORT"
)

// RPC请求的状态码
const (
	Success = 200
	Error   = 500
)
