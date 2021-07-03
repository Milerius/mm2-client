package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mm2_client/config"
	"net/http"
)

type ElectrumRequest struct {
	Userpass  string `json:"userpass"`
	Method    string `json:"method"`
	Coin      string `json:"coin"`
	TxHistory bool   `json:"tx_history"`
	Servers   []struct {
		URL                     string `json:"url"`
		Protocol                string `json:"protocol,omitempty"`
		DisableCertVerification bool   `json:"disable_cert_verification,omitempty"`
	} `json:"servers"`
	SwapContractAddress  string `json:"swap_contract_address,omitempty"`
	FallbackSwapContract string `json:"fallback_swap_contract,omitempty"`
}

func NewElectrumRequest(cfg *config.DesktopCFG) *ElectrumRequest {
	genReq := NewGenericRequest("electrum")
	req := &ElectrumRequest{Userpass: genReq.Userpass, Method: genReq.Method}
	req.Coin = cfg.Coin
	req.TxHistory = true
	req.Servers = []struct {
		URL                     string `json:"url"`
		Protocol                string `json:"protocol,omitempty"`
		DisableCertVerification bool   `json:"disable_cert_verification,omitempty"`
	}(cfg.Electrum)
	req.SwapContractAddress, req.FallbackSwapContract = cfg.RetrieveContracts()
	return req
}

func (req *ElectrumRequest) ToJson() string {
	b, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}

func Electrum(coin string) bool {
	if val, ok := config.GCFGRegistry[coin]; ok {
		req := NewElectrumRequest(val).ToJson()
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
