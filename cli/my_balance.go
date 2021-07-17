package cli

import (
	"encoding/json"
	"fmt"
	"mm2_client/config"
	"mm2_client/http"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"mm2_client/mm2_tools_generics/mm2_http_request"
)

func MyBalance(coin string) {
	resp, err := mm2_http_request.MyBalance(coin)
	if resp != nil {
		resp.ToTable()
	} else {
		fmt.Println(err)
	}
}

func MyBalanceMultipleCoins(coins []string) {
	var outBatch []interface{}
	for _, v := range coins {
		if val, ok := config.GCFGRegistry[v]; ok {
			if req := mm2_http_request.NewMyBalanceCoinRequest(val); req != nil {
				outBatch = append(outBatch, req)
			}
		} else {
			fmt.Printf("coin %s doesn't exist - skipping\n", v)
		}
	}

	resp := http.BatchRequest(outBatch)
	if len(resp) > 0 {
		var outResp []mm2_data_structure.MyBalanceAnswer
		err := json.Unmarshal([]byte(resp), &outResp)
		if err != nil {
			fmt.Printf("Err: %v\n", err)
		} else {
			mm2_data_structure.ToTableMyBalanceAnswers(outResp)
		}
	}
}

func MyBalanceMultipleCoinsSilent(coins []string) []mm2_data_structure.MyBalanceAnswer {
	var outBatch []interface{}
	for _, v := range coins {
		if val, ok := config.GCFGRegistry[v]; ok {
			if req := mm2_http_request.NewMyBalanceCoinRequest(val); req != nil {
				outBatch = append(outBatch, req)
			}
		} else {
			fmt.Printf("coin %s doesn't exist - skipping\n", v)
		}
	}

	resp := http.BatchRequest(outBatch)
	if len(resp) > 0 {
		var outResp []mm2_data_structure.MyBalanceAnswer
		err := json.Unmarshal([]byte(resp), &outResp)
		if err != nil {
			fmt.Printf("Err: %v\n", err)
		} else {
			return outResp
		}
	}
	return nil
}
