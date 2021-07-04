package services

import (
	"fmt"
	"github.com/adshao/go-binance/v2"
	"mm2_client/config"
	"sync"
)

const BinanceWebsocketEndpoint = "wss://stream.binance.com:9443"

var BinancePriceRegistry sync.Map

//! <symbol>@ticker

func StartBinanceWebsocketService() {
	//fmt.Println("StartBinanceWebsocketService")
	var out []string
	for _, v := range config.GCFGRegistry {
		if len(v.BinanceAvailablePairs) > 0 {
			for _, cur := range v.BinanceAvailablePairs {
				//fmt.Printf("%s/%s supported processing\n", v.Coin, cur)
				out = append(out, v.Coin+cur)
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
		/*go func() {
			time.Sleep(5 * time.Second)
			stopC <- struct{}{}
		}()
		<-doneC*/
	}
}
