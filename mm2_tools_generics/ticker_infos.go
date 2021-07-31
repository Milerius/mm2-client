package mm2_tools_generics

import (
	"encoding/json"
	"fmt"
	"mm2_client/external_services"
)

type TickerInfosRequest struct {
	Ticker string `json:"ticker"`
}

type TickerInfosAnswer struct {
	Ticker            string     `json:"ticker"`
	LastPrice         string     `json:"last_price"`
	LastUpdated       string     `json:"last_updated"`
	Volume24h         string     `json:"volume24h"`
	PriceProvider     string     `json:"price_provider"`
	VolumeProvider    string     `json:"volume_provider"`
	Sparkline7D       *[]float64 `json:"sparkline_7d"`
	SparklineProvider string     `json:"sparkline_provider"`
	Change24h         string     `json:"change_24h"`
	Change24hProvider string     `json:"change_24h_provider"`
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
	val, date, provider := external_services.RetrieveUSDValIfSupported(ticker)
	volume, _, volumeProvider := external_services.RetrieveVolume24h(ticker)
	sparkline7d, _, sparklineProvider := external_services.RetrieveSparkline7D(ticker)
	change24h, _, change24hProvider := external_services.RetrievePercentChange24h(ticker)
	return &TickerInfosAnswer{Ticker: ticker, LastPrice: val, LastUpdated: date,
		PriceProvider: provider,
		Volume24h:     volume, VolumeProvider: volumeProvider,
		Sparkline7D: sparkline7d, SparklineProvider: sparklineProvider,
		Change24h: change24h, Change24hProvider: change24hProvider}
}
