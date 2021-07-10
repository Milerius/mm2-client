package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CancelOrderRequest struct {
	Userpass string `json:"userpass"`
	Method   string `json:"method"`
	Uuid     string `json:"uuid"`
}

type CancelOrderAnswer struct {
	Result string `json:"result"`
}

func NewCancelOrderRequest(uuid string) *CancelOrderRequest {
	genReq := NewGenericRequest("cancel_order")
	req := &CancelOrderRequest{Userpass: genReq.Userpass, Method: genReq.Method, Uuid: uuid}
	return req
}

func (req *CancelOrderRequest) ToJson() string {
	b, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}

func CancelOrder(uuid string) *CancelOrderAnswer {
	req := NewCancelOrderRequest(uuid).ToJson()
	resp, err := http.Post(GMM2Endpoint, "application/json", bytes.NewBuffer([]byte(req)))
	if err != nil {
		fmt.Printf("Err: %v\n", err)
		return nil
	}
	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		var answer = &CancelOrderAnswer{}
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
