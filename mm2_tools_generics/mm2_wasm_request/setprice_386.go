package mm2_wasm_request

import (
	"errors"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
)

func SetPrice(base string, rel string, price string, volume *string, max *bool, cancelPrevious bool, minVolume *string,
	baseConfs *int, baseNota *bool, relConfs *int, relNota *bool) (*mm2_data_structure.SetPriceAnswer, error) {
	return nil, errors.New("not implemented")
}
