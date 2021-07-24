package mm2_wasm_request

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kpango/glg"
	"mm2_client/config"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"syscall/js"
)

func Orderbook(base string, rel string) (*mm2_data_structure.OrderbookAnswer, error) {
	_, baseOk := config.GCFGRegistry[base]
	_, relOk := config.GCFGRegistry[rel]
	if relOk && baseOk {
		req := mm2_data_structure.NewOrderbookRequest(base, rel).ToJson()
		balVal, errVal := Await(js.Global().Call("rpc_request", req))
		if errVal != nil {
			return nil, errors.New(errVal[0].String())
		} else {
			var answer = &mm2_data_structure.OrderbookAnswer{}
			decodeErr := json.Unmarshal([]byte(balVal[0].String()), answer)
			if decodeErr != nil {
				_ = glg.Errorf("Err: %v", decodeErr)
				return nil, decodeErr
			}
			return answer, nil
		}
	} else {
		errStr := fmt.Sprintf("coin: %s or %s doesn't exist or is not present in the desktop configuration", base, rel)
		return nil, errors.New(errStr)
	}
}
