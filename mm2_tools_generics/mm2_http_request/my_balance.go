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

func MyBalance(coin string) (*mm2_data_structure.MyBalanceAnswer, error) {
	if val, ok := config.GCFGRegistry[coin]; ok {
		req := mm2_data_structure.NewMyBalanceCoinRequest(val).ToJson()
		resp, err := http.Post(mm2_data_structure.GMM2Endpoint, "application/json", bytes.NewBuffer([]byte(req)))
		if err != nil {
			glg.Errorf("Err: %v", err)
			return nil, err
		}
		if resp.StatusCode == http.StatusOK {
			defer resp.Body.Close()
			var answer = &mm2_data_structure.MyBalanceAnswer{}
			decodeErr := json.NewDecoder(resp.Body).Decode(answer)
			if decodeErr != nil {
				glg.Errorf("Err: %v", err)
				return nil, decodeErr
			}
			return answer, nil
		} else {
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			fmt.Printf("Err: %s\n", bodyBytes)
		}
	} else {
		err := fmt.Sprintf("coin: %s doesn't exist or is not present in the desktop configuration", coin)
		return nil, errors.New(err)
	}
	return nil, errors.New("unknown error")
}
