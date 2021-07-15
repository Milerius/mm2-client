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

type CexRatesRequest struct {
	Base string `json:"base"`
	Rel  string `json:"rel"`
}

type CexRateAnswer struct {
	Base        string `json:"base"`
	Rel         string `json:"rel"`
	LastPrice   string `json:"last_price"`
	LastUpdated string `json:"last_updated"`
	Provider    string `json:"provider"`
	Calculated  bool   `json:"calculated"`
}

func (req *CexRateAnswer) ToJson() string {
	b, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}

func CexRates(ctx *fasthttp.RequestCtx) {
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

	out := &CexRatesRequest{}
	r := bytes.NewReader(ctx.PostBody())
	err := json.NewDecoder(r).Decode(out)
	if err != nil {
		_ = glg.Errorf("%v", err)
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.SetBodyString(err.Error())
		return
	}

	val, calculated, date, provider := services.RetrieveCEXRatesFromPair(out.Base, out.Rel)
	resp := &CexRateAnswer{Base: out.Base, Rel: out.Rel, LastPrice: val, LastUpdated: date, Provider: provider, Calculated: calculated}
	ctx.SetStatusCode(200)
	ctx.SetBodyString(resp.ToJson())
	ctx.SetContentType("application/json")
}
