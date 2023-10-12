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

const customTxEndpoint = "https://etherscan-proxy.komodo.earth/api/"

func MyTxHistory(coin string, defaultNbTx int, defaultPage int, withFiatValue bool, isMax bool) (*mm2_data_structure.MyTxHistoryAnswer, error) {
	if _, ok := config.GCFGRegistry[coin]; ok {
		req := mm2_data_structure.NewMyTxHistoryRequest(coin, defaultNbTx, defaultPage, isMax).ToJson()
		resp, err := http.Post(mm2_data_structure.GMM2Endpoint, "application/json", bytes.NewBuffer([]byte(req)))
		if err != nil {
			_ = glg.Errorf("%v", err)
			return nil, err
		}
		if resp.StatusCode == http.StatusOK {
			defer resp.Body.Close()
			var answer = &mm2_data_structure.MyTxHistoryAnswer{}
			decodeErr := json.NewDecoder(resp.Body).Decode(answer)
			if decodeErr != nil {
				glg.Errorf("%v", err)
				return nil, decodeErr
			}
			return answer, nil
		} else {
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			errStr := fmt.Sprintf("err: %s", bodyBytes)
			return nil, errors.New(errStr)
		}
	} else {
		errStr := fmt.Sprintf("coin: %s doesn't exist or is not present in the desktop configuration", coin)
		_ = glg.Errorf("%s", errStr)
		return nil, errors.New(errStr)
	}
}

func CustomMyTxHistory(coin string, defaultNbTx int, defaultPage int, withFiatValue bool, isMax bool, contract string,
	query string, address string, coinType string) (*mm2_data_structure.MyTxHistoryAnswer, error) {
	endpoint := customTxEndpoint
	if contract != "" {
		endpoint = endpoint + "v2/" + query + "/" + contract + "/" + address
	} else {
		endpoint = endpoint + "v1/" + query + "/" + address
	}
	resp, err := http.Get(endpoint)
	if err != nil {
		_ = glg.Errorf("error occured: %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	var cResp = new(mm2_data_structure.MyTxHistoryAnswer)
	if decodeErr := json.NewDecoder(resp.Body).Decode(cResp); decodeErr != nil {
		_ = glg.Errorf("error occured: %v", decodeErr)
		return nil, decodeErr
	}
	cResp.CoinType = coinType
	return cResp, nil
}
