package market_making

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mm2_client/constants"
	"mm2_client/helpers"
	"mm2_client/http"
	"mm2_client/services"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type SimplePairMarketMakerConf struct {
	Base           string  `json:"base"`
	Rel            string  `json:"rel"`
	Max            bool    `json:"max,omitempty"`
	BalancePercent string  `json:"balance_percent,omitempty"`
	MinVolume      *string `json:"min_volume,omitempty"`
	Spread         string  `json:"spread"`
	BaseConfs      int     `json:"base_confs"`
	BaseNota       bool    `json:"base_nota"`
	RelConfs       int     `json:"rel_confs"`
	RelNota        bool    `json:"rel_nota"`
	Enable         bool    `json:"enable"`
}

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

var gSimpleMarketMakerRegistry = make(map[string]SimplePairMarketMakerConf)
var gQuitMarketMakerBot chan struct{}

func init() {
	_ = os.MkdirAll(filepath.Join(constants.GetAppDataPath(), "logs"), os.ModePerm)
	file, err := os.OpenFile(constants.GetAppDataPath()+"/logs/simple.market.maker.logs", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func NewMarketMakerConfFromFile(targetPath string) bool {
	file, _ := ioutil.ReadFile(targetPath)
	err := json.Unmarshal(file, &gSimpleMarketMakerRegistry)
	if err != nil {
		fmt.Println("Couldn't parse cfg file - aborting")
		return false
	}
	return true
}

func updateOrderFromCfg(cfg SimplePairMarketMakerConf, makerOrder http.MakerOrderContent) {
	cexPrice, calculated, date := services.RetrieveCEXRatesFromPair(cfg.Base, cfg.Rel)
	if elapsed := helpers.DateToTimeElapsed(date); elapsed < (5 * time.Minute).Seconds() {
		price := helpers.BigFloatMultiply(cexPrice, cfg.Spread, 8)
		resp := http.UpdateMakerOrder(makerOrder.Uuid, &price, nil, &cfg.Max, &makerOrder.MinBaseVol, &cfg.BaseConfs, &cfg.BaseNota, &cfg.RelConfs, &cfg.RelNota)
		if resp != nil {
			InfoLogger.Printf("Successfully updated order %s - cex_price: [%s] - new_price: [%s] - calculated: [%t] elapsed_since_price: %f seconds\n resp: [%v]", makerOrder.Uuid, cexPrice, price, calculated, elapsed, resp)
		}
	} else {
		cancelResp := http.CancelOrder(makerOrder.Uuid)
		if cancelResp != nil {
			WarningLogger.Printf("Cancelled order %s - reason: price elapsed > 5min\n", makerOrder.Uuid)
		}
	}
}

func createOrderFromConf(cfg SimplePairMarketMakerConf) {
	cexPrice, calculated, date := services.RetrieveCEXRatesFromPair(cfg.Base, cfg.Rel)
	if elapsed := helpers.DateToTimeElapsed(date); elapsed < (5 * time.Minute).Seconds() {
		if helpers.AsFloat(cexPrice) > 0 {
			price := helpers.BigFloatMultiply(cexPrice, cfg.Spread, 8)
			var max *bool = nil
			var volume *string = nil
			var minVolume *string = nil
			var maxBalance = http.MyBalance(cfg.Base).Balance
			if cfg.Max {
				max = helpers.BoolAddr(true)
			} else {
				vol := helpers.BigFloatMultiply(maxBalance, cfg.BalancePercent, 8)
				volume = &vol
			}
			if cfg.MinVolume != nil {
				if cfg.Max {
					minVol := helpers.BigFloatMultiply(maxBalance, *cfg.MinVolume, 8)
					minVolume = &minVol
				} else if !cfg.Max && volume != nil {
					minVol := helpers.BigFloatMultiply(*volume, *cfg.MinVolume, 8)
					minVolume = &minVol
				}
			}
			resp := http.SetPrice(cfg.Base, cfg.Rel, price, volume, max, true, minVolume,
				&cfg.BaseConfs, &cfg.BaseNota, &cfg.RelConfs, &cfg.RelNota)
			if resp != nil {
				InfoLogger.Printf("Successfully placed the order: %s, calculated: %t cex_price: [%s] - our price: [%s] - elapsed since last price update: %f seconds\n", resp.Result.Uuid, calculated, cexPrice, price, elapsed)
			} else {
				ErrorLogger.Printf("Couldn't place the order for %s/%s\n", cfg.Base, cfg.Rel)
			}
		} else {
			WarningLogger.Printf("Price is 0 for %s/%s - skipping order creation\n", cfg.Base, cfg.Rel)
		}
	} else {
		WarningLogger.Printf("Last Price update is too far %f seconds, need to be under 5 minute to create an order\n", helpers.DateToTimeElapsed(date))
	}
}

func marketMakerProcess() {
	InfoLogger.Println("process market maker")

	hitRegistry := make(map[string]bool)

	InfoLogger.Println("Retrieving my orders for the update")
	//! Need to iterate through existing orders and update them
	wgUpdate := sync.WaitGroup{}
	updateFunctor := func(cfg SimplePairMarketMakerConf, makerOrder http.MakerOrderContent, combination string) {
		defer wgUpdate.Done()
		updateOrderFromCfg(cfg, makerOrder)
	}
	if resp := http.MyOrders(); resp != nil {
		for _, curMakerOrder := range resp.Result.MakerOrders {
			combination := curMakerOrder.Base + "/" + curMakerOrder.Rel
			if val, ok := gSimpleMarketMakerRegistry[combination]; ok && val.Enable {
				wgUpdate.Add(1)
				go updateFunctor(val, curMakerOrder, combination)
				hitRegistry[combination] = true
			}
		}
	}
	wgUpdate.Wait()
	InfoLogger.Println("Orders updated")

	InfoLogger.Println("Iterating over order that have not been updated and that need a creation")
	//! If i didn't visit one of the supported coin i need to create an order
	wgCreate := sync.WaitGroup{}
	creatorFunctor := func(cfg SimplePairMarketMakerConf, combination string) {
		defer wgCreate.Done()
		InfoLogger.Printf("Need to create order for pair: [%s]\n", combination)
		createOrderFromConf(cfg)
	}

	for curCombination, cfg := range gSimpleMarketMakerRegistry {
		if _, ok := hitRegistry[curCombination]; !ok && cfg.Enable {
			wgCreate.Add(1)
			go creatorFunctor(cfg, curCombination)
		}
	}
	wgCreate.Wait()
	InfoLogger.Println("Orders created")
}

func startSimpleMarketMakerBotService() {
	for {
		select {
		case <-gQuitMarketMakerBot:
			InfoLogger.Println("Simple Market Maker Bot service successfully stopped")
			constants.GSimpleMarketMakerBotRunning = false
			close(gQuitMarketMakerBot)
			constants.GSimpleMarketMakerBotStopping = false
			cancelPendingOrders()
			return
		default:
			marketMakerProcess()
			time.Sleep(30 * time.Second)
		}
	}
}

func StopSimpleMarketMakerBotService() {
	if !constants.GSimpleMarketMakerBotStopping && constants.GSimpleMarketMakerBotRunning {
		constants.GSimpleMarketMakerBotStopping = true
		//! Also need to cancel all existing orders (Could use by UUID meanwhile this time)
		InfoLogger.Println("Stopping Simple Market Maker Bot Service - may take up to 30 seconds")
		cancelPendingOrders()
		go func() {
			gQuitMarketMakerBot <- struct{}{}
		}()
	} else {
		fmt.Println("Simple market maker is still shutting down or not running")
	}
}

func cancelPendingOrders() {
	var outBatch []interface{}
	if resp := http.MyOrders(); resp != nil {
		for _, cur := range resp.Result.MakerOrders {
			if req := http.NewCancelOrderRequest(cur.Uuid); req != nil {
				outBatch = append(outBatch, req)
			}
		}
	}

	resp := http.BatchRequest(outBatch)
	if len(resp) > 0 {
		var outResp []http.CancelOrderAnswer
		err := json.Unmarshal([]byte(resp), &outResp)
		if err != nil {
			WarningLogger.Println("Couldn't cancel all pending orders")
		} else {
			InfoLogger.Println("Successfully cancelled all pending orders")
		}
	}
}

func StartSimpleMarketMakerBot() {
	if constants.GMM2Running {
		if constants.GSimpleMarketMakerBotRunning {
			fmt.Println("Simple Market Maker bot is already running (or being stopped) - skipping")
		} else {
			if resp := NewMarketMakerConfFromFile(constants.GSimpleMarketMakerConf); resp {
				cancelPendingOrders()
				InfoLogger.Printf("Starting simple market maker bot with %d coin(s)\n", len(gSimpleMarketMakerRegistry))
				constants.GSimpleMarketMakerBotRunning = true
				gQuitMarketMakerBot = make(chan struct{})
				go startSimpleMarketMakerBotService()
			} else {
				fmt.Println("Couldn't start simple market maker without valid conf")
			}
		}
	} else {
		fmt.Println("MM2 need to be started before starting the simple market maker bot")
	}
}
