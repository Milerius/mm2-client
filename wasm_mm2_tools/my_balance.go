package main

import (
	"fmt"
	"github.com/kpango/glg"
	"mm2_client/mm2_tools_generics"
	"syscall/js"
)

func internalMyBalance(coin string) {
	resp, err := mm2_tools_generics.MyBalance(coin)
	if resp != nil {
		resp.ToTable()
	} else {
		fmt.Println(err)
	}
}

func MyBalance() js.Func {
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 1 {
			usage := "invalid nb args - usage: my_balance(coin1, coin2, coin3)"
			_ = glg.Error(usage)
			result := map[string]interface{}{
				"error": usage,
			}
			return result
		}
		if len(args) == 1 {
			go internalMyBalance(args[0].String())
		} else {
			glg.Infof("Not implemented yet")
		}
		return nil
	})
	return jsfunc
}
