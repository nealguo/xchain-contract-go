package config

// Docker中的环境变量
const (
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

// 路径相关
const (
	WorkDir           = "/opt/app/"
	TlsCertFilePath   = WorkDir + "server.crt"
	TlsPrivateKeyPath = WorkDir + "server.pem"
	PeerHostJson      = WorkDir + "peer_host.json"
	ContractFileName  = "contract.go"
	LibraryFileName   = "contract.so"
)
