package mm2_data_structure

import (
	"encoding/json"
	"fmt"
)

type DataCancel struct {
	Base   *string `json:"base,omitempty"`
	Rel    *string `json:"rel,omitempty"`
	Ticker *string `json:"ticker,omitempty"`
}

type CancelAllOrdersRequest struct {
	Userpass string `json:"userpass"`
	Method   string `json:"method"`
	CancelBy struct {
		Type string      `json:"type"`
		Data *DataCancel `json:"data,omitempty"`
	} `json:"cancel_by"`
}

type CancelAllOrdersAnswer struct {
	Result struct {
		Cancelled         []string `json:"cancelled"`
		CurrentlyMatching []string `json:"currently_matching"`
	} `json:"result"`
}

func (req *CancelAllOrdersRequest) ToJson() string {
	b, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}
