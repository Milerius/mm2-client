package main

import (
	"fmt"
	"github.com/kpango/glg"
	"mm2_client/mm2_tools_generics"
	"syscall/js"
)

func withdraw() js.Func {
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		//! withdraw("KMD", 1, address)
		glg.Infof("len args: %d", len(args))
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
			resp, err := mm2_tools_generics.WithdrawCLI(args[0].String(), args[1].String(), args[2].String(), fees)
			if err != nil {
				_ = glg.Errorf("%v", err)
			} else {
				resp.ToTable()
				res := js.Global().Call("confirm", "Do you want to broadcast the transaction")
				if res.Bool() {
					mm2_tools_generics.BroadcastCLI(resp.Coin, resp.TxHex)
				} else {
					fmt.Println(err)
				}
			}
		}()
		return "done"
	})
	return jsfunc
}
