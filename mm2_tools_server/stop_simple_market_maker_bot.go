package mm2_tools_server

import (
	"github.com/kpango/glg"
	"github.com/valyala/fasthttp"
	"mm2_client/http"
	"mm2_client/market_making"
)

func StopSimpleMarketMakerBot(ctx *fasthttp.RequestCtx) {
	err := market_making.StopSimpleMarketMakerBotService()
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		_ = glg.Errorf("Error during initialization: %v", err)
		ctx.SetBodyString(err.Error())
		http.GRuntimeUserpass = ""
		return
	}
	ctx.SetStatusCode(200)
	ctx.SetBodyString("Successfully stopped")
	http.GRuntimeUserpass = ""
}
