package main

import (
	"fmt"
	"mm2_client/mm2_tools_generics"
	"strconv"
	"syscall/js"
)

func myOrders() js.Func {
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		go func() {
			if len(args) == 0 {
				mm2_tools_generics.MyOrdersCLI(false)
			} else if len(args) == 1 {
				val, err := strconv.ParseBool(args[0].String())
				if err == nil {
					mm2_tools_generics.MyOrdersCLI(val)
				} else {
					fmt.Println(err)
				}
			}
		}()
		return "done"
	})
	return jsfunc
}
