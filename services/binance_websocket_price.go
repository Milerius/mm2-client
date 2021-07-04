package services

import (
	"fmt"
	"github.com/adshao/go-binance/v2"
	"mm2_client/config"
	"mm2_client/helpers"
	"strings"
	"sync"
	"time"
)

const BinanceWebsocketEndpoint = "wss://stream.binance.com:9443"

var BinancePriceRegistry sync.Map

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
	fmt.Printf("Starting websocket service for symbol: %s\n", cur)
	_, _, err := binance.WsMarketStatServe(cur, wsMarketHandler, errHandler)
	if err != nil {
		fmt.Printf("err for %s: %v\n", cur, err)
		time.Sleep(1 * time.Second)
		if strings.Contains(err.Error(), "EOF") {
			go startWebsocketForSymbol(cur)
		}
	}
}
