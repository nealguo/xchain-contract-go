package base

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"sync"
	"xchain-contract-go/src/main/go/contract"
	"xchain-contract-go/src/xchain/client"
	"xchain-contract-go/src/xchain/config"
	pb "xchain-contract-go/src/xchain/proto/contract"
	"xchain-contract-go/src/xchain/proto/peer"
	"xchain-contract-go/src/xchain/service"
)

// 加载用户自定义的合约
var peerContractService service.Contract = &contract.UserContract{}

func Start(server, remote *config.Node) {
	// 初始化配置
	tlsCertFilePath := "data/cert/tls/server.crt"
	tlsPrivateKeyPath := "data/cert/tls/server.pem"
	tls := &config.Tls{TlsCertFilePath: tlsCertFilePath, TlsPrivateKeyPath: tlsPrivateKeyPath}

	// 开始等待服务端和客户端启动完毕
	wg := sync.WaitGroup{}
	wg.Add(2)

	// 启动服务端
	go startServer(server, tls, &wg)

	// 启动客户端
	go startClient(remote, tls, &wg)

	// 结束等待服务端和客户端启动完毕
	wg.Wait()

	log.Printf("Succeeded to start Server and Client")
}

// Start Server side of Contract
func startServer(node *config.Node, tls *config.Tls, wg *sync.WaitGroup) {
	// 监听指定地址和端口
	addr := node.Address()
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen:%s, err:%v", addr, err)
	}

	// 增加TLS认证
	creds, err := credentials.NewServerTLSFromFile(tls.TlsCertFilePath, tls.TlsPrivateKeyPath)
	if err != nil {
		log.Fatalf("Failed to generate credentials:%v", err)
	}

	// 实例化Server, 并开启TLS认证
	server := grpc.NewServer(grpc.Creds(creds))

	// 注册PeerContractService服务
	pb.RegisterPeerContractServiceServer(server, peerContractService)
	log.Printf("Server started, listen on " + addr + " with TLS")

	wg.Done()

	server.Serve(listen)
}

// Start Client and connect to Peer
func startClient(node *config.Node, tls *config.Tls, wg *sync.WaitGroup) {
	// 创建TLS连接
	creds, err := credentials.NewClientTLSFromFile(tls.TlsCertFilePath, "localhost")
	if err != nil {
		log.Fatalf("Failed to create TLS credentials:%v", err)
	}

	// 监听指定地址和端口
	addr := node.Address()
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to connect to addr:%s, err:%v", addr, err)
	}

	// 初始化客户端
	c := peer.NewContractPeerServiceClient(conn)
	client.SetPeerClient(c)

	log.Println("Client started, connected to " + addr + " With TLS")

	wg.Done()
}
