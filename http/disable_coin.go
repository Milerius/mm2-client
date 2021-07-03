package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mm2_client/config"
	"net/http"
)

type DisableCoinRequest struct {
	Userpass string `json:"userpass"`
	Method   string `json:"method"`
	Coin     string `json:"coin"`
}

type DisableCoinAnswer struct {
	Result struct {
		CancelledOrders []string `json:"cancelled_orders,omitempty"`
		Coin            string   `json:"coin"`
	} `json:"result,omitempty"`
	Error string `json:"error,omitempty"`
}

func NewDisableCoinRequest(cfg *config.DesktopCFG) *DisableCoinRequest {
	genReq := NewGenericRequest("disable_coin")
	req := &DisableCoinRequest{Userpass: genReq.Userpass, Method: genReq.Method}
	req.Coin = cfg.Coin
	return req
}

func (req *DisableCoinRequest) ToJson() string {
	b, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}

func DisableCoin(coin string) *DisableCoinAnswer {
	if val, ok := config.GCFGRegistry[coin]; ok {
		req := NewDisableCoinRequest(val).ToJson()
		resp, err := http.Post(GMM2Endpoint, "application/json", bytes.NewBuffer([]byte(req)))
		if err != nil {
			fmt.Printf("Err: %v\n", err)
			return nil
		}
		if resp.StatusCode == http.StatusOK {
			defer resp.Body.Close()
			var answer = &DisableCoinAnswer{}
			decodeErr := json.NewDecoder(resp.Body).Decode(answer)
			if decodeErr != nil {
				fmt.Printf("Err: %v\n", err)
				return nil
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
