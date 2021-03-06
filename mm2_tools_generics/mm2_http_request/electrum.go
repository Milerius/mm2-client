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

func Electrum(coin string) (*mm2_data_structure.GenericEnableAnswer, error) {
	if val, ok := config.GCFGRegistry[coin]; ok {
		reqRaw := mm2_data_structure.NewElectrumRequest(val)
		if reqRaw != nil {
			req := reqRaw.ToJson()
			resp, err := http.Post(mm2_data_structure.GMM2Endpoint, "application/json", bytes.NewBuffer([]byte(req)))
			if err != nil {
				glg.Errorf("Err: %v", err)
				return nil, err
			}
			if resp.StatusCode == http.StatusOK {
				defer resp.Body.Close()
				var answer = &mm2_data_structure.GenericEnableAnswer{}
				decodeErr := json.NewDecoder(resp.Body).Decode(answer)
				if decodeErr != nil {
					glg.Errorf("Err: %v", decodeErr)
					return nil, decodeErr
				}
				return answer, nil
			} else {
				bodyBytes, _ := ioutil.ReadAll(resp.Body)
				errStr := fmt.Sprintf("err: %s", bodyBytes)
				return nil, errors.New(errStr)
			}
		} else {
			_ = glg.Errorf("coins: %s have no valid electrum for your servers - skipping", coin)
			return nil, errors.New("coins: " + coin + " have no valid electrum for your servers - skipping")
		}
	} else {
		_ = glg.Errorf("coin: %s doesn't exist or is not present in the desktop configuration", coin)
		return nil, errors.New("coin: " + coin + "doesn't exist or is not present in the desktop configuration")
	}
}
