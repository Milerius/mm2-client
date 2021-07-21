package main

import (
	"github.com/kpango/glg"
	"mm2_client/mm2_tools_generics"
	"syscall/js"
)

func disableCoin() js.Func {
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			usage := "invalid nb args - usage: disable_coin(coin1, coin2, coin3...)"
			_ = glg.Error(usage)
			result := map[string]interface{}{
				"error": usage,
			}
			return result
		}
		if len(args) == 1 {
			go func() { mm2_tools_generics.DisableCoinCLI(args[0].String()) }()
		} else {
			var out []string
			for _, cur := range args {
				out = append(out, cur.String())
			}
			go func() { mm2_tools_generics.DisableCoins(out) }()
		}
		return "done"
	})
	return jsfunc
}

func disableEnabledCoins() js.Func {
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		go func() {
			val, _ := mm2_tools_generics.GetEnabledCoins()
			mm2_tools_generics.DisableCoins(val.ToSlice())
		}()
		return "done"
	})
	return jsfunc
}
