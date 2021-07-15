package mm2_tools_server

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kpango/glg"
	"github.com/valyala/fasthttp"
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

func getTickerInfos(ticker string) (*TickerInfosAnswer, error) {
	val, date, provider := services.RetrieveUSDValIfSupported(ticker)
	if val != "0" {
		return &TickerInfosAnswer{Ticker: ticker, LastPrice: val, LastUpdated: date, Provider: provider}, nil
	}
	return nil, errors.New("couldn't fetch price of " + ticker)
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
	out := &TickerInfosRequest{}
	r := bytes.NewReader(ctx.PostBody())
	err := json.NewDecoder(r).Decode(out)
	if err != nil {
		_ = glg.Errorf("%v", err)
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.SetBodyString(err.Error())
		return
	}
	resp, errResp := getTickerInfos(out.Ticker)
	if resp != nil {
		ctx.SetStatusCode(200)
		ctx.SetBodyString(resp.ToJson())
		ctx.SetContentType("application/json")
	} else {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		_ = glg.Errorf("Error: %v", errResp)
		ctx.SetBodyString(errResp.Error())
	}
}
