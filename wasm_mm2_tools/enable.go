package main

import (
	"github.com/kpango/glg"
	"mm2_client/config"
	"mm2_client/mm2_tools_generics"
	"syscall/js"
)

func enableActiveCoins() js.Func {
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		go func() { mm2_tools_generics.EnableMultipleCoins(config.RetrieveActiveCoins()) }()
		return "done"
	})
	return jsfunc
}

func enable() js.Func {
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			usage := "invalid nb args - usage: enable(coin1, coin2, coin3...)"
			_ = glg.Error(usage)
			result := map[string]interface{}{
				"error": usage,
			}
			return result
		}
		if len(args) == 1 {
			go func() { mm2_tools_generics.EnableCLI(args[0].String()) }()
		} else {
			var out []string
			for _, cur := range args {
				out = append(out, cur.String())
			}
			go func() { mm2_tools_generics.EnableMultipleCoins(out) }()
		}
		return "done"
	})
	return jsfunc
}
