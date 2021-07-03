package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mm2_client/config"
	"net/http"
)

type EnableRequest struct {
	Coin                 string   `json:"coin"`
	FallbackSwapContract string   `json:"fallback_swap_contract"`
	Method               string   `json:"method"`
	SwapContractAddress  string   `json:"swap_contract_address"`
	TxHistory            bool     `json:"tx_history"`
	Urls                 []string `json:"urls"`
	Userpass             string   `json:"userpass"`
}

func newEnableRequest(cfg *config.DesktopCFG) *EnableRequest {
	genReq := NewGenericRequest("enable")
	req := &EnableRequest{Userpass: genReq.Userpass, Method: genReq.Method}
	//cfg := config.GCFGRegistry[coin]
	req.Coin = cfg.Coin
	req.TxHistory = true
	req.Urls = cfg.Nodes
	req.SwapContractAddress, req.FallbackSwapContract = cfg.RetrieveContracts()
	return req
}

func (req *EnableRequest) toJson() string {
	b, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}

func Enable(coin string) bool {
	if val, ok := config.GCFGRegistry[coin]; ok {
		req := newEnableRequest(val).toJson()
		resp, err := http.Post(GMM2Endpoint, "application/json", bytes.NewBuffer([]byte(req)))
		if err != nil {
			fmt.Printf("Err: %v\n", err)
			return false
		}
		if resp.StatusCode == http.StatusOK {
			defer resp.Body.Close()
			var answer = &GenericEnableAnswer{}
			decodeErr := json.NewDecoder(resp.Body).Decode(answer)
			if decodeErr != nil {
				fmt.Printf("Err: %v\n", err)
				return false
			}
			answer.ToTable()
			return answer.Result == "success"
		} else {
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			fmt.Printf("Err: %s\n", bodyBytes)
		}
	} else {
		fmt.Printf("coin: %s doesn't exist or is not present in the desktop configuration\n", coin)
		return false
	}
	return false
}
