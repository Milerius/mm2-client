package mm2_tools_generics

import (
	"fmt"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"mm2_client/mm2_tools_generics/mm2_http_request"
	"mm2_client/mm2_tools_generics/mm2_wasm_request"
	"runtime"
)

func MyOrders() (*mm2_data_structure.MyOrdersAnswer, error) {
	if runtime.GOARCH == "wasm" {
		return mm2_wasm_request.MyOrders()
	} else {
		return mm2_http_request.MyOrders()
	}
}

func MyOrdersCLI(withFees bool) {
	if resp, err := MyOrders(); resp != nil {
		resp.ToTable(withFees)
	} else {
		fmt.Println(err)
	}
}
