package mm2_tools_server

import (
	"encoding/json"
	"github.com/kpango/glg"
	"github.com/valyala/fasthttp"
	"mm2_client/config"
	"mm2_client/constants"
	"mm2_client/helpers"
	"mm2_client/mm2_tools_generics"
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

	args := ctx.QueryArgs()
	expirationPriceValidity := 0
	if args.Len() > 0 {
		glg.Info("len is above 0, there is an arg to this get request")
		expireAt, err := args.GetUint("expire_at")
		if err != nil {
			_ = glg.Errorf("error with request args: %v", err)
			ctx.SetStatusCode(http.StatusBadRequest)
			ctx.SetBodyString("invalid request, please follow this format: api/v1/tickers?expire_at=60")
			return
		} else {
			_ = glg.Infof("Expire at arg is: %d", expireAt)
			expirationPriceValidity = expireAt
		}
	}

	var out = make(map[string]*mm2_tools_generics.TickerInfosAnswer)
	for _, cur := range config.GCFGRegistry {
		resp := mm2_tools_generics.GetTickerInfos(cur.Coin, expirationPriceValidity)
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

func TickerAllInfosV2(ctx *fasthttp.RequestCtx) {
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

	args := ctx.QueryArgs()
	expirationPriceValidity := 0
	if args.Len() > 0 {
		glg.Info("len is above 0, there is an arg to this get request")
		expireAt, err := args.GetUint("expire_at")
		if err != nil {
			_ = glg.Errorf("error with request args: %v", err)
			ctx.SetStatusCode(http.StatusBadRequest)
			ctx.SetBodyString("invalid request, please follow this format: api/v1/tickers?expire_at=60")
			return
		} else {
			_ = glg.Infof("Expire at arg is: %d", expireAt)
			expirationPriceValidity = expireAt
		}
	}

	var out = make(map[string]*mm2_tools_generics.TickerInfosAnswer)
	var memoization = make(map[string]bool)
	for _, cur := range config.GCFGRegistry {
		if memoization[helpers.RetrieveMainTicker(cur.Coin)] == false && cur.IsTestNet == false {
			resp := mm2_tools_generics.GetTickerInfos(cur.Coin, expirationPriceValidity)
			resp.Ticker = helpers.RetrieveMainTicker(resp.Ticker)
			resp.Sparkline7D = nil
			if helpers.AsFloat(resp.LastPrice) > 0 {
				out[resp.Ticker] = resp
				memoization[resp.Ticker] = true
			}
		}
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
