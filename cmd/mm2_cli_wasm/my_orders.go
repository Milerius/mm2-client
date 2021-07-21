package main

import (
	"mm2_client/mm2_tools_generics"
	"syscall/js"
)

func myOrders() js.Func {
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		go func() { mm2_tools_generics.MyOrdersCLI() }()
		return "done"
	})
	return jsfunc
}
