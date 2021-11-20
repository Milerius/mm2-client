package cli

import (
	"encoding/json"
	"fmt"
	"mm2_client/config"
	"mm2_client/external_services"
	"mm2_client/helpers"
	"mm2_client/mm2_tools_generics"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"sync"
	"sync/atomic"
)

func queryLiquidityAll() {
	var wg = sync.WaitGroup{}
	var totalLiquidityAtomic = atomic.Value{}
	totalLiquidityAtomic.Store("0")
	for ticker, cfg := range config.GCFGRegistry {
		if cfg.IsTestNet == false && cfg.WalletOnly == false {
			value, _, _ := external_services.RetrieveUSDValIfSupported(cfg.Coin, 0)
			if helpers.AsFloat(value) > 0 {
				wg.Add(1)
				copyTicker := ticker
				go func() {
					defer wg.Done()
					totalLiquidity := queryLiquidity(copyTicker, false)
					total := helpers.BigFloatAdd(totalLiquidityAtomic.Load().(string), totalLiquidity, 8)
					totalLiquidityAtomic.Store(total)
				}()
			}
		}
	}
	wg.Wait()
	fmt.Printf("Total liquidity is: %s $\n", totalLiquidityAtomic.Load())
}

func queryLiquidity(coin string, verbose bool) string {
	totalLiquidity := "0"
	coinValue := "0"
	var outBatch []interface{}
	if val, ok := config.GCFGRegistry[coin]; ok {
		value, _, _ := external_services.RetrieveUSDValIfSupported(val.Coin, 0)
		coinValue = value
		if helpers.AsFloat(value) > 0 {
			for ticker, cfg := range config.GCFGRegistry {
				if ticker != val.Coin && cfg.WalletOnly == false {
					curValue, _, _ := external_services.RetrieveUSDValIfSupported(ticker, 0)
					if helpers.AsFloat(curValue) > 0 {
						req := mm2_data_structure.NewOrderbookRequest(coin, ticker)
						outBatch = append(outBatch, req)
					}
				}
			}
		}
	}
	resp := mm2_tools_generics.BatchRequest(outBatch)
	if len(resp) > 0 {
		var outResp []mm2_data_structure.OrderbookAnswer
		err := json.Unmarshal([]byte(resp), &outResp)
		if len(outResp) > 0 && err == nil {
			for _, curResp := range outResp {
				curLiquidity := helpers.BigFloatMultiply(curResp.TotalAsksBaseVol, coinValue, 8)
				if helpers.AsFloat(curResp.TotalAsksRelVol) > 0 && verbose {
					fmt.Printf("Liquidity for %s/%s asks is %s $\n", coin, curResp.Rel, curLiquidity)
				}
				totalLiquidity = helpers.BigFloatAdd(totalLiquidity, curLiquidity, 8)
			}
		}
	}
	fmt.Printf("Total liquidity for %s is: %s $\n", coin, totalLiquidity)
	return totalLiquidity
}
