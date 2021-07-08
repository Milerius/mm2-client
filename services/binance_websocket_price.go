package services

import (
	"fmt"
	"github.com/adshao/go-binance/v2"
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

func contains(coin string, stablecoin string) (string, bool) {
	val, ok := BinancePriceRegistry.Load(helpers.RetrieveMainTicker(coin) + stablecoin)
	valStr := "0"
	if ok {
		valStr = val.(string)
	}
	return valStr, ok
}

func RetrieveUSDValIfSupported(coin string) string {

	if valUSD, okUSD := contains(coin, "USD"); okUSD {
		return valUSD
	} else if valUSDT, okUSDT := contains(coin, "USDT"); okUSDT {
		return valUSDT
	} else if valBUSD, okBUSD := contains(coin, "BUSD"); okBUSD {
		return valBUSD
	} else if valUSDC, okUSDC := contains(coin, "USDC"); okUSDC {
		return valUSDC
	} else if valDAI, okDAI := contains(coin, "DAI"); okDAI {
		return valDAI
	} else if helpers.IsAStableCoin(coin) {
		return "1"
	}
	return "0"
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
		_, okErc := config.GCFGRegistry[v.BaseAsset+"-ERC20"]
		_, okBep := config.GCFGRegistry[v.BaseAsset+"-BEP20"]
		_, okQrc := config.GCFGRegistry[v.BaseAsset+"-QRC20"]
		if (ok || okErc || okBep || okQrc) && helpers.IsAStableCoin(v.QuoteAsset) && v.Status == "TRADING" {
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
		go startWebsocketForSymbol(cur)
	}
}

func startWebsocketForSymbol(cur string) {
	wsMarketHandler := func(event *binance.WsMarketStatEvent) {
		BinancePriceRegistry.Store(event.Symbol, event.LastPrice)
	}
	errHandler := func(err error) {
		if strings.Contains(err.Error(), "websocket: close 1006 (abnormal closure)") {
			go startWebsocketForSymbol(cur)
		}
	}
	//fmt.Printf("Starting websocket service for symbol: %s\n", cur)
	_, _, err := binance.WsMarketStatServe(cur, wsMarketHandler, errHandler)
	if err != nil {
		//fmt.Printf("err for %s: %v\n", cur, err)
		time.Sleep(1 * time.Second)
		if strings.Contains(err.Error(), "EOF") {
			go startWebsocketForSymbol(cur)
		}
	}
}

func GetBinanceSupportedPairs() []string {
	var out []string
	for key, _ := range BinanceSupportedTickers {
		for alt, _ := range BinanceSupportedTickers {
			if key != alt {
				out = append(out, key+"-"+alt)
			}
		}
	}

	var data [][]string

	for _, curPair := range out {
		splitted := strings.Split(curPair, "-")
		base := splitted[0]
		rel := splitted[1]
		basePrice := RetrieveUSDValIfSupported(base)
		relPrice := RetrieveUSDValIfSupported(rel)
		combined := base + rel
		price := helpers.BigFloatDivide(basePrice, relPrice, 8)
		calculated := true
		if val, ok := BinancePriceRegistry.Load(combined); ok {
			price = val.(string)
			calculated = false
		}
		var cur []string
		if !calculated {
			cur = []string{base, basePrice, rel, relPrice, price, emoji.Sprintf(":x:")}
		} else {
			cur = []string{base, basePrice, rel, relPrice, price, emoji.Sprintf(":white_check_mark:")}
		}
		data = append(data, cur)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"Base", "BasePrice", "Rel", "RelPrice", "PriceOfThePair", "Calculated"})
	table.SetFooter([]string{"Base", "BasePrice", "Rel", "RelPrice", "PriceOfThePair", "Calculated"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data) // Add Bulk Data
	table.Render()

	return out
}
