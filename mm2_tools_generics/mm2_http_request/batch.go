package mm2_http_request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	http2 "mm2_client/http"
	"net/http"
)

func BatchRequest(batch []interface{}) string {
	p, err := json.Marshal(batch)
	//fmt.Println(string(p))
	if err == nil {
		resp, reqErr := http.Post(http2.GMM2Endpoint, "application/json", bytes.NewBuffer(p))
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
