package main

import (
	"mm2_client/mm2_tools_generics"
	"syscall/js"
)

func cancelOrder() js.Func {
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			result := map[string]interface{}{
				"help":  mm2_tools_generics.CancelOrderHelp,
				"usage": mm2_tools_generics.CancelOrderUsage,
			}
			return result
		}

		go func() {
			mm2_tools_generics.CancelOrderCLI(args[0].String())
		}()
		return "done"
	})
	return jsfunc
}
