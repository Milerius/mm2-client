package mm2_wasm_request

import (
	"encoding/json"
	"errors"
	"github.com/kpango/glg"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"syscall/js"
)

func MyRecentSwaps(limit string, pageNumber string, baseCoin string, relCoin string, from string, to string) (*mm2_data_structure.MyRecentSwapsAnswer, error) {
	req := mm2_data_structure.NewMyRecentSwapsRequest(limit, pageNumber, baseCoin, relCoin, from, to).ToJson()
	balVal, errVal := Await(js.Global().Call("rpc_request", req))
	if errVal != nil {
		return nil, errors.New(errVal[0].String())
	} else {
		var answer = &mm2_data_structure.MyRecentSwapsAnswer{}
		decodeErr := json.Unmarshal([]byte(balVal[0].String()), answer)
		if decodeErr != nil {
			_ = glg.Errorf("Err: %v", decodeErr)
			return nil, decodeErr
		}
		return answer, nil
	}
}
