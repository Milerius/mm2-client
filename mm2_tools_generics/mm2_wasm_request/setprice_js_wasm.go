package mm2_wasm_request

import (
	"encoding/json"
	"errors"
	"github.com/kpango/glg"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"syscall/js"
)

func SetPrice(base string, rel string, price string, volume *string, max *bool, cancelPrevious bool, minVolume *string,
	baseConfs *int, baseNota *bool, relConfs *int, relNota *bool) (*mm2_data_structure.SetPriceAnswer, error) {
	req := mm2_data_structure.NewSetPriceRequest(base, rel, price, volume, max, cancelPrevious, minVolume, baseConfs, baseNota, relConfs, relNota).ToJson()
	balVal, errVal := Await(js.Global().Call("rpc_request", req))
	if errVal != nil {
		return nil, errors.New(errVal[0].String())
	} else {
		var answer = &mm2_data_structure.SetPriceAnswer{}
		decodeErr := json.Unmarshal([]byte(balVal[0].String()), answer)
		if decodeErr != nil {
			_ = glg.Errorf("err: %v", decodeErr)
			return nil, decodeErr
		}
		return answer, nil
	}
}
