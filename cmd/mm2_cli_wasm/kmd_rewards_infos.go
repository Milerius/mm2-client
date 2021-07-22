package main

import (
	"fmt"
	"mm2_client/mm2_tools_generics"
	"syscall/js"
)

func kmdRewardsInfos() js.Func {
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		go func() {
			if mm2_tools_generics.KmdRewardsInfoCLI() {
				res := js.Global().Call("confirm", "Do you want to claim the rewards")
				if res.Bool() {
					if respBalance, err := mm2_tools_generics.MyBalance("KMD"); respBalance != nil {
						mm2_tools_generics.Send("KMD", "max", respBalance.Address, []string{})
					} else {
						fmt.Println(err)
					}
				}
			}
		}()
		return "done"
	})
	return jsfunc
}
