package contract

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"sync"
	"xchain-contract-go/src/xchain/config"
	pb "xchain-contract-go/src/xchain/proto"
)

func Start(server, remote *config.Node) {
	// 初始化配置
	tls := &config.Tls{TlsCertFilePath: config.TlsCertFilePath, TlsPrivateKeyPath: config.TlsPrivateKeyPath}

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

func RestartClient(remote *config.Node) {
	// 初始化配置
	tls := &config.Tls{TlsCertFilePath: config.TlsCertFilePath, TlsPrivateKeyPath: config.TlsPrivateKeyPath}

	// 开始等待客户端启动完毕
	wg := sync.WaitGroup{}
	wg.Add(1)

	// 启动客户端
	go startClient(remote, tls, &wg)

	// 结束等待客户端启动完毕
	wg.Wait()

	log.Printf("Succeeded to restart Client")
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
	pb.RegisterPeerContractServiceServer(server, GetService())
	log.Printf("Server started, listen on %s with TLS\n", addr)

	wg.Done()

	server.Serve(listen)
}

// Start Client and connect to Peer
func startClient(node *config.Node, tls *config.Tls, wg *sync.WaitGroup) {
	// 创建TLS连接，serverNameOverride对应RSA证书生成时填写的"Common Name"值
	creds, err := credentials.NewClientTLSFromFile(tls.TlsCertFilePath, "xchain")
	if err != nil {
		log.Fatalf("Failed to create TLS credentials:%v", err)
	}

	// 监听指定地址和端口
	addr := node.Address()
	if addr == "" {
		log.Fatalf("Failed to get peer addr, addr is empty\n")
	}
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to connect to addr:%s, err:%v", addr, err)
	}

	// 初始化客户端
	c := pb.NewContractPeerServiceClient(conn)
	SetPeerClient(c)

	log.Printf("Client started, connected to %s With TLS\n", addr)

	wg.Done()
}
