package main

import (
	"mm2_client/mm2_tools_generics"
	"syscall/js"
)

func send() js.Func {
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 3 {
			result := map[string]interface{}{
				"help":  mm2_tools_generics.WithdrawHelp,
				"usage": mm2_tools_generics.WithdrawUsage,
			}
			return result
		}

		var fees []string
		if len(args) > 3 {
			for _, cur := range args[3:] {
				fees = append(fees, cur.String())
			}
		}
		go func() {
			mm2_tools_generics.Send(args[0].String(), args[1].String(), args[2].String(), fees)
		}()
		return "done"
	})
	return jsfunc
}
