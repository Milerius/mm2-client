package mm2_tools_generics

import (
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"mm2_client/mm2_tools_generics/mm2_http_request"
	"mm2_client/mm2_tools_generics/mm2_wasm_request"
	"runtime"
)

func Enable(coin string) (*mm2_data_structure.GenericEnableAnswer, error) {
	if runtime.GOARCH == "wasm" {
		return mm2_wasm_request.Enable(coin)
	} else {
		return mm2_http_request.Enable(coin)
	}
}
