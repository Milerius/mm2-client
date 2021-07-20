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

type Volume24hRequest struct {
	Coin string `json:"coin"`
}

type Volume24hAnswer struct {
	Coin        string `json:"string"`
	Volume24h   string `json:"volume24h"`
	LastUpdated string `json:"last_updated"`
	Provider    string `json:"provider"`
}

func (req *Volume24hAnswer) ToJson() string {
	b, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}

func Volume24h(ctx *fasthttp.RequestCtx) {
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
	out := &Volume24hRequest{}
	r := bytes.NewReader(ctx.PostBody())
	err := json.NewDecoder(r).Decode(out)
	if err != nil {
		_ = glg.Errorf("%v", err)
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.SetBodyString(err.Error())
		return
	}

	val, date, provider := services.RetrieveVolume24h(out.Coin)
	resp := &Volume24hAnswer{Coin: out.Coin, Volume24h: val, LastUpdated: date, Provider: provider}
	ctx.SetStatusCode(200)
	ctx.SetBodyString(resp.ToJson())
	ctx.SetContentType("application/json")
}
