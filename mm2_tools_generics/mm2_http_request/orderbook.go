package mm2_http_request

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kpango/glg"
	"io/ioutil"
	"mm2_client/config"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"net/http"
)

func Orderbook(base string, rel string) (*mm2_data_structure.OrderbookAnswer, error) {
	_, baseOk := config.GCFGRegistry[base]
	_, relOk := config.GCFGRegistry[rel]
	if relOk && baseOk {
		req := mm2_data_structure.NewOrderbookRequest(base, rel).ToJson()
		resp, err := http.Post(mm2_data_structure.GMM2Endpoint, "application/json", bytes.NewBuffer([]byte(req)))
		if err != nil {
			_ = glg.Errorf("%v", err)
			return nil, err
		}
		if resp.StatusCode == http.StatusOK {
			defer resp.Body.Close()
			var answer = &mm2_data_structure.OrderbookAnswer{}
			decodeErr := json.NewDecoder(resp.Body).Decode(answer)
			if decodeErr != nil {
				_ = glg.Errorf("%v", err)
				return nil, decodeErr
			}
			return answer, nil
		} else {
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			errStr := fmt.Sprintf("%s", bodyBytes)
			return nil, errors.New(errStr)
		}
	} else {
		errStr := fmt.Sprintf("coin: %s or %s doesn't exist or is not present in the desktop configuration\n", base, rel)
		_ = glg.Errorf("%v", errStr)
		return nil, errors.New(errStr)
	}
}
