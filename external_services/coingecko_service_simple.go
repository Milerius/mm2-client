package external_services

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

type CoingeckoSparkLineData struct {
    Price *[]float64 `json:"price,omitempty"`
}

type CoingeckoAnswer struct {
    Id                           string    `json:"id"`
    Symbol                       string    `json:"symbol"`
    Name                         string    `json:"name"`
    Image                        string    `json:"image"`
    CurrentPrice                 float64   `json:"current_price"`
    MarketCap                    float64   `json:"market_cap"`
    MarketCapRank                *int      `json:"market_cap_rank"`
    FullyDilutedValuation        *float64  `json:"fully_diluted_valuation"`
    TotalVolume                  float64   `json:"total_volume"`
    High24H                      *float64  `json:"high_24h"`
    Low24H                       *float64  `json:"low_24h"`
    PriceChange24H               *float64  `json:"price_change_24h"`
    PriceChangePercentage24H     *float64  `json:"price_change_percentage_24h"`
    MarketCapChange24H           *float64  `json:"market_cap_change_24h"`
    MarketCapChangePercentage24H *float64  `json:"market_cap_change_percentage_24h"`
    CirculatingSupply            float64   `json:"circulating_supply"`
    TotalSupply                  *float64  `json:"total_supply"`
    MaxSupply                    *float64  `json:"max_supply,omitempty"`
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
    LastUpdated                        string                  `json:"last_updated"`
    SparklineIn7D                      *CoingeckoSparkLineData `json:"sparkline_in_7d,omitempty"`
    PriceChangePercentage24HInCurrency *float64                `json:"price_change_percentage_24h_in_currency"`
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
    url += "&order=market_cap_desc&price_change_percentage=24h&sparkline=true&per_page=250"
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

func StartCoingeckoService() {
    for {
        if resp := processCoingecko(); resp != nil {
            glg.Info("Coingecko request successfully processed")
            for _, cur := range *resp {
                CoingeckoPriceRegistry.Store(cur.Id, cur)
            }
        } else {
            glg.Error("Something went wrong when processing coingecko request")
        }
        time.Sleep(time.Second * 30)
    }
}

func CoingeckoRetrieveUSDValIfSupported(coin string) (string, string, string) {
    dateStr := helpers.GetDateFromTimestampStandard(time.Now().UnixNano())
    valStr := "0"
    if cfg, cfgOk := config.GCFGRegistry[coin]; cfgOk {
        val, ok := CoingeckoPriceRegistry.Load(cfg.CoingeckoID)
        if ok {
            resp := val.(CoingeckoAnswer)
            valStr = fmt.Sprintf("%f", resp.CurrentPrice)
            dateStr = resp.LastUpdated
        }
        return valStr, dateStr, "coingecko"
    }
    return valStr, dateStr, "unknown"
}

func CoingeckoRetrieveCEXRatesFromPair(base string, rel string) (string, bool, string, string) {
    basePrice, baseDate, _ := CoingeckoRetrieveUSDValIfSupported(base)
    relPrice, relDate, _ := CoingeckoRetrieveUSDValIfSupported(rel)
    price := helpers.BigFloatDivide(basePrice, relPrice, 8)
    date := helpers.GetDateFromTimestampStandard(time.Now().UnixNano())
    if helpers.RFC3339ToTimestamp(baseDate) <= helpers.RFC3339ToTimestamp(relDate) {
        date = baseDate
    } else {
        date = relDate
    }
    return price, true, date, "coingecko"
}

func CoingeckoGetTotalVolume(coin string) (string, string, string) {
    totalVolumeStr := "0"
    dateStr := helpers.GetDateFromTimestampStandard(time.Now().UnixNano())
    if cfg, cfgOk := config.GCFGRegistry[coin]; cfgOk {
        val, ok := CoingeckoPriceRegistry.Load(cfg.CoingeckoID)
        if ok {
            resp := val.(CoingeckoAnswer)
            totalVolumeStr = fmt.Sprintf("%f", resp.TotalVolume)
            dateStr = resp.LastUpdated
        }
        return totalVolumeStr, dateStr, "coingecko"
    }
    return totalVolumeStr, dateStr, "unknown"
}

func CoingeckoGetSparkline7D(coin string) (*[]float64, string, string) {
    var out *[]float64
    dateStr := helpers.GetDateFromTimestampStandard(time.Now().UnixNano())
    if cfg, cfgOk := config.GCFGRegistry[coin]; cfgOk {
        val, ok := CoingeckoPriceRegistry.Load(cfg.CoingeckoID)
        if ok {
            resp := val.(CoingeckoAnswer)
            if resp.SparklineIn7D != nil && resp.SparklineIn7D.Price != nil {
                out = resp.SparklineIn7D.Price
            }
            dateStr = resp.LastUpdated
        }
        return out, dateStr, "coingecko"
    }
    return out, dateStr, "unknown"
}

func CoingeckoGetChange24h(coin string) (string, string, string) {
    changePercent24h := "0"
    dateStr := helpers.GetDateFromTimestampStandard(time.Now().UnixNano())
    if cfg, cfgOk := config.GCFGRegistry[coin]; cfgOk {
        val, ok := CoingeckoPriceRegistry.Load(cfg.CoingeckoID)
        if ok {
            resp := val.(CoingeckoAnswer)
            if resp.PriceChangePercentage24H != nil {
                changePercent24h = fmt.Sprintf("%f", *resp.PriceChangePercentage24HInCurrency)
            }
            dateStr = resp.LastUpdated
        }
        return changePercent24h, dateStr, "coingecko"
    }
    return changePercent24h, dateStr, "unknown"
}
