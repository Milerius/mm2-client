package mm2_tools_generics

import (
	"errors"
	"fmt"
	"mm2_client/config"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"mm2_client/mm2_tools_generics/mm2_http_request"
	"mm2_client/mm2_tools_generics/mm2_wasm_request"
	"runtime"
)

func Withdraw(coin string, amount string, address string, fees []string, coinType string) (*mm2_data_structure.WithdrawAnswer, error) {
	if runtime.GOARCH == "wasm" {
		return mm2_wasm_request.Withdraw(coin, amount, address, fees, coinType)
	} else {
		return mm2_http_request.Withdraw(coin, amount, address, fees, coinType)
	}
}

func WithdrawCLI(coin string, amount string, address string, fees []string) (*mm2_data_structure.WithdrawAnswer, error) {
	if val, ok := config.GCFGRegistry[coin]; ok {
		var resp *mm2_data_structure.WithdrawAnswer = nil
		var err error = nil
		if len(fees) > 0 {
			switch val.Type {
			case "ERC-20", "BEP-20", "QRC-20":
				if len(fees) == 3 {
					resp, err = Withdraw(coin, amount, address, fees, val.Type)
				} else {
					ShowCommandHelp("withdraw")
				}
			case "UTXO", "Smart Chain":
				if len(fees) == 2 {
					resp, err = Withdraw(coin, amount, address, fees, val.Type)
				} else {
					ShowCommandHelp("withdraw")
				}
			}
		} else {
			resp, err = Withdraw(coin, amount, address, []string{}, "")
		}
		return resp, err
	} else {
		fmt.Printf("%s is not present in the cfg - skipping\n", coin)
		return nil, errors.New(coin + " is not present in the cfg - skipping")
	}
}
