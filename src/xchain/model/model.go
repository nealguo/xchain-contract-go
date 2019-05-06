package model

// 版本模型
type Version struct {
	BlockNum int `json:"blockNum"`
	TxNum    int `json:"txNum"`
}

// 写集合模型
type RwSetWrite struct {
	Key        string `json:"key"`
	Value      string `json:"value"`
	Delete     bool   `json:"delete"`
	Collection string `json:"collection"`
	AppHash    string `json:"appHash"`
}

// 读集合模型
type RwSetRead struct {
	Key        string  `json:"key"`
	Version    Version `json:"version"`
	Collection string  `json:"collection"`
}

// 交易模型
type Transaction struct {
	TxId        int64         `json:"txId"`
	Invoke      string        `json:"invoke"`
	Contract    string        `json:"contract"`
	Writes      []*RwSetWrite `json:"writes"`
	Reads       []*RwSetRead  `json:"reads"`
	Version     Version       `json:"version"`
	ChannelName string        `json:"channelName"`
	Offset      int64         `json:"offset"`
	Hash        string        `json:"hash"`
}

// 合约调用后返回值对应的模型
type ResponseModel struct {
	Value       string       `json:"value"`
	Transaction *Transaction `json:"transaction"`
}

// 存储在状态库中的状态的模型
type StateModel struct {
	Value   string  `json:"value"`
	Version Version `json:"version"`
}
