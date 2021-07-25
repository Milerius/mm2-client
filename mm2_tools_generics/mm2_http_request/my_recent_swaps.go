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

func MyRecentSwaps(limit string, pageNumber string, baseCoin string, relCoin string, from string, to string) (*mm2_data_structure.MyRecentSwapsAnswer, error) {
	req := mm2_data_structure.NewMyRecentSwapsRequest(limit, pageNumber, baseCoin, relCoin, from, to).ToJson()
	resp, err := http.Post(mm2_data_structure.GMM2Endpoint, "application/json", bytes.NewBuffer([]byte(req)))
	if err != nil {
		_ = glg.Errorf("%v", err)
		return nil, err
	}
	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		var answer = &mm2_data_structure.MyRecentSwapsAnswer{}
		decodeErr := json.NewDecoder(resp.Body).Decode(answer)
		if decodeErr != nil {
			_ = glg.Errorf("%v", decodeErr)
			return nil, decodeErr
		}
		return answer, nil
	} else {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		errStr := fmt.Sprintf("Err: %s\n", bodyBytes)
		return nil, errors.New(errStr)
	}
}
