package main

import (
	"log"
	"xchain-contract-go/src/xchain/config"
	"xchain-contract-go/src/xchain/contract"
)

var server *config.Node
var client *config.Node

func init() {
	log.Printf("Env is preparing...")

	// 初始化环境变量
	contract.Prepare()

	// 初始化合约作为服务端
	contractPort := contract.GetContractPort()
	server = &config.Node{Host: "", Port: contractPort}

	// 初始化Peer的客户端
	client = contract.GetPeerNode()

	log.Printf("Env is ready, server:%v, client:%v", server, client)
}

func main() {
	log.Printf("Contract is initializing...")
	contract.Load()

	log.Printf("Server and Client is staring...")
	contract.Start(server, client)

	log.Printf("Contract library is building...")

	select {}
}
