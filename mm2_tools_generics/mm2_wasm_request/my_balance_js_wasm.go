package mm2_wasm_request

import (
	"encoding/json"
	"errors"
	"github.com/kpango/glg"
	"mm2_client/config"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"syscall/js"
)

func MyBalance(coin string) (*mm2_data_structure.MyBalanceAnswer, error) {
	if val, ok := config.GCFGRegistry[coin]; ok {
		req := mm2_data_structure.NewMyBalanceCoinRequest(val).ToJson()
		balVal, errVal := Await(js.Global().Call("rpc_request", req))
		if errVal != nil {
			return nil, errors.New(errVal[0].String())
		} else {
			var answer = &mm2_data_structure.MyBalanceAnswer{}
			decodeErr := json.Unmarshal([]byte(balVal[0].String()), answer)
			if decodeErr != nil {
				glg.Errorf("Err: %v", decodeErr)
				return nil, decodeErr
			}
			return answer, nil
		}
	} else {
		return nil, errors.New("coin " + coin + " not found in config - skipping")
	}
}
