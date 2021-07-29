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

func TradePreimage(base string, rel string, price string, method string, volumeOrMax string) (*mm2_data_structure.TradePreimageAnswer, error) {
	req := mm2_data_structure.NewTradePreimageRequest(base, rel, price, method, volumeOrMax).ToJson()
	resp, err := http.Post(mm2_data_structure.GMM2Endpoint, "application/json", bytes.NewBuffer([]byte(req)))
	if err != nil {
		_ = glg.Errorf("%v", err)
		return nil, err
	}
	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		var answer = &mm2_data_structure.TradePreimageAnswer{}
		decodeErr := json.NewDecoder(resp.Body).Decode(answer)
		if decodeErr != nil {
			_ = glg.Errorf("decode err: %v", err)
			return nil, decodeErr
		}
		return answer, nil
	} else {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		errStr := fmt.Sprintf("%s", bodyBytes)
		return nil, errors.New(errStr)
	}
}
