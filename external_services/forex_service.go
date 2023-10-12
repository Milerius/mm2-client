package external_services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mm2_client/config"
	"mm2_client/constants"
	"mm2_client/helpers"
	"net/http"
	"sync"
	"time"

	"github.com/kpango/glg"
)

type ForexAnswer struct {
	Success   bool               `json:"success"`
	Timestamp int64              `json:"timestamp"`
	Base      string             `json:"base"`
	Date      string             `json:"date"`
	Rates     map[string]float64 `json:"rates"`
}

var gForexEndpoint = "https://rates.komodo.earth/api/v1/usd_rates"

var ForexPriceRegistry sync.Map

func processForex() *ForexAnswer {
	url := gForexEndpoint
	_ = glg.Infof("Processing forex request: %s", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		var answer = &ForexAnswer{}
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

func StartForexService() {
	for {
		if resp := processForex(); resp != nil {
			glg.Info("Forex request successfully processed")
			ForexPriceRegistry.Store("Forex", resp)
		} else {
			glg.Error("Something went wrong when processing forex request")
		}
		time.Sleep(constants.GPricesLoopTime)
	}
}

func ForexRetrieveUSDValIfSupported(coin string) (string, string, string) {
	dateStr := helpers.GetDateFromTimestampStandard(time.Now().UnixNano())
	valStr := "0"
	if cfg, cfgOk := config.GCFGRegistry[coin]; cfgOk && cfg.ForexId != nil {
		val, ok := ForexPriceRegistry.Load("Forex")
		if ok {
			resp := val.(*ForexAnswer)
			if price, priceOk := resp.Rates[*cfg.ForexId]; priceOk {
				valStr = fmt.Sprintf("%f", 1.0/price)
				dateStr = helpers.GetDateFromTimestampStandardSeconds(resp.Timestamp)
			} else {
				return valStr, dateStr, "unknown"
			}
		}
		return valStr, dateStr, "forex"
	}
	return valStr, dateStr, "unknown"
}
