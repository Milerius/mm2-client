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

func Broadcast(coin string, txHex string) (*mm2_data_structure.BroadcastAnswer, error) {
	if val, ok := config.GCFGRegistry[coin]; ok {
		req := mm2_data_structure.NewBroadcastRequest(coin, txHex).ToJson()
		resp, err := http.Post(mm2_data_structure.GMM2Endpoint, "application/json", bytes.NewBuffer([]byte(req)))
		if err != nil {
			_ = glg.Errorf("%v", err)
			return nil, err
		}
		if resp.StatusCode == http.StatusOK {
			defer resp.Body.Close()
			var answer = &mm2_data_structure.BroadcastAnswer{}
			decodeErr := json.NewDecoder(resp.Body).Decode(answer)
			if decodeErr != nil {
				_ = glg.Errorf("%v", decodeErr)
				return nil, decodeErr
			}

			if val.ExplorerTxURL != "" {
				answer.TxUrl = val.ExplorerURL + val.ExplorerTxURL + answer.TxHash
			} else {
				answer.TxUrl = val.ExplorerURL + "tx/" + answer.TxHash
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
