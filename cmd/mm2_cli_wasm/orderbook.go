package main

import (
	"mm2_client/mm2_tools_generics"
	"syscall/js"
)

func orderbook() js.Func {
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 2 {
			result := map[string]interface{}{
				"help":  mm2_tools_generics.OrderbookHelp,
				"usage": mm2_tools_generics.OrderbookUsage,
			}
			return result
		}

		go func() {
			mm2_tools_generics.OrderbookCLI(args[0].String(), args[1].String())
		}()
		return "done"
	})
	return jsfunc
}
