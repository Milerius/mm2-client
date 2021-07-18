package mm2_tools_generics

import (
	"encoding/json"
	"fmt"
	"mm2_client/config"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"mm2_client/mm2_tools_generics/mm2_http_request"
	"mm2_client/mm2_tools_generics/mm2_wasm_request"
	"runtime"
)

func MyBalance(coin string) (*mm2_data_structure.MyBalanceAnswer, error) {
	if runtime.GOARCH == "wasm" {
		return mm2_wasm_request.MyBalance(coin)
	} else {
		return mm2_http_request.MyBalance(coin)
	}
}

func MyBalanceCLI(coin string) {
	resp, err := MyBalance(coin)
	if resp != nil {
		resp.ToTable()
	} else {
		fmt.Println(err)
	}
}

func MyBalanceMultipleCoinsSilent(coins []string) []mm2_data_structure.MyBalanceAnswer {
	var outBatch []interface{}
	for _, v := range coins {
		if val, ok := config.GCFGRegistry[v]; ok {
			if req := mm2_data_structure.NewMyBalanceCoinRequest(val); req != nil {
				outBatch = append(outBatch, req)
			}
		} else {
			fmt.Printf("coin %s doesn't exist - skipping\n", v)
		}
	}

	resp := BatchRequest(outBatch)
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

func MyBalanceMultipleCoinsCLI(coins []string) {
	resp := MyBalanceMultipleCoinsSilent(coins)
	if len(resp) > 0 {
		mm2_data_structure.ToTableMyBalanceAnswers(resp)
	}
}
