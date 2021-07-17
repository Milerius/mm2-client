package mm2_tools_generics

import (
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"mm2_client/mm2_tools_generics/mm2_http_request"
	"mm2_client/mm2_tools_generics/mm2_wasm_request"
	"runtime"
)

func CancelAllOrders(kind string, args []string) (*mm2_data_structure.CancelAllOrdersAnswer, error) {
	if runtime.GOARCH == "wasm" {
		return mm2_wasm_request.CancelAllOrders(kind, args)
	} else {
		return mm2_http_request.CancelAllOrders(kind, args)
	}
}
