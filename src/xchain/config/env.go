package config

import (
	"log"
	"os"
	"strconv"
	"xchain-contract-go/src/xchain/proto/peer"
)

var peerNode *Node
var contractPort uint16
var contractID *peer.ContractID

func Prepare() {
	// contract port
	cPort, err := strconv.Atoi(os.Getenv(ContractPort))
	if err != nil {
		log.Fatalf("wrong port of Contract:%s", err)
	}
	contractPort = uint16(cPort)

	// contract info
	contractID = &peer.ContractID{
		Path:         os.Getenv(ContractPath),
		Name:         os.Getenv(ContractName),
		Version:      os.Getenv(ContractVersion),
		Type:         "2", // 用户合约
		LanguageType: "1"} // Golang

	// peer info
	peerHost := os.Getenv(PeerAddress)
	if peerHost == "" {
		peerHost = "127.0.0.1"
		log.Println("Use 127.0.0.1 as Peer host")
	}

	pPort := os.Getenv(PeerPort)
	if pPort == "" {
		pPort = "50051"
		log.Println("Use 50051 as Peer port")
	}
	peerPort, err := strconv.Atoi(pPort)
	if err != nil {
		log.Fatalf("Wrong port of Peer:%s", err)
	}
	peerNode = &Node{Host: peerHost, Port: uint16(peerPort)}
}

func GetPeerNode() *Node {
	return peerNode
}

func GetContractPort() uint16 {
	return contractPort
}

func GetContractID() *peer.ContractID {
	return contractID
}
