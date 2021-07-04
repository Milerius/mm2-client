package cli

import (
	"fmt"
	"mm2_client/config"
	"mm2_client/http"
)

func Withdraw(coin string, address string, amount string, fees []string) {
	if val, ok := config.GCFGRegistry[coin]; ok {
		if len(fees) > 0 {
			switch val.Type {
			case "ERC-20", "BEP-20", "QRC-20":
				if len(fees) == 3 {
					http.Withdraw(coin, address, amount, fees, val.Type)
				} else {
					ShowCommandHelp("withdraw")
				}
			case "UTXO", "Smart Chain":
				if len(fees) == 2 {
					http.Withdraw(coin, address, amount, fees, val.Type)
				} else {
					ShowCommandHelp("withdraw")
				}
			}
		} else {
			http.Withdraw(coin, address, amount, []string{}, "")
		}
	} else {
		fmt.Printf("%s is not present in the cfg - skipping\n", coin)
	}
}
