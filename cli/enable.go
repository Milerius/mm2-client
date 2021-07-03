package cli

import (
	"fmt"
	"mm2_client/config"
	"mm2_client/http"
)

func Enable(coin string) {
	if val, ok := config.GCFGRegistry[coin]; ok {
		switch val.Type {
		case "BEP-20", "ERC-20":
			http.Enable(coin)
		case "UTXO", "QRC-20", "Smart Chain":
			http.Electrum(coin)
		default:
			fmt.Println("Not supported yet")
		}
	} else {
		fmt.Printf("Cannot enable %s, did you start MM2 with the command <start> ?\n", coin)
	}
}

func EnableMultipleCoins(coins []string) {
	var outCoins []string
	for _, v := range coins {
		if _, ok := config.GCFGRegistry[v]; ok {
			outCoins = append(outCoins, v)
		} else {
			fmt.Printf("coin %s doesn't exist - skipping\n", v)
		}
	}
	fmt.Printf("Will enable %v\n", outCoins)
}
