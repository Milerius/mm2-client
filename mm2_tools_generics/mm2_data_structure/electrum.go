package mm2_data_structure

import (
	"encoding/json"
	"fmt"
	"mm2_client/config"
)

type ElectrumRequest struct {
	Userpass             string                `json:"userpass"`
	Method               string                `json:"method"`
	Coin                 string                `json:"coin"`
	TxHistory            bool                  `json:"tx_history"`
	Servers              []config.ElectrumData `json:"servers"`
	SwapContractAddress  string                `json:"swap_contract_address,omitempty"`
	FallbackSwapContract string                `json:"fallback_swap_contract,omitempty"`
}

func NewElectrumRequest(cfg *config.DesktopCFG) *ElectrumRequest {
	genReq := NewGenericRequest("electrum")
	req := &ElectrumRequest{Userpass: genReq.Userpass, Method: genReq.Method}
	req.Coin = cfg.Coin
	req.TxHistory = true
	req.Servers = cfg.RetrieveElectrums()
	if len(req.Servers) == 0 {
		return nil
	}
	for i, cur := range req.Servers {
		if cur.Protocol != nil && *cur.Protocol == "" {
			req.Servers[i].Protocol = nil
		}
	}
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
