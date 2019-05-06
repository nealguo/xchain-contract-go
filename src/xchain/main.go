package main

import (
	"log"
	"xchain-contract-go/src/main/go/contract"
	"xchain-contract-go/src/xchain/base"
	"xchain-contract-go/src/xchain/config"
)

var server *config.Node
var client *config.Node

func init() {
	log.Printf("Env is preparing...")

	// 初始化环境变量
	config.Prepare()

	// 初始化合约作为服务端
	contractPort := config.GetContractPort()
	server = &config.Node{Host: "", Port: contractPort}

	// 初始化Peer的客户端
	peerNode := config.GetPeerNode()
	client = &config.Node{Host: peerNode.Host, Port: peerNode.Port}

	log.Printf("Env is ready, server:%v, client:%v", server, client)
}

func main() {
	log.Printf("Contract is initializing...")
	contract.Init()

	log.Printf("Server and Client is staring...")
	base.Start(server, client)

	log.Printf("--- Success, user contract is ready now ---")

	select {}
}
