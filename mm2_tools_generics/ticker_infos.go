package mm2_tools_generics

import (
	"encoding/json"
	"fmt"
	"mm2_client/services"
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

func (req *TickerInfosAnswer) ToJson() string {
	b, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}

func (req *TickerInfosAnswer) ToWeb() map[string]interface{} {
	out := make(map[string]interface{})
	if req != nil {
		_ = json.Unmarshal([]byte(req.ToJson()), &out)
		return out
	}
	return nil
}

func GetTickerInfos(ticker string) *TickerInfosAnswer {
	val, date, provider := services.RetrieveUSDValIfSupported(ticker)
	return &TickerInfosAnswer{Ticker: ticker, LastPrice: val, LastUpdated: date, Provider: provider}
}
