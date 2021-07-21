package mm2_wasm_request

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kpango/glg"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"syscall/js"
)

func KmdRewardsInfo() (*mm2_data_structure.KMDRewardsInfoAnswer, error) {
	answerEnabled, _ := GetEnabledCoins()
	if answerEnabled != nil && answerEnabled.Contains("KMD") {
		req := mm2_data_structure.NewGenericRequest("get_enabled_coins").ToJson()
		balVal, errVal := Await(js.Global().Call("rpc_request", req))
		if errVal != nil {
			return nil, errors.New(errVal[0].String())
		} else {
			var answer = &mm2_data_structure.KMDRewardsInfoAnswer{}
			decodeErr := json.Unmarshal([]byte(balVal[0].String()), answer)
			if decodeErr != nil {
				_ = glg.Errorf("err: %v", decodeErr)
				return nil, decodeErr
			}
			return answer, nil
		}
	} else {
		fmt.Println("KMD need to be enabled in order to call KmdRewardsInfo")
		return nil, errors.New("KMD need to be enabled in order to call KmdRewardsInfo")
	}
}
