package mm2_tools_generics

import (
	"fmt"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"mm2_client/mm2_tools_generics/mm2_http_request"
	"mm2_client/mm2_tools_generics/mm2_wasm_request"
	"runtime"
)

func Orderbook(base string, rel string) (*mm2_data_structure.OrderbookAnswer, error) {
	if runtime.GOARCH == "wasm" {
		return mm2_wasm_request.Orderbook(base, rel)
	} else {
		return mm2_http_request.Orderbook(base, rel)
	}
}

func OrderbookCLI(base string, rel string) {
	if resp, err := mm2_http_request.Orderbook(base, rel); resp != nil {
		resp.ToTable(base, rel)
	} else {
		fmt.Println(err)
	}
}
