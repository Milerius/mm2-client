package mm2_http_request

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kpango/glg"
	"io/ioutil"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"net/http"
)

func NewCancelAllOrdersRequest(kind string, args []string) *mm2_data_structure.CancelAllOrdersRequest {
	genReq := mm2_data_structure.NewGenericRequest("cancel_all_orders")
	req := &mm2_data_structure.CancelAllOrdersRequest{Userpass: genReq.Userpass, Method: genReq.Method}
	switch kind {
	case "all":
		req.CancelBy.Type = "All"
	case "by_pair":
		req.CancelBy.Type = "Pair"
		req.CancelBy.Data = &mm2_data_structure.DataCancel{Base: &args[0], Rel: &args[1]}
	case "by_coin":
		req.CancelBy.Type = "Coin"
		req.CancelBy.Data = &mm2_data_structure.DataCancel{Ticker: &args[0]}
	}
	return req
}

func CancelAllOrders(kind string, args []string) (*mm2_data_structure.CancelAllOrdersAnswer, error) {
	req := NewCancelAllOrdersRequest(kind, args).ToJson()
	resp, err := http.Post(mm2_data_structure.GMM2Endpoint, "application/json", bytes.NewBuffer([]byte(req)))
	if err != nil {
		glg.Errorf("Err: %v", err)
		return nil, err
	}
	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		var answer = &mm2_data_structure.CancelAllOrdersAnswer{}
		decodeErr := json.NewDecoder(resp.Body).Decode(answer)
		if decodeErr != nil {
			glg.Errorf("Err: %v", decodeErr)
			return nil, decodeErr
		}
		return answer, nil
	} else {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		errStr := fmt.Sprintf("err: %s\n", bodyBytes)
		return nil, errors.New(errStr)
	}
}
