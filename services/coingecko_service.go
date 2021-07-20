package services

import (
	"encoding/json"
	"fmt"
	"github.com/kpango/glg"
	"io/ioutil"
	"mm2_client/config"
	"mm2_client/helpers"
	"net/http"
	"strings"
	"sync"
	"time"
)

type CoingeckoAnswer struct {
	Id                           string    `json:"id"`
	Symbol                       string    `json:"symbol"`
	Name                         string    `json:"name"`
	Image                        string    `json:"image"`
	CurrentPrice                 float64   `json:"current_price"`
	MarketCap                    float64   `json:"market_cap"`
	MarketCapRank                *int      `json:"market_cap_rank"`
	FullyDilutedValuation        *int64    `json:"fully_diluted_valuation"`
	TotalVolume                  float64   `json:"total_volume"`
	High24H                      *float64  `json:"high_24h"`
	Low24H                       *float64  `json:"low_24h"`
	PriceChange24H               *float64  `json:"price_change_24h"`
	PriceChangePercentage24H     *float64  `json:"price_change_percentage_24h"`
	MarketCapChange24H           *float64  `json:"market_cap_change_24h"`
	MarketCapChangePercentage24H *float64  `json:"market_cap_change_percentage_24h"`
	CirculatingSupply            float64   `json:"circulating_supply"`
	TotalSupply                  *float64  `json:"total_supply"`
	MaxSupply                    *float64  `json:"max_supply"`
	Ath                          float64   `json:"ath"`
	AthChangePercentage          float64   `json:"ath_change_percentage"`
	AthDate                      time.Time `json:"ath_date"`
	Atl                          float64   `json:"atl"`
	AtlChangePercentage          float64   `json:"atl_change_percentage"`
	AtlDate                      time.Time `json:"atl_date"`
	Roi                          *struct {
		Times      float64 `json:"times"`
		Currency   string  `json:"currency"`
		Percentage float64 `json:"percentage"`
	} `json:"roi"`
	LastUpdated   string `json:"last_updated"`
	SparklineIn7D struct {
		Price []float64 `json:"price"`
	} `json:"sparkline_in_7d"`
	PriceChangePercentage24HInCurrency *float64 `json:"price_change_percentage_24h_in_currency"`
}

const gCoingeckoEndpoint = "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids="

var CoingeckoPriceRegistry sync.Map

func NewCoingeckoRequest() string {
	url := gCoingeckoEndpoint
	for _, cur := range config.GCFGRegistry {
		if cur.CoingeckoID != "test-coin" {
			url += cur.CoingeckoID + ","
		}
	}
	url = strings.TrimSuffix(url, ",")
	url += "&order=market_cap_desc&price_change_percentage=24h"
	return url
}

func processCoingecko() *[]CoingeckoAnswer {
	url := NewCoingeckoRequest()
	_ = glg.Infof("Processing coingecko request: %s", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		var answer = &[]CoingeckoAnswer{}
		decodeErr := json.NewDecoder(resp.Body).Decode(answer)
		if decodeErr != nil {
			fmt.Printf("Err: %v\n", err)
			return nil
		}
		return answer
	} else {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		glg.Errorf("Http status not OK: %s", bodyBytes)
		return nil
	}
}

func StartCoingeckoService() {
	for {
		if resp := processCoingecko(); resp != nil {
			glg.Info("Coingecko request successfully processed")
			for _, cur := range *resp {
				CoingeckoPriceRegistry.Store(strings.ToUpper(cur.Symbol), cur)
			}
		} else {
			glg.Error("Something went wrong when processing coingecko request")
		}
		time.Sleep(time.Second * 30)
	}
}

func CoingeckoRetrieveUSDValIfSupported(coin string) (string, string, string) {
	coin = helpers.RetrieveMainTicker(coin)
	val, ok := CoingeckoPriceRegistry.Load(coin)
	valStr := "0"
	dateStr := helpers.GetDateFromTimestampStandard(time.Now().UnixNano())
	if ok {
		resp := val.(CoingeckoAnswer)
		valStr = fmt.Sprintf("%f", resp.CurrentPrice)
		dateStr = resp.LastUpdated
	}
	return valStr, dateStr, "coingecko"
}

func CoingeckoRetrieveCEXRatesFromPair(base string, rel string) (string, bool, string, string) {
	basePrice, baseDate, _ := CoingeckoRetrieveUSDValIfSupported(base)
	relPrice, relDate, _ := CoingeckoRetrieveUSDValIfSupported(rel)
	price := helpers.BigFloatDivide(basePrice, relPrice, 8)
	calculated := true
	date := helpers.GetDateFromTimestampStandard(time.Now().UnixNano())
	if helpers.RFC3339ToTimestamp(baseDate) <= helpers.RFC3339ToTimestamp(relDate) {
		date = baseDate
	} else {
		date = relDate
	}
	return price, calculated, date, "coingecko"
}

func CoingeckoGetTotalVolume(coin string) (string, string, string) {
	coin = helpers.RetrieveMainTicker(coin)
	val, ok := CoingeckoPriceRegistry.Load(coin)
	totalVolumeStr := "0"
	dateStr := helpers.GetDateFromTimestampStandard(time.Now().UnixNano())
	if ok {
		resp := val.(CoingeckoAnswer)
		totalVolumeStr = fmt.Sprintf("%f", resp.TotalVolume)
		dateStr = resp.LastUpdated
	}
	return totalVolumeStr, dateStr, "coingecko"
}
