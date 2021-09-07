package mm2_tools_server

import (
	"bytes"
	"encoding/json"
	"github.com/kpango/glg"
	"github.com/valyala/fasthttp"
	"mm2_client/constants"
	"mm2_client/mm2_tools_generics"
	"net/http"
)

func TickerInfos(ctx *fasthttp.RequestCtx) {
	if !constants.GPricesServicesRunning {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.SetBodyString("You need to start price services first")
		glg.Warn("You need to start service price first")
		return
	}

	if !constants.GDesktopCfgLoaded {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.SetBodyString("You need to load desktop cfg first (you can do that through start_price_service)")
		glg.Warn("You need to load desktop cfg first")
		return
	}
	out := &mm2_tools_generics.TickerInfosRequest{}
	r := bytes.NewReader(ctx.PostBody())
	err := json.NewDecoder(r).Decode(out)
	if err != nil {
		_ = glg.Errorf("%v", err)
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.SetBodyString(err.Error())
		return
	}
	resp := mm2_tools_generics.GetTickerInfos(out.Ticker, 0)
	ctx.SetStatusCode(200)
	ctx.SetBodyString(resp.ToJson())
	ctx.SetContentType("application/json")
}
