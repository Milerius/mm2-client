package mm2_tools_generics

import (
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"mm2_client/mm2_tools_generics/mm2_http_request"
	"mm2_client/mm2_tools_generics/mm2_wasm_request"
	"runtime"
)

func Broadcast(coin string, txHex string) (*mm2_data_structure.BroadcastAnswer, error) {
	if runtime.GOARCH == "wasm" {
		return mm2_wasm_request.Broadcast(coin, txHex)
	} else {
		return mm2_http_request.Broadcast(coin, txHex)
	}
}
