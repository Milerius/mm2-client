package main

import (
	"github.com/kpango/glg"
	"mm2_client/config"
	"mm2_client/mm2_tools_generics"
	"syscall/js"
)

func internalMultipleBalance(coins []string) {
	mm2_tools_generics.MyBalanceMultipleCoinsCLI(coins)
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
			go mm2_tools_generics.MyBalanceCLI(args[0].String())
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
