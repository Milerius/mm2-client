package mm2_tools_generics

import (
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"mm2_client/mm2_tools_generics/mm2_wasm_request"
	"runtime"
)

func TradePreimage(base string, rel string, price string, method string, volumeOrMax string) (*mm2_data_structure.TradePreimageAnswer, error) {
	if runtime.GOARCH == "wasm" {
		return mm2_wasm_request.TradePreimage(base, rel, price, method, volumeOrMax)
	} else {
		return mm2_wasm_request.TradePreimage(base, rel, price, method, volumeOrMax)
	}
}
