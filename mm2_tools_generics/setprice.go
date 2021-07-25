package mm2_tools_generics

import (
	"fmt"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"mm2_client/mm2_tools_generics/mm2_http_request"
	"mm2_client/mm2_tools_generics/mm2_wasm_request"
	"runtime"
)

func SetPrice(base string, rel string, price string, volume *string, max *bool, cancelPrevious bool, minVolume *string,
	baseConfs *int, baseNota *bool, relConfs *int, relNota *bool) (*mm2_data_structure.SetPriceAnswer, error) {
	if runtime.GOARCH == "wasm" {
		return mm2_wasm_request.SetPrice(base, rel, price, volume, max, cancelPrevious, minVolume, baseConfs, baseNota, relConfs, relNota)
	} else {
		return mm2_http_request.SetPrice(base, rel, price, volume, max, cancelPrevious, minVolume, baseConfs, baseNota, relConfs, relNota)
	}
}

func SetPriceCLI(base string, rel string, price string, volume *string, max *bool, cancelPrevious bool, minVolume *string, baseConfs *int, baseNota *bool, relConfs *int, relNota *bool) {
	if resp, err := SetPrice(base, rel, price, volume, max, cancelPrevious, minVolume, baseConfs, baseNota, relConfs, relNota); resp != nil {
		resp.ToTable()
	} else {
		fmt.Println(err)
	}
}
