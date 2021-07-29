package mm2_wasm_request

import (
	"encoding/json"
	"errors"
	"github.com/kpango/glg"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"syscall/js"
)

func TradePreimage(base string, rel string, price string, method string, volumeOrMax string) (*mm2_data_structure.TradePreimageAnswer, error) {
	req := mm2_data_structure.NewTradePreimageRequest(base, rel, price, method, volumeOrMax).ToJson()
	tradePreimageVal, errVal := Await(js.Global().Call("rpc_request", req))
	if errVal != nil {
		return nil, errors.New(errVal[0].String())
	} else {
		var answer = &mm2_data_structure.TradePreimageAnswer{}
		decodeErr := json.Unmarshal([]byte(tradePreimageVal[0].String()), answer)
		if decodeErr != nil {
			_ = glg.Errorf("Err: %v", decodeErr)
			return nil, decodeErr
		}
		return answer, nil
	}
}
