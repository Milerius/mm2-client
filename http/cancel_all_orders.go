package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CancelAllOrdersRequest struct {
	Userpass string `json:"userpass"`
	Method   string `json:"method"`
	CancelBy struct {
		Type string `json:"type"`
		Data *struct {
			Base   *string `json:"base,omitempty"`
			Rel    *string `json:"rel,omitempty"`
			Ticker *string `json:"ticker,omitempty"`
		} `json:"data,omitempty"`
	} `json:"cancel_by"`
}

type CancelAllOrdersAnswer struct {
	Result struct {
		Cancelled         []string `json:"cancelled"`
		CurrentlyMatching []string `json:"currently_matching"`
	} `json:"result"`
}

func NewCancelAllOrdersRequest(kind string, args []string) *CancelAllOrdersRequest {
	genReq := NewGenericRequest("cancel_all_orders")
	req := &CancelAllOrdersRequest{Userpass: genReq.Userpass, Method: genReq.Method}
	return req
}

func (req *CancelAllOrdersRequest) ToJson() string {
	b, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}

func CancelAllOrders(kind string, args []string) *CancelAllOrdersAnswer {
	req := NewCancelAllOrdersRequest(kind, args).ToJson()
	resp, err := http.Post(GMM2Endpoint, "application/json", bytes.NewBuffer([]byte(req)))
	if err != nil {
		fmt.Printf("Err: %v\n", err)
		return nil
	}
	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		var answer = &CancelAllOrdersAnswer{}
		decodeErr := json.NewDecoder(resp.Body).Decode(answer)
		if decodeErr != nil {
			fmt.Printf("Err: %v\n", err)
			return nil
		}
		return answer
	} else {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("Err: %s\n", bodyBytes)
		return nil
	}
}
