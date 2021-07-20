package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"net/http"
)

type StopAnswer struct {
	Result string `json:"result"`
}

func Stop() bool {
	req := mm2_data_structure.NewGenericRequest("stop").ToJson()
	resp, err := http.Post(mm2_data_structure.GMM2Endpoint, "application/json", bytes.NewBuffer([]byte(req)))
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
		mm2_data_structure.GRuntimeUserpass = ""
		return answer.Result == "success"
	}
	return true
}
