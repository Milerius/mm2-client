package mm2_wasm_request

import (
	"encoding/json"
	"errors"
	"github.com/kpango/glg"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"syscall/js"
)

func CancelOrder(uuid string) (*mm2_data_structure.CancelOrderAnswer, error) {
	req := mm2_data_structure.NewCancelOrderRequest(uuid).ToJson()
	balVal, errVal := Await(js.Global().Call("rpc_request", req))
	if errVal != nil {
		return nil, errors.New(errVal[0].String())
	} else {
		var answer = &mm2_data_structure.CancelOrderAnswer{}
		decodeErr := json.Unmarshal([]byte(balVal[0].String()), answer)
		if decodeErr != nil {
			_ = glg.Errorf("Err: %v", decodeErr)
			return nil, decodeErr
		}
		return answer, nil
	}
}
