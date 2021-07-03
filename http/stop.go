package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Stop() bool {
	req := NewGenericRequest("stop").ToJson()
	resp, err := http.Post("http://127.0.0.1:7783", "application/json", bytes.NewBuffer([]byte(req)))
	if err != nil {
		fmt.Printf("Err: %v\n", err)
		return false
	}
	if resp.StatusCode == http.StatusOK {
		bodyBytes, bodyErr := ioutil.ReadAll(resp.Body)
		if bodyErr != nil {
			fmt.Printf("Err: %v\n", bodyErr)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
		return true
	}
	return true
}
