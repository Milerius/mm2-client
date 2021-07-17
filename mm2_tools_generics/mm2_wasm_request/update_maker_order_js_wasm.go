package mm2_wasm_request

import "mm2_client/mm2_tools_generics/mm2_data_structure"

func UpdateMakerOrder(uuid string, newPrice *string, volumeDelta *string, max *bool, minVolume *string,
	baseConfs *int, baseNota *bool, relConfs *int, relNota *bool) (*mm2_data_structure.UpdateMakerOrderAnswer, error) {
	return nil, errors.New("not implemented on this platform")
}
