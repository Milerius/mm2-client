package mm2_tools_server

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func setResponseHeader(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
		h(ctx)
		return
	}
}

func InitRooter() *router.Router {
	r := router.New()
	r.POST("/api/v1/start_simple_market_maker_bot", setResponseHeader(StartSimpleMarketMakerBot))
	r.POST("/api/v1/stop_simple_market_maker_bot", setResponseHeader(StopSimpleMarketMakerBot))
	r.POST("/api/v1/start_price_service", setResponseHeader(StartPriceService))
	r.POST("/api/v1/ticker_infos", setResponseHeader(TickerInfos))
	r.GET("/api/v1/tickers", setResponseHeader(TickerAllInfos))
	r.POST("/api/v1/cex_rates", setResponseHeader(CexRates))
	r.POST("/api/v1/volume24h", setResponseHeader(Volume24h))
	return r
}
