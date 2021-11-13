package main

import (
	"github.com/kpango/glg"
	"mm2_client/config"
	"mm2_client/mm2_tools_generics"
	"syscall/js"
)

func startSimpleMarketMakerBot() js.Func {
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		go func() {
			err := config.ParseMarketMakerConfFromUrl("http://localhost:8080/static/assets/mm2_market_maker.json")
			if err != nil {
				_ = glg.Errorf("err when starting simple market maker bot: %v", err)
			}
			mm2_tools_generics.StartSimpleMarketMakerBotCLI()
		}()
		return "done"
	})
	return jsfunc
}

func stopSimpleMarketMakerBot() js.Func {
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		go func() {
			mm2_tools_generics.StopSimpleMarketMakerBotCLI()
		}()
		return "done"
	})
	return jsfunc
}
