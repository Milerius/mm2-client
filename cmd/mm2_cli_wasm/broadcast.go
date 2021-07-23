package main

import (
	"mm2_client/mm2_tools_generics"
	"syscall/js"
)

func broadcast() js.Func {
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		//! withdraw("KMD", 1, address)
		if len(args) != 2 {
			result := map[string]interface{}{
				"help":  mm2_tools_generics.BroadcastHelp,
				"usage": mm2_tools_generics.BroadcastUsage,
			}
			return result
		}

		go func() {
			mm2_tools_generics.BroadcastCLI(args[0].String(), args[1].String())
		}()
		return "done"
	})
	return jsfunc
}
