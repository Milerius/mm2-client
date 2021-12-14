package external_services

import (
	"encoding/json"
	"fmt"
	"github.com/kpango/glg"
	"io/ioutil"
	"mm2_client/config"
	"mm2_client/helpers"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type NomicsAnswer struct {
	Id                 string    `json:"id"`
	Currency           string    `json:"currency"`
	Symbol             string    `json:"symbol"`
	Name               string    `json:"name"`
	LogoUrl            string    `json:"logo_url"`
	Status             string    `json:"status"`
	Price              string    `json:"price"`
	PriceDate          time.Time `json:"price_date"`
	PriceTimestamp     time.Time `json:"price_timestamp"`
	CirculatingSupply  string    `json:"circulating_supply"`
	MaxSupply          string    `json:"max_supply"`
	MarketCap          string    `json:"market_cap"`
	MarketCapDominance string    `json:"market_cap_dominance"`
	NumExchanges       string    `json:"num_exchanges"`
	NumPairs           string    `json:"num_pairs"`
	NumPairsUnmapped   string    `json:"num_pairs_unmapped"`
	FirstCandle        time.Time `json:"first_candle"`
	FirstTrade         time.Time `json:"first_trade"`
	FirstOrderBook     time.Time `json:"first_order_book"`
	Rank               string    `json:"rank"`
	RankDelta          string    `json:"rank_delta"`
	High               string    `json:"high"`
	HighTimestamp      time.Time `json:"high_timestamp"`
	D                  struct {
		Volume             string `json:"volume"`
		PriceChange        string `json:"price_change"`
		PriceChangePct     string `json:"price_change_pct"`
		VolumeChange       string `json:"volume_change"`
		VolumeChangePct    string `json:"volume_change_pct"`
		MarketCapChange    string `json:"market_cap_change"`
		MarketCapChangePct string `json:"market_cap_change_pct"`
	} `json:"1d"`
}

var gNomicsEndpoint = "https://api.nomics.com/v1/currencies/ticker?key=" + os.Getenv("NOMICS_API_KEY")
var gNomicsOptions = "&interval=1d&status=active"

var NomicsPriceRegistry sync.Map

func NewNomicsRequest() string {
	url := gNomicsEndpoint + "&ids="
	for _, cur := range config.GCFGRegistry {
		if !cur.IsTestNet && cur.NomicsId != nil {
			url += *cur.NomicsId + ","
		}
	}
	url = strings.TrimSuffix(url, ",")
	url += gNomicsOptions
	return url
}

func processNomics() *[]NomicsAnswer {
	url := NewNomicsRequest()
	_ = glg.Infof("Processing nomics request: %s", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		var answer = &[]NomicsAnswer{}
		decodeErr := json.NewDecoder(resp.Body).Decode(answer)
		if decodeErr != nil {
			fmt.Printf("Err: %v\n", decodeErr)
			return nil
		}
		return answer
	} else {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		glg.Errorf("Http status not OK: %s", bodyBytes)
		return nil
	}
}

func StartNomicsService() {
	for {
		if resp := processNomics(); resp != nil {
			glg.Info("Nomics request successfully processed")
			for _, cur := range *resp {
				NomicsPriceRegistry.Store(cur.Id, cur)
			}
		} else {
			glg.Error("Something went wrong when processing nomics request")
		}
		time.Sleep(time.Second * 30)
	}
}

func NomicsRetrieveUSDValIfSupported(coin string) (string, string, string) {
	dateStr := helpers.GetDateFromTimestampStandard(time.Now().UnixNano())
	valStr := "0"
	if cfg, cfgOk := config.GCFGRegistry[coin]; cfgOk && cfg.NomicsId != nil {
		val, ok := NomicsPriceRegistry.Load(*cfg.NomicsId)
		if ok {
			resp := val.(NomicsAnswer)
			valStr = fmt.Sprintf("%s", resp.Price)
			dateStr = helpers.GetDateFromTime(resp.PriceTimestamp)
		}
		return valStr, dateStr, "nomics"
	}
	return valStr, dateStr, "unknown"
}

func NomicsRetrieveCEXRatesFromPair(base string, rel string) (string, bool, string, string) {
	basePrice, baseDate, _ := NomicsRetrieveUSDValIfSupported(base)
	relPrice, relDate, _ := NomicsRetrieveUSDValIfSupported(rel)
	price := helpers.BigFloatDivide(basePrice, relPrice, 8)
	date := helpers.GetDateFromTimestampStandard(time.Now().UnixNano())
	if helpers.RFC3339ToTimestamp(baseDate) <= helpers.RFC3339ToTimestamp(relDate) {
		date = baseDate
	} else {
		date = relDate
	}
	return price, true, date, "nomics"
}

func NomicsGetChange24h(coin string) (string, string, string) {
	changePercent24h := "0"
	dateStr := helpers.GetDateFromTimestampStandard(time.Now().UnixNano())
	if cfg, cfgOk := config.GCFGRegistry[coin]; cfgOk && cfg.NomicsId != nil {
		val, ok := NomicsPriceRegistry.Load(*cfg.NomicsId)
		if ok {
			resp := val.(NomicsAnswer)
			changePercent24h = fmt.Sprintf("%s", resp.D.PriceChangePct)
			changePercent24h = helpers.BigFloatMultiply(changePercent24h, "100", 3)
			dateStr = helpers.GetDateFromTime(resp.PriceTimestamp)
		}
		return changePercent24h, dateStr, "nomics"
	}
	return changePercent24h, dateStr, "unknown"
}
