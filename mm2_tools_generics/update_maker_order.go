package mm2_tools_generics

import (
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"mm2_client/mm2_tools_generics/mm2_http_request"
	"mm2_client/mm2_tools_generics/mm2_wasm_request"
	"runtime"
)

func UpdateMakerOrder(uuid string, newPrice *string, volumeDelta *string, max *bool, minVolume *string,
	baseConfs *int, baseNota *bool, relConfs *int, relNota *bool) (*mm2_data_structure.UpdateMakerOrderAnswer, error) {
	if runtime.GOARCH == "wasm" {
		return mm2_wasm_request.UpdateMakerOrder(uuid, newPrice, volumeDelta, max, minVolume, baseConfs, baseNota, relConfs, relNota)
	} else {
		return mm2_http_request.UpdateMakerOrder(uuid, newPrice, volumeDelta, max, minVolume, baseConfs, baseNota, relConfs, relNota)
	}
}
