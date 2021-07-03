package cli

import (
	"encoding/json"
	"fmt"
	"mm2_client/config"
	"mm2_client/helpers"
	"mm2_client/http"
)

func DisableCoin(coin string) {
	resp := http.DisableCoin(coin)
	if resp != nil {
		config.GCFGRegistry[coin].Active = false
		go config.Update(http.GetLastDesktopVersion())
		helpers.PrintCheck(coin+" successfully disabled", true)
	}
}

func DisableCoins(coins []string) {
	var outBatch []interface{}
	for _, v := range coins {
		if val, ok := config.GCFGRegistry[v]; ok {
			config.GCFGRegistry[v].Active = false
			outBatch = append(outBatch, http.NewDisableCoinRequest(val))
		} else {
			fmt.Printf("coin %s doesn't exist - skipping\n", v)
		}
	}
	if len(outBatch) > 0 {
		resp := http.BatchRequest(outBatch)
		if len(resp) > 0 {
			var outResp []http.DisableCoinAnswer
			err := json.Unmarshal([]byte(resp), &outResp)
			if err != nil {
				fmt.Printf("Err: %v\n", err)
			} else {
				go config.Update(http.GetLastDesktopVersion())
				for _, cur := range outResp {
					if len(cur.Error) == 0 {
						helpers.PrintCheck(cur.Result.Coin+" successfully disabled", true)
					} else {
						fmt.Println(cur.Error)
					}
				}
			}
		}
	} else {
		fmt.Println("None of the desired coins exists - skipping")
	}
}
