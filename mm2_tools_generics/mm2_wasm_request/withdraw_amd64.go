package mm2_wasm_request

import (
	"errors"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
)

func Withdraw(coin string, amount string, address string, fees []string, coinType string) (*mm2_data_structure.WithdrawAnswer, error) {
	return nil, errors.New("not implemented on this platform")
}
