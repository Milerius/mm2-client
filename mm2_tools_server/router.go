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
	//r.GET("/api/v1/eth_tx_history/{address}", setResponseHeader(EthTransactionsHistory))
	//r.GET("/api/v1/ohlc/tickers_list", OHLCGetAvailablePairs)
	return r
}
