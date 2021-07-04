package services

import (
	"fmt"
	"github.com/adshao/go-binance/v2"
	"mm2_client/config"
	"mm2_client/helpers"
	"sync"
)

const BinanceWebsocketEndpoint = "wss://stream.binance.com:9443"

var BinancePriceRegistry sync.Map

//! <symbol>@ticker

func StartBinanceWebsocketService() {
	var out []string
	keys := make(map[string]bool)
	for _, v := range config.GCFGRegistry {
		if len(v.BinanceAvailablePairs) > 0 {
			for _, cur := range v.BinanceAvailablePairs {
				symbol := helpers.RetrieveMainTicker(v.Coin) + cur
				if _, value := keys[symbol]; !value {
					keys[symbol] = true
					out = append(out, symbol)
				}
			}
		}
	}

	for _, cur := range out {
		wsMarketHandler := func(event *binance.WsMarketStatEvent) {
			//fmt.Printf("lastprice %s: %s\n", event.Symbol, event.LastPrice)
			BinancePriceRegistry.Store(event.Symbol, event.LastPrice)
		}
		errHandler := func(err error) {
			fmt.Println(err)
		}
		_, _, err := binance.WsMarketStatServe(cur, wsMarketHandler, errHandler)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
