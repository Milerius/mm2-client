package cli

import (
	"encoding/json"
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
	var outBatch []interface{}
	for _, v := range coins {
		if val, ok := config.GCFGRegistry[v]; ok {
			switch val.Type {
			case "BEP-20", "ERC-20":
				outBatch = append(outBatch, http.NewEnableRequest(val))
			case "UTXO", "QRC-20", "Smart Chain":
				outBatch = append(outBatch, http.NewElectrumRequest(val))
			default:
				fmt.Println("Not supported yet")
			}
		} else {
			fmt.Printf("coin %s doesn't exist - skipping\n", v)
		}
	}

	resp := http.BatchRequest(outBatch)
	if len(resp) > 0 {
		var outResp []http.GenericEnableAnswer
		err := json.Unmarshal([]byte(resp), &outResp)
		if err != nil {
			fmt.Printf("Err: %v\n", err)
		} else {
			http.ToTable(outResp)
		}
	}
}
