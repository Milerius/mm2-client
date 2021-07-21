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

func Withdraw(coin string, amount string, address string, fees []string, coinType string) (*mm2_data_structure.WithdrawAnswer, error) {
	if _, ok := config.GCFGRegistry[coin]; ok {
		req := mm2_data_structure.NewWithdrawRequest(coin, amount, address, fees, coinType).ToJson()
		resp, err := http.Post(mm2_data_structure.GMM2Endpoint, "application/json", bytes.NewBuffer([]byte(req)))
		if err != nil {
			fmt.Printf("Err: %v\n", err)
			return nil, err
		}
		if resp.StatusCode == http.StatusOK {
			defer resp.Body.Close()
			var answer = &mm2_data_structure.WithdrawAnswer{}
			decodeErr := json.NewDecoder(resp.Body).Decode(answer)
			if decodeErr != nil {
				fmt.Printf("Err: %v\n", err)
				return nil, decodeErr
			}
			return answer, nil
		} else {
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			errStr := fmt.Sprintf("Err: %s\n", bodyBytes)
			return nil, errors.New(errStr)
		}
	} else {
		errStr := fmt.Sprintf("coin: %s doesn't exist or is not present in the desktop configuration\n", coin)
		_ = glg.Errorf("%s", errStr)
		return nil, errors.New(errStr)
	}
}
