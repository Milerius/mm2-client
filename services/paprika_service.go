package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mm2_client/helpers"
	"net/http"
	"sync"
	"time"
)

type CoinpaprikaAnswer struct {
	Id                string    `json:"id"`
	Name              string    `json:"name"`
	Symbol            string    `json:"symbol"`
	Rank              int       `json:"rank"`
	CirculatingSupply int64     `json:"circulating_supply"`
	TotalSupply       int64     `json:"total_supply"`
	MaxSupply         int64     `json:"max_supply"`
	BetaValue         float64   `json:"beta_value"`
	FirstDataAt       time.Time `json:"first_data_at"`
	LastUpdated       string    `json:"last_updated"`
	Quotes            struct {
		USD struct {
			Price               float64   `json:"price"`
			Volume24H           float64   `json:"volume_24h"`
			Volume24HChange24H  float64   `json:"volume_24h_change_24h"`
			MarketCap           int64     `json:"market_cap"`
			MarketCapChange24H  float64   `json:"market_cap_change_24h"`
			PercentChange15M    float64   `json:"percent_change_15m"`
			PercentChange30M    float64   `json:"percent_change_30m"`
			PercentChange1H     float64   `json:"percent_change_1h"`
			PercentChange6H     float64   `json:"percent_change_6h"`
			PercentChange12H    float64   `json:"percent_change_12h"`
			PercentChange24H    float64   `json:"percent_change_24h"`
			PercentChange7D     float64   `json:"percent_change_7d"`
			PercentChange30D    float64   `json:"percent_change_30d"`
			PercentChange1Y     float64   `json:"percent_change_1y"`
			AthPrice            float64   `json:"ath_price"`
			AthDate             time.Time `json:"ath_date"`
			PercentFromPriceAth float64   `json:"percent_from_price_ath"`
		} `json:"USD"`
	} `json:"quotes"`
}

const gPaprikaEndpoint = "https://api.coinpaprika.com/v1/tickers"

var CoinpaprikaRegistry sync.Map

func processCoinpaprika() *[]CoinpaprikaAnswer {
	url := gPaprikaEndpoint
	infoLogger.Printf("Processing coinpaprika request: %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		var answer = &[]CoinpaprikaAnswer{}
		decodeErr := json.NewDecoder(resp.Body).Decode(answer)
		if decodeErr != nil {
			fmt.Printf("Err: %v\n", err)
			return nil
		}
		return answer
	} else {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		warningLogger.Printf("Http status not OK: %s\n", bodyBytes)
		return nil
	}
}

func StartCoinpaprikaService() {
	for {
		if resp := processCoinpaprika(); resp != nil {
			infoLogger.Println("Coinpaprika request successfully processed")
			for _, cur := range *resp {
				CoinpaprikaRegistry.Store(cur.Symbol, cur)
			}
		} else {
			warningLogger.Println("Something went wrong when processing coinpaprika request")
		}
		time.Sleep(time.Second * 30)
	}
}

func CoinpaprikaRetrieveUSDValIfSupported(coin string) (string, string, string) {
	coin = helpers.RetrieveMainTicker(coin)
	val, ok := CoinpaprikaRegistry.Load(coin)
	valStr := "0"
	dateStr := helpers.GetDateFromTimestampStandard(time.Now().UnixNano())
	if ok {
		resp := val.(CoinpaprikaAnswer)
		valStr = fmt.Sprintf("%f", resp.Quotes.USD.Price)
		dateStr = resp.LastUpdated
	}
	return valStr, dateStr, "coinpaprika"
}

func CoinpaprikaRetrieveCEXRatesFromPair(base string, rel string) (string, bool, string, string) {
	basePrice, baseDate, _ := CoinpaprikaRetrieveUSDValIfSupported(base)
	relPrice, relDate, _ := CoinpaprikaRetrieveUSDValIfSupported(rel)
	price := helpers.BigFloatDivide(basePrice, relPrice, 8)
	date := helpers.GetDateFromTimestampStandard(time.Now().UnixNano())
	if helpers.RFC3339ToTimestamp(baseDate) <= helpers.RFC3339ToTimestamp(relDate) {
		date = baseDate
	} else {
		date = relDate
	}
	return price, true, date, "coinpaprika"
}
