package mm2_wasm_request

import (
	"encoding/json"
	"github.com/kpango/glg"
	"syscall/js"
)

func BatchRequest(batch []interface{}) string {
	p, _ := json.Marshal(batch)
	respVal, errVal := Await(js.Global().Call("rpc_request", string(p)))
	if respVal != nil {
		resp := respVal[0].String()
		return resp
	} else {
		glg.Errorf("%v", errVal)
		return ""
	}
}
