package main

import (
	"github.com/kpango/glg"
	"mm2_client/market_making"
	"syscall/js"
)

func startSimpleMarketMakerBot() js.Func {
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		go func() {
			err := market_making.StartSimpleMarketMakerBot("http://localhost:8080/static/assets/simple_market_bot.json", "url")
			if err != nil {
				_ = glg.Errorf("err when starting simple market maker bot: %v", err)
			}
		}()
		return "done"
	})
	return jsfunc
}

func stopSimpleMarketMakerBot() js.Func {
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		go func() {
			err := market_making.StopSimpleMarketMakerBotService()
			if err != nil {
				_ = glg.Errorf("err when stoping simple market maker bot: %v", err)
			}
		}()
		return "done"
	})
	return jsfunc
}
