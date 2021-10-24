package external_services

import (
	"fmt"
	"github.com/adshao/go-binance/v2"
	"github.com/kpango/glg"
	"github.com/kyokomi/emoji/v2"
	"github.com/olekukonko/tablewriter"
	"mm2_client/config"
	"mm2_client/helpers"
	"os"
	"strings"
	"sync"
	"time"
)

var BinancePriceRegistry sync.Map
var BinanceSupportedTickers = make(map[string]bool)

//! <symbol>@ticker

func contains(coin string, stablecoin string) (string, string, string, bool) {
	val, ok := BinancePriceRegistry.Load(helpers.RetrieveMainTicker(coin) + stablecoin)
	valStr := "0"
	dateStr := helpers.GetDateFromTimestampStandard(time.Now().UnixNano())
	priceChangePercent := "0"
	if ok {
		valStr = val.([]string)[0]
		dateStr = val.([]string)[1]
		priceChangePercent = val.([]string)[2]
	}
	return valStr, dateStr, priceChangePercent, ok
}

func BinanceRetrieveUSDValIfSupported(coin string) (string, string, string, string) {
	coin = helpers.RetrieveMainTicker(coin)
	if valUSD, dateUSD, priceChangePercentUSD, okUSD := contains(coin, "USD"); okUSD {
		return valUSD, dateUSD, priceChangePercentUSD, "binance"
	} else if valUSDT, dateUSDT, priceChangePercentUSDT, okUSDT := contains(coin, "USDT"); okUSDT {
		return valUSDT, dateUSDT, priceChangePercentUSDT, "binance"
	} else if valBUSD, dateBUSD, priceChangePercentBUSD, okBUSD := contains(coin, "BUSD"); okBUSD {
		return valBUSD, dateBUSD, priceChangePercentBUSD, "binance"
	} else if valUSDC, dateUSDC, priceChangePercentUSDC, okUSDC := contains(coin, "USDC"); okUSDC {
		return valUSDC, dateUSDC, priceChangePercentUSDC, "binance"
	} else if valDAI, dateDAI, priceChangePercentDAI, okDAI := contains(coin, "DAI"); okDAI {
		return valDAI, dateDAI, priceChangePercentDAI, "binance"
	} else if helpers.IsAStableCoin(coin) {
		return "1", helpers.GetDateFromTimestampStandard(time.Now().UnixNano()), "0", "binance"
	}
	return "0", helpers.GetDateFromTimestampStandard(time.Now().UnixNano()), "0", "binance"
}

func StartBinanceWebsocketService() {
	var out []string
	keys := make(map[string]bool)
	infos, err := GetBinanceExchangeInfos()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range infos.Symbols {
		_, ok := config.GCFGRegistry[v.BaseAsset]
		_, okArb := config.GCFGRegistry[v.BaseAsset+"-ARB20"]
		_, okAvx := config.GCFGRegistry[v.BaseAsset+"-AVX20"]
		_, okBep := config.GCFGRegistry[v.BaseAsset+"-BEP20"]
		_, okErc := config.GCFGRegistry[v.BaseAsset+"-ERC20"]
		_, okFtm := config.GCFGRegistry[v.BaseAsset+"-FTM20"]
		_, okHco := config.GCFGRegistry[v.BaseAsset+"-HCO20"]
		_, okKrc := config.GCFGRegistry[v.BaseAsset+"-KRC20"]
		_, okMvr := config.GCFGRegistry[v.BaseAsset+"-MVR20"]
		_, okPlg := config.GCFGRegistry[v.BaseAsset+"-PLG20"]
		_, okQrc := config.GCFGRegistry[v.BaseAsset+"-QRC20"]
		if (ok || okErc || okBep || okQrc || okPlg || okAvx || okArb || okMvr || okHco || okKrc || okFtm) && helpers.IsAStableCoin(v.QuoteAsset) && v.Status == "TRADING" {
			if _, value := keys[v.Symbol]; !value {
				keys[v.Symbol] = true
				out = append(out, v.Symbol)
				//fmt.Println(v.BaseAsset)
				BinanceSupportedTickers[v.BaseAsset] = true
			}
		}
	}

	//fmt.Printf("%v\n", out)
	for _, cur := range out {
		//fmt.Println(cur)
		go startWebsocketForSymbol(cur)
	}
}

func startWebsocketForSymbol(cur string) {
	wsMarketHandler := func(event *binance.WsMarketStatEvent) {
		//fmt.Println(event.Time)
		BinancePriceRegistry.Store(event.Symbol, []string{event.LastPrice, helpers.GetDateFromTimestampStandard(event.Time * int64(time.Millisecond)), event.PriceChangePercent})
	}
	errHandler := func(err error) {
		glg.Errorf("err: %s", err)
		if strings.Contains(err.Error(), "websocket: close 1006 (abnormal closure)") {
			go startWebsocketForSymbol(cur)
		}
	}
	//fmt.Printf("Starting websocket service for symbol: %s\n", cur)
	doneC, _, err := binance.WsMarketStatServe(cur, wsMarketHandler, errHandler)
	if err != nil {
		glg.Errorf("err: %s", err)
		time.Sleep(1 * time.Second)
		if strings.Contains(err.Error(), "EOF") {
			go startWebsocketForSymbol(cur)
		}
	}
	<-doneC
	glg.Infof("Disconnected, reconnecting now")
	go startWebsocketForSymbol(cur)
}

func retrievePossibilities(cur string) []string {
	var base []string
	var rel []string
	var out []string

	res := strings.Split(cur, "/")
	curBase := res[0]
	curRel := res[1]
	functorAppendIfExist := func(out *[]string, ticker string) {
		if val, ok := config.GCFGRegistry[ticker]; ok {
			*out = append(*out, val.Coin)
		}
	}

	functorAppendIfExist(&base, curBase)
	functorAppendIfExist(&base, curBase+"-ARB20")
	functorAppendIfExist(&base, curBase+"-AVX20")
	functorAppendIfExist(&base, curBase+"-BEP20")
	functorAppendIfExist(&base, curBase+"-ERC20")
	functorAppendIfExist(&base, curBase+"-FTM20")
	functorAppendIfExist(&base, curBase+"-HCO20")
	functorAppendIfExist(&base, curBase+"-KRC20")
	functorAppendIfExist(&base, curBase+"-MVR20")
	functorAppendIfExist(&base, curBase+"-PLG20")
	functorAppendIfExist(&base, curBase+"-QRC20")
	functorAppendIfExist(&rel, curRel)
	functorAppendIfExist(&rel, curRel+"-ARB20")
	functorAppendIfExist(&rel, curRel+"-AVX20")
	functorAppendIfExist(&rel, curRel+"-BEP20")
	functorAppendIfExist(&rel, curRel+"-ERC20")
	functorAppendIfExist(&rel, curRel+"-FTM20")
	functorAppendIfExist(&rel, curRel+"-HCO20")
	functorAppendIfExist(&rel, curRel+"-KRC20")
	functorAppendIfExist(&rel, curRel+"-MVR20")
	functorAppendIfExist(&rel, curRel+"-PLG20")
	functorAppendIfExist(&rel, curRel+"-QRC20")

	for _, b := range base {
		for _, r := range rel {
			out = append(out, b+"/"+r)
		}
	}

	return out
}

func GetBinanceSupportedPairsInternals() []string {
	var out []string
	for base, _ := range BinanceSupportedTickers {
		for rel, _ := range BinanceSupportedTickers {
			if base != rel {
				combination := base + "/" + rel
				possibilities := retrievePossibilities(combination)
				out = append(out, possibilities...)
			}
		}
	}
	return out
}

func GetBinanceSupportedPairs(ticker string) []string {
	var out = GetBinanceSupportedPairsInternals()

	var data [][]string

	for _, curPair := range out {
		splitted := strings.Split(curPair, "/")
		base := splitted[0]
		if len(ticker) > 0 && !strings.Contains(base, ticker) {
			continue
		}
		rel := splitted[1]
		basePrice, dateBase, _, _ := BinanceRetrieveUSDValIfSupported(base)
		relPrice, dateRel, _, _ := BinanceRetrieveUSDValIfSupported(rel)
		combined := helpers.RetrieveMainTicker(base) + helpers.RetrieveMainTicker(rel)
		price := helpers.BigFloatDivide(basePrice, relPrice, 8)
		calculated := true
		date := helpers.GetDateFromTimestampStandard(time.Now().UnixNano())
		if val, ok := BinancePriceRegistry.Load(combined); ok {
			price = val.([]string)[0]
			date = val.([]string)[1]
			calculated = false
		}
		var cur []string
		if !calculated {
			cur = []string{base, basePrice, rel, relPrice, price, emoji.Sprintf(":x:"), date, date}
		} else {
			cur = []string{base, basePrice, rel, relPrice, price, emoji.Sprintf(":white_check_mark:"), dateBase, dateRel}
		}
		data = append(data, cur)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"Base", "BasePrice", "Rel", "RelPrice", "PriceOfThePair", "Calculated", "BaseLastUpdate", "RelLastUpdate"})
	table.SetFooter([]string{"Base", "BasePrice", "Rel", "RelPrice", "PriceOfThePair", "Calculated", "BaseLastUpdate", "RelLastUpdate"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data) // Add Bulk Data
	table.Render()

	return out
}

func BinanceRetrieveCEXRatesFromPair(base string, rel string) (string, bool, string, string) {
	basePrice, baseDate, _, _ := BinanceRetrieveUSDValIfSupported(base)
	relPrice, relDate, _, _ := BinanceRetrieveUSDValIfSupported(rel)
	combined := helpers.RetrieveMainTicker(base) + helpers.RetrieveMainTicker(rel)
	price := helpers.BigFloatDivide(basePrice, relPrice, 8)
	calculated := true
	date := helpers.GetDateFromTimestampStandard(time.Now().UnixNano())
	if helpers.RFC3339ToTimestamp(baseDate) <= helpers.RFC3339ToTimestamp(relDate) {
		//fmt.Printf("Pick [%s date: %s] instead of [%s date: %s]\n", base, baseDate, rel, relDate)
		date = baseDate
	} else {
		//fmt.Printf("Pick [%s date: %s] instead of [%s date: %s]\n", rel, relDate, base, baseDate)
		date = relDate
	}
	if val, ok := BinancePriceRegistry.Load(combined); ok {
		price = val.([]string)[0]
		date = val.([]string)[1]
		calculated = false
	}
	return price, calculated, date, "binance"
}

/*func BinanceGetChange24h(coin string) (string, string, string) {

}*/
