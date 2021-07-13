package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type StopAnswer struct {
	Result string `json:"result"`
}

func Stop() bool {
	req := NewGenericRequest("stop").ToJson()
	resp, err := http.Post(GMM2Endpoint, "application/json", bytes.NewBuffer([]byte(req)))
	if err != nil {
		fmt.Printf("Err: %v\n", err)
		return false
	}
	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		var answer = &StopAnswer{}
		decodeErr := json.NewDecoder(resp.Body).Decode(answer)
		if decodeErr != nil {
			fmt.Printf("Err: %v\n", err)
			return false
		}
		GRuntimeUserpass = ""
		return answer.Result == "success"
	}
	return true
}
