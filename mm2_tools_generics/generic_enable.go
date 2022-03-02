package mm2_tools_generics

import (
	"encoding/json"
	"fmt"
	"mm2_client/config"
	"mm2_client/http"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
)

func EnableCLI(coin string) {
	var resp *mm2_data_structure.GenericEnableAnswer = nil
	var err error = nil
	if val, ok := config.GCFGRegistry[coin]; ok {
		switch val.Type {
		case "BEP-20", "ERC-20", "Arbitrum", "Optimism", "Matic", "FTM-20", "HRC-20", "AVX-20", "HecoChain", "Moonriver", "Moonbeam", "KRC-20", "Ubiq", "Ethereum Classic", "SmartBCH":
			resp, err = Enable(coin)
		case "UTXO", "QRC-20", "Smart Chain":
			resp, err = Electrum(coin)
		default:
			fmt.Println("Not supported yet")
		}
		if !val.Active && err == nil {
			val.Active = true
			config.GCFGRegistry[coin] = val
			go config.Update(http.GetLastDesktopVersion())
		}
		if resp != nil {
			resp.ToTable()
		} else if err != nil {
			fmt.Printf("Err when enabling coin %s: %v\n", coin, err)
		}
	} else {
		fmt.Printf("Cannot enable %s, did you start MM2 with the command <start> ?\n", coin)
	}
}

func EnableMultipleCoins(coins []string) {
	var outBatch []interface{}
	requireUpdate := false
	for _, v := range coins {
		if val, ok := config.GCFGRegistry[v]; ok {
			switch val.Type {
			case "BEP-20", "ERC-20", "Arbitrum", "Optimism", "Matic", "FTM-20", "HRC-20", "AVX-20", "HecoChain", "Moonriver", "Moonbeam", "KRC-20", "Ubiq", "Ethereum Classic", "SmartBCH":
				req := mm2_data_structure.NewEnableRequest(val)
				//fmt.Println(req)
				outBatch = append(outBatch, req)
				if !val.Active {
					val.Active = true
					config.GCFGRegistry[v] = val
					requireUpdate = true
				}
			case "UTXO", "QRC-20", "Smart Chain":
				req := mm2_data_structure.NewElectrumRequest(val)
				if req != nil {
					//fmt.Println(req.ToJson())
					outBatch = append(outBatch, req)
					if !val.Active {
						val.Active = true
						config.GCFGRegistry[v] = val
						requireUpdate = true
					}
				} else {
					fmt.Printf("Skipping for %s - because no available electrum servers for your platform\n", val.Coin)
				}
			default:
				fmt.Println("Not supported yet")
			}
		} else {
			fmt.Printf("coin %s doesn't exist - skipping\n", v)
		}
	}

	if requireUpdate {
		go config.Update(http.GetLastDesktopVersion())
	}

	resp := BatchRequest(outBatch)
	if len(resp) > 0 {
		var outResp []mm2_data_structure.GenericEnableAnswer
		err := json.Unmarshal([]byte(resp), &outResp)
		if err != nil {
			fmt.Printf("Err: %v\n", err)
		} else {
			mm2_data_structure.ToTableGenericEnableAnswers(outResp)
		}
	}
}
