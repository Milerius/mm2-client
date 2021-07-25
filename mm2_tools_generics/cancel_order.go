package mm2_tools_generics

import (
	"fmt"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"mm2_client/mm2_tools_generics/mm2_http_request"
	"mm2_client/mm2_tools_generics/mm2_wasm_request"
	"runtime"
)

func CancelOrder(uuid string) (*mm2_data_structure.CancelOrderAnswer, error) {
	if runtime.GOARCH == "wasm" {
		return mm2_wasm_request.CancelOrder(uuid)
	} else {
		return mm2_http_request.CancelOrder(uuid)
	}
}

func CancelOrderCLI(uuid string) {
	if resp, cancelErr := CancelOrder(uuid); resp != nil {
		fmt.Println(resp.Result)
	} else {
		fmt.Println(cancelErr)
	}
}
