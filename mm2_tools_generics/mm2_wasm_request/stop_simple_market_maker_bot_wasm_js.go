package mm2_wasm_request

import (
	"encoding/json"
	"github.com/kpango/glg"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"syscall/js"
)
import "errors"

func StopSimpleMarketMakerBot() (*mm2_data_structure.StopSimpleMarketMakerAnswer, error) {
	req := mm2_data_structure.NewStopSimpleMarketMakerRequest().ToJson()
	balVal, errVal := Await(js.Global().Call("rpc_request", req))
	if errVal != nil {
		return nil, errors.New(errVal[0].String())
	} else {
		var answer = &mm2_data_structure.StopSimpleMarketMakerAnswer{}
		decodeErr := json.Unmarshal([]byte(balVal[0].String()), answer)
		if decodeErr != nil {
			_ = glg.Errorf("err: %v", decodeErr)
			return nil, decodeErr
		}
		return answer, nil
	}
}
