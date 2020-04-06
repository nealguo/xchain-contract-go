package contract

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"xchain-contract-go/src/xchain/config"
)

func CheckPeerHost(ip string) {
	if ip == "" {
		return
	}
	peer := GetPeerNode()
	if peer.Host == ip {
		return
	}
	err := SetPeerHostSaved(ip)
	if err != nil {
		log.Printf("wrong when updating peer IP, err:%v\n", err)
		return
	}
	// Peer的IP变更后，需要更新配置并重新初始化client
	Prepare()
	peerNew := GetPeerNode()
	RestartClient(peerNew)
	log.Printf("peer IP has changed, restart contract client")
}

var lock sync.Mutex

type PeerHost struct {
	IP string `json:"ip"`
}

func SetPeerHostSaved(ip string) error {
	// sync
	lock.Lock()
	defer lock.Unlock()
	// write to file
	peer := &PeerHost{IP: ip}
	file, err := os.Create(config.PeerHostJson)
	defer file.Close()
	if err != nil {
		log.Printf("open %s wrong\n", config.PeerHostJson)
		return err
	}
	bytes, err := json.Marshal(peer)
	if err != nil {
		log.Printf("serialize json wrong, struct:%v\n", peer)
		return err
	}
	err = ioutil.WriteFile(config.PeerHostJson, bytes, os.ModePerm)
	if err != nil {
		log.Printf("write json in %s wrong\n", config.PeerHostJson)
		return err
	}
	return nil
}

func GetPeerHostSaved() string {
	// sync
	lock.Lock()
	defer lock.Unlock()
	// read from file
	var peer PeerHost
	bytes, err := ioutil.ReadFile(config.PeerHostJson)
	if err != nil {
		log.Printf("read %s wrong, err:%v\n", config.PeerHostJson, err)
		return ""
	}
	err = json.Unmarshal(bytes, &peer)
	if err != nil {
		log.Printf("unserialize json in %s wrong, err:%v\n", config.PeerHostJson, err)
		return ""
	}
	return peer.IP
}
