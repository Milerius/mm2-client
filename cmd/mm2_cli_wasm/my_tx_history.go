package main

import (
	"mm2_client/mm2_tools_generics"
	"syscall/js"
)

func myTxHistory() js.Func {
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 1 {
			result := map[string]interface{}{
				"help":  mm2_tools_generics.MyOrdersHelp,
				"usage": mm2_tools_generics.MyOrdersUsage,
			}
			return result
		}

		var argsGo []string
		if len(args) > 1 {
			for _, cur := range args[1:] {
				argsGo = append(argsGo, cur.String())
			}
		}
		go func() {
			mm2_tools_generics.MyTxHistoryCLI(args[0].String(), argsGo)
		}()
		return "done"
	})
	return jsfunc
}
