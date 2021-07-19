package mm2_tools_generics

import (
	"encoding/json"
	"fmt"
	"mm2_client/config"
	"mm2_client/helpers"
	"mm2_client/http"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"mm2_client/mm2_tools_generics/mm2_http_request"
	"mm2_client/mm2_tools_generics/mm2_wasm_request"
	"runtime"
)

func DisableCoin(coin string) (*mm2_data_structure.DisableCoinAnswer, error) {
	if runtime.GOARCH == "wasm" {
		return mm2_wasm_request.DisableCoin(coin)
	} else {
		return mm2_http_request.DisableCoin(coin)
	}
}

func DisableCoinCLI(coin string) {
	resp, err := mm2_http_request.DisableCoin(coin)
	if resp != nil {
		config.GCFGRegistry[coin].Active = false
		go config.Update(http.GetLastDesktopVersion())
		helpers.PrintCheck(coin+" successfully disabled", true)
	} else if err != nil {
		fmt.Println(err)
	}
}

func DisableCoins(coins []string) {
	var outBatch []interface{}
	for _, v := range coins {
		if val, ok := config.GCFGRegistry[v]; ok {
			//fmt.Printf("%s became inactive\n", v)
			config.GCFGRegistry[v].Active = false
			outBatch = append(outBatch, mm2_data_structure.NewDisableCoinRequest(val))
		} else {
			fmt.Printf("coin %s doesn't exist - skipping\n", v)
		}
	}

	if len(outBatch) > 0 {
		resp := mm2_http_request.BatchRequest(outBatch)
		if len(resp) > 0 {
			var outResp []mm2_data_structure.DisableCoinAnswer
			err := json.Unmarshal([]byte(resp), &outResp)
			if err != nil {
				fmt.Printf("Err: %v\n", err)
			} else {
				for _, cur := range outResp {
					if len(cur.Error) == 0 {
						helpers.PrintCheck(cur.Result.Coin+" successfully disabled", true)
					} else {
						fmt.Println(cur.Error)
					}
				}
				config.Update(http.GetLastDesktopVersion())
			}
		}
	} else {
		fmt.Println("None of the desired coins exists - skipping")
	}
}
