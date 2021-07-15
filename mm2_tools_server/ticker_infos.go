package mm2_tools_server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kpango/glg"
	"github.com/valyala/fasthttp"
	"mm2_client/constants"
	"mm2_client/services"
	"net/http"
)

type TickerInfosRequest struct {
	Ticker string `json:"ticker"`
}

type TickerInfosAnswer struct {
	Ticker      string `json:"ticker"`
	LastPrice   string `json:"last_price"`
	LastUpdated string `json:"last_updated"`
	Provider    string `json:"provider"`
}

func GetTickerInfos(ticker string) *TickerInfosAnswer {
	val, date, provider := services.RetrieveUSDValIfSupported(ticker)
	return &TickerInfosAnswer{Ticker: ticker, LastPrice: val, LastUpdated: date, Provider: provider}
}

func (req *TickerInfosAnswer) ToJson() string {
	b, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}

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
	out := &TickerInfosRequest{}
	r := bytes.NewReader(ctx.PostBody())
	err := json.NewDecoder(r).Decode(out)
	if err != nil {
		_ = glg.Errorf("%v", err)
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.SetBodyString(err.Error())
		return
	}
	resp := GetTickerInfos(out.Ticker)
	ctx.SetStatusCode(200)
	ctx.SetBodyString(resp.ToJson())
	ctx.SetContentType("application/json")
}
