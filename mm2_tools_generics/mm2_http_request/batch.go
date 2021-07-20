package mm2_http_request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"net/http"
)

func BatchRequest(batch []interface{}) string {
	p, err := json.Marshal(batch)
	//fmt.Println(string(p))
	if err == nil {
		resp, reqErr := http.Post(mm2_data_structure.GMM2Endpoint, "application/json", bytes.NewBuffer(p))
		if reqErr != nil {
			fmt.Printf("Err: %v\n", reqErr)
			return ""
		}
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return string(bodyBytes)
	} else {
		fmt.Printf("Err: %v\n", err)
		return ""
	}
}
