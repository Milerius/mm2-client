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

func GetEnabledCoins() (*mm2_data_structure.GetEnabledCoinsAnswer, error) {
	req := mm2_data_structure.NewGenericRequest("get_enabled_coins").ToJson()
	resp, err := http.Post(mm2_data_structure.GMM2Endpoint, "application/json", bytes.NewBuffer([]byte(req)))
	if err != nil {
		_ = glg.Errorf("err: %v", err)
		return nil, err
	}
	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		res := &mm2_data_structure.GetEnabledCoinsAnswer{}
		decodeErr := json.NewDecoder(resp.Body).Decode(res)
		if decodeErr != nil {
			_ = glg.Errorf("err: %v", err)
			return nil, decodeErr
		}
		return res, nil
	} else {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		errStr := fmt.Sprintf("err: %s\n", bodyBytes)
		return nil, errors.New(errStr)
	}
}
