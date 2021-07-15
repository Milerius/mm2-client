package mm2_tools_server

import (
	"encoding/json"
	"github.com/kpango/glg"
	"github.com/valyala/fasthttp"
	"mm2_client/config"
	"mm2_client/constants"
	"net/http"
)

func TickerAllInfos(ctx *fasthttp.RequestCtx) {
	if !constants.GPricesServicesRunning {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.SetBodyString("You need to start price services first")
		return
	}

	if !constants.GDesktopCfgLoaded {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.SetBodyString("You need to load desktop cfg first (you can do that through start_price_service)")
		return
	}

	var out = make(map[string]*TickerInfosAnswer)
	for _, cur := range config.GCFGRegistry {
		resp := GetTickerInfos(cur.Coin)
		out[cur.Coin] = resp
	}
	b, err := json.Marshal(out)
	if err != nil {
		_ = glg.Errorf("err during ticker_all_infos_request: %v", err)
		ctx.SetStatusCode(http.StatusInternalServerError)
		ctx.SetBodyString(err.Error())
		return
	}
	ctx.SetStatusCode(200)
	ctx.SetBodyString(string(b))
	ctx.SetContentType("application/json")
}
