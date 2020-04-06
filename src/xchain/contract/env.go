package contract

import (
	"log"
	"os"
	"strconv"
	"xchain-contract-go/src/xchain/config"
	pb "xchain-contract-go/src/xchain/proto"
)

var peerNode *config.Node
var contractPort uint16
var contractID *pb.ContractID

func Prepare() {
	// contract port
	cPort, err := strconv.Atoi(os.Getenv(config.ContractPort))
	if err != nil {
		log.Fatalf("Wrong port of Contract:%s", err)
	}
	contractPort = uint16(cPort)

	// contract info
	contractID = &pb.ContractID{
		Name:         os.Getenv(config.ContractName),
		Version:      os.Getenv(config.ContractVersion),
		Type:         "2", // 用户合约
		LanguageType: "1", // Golang
	}

	// peer info
	peerHost := os.Getenv(config.PeerAddress)
	if peerHost == "" {
		log.Fatalf("IP of Peer is empty, exit now")
	}
	pPort := os.Getenv(config.PeerPort)
	if pPort == "" {
		log.Fatalf("Port of Peer is empty, exit now")
	}
	peerPort, err := strconv.Atoi(pPort)
	if err != nil {
		log.Fatalf("Wrong port of Peer:%s", err)
	}
	peerHostSaved := GetPeerHostSaved()
	if peerHostSaved == "" {
		SetPeerHostSaved(peerHost)
	} else {
		peerHost = peerHostSaved
	}
	peerNode = &config.Node{Host: peerHost, Port: uint16(peerPort)}
}

func GetPeerNode() *config.Node {
	return peerNode
}

func GetContractPort() uint16 {
	return contractPort
}

func GetContractID() *pb.ContractID {
	return contractID
}
