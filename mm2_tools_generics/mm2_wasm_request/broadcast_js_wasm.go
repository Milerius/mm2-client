package mm2_wasm_request

import (
	"encoding/json"
	"errors"
	"github.com/kpango/glg"
	"mm2_client/config"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"syscall/js"
)

func Broadcast(coin string, txHex string) (*mm2_data_structure.BroadcastAnswer, error) {
	if val, ok := config.GCFGRegistry[coin]; ok {
		req := mm2_data_structure.NewBroadcastRequest(coin, txHex).ToJson()
		balVal, errVal := Await(js.Global().Call("rpc_request", req))
		if errVal != nil {
			return nil, errors.New(errVal[0].String())
		} else {
			var answer = &mm2_data_structure.BroadcastAnswer{}
			decodeErr := json.Unmarshal([]byte(balVal[0].String()), answer)
			if decodeErr != nil {
				_ = glg.Errorf("Err: %v", decodeErr)
				return nil, decodeErr
			}
			if val.ExplorerTxURL != "" {
				answer.TxUrl = val.ExplorerURL + val.ExplorerTxURL + answer.TxHash
			} else {
				answer.TxUrl = val.ExplorerURL + "tx/" + answer.TxHash
			}
			return answer, nil
		}
	} else {
		return nil, errors.New("coin " + coin + " not found in config - skipping")
	}
}
