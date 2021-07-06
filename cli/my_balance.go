package cli

import (
	"encoding/json"
	"fmt"
	"mm2_client/config"
	"mm2_client/http"
)

func MyBalance(coin string) {
	resp := http.MyBalance(coin)
	if resp != nil {
		resp.ToTable()
	}
}

func MyBalanceMultipleCoins(coins []string) {
	var outBatch []interface{}
	for _, v := range coins {
		if val, ok := config.GCFGRegistry[v]; ok {
			if req := http.NewMyBalanceCoinRequest(val); req != nil {
				outBatch = append(outBatch, req)
			}
		} else {
			fmt.Printf("coin %s doesn't exist - skipping\n", v)
		}
	}

	resp := http.BatchRequest(outBatch)
	if len(resp) > 0 {
		var outResp []http.MyBalanceAnswer
		err := json.Unmarshal([]byte(resp), &outResp)
		if err != nil {
			fmt.Printf("Err: %v\n", err)
		} else {
			http.ToTableMyBalanceAnswers(outResp)
		}
	}
}

func MyBalanceMultipleCoinsSilent(coins []string) []http.MyBalanceAnswer {
	var outBatch []interface{}
	for _, v := range coins {
		if val, ok := config.GCFGRegistry[v]; ok {
			if req := http.NewMyBalanceCoinRequest(val); req != nil {
				outBatch = append(outBatch, req)
			}
		} else {
			fmt.Printf("coin %s doesn't exist - skipping\n", v)
		}
	}

	resp := http.BatchRequest(outBatch)
	if len(resp) > 0 {
		var outResp []http.MyBalanceAnswer
		err := json.Unmarshal([]byte(resp), &outResp)
		if err != nil {
			fmt.Printf("Err: %v\n", err)
		} else {
			return outResp
		}
	}
	return nil
}
