package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mm2_client/config"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"net/http"
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
	genReq := mm2_data_structure.NewGenericRequest("send_raw_transaction")
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

func Broadcast(coin string, txHex string) *BroadcastAnswer {
	if val, ok := config.GCFGRegistry[coin]; ok {
		req := NewBroadcastRequest(coin, txHex).ToJson()
		resp, err := http.Post(mm2_data_structure.GMM2Endpoint, "application/json", bytes.NewBuffer([]byte(req)))
		if err != nil {
			fmt.Printf("Err: %v\n", err)
			return nil
		}
		if resp.StatusCode == http.StatusOK {
			defer resp.Body.Close()
			var answer = &BroadcastAnswer{}
			decodeErr := json.NewDecoder(resp.Body).Decode(answer)
			if decodeErr != nil {
				fmt.Printf("Err: %v\n", err)
				return nil
			}

			if val.ExplorerTxURL != "" {
				answer.TxUrl = val.ExplorerURL[0] + val.ExplorerTxURL + answer.TxHash
			} else {
				answer.TxUrl = val.ExplorerURL[0] + "tx/" + answer.TxHash
			}
			return answer
		} else {
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			fmt.Printf("Err: %s\n", bodyBytes)
		}
	} else {
		fmt.Printf("coin: %s doesn't exist or is not present in the desktop configuration\n", coin)
		return nil
	}
	return nil
}
