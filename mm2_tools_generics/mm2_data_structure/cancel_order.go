package mm2_data_structure

import (
	"encoding/json"
	"fmt"
)

type CancelOrderRequest struct {
	Userpass string `json:"userpass"`
	Method   string `json:"method"`
	Uuid     string `json:"uuid"`
}

type CancelOrderAnswer struct {
	Result string `json:"result"`
}

func (req *CancelOrderRequest) ToJson() string {
	b, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}
