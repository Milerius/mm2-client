package mm2_tools_generics

import (
	"mm2_client/mm2_tools_generics/mm2_http_request"
	"mm2_client/mm2_tools_generics/mm2_wasm_request"
	"runtime"
)

func BatchRequest(batch []interface{}) string {
	if runtime.GOARCH == "wasm" {
		return mm2_wasm_request.BatchRequest(batch)
	} else {
		return mm2_http_request.BatchRequest(batch)
	}
}
