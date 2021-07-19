package mm2_http_request

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kpango/glg"
	"io/ioutil"
	"mm2_client/config"
	http2 "mm2_client/http"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"net/http"
)

func DisableCoin(coin string) (*mm2_data_structure.DisableCoinAnswer, error) {
	if val, ok := config.GCFGRegistry[coin]; ok {
		req := mm2_data_structure.NewDisableCoinRequest(val).ToJson()
		resp, err := http.Post(http2.GMM2Endpoint, "application/json", bytes.NewBuffer([]byte(req)))
		if err != nil {
			_ = glg.Errorf("Err: %v", err)
			return nil, err
		}
		if resp.StatusCode == http.StatusOK {
			defer resp.Body.Close()
			var answer = &mm2_data_structure.DisableCoinAnswer{}
			decodeErr := json.NewDecoder(resp.Body).Decode(answer)
			if decodeErr != nil {
				_ = glg.Errorf("Err: %v", decodeErr)
				return nil, decodeErr
			}
			return answer, nil
		} else {
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			_ = glg.Errorf("Err: %s", bodyBytes)
			return nil, errors.New(string(bodyBytes))
		}
	} else {
		errStr := fmt.Sprintf("coin: %s doesn't exist or is not present in the desktop configuration", coin)
		return nil, errors.New(errStr)
	}
}
