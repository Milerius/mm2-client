package main

import (
	"mm2_client/mm2_tools_generics"
	"syscall/js"
)

func myRecentSwaps() js.Func {
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) > 6 {
			result := map[string]interface{}{
				"help":  mm2_tools_generics.MyRecentSwapsHelp,
				"usage": mm2_tools_generics.MyRecentSwapsUsage,
			}
			return result
		}

		functor := func(limit string, pageNumber string, extraArgs []string) {
			mm2_tools_generics.MyRecentSwapsCLI(limit, pageNumber, extraArgs)
		}

		if len(args) == 0 {
			go functor("50", "1", []string{})
		} else if len(args) == 1 {
			go functor(args[0].String(), "1", []string{})
		} else if len(args) == 2 {
			go functor(args[0].String(), args[1].String(), []string{})
		} else if len(args) > 2 {
			var extraArgs []string
			for _, cur := range args[2:] {
				extraArgs = append(extraArgs, cur.String())
			}
			go functor(args[0].String(), args[1].String(), extraArgs)
		}
		return "done"
	})
	return jsfunc
}
