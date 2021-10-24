package mm2_data_structure

import (
	"encoding/json"
	"fmt"
	"mm2_client/config"
)

type EnableRequest struct {
	Coin                 string   `json:"coin"`
	FallbackSwapContract string   `json:"fallback_swap_contract"`
	Method               string   `json:"method"`
	SwapContractAddress  string   `json:"swap_contract_address"`
	TxHistory            bool     `json:"tx_history"`
	Urls                 []string `json:"urls"`
	Userpass             string   `json:"userpass"`
	GasStationUrl        string   `json:"gas_station_url,omitempty"`
}

func NewEnableRequest(cfg *config.DesktopCFG) *EnableRequest {
	genReq := NewGenericRequest("enable")
	req := &EnableRequest{Userpass: genReq.Userpass, Method: genReq.Method}
	req.Coin = cfg.Coin
	req.TxHistory = false
	req.Urls = cfg.Nodes
	req.SwapContractAddress, req.FallbackSwapContract = cfg.RetrieveContracts()
	req.GasStationUrl = cfg.RetrieveGasStationUrl()
	return req
}

func (req *EnableRequest) ToJson() string {
	b, err := json.Marshal(req)
	if err != nil {
		fmt.Printf("func (req *EnableRequest) ToJson() Err: %v\n", err)
		return ""
	}
	return string(b)
}
