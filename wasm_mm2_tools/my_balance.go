package main

import (
	"encoding/json"
	"fmt"
	"github.com/kpango/glg"
	"mm2_client/config"
	"mm2_client/mm2_tools_generics"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"mm2_client/mm2_tools_generics/mm2_wasm_request"
	"syscall/js"
)

func internalMyBalance(coin string) {
	resp, err := mm2_tools_generics.MyBalance(coin)
	if resp != nil {
		resp.ToTable()
	} else {
		fmt.Println(err)
	}
}

func internalMultipleBalance(coins []string) {
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
	p, _ := json.Marshal(outBatch)
	respVal, errVal := mm2_wasm_request.Await(js.Global().Call("rpc_request", string(p)))
	if respVal != nil {
		resp := respVal[0].String()
		if len(resp) > 0 {
			var outResp []mm2_data_structure.MyBalanceAnswer
			err := json.Unmarshal([]byte(resp), &outResp)
			if err != nil {
				fmt.Printf("Err: %v\n", err)
			} else {
				mm2_data_structure.ToTableMyBalanceAnswers(outResp)
			}
		}
	} else {
		fmt.Println(errVal)
	}
}

func MyBalance() js.Func {
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 1 {
			usage := "invalid nb args - usage: my_balance(coin1, coin2, coin3)"
			_ = glg.Error(usage)
			result := map[string]interface{}{
				"error": usage,
			}
			return result
		}
		if len(args) == 1 {
			go internalMyBalance(args[0].String())
		} else {
			glg.Infof("Not implemented yet")
			var coins []string
			for _, cur := range args {
				coins = append(coins, cur.String())
			}
			go internalMultipleBalance(coins)
		}
		return nil
	})
	return jsfunc
}

func MyBalanceAll() js.Func {
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var coins []string
		for _, cur := range config.GCFGRegistry {
			if cur.Active {
				coins = append(coins, cur.Coin)
			}
		}
		go internalMultipleBalance(coins)
		return nil
	})
	return jsfunc
}
