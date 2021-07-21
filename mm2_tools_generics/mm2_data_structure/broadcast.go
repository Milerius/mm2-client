package mm2_data_structure

import (
	"encoding/json"
	"fmt"
)

type BroadcastRequest struct {
	Method   string `json:"method"`
	Coin     string `json:"coin"`
	TxHex    string `json:"tx_hex"`
	Userpass string `json:"userpass"`
}

type BroadcastAnswer struct {
	TxHash string `json:"tx_hash"`
	TxUrl  string
}

func NewBroadcastRequest(coin string, txHex string) *BroadcastRequest {
	genReq := NewGenericRequest("send_raw_transaction")
	req := &BroadcastRequest{Userpass: genReq.Userpass, Method: genReq.Method, Coin: coin, TxHex: txHex}
	return req
}

func (req *BroadcastRequest) ToJson() string {
	b, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}
