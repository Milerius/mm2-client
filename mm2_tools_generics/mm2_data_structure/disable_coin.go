package mm2_data_structure

import (
	"encoding/json"
	"fmt"
	"mm2_client/config"
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
