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

func KmdRewardsInfo() (*mm2_data_structure.KMDRewardsInfoAnswer, error) {
	answerEnabled, _ := GetEnabledCoins()
	if answerEnabled != nil && answerEnabled.Contains("KMD") {
		req := mm2_data_structure.NewGenericRequest("kmd_rewards_info").ToJson()
		resp, err := http.Post(mm2_data_structure.GMM2Endpoint, "application/json", bytes.NewBuffer([]byte(req)))
		if err != nil {
			_ = glg.Errorf("err: %v\n", err)
			return nil, err
		}
		if resp.StatusCode == http.StatusOK {
			defer resp.Body.Close()
			var answer = &mm2_data_structure.KMDRewardsInfoAnswer{}
			decodeErr := json.NewDecoder(resp.Body).Decode(answer)
			if decodeErr != nil {
				_ = glg.Errorf("err: %v", err)
				return nil, decodeErr
			}
			return answer, nil
		} else {
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			errStr := fmt.Sprintf("err: %s", bodyBytes)
			return nil, errors.New(errStr)
		}
	} else {
		fmt.Println("KMD need to be enabled in order to call KmdRewardsInfo")
		return nil, errors.New("KMD need to be enabled in order to call KmdRewardsInfo")
	}
}
