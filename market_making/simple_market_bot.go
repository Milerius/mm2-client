package market_making

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kpango/glg"
	"io/ioutil"
	"mm2_client/constants"
	"mm2_client/helpers"
	"mm2_client/mm2_tools_generics"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"mm2_client/services"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type SimplePairMarketMakerConf struct {
	Base                 string   `json:"base"`
	Rel                  string   `json:"rel"`
	Max                  bool     `json:"max,omitempty"`
	BalancePercent       string   `json:"balance_percent,omitempty"`
	MinVolume            *string  `json:"min_volume,omitempty"`
	Spread               string   `json:"spread"`
	BaseConfs            int      `json:"base_confs"`
	BaseNota             bool     `json:"base_nota"`
	RelConfs             int      `json:"rel_confs"`
	RelNota              bool     `json:"rel_nota"`
	Enable               bool     `json:"enable"`
	PriceElapsedValidity *float64 `json:"price_elapsed_validity,omitempty"`
}

var gSimpleMarketMakerRegistry = make(map[string]SimplePairMarketMakerConf)
var gQuitMarketMakerBot chan struct{}

func init() {
	_ = os.MkdirAll(filepath.Join(constants.GetAppDataPath(), "logs"), os.ModePerm)
}

func NewMarketMakerConfFromFile(targetPath string) bool {
	file, _ := ioutil.ReadFile(targetPath)
	err := json.Unmarshal(file, &gSimpleMarketMakerRegistry)
	if err != nil {
		fmt.Println("Couldn't parse cfg file - aborting")
		return false
	}
	for key, cur := range gSimpleMarketMakerRegistry {
		if cur.PriceElapsedValidity == nil {
			validity := 300.0
			cur.PriceElapsedValidity = &validity
			gSimpleMarketMakerRegistry[key] = cur
			glg.Infof("Overriding price elapsed validity settings for %s/%s with %.1f - because it's not present in the json configuration", cur.Base, cur.Rel, 300.0)
		}
	}
	return true
}

func updateOrderFromCfg(cfg SimplePairMarketMakerConf, makerOrder mm2_data_structure.MakerOrderContent) {
	cexPrice, calculated, date, provider := services.RetrieveCEXRatesFromPair(cfg.Base, cfg.Rel)
	if elapsed := helpers.DateToTimeElapsed(date); elapsed < *cfg.PriceElapsedValidity {
		price := helpers.BigFloatMultiply(cexPrice, cfg.Spread, 8)
		resp, err := mm2_tools_generics.UpdateMakerOrder(makerOrder.Uuid, &price, nil, &cfg.Max, &makerOrder.MinBaseVol, &cfg.BaseConfs, &cfg.BaseNota, &cfg.RelConfs, &cfg.RelNota)
		if resp != nil {
			glg.Infof("Successfully updated %s/%s order %s - cex_price: [%s] - new_price: [%s] - calculated: [%t] elapsed_since_price: %f seconds - provider: %s",
				makerOrder.Base, makerOrder.Rel, makerOrder.Uuid,
				cexPrice, price, calculated, elapsed, provider)
			glg.Get().EnableJSON().Info(resp)
			glg.Get().DisableJSON()
		} else {
			glg.Warnf("err update_maker_order: %v", err)
		}
	} else {
		cancelResp, cancelErr := mm2_tools_generics.CancelOrder(makerOrder.Uuid)
		if cancelResp != nil {
			glg.Warnf("Cancelled %s/%s order %s - reason: price elapsed > %.1f seconds", makerOrder.Base, makerOrder.Rel, makerOrder.Uuid, *cfg.PriceElapsedValidity)
		} else {
			glg.Warnf("Error when cancelling order: %v", cancelErr)
		}
	}
}

func createOrderFromConf(cfg SimplePairMarketMakerConf) {
	cexPrice, calculated, date, provider := services.RetrieveCEXRatesFromPair(cfg.Base, cfg.Rel)
	if elapsed := helpers.DateToTimeElapsed(date); elapsed < *cfg.PriceElapsedValidity {
		if helpers.AsFloat(cexPrice) > 0 {
			price := helpers.BigFloatMultiply(cexPrice, cfg.Spread, 8)
			var max *bool = nil
			var volume *string = nil
			var minVolume *string = nil
			respBalance, err := mm2_tools_generics.MyBalance(cfg.Base)
			if respBalance != nil {
				var maxBalance = respBalance.Balance
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
				resp, setPriceErr := mm2_tools_generics.SetPrice(cfg.Base, cfg.Rel, price, volume, max, true, minVolume,
					&cfg.BaseConfs, &cfg.BaseNota, &cfg.RelConfs, &cfg.RelNota)
				if resp != nil {
					glg.Infof("Successfully placed the %s/%s order: %s, calculated: %t cex_price: [%s] - our price: [%s] - elapsed since last price update: %f seconds - provider: %s", cfg.Base, cfg.Rel, resp.Result.Uuid, calculated, cexPrice, price, elapsed, provider)
					glg.Get().EnableJSON().Info(resp)
					glg.Get().DisableJSON()
				} else {
					glg.Errorf("Couldn't place the order for %s/%s: %v", cfg.Base, cfg.Rel, setPriceErr)
				}
			} else {
				glg.Errorf("Cannot retrieve balance of %s - skipping: %v", cfg.Base, err)
			}
		} else {
			glg.Warnf("Price is 0 for %s/%s - skipping order creation", cfg.Base, cfg.Rel)
		}
	} else {
		glg.Warnf("Last Price update for %s/%s is too far %f seconds, need to be under %.1f seconds to create an order", cfg.Base, cfg.Rel, helpers.DateToTimeElapsed(date), *cfg.PriceElapsedValidity)
	}
}

func marketMakerProcess() {
	glg.Info("process market maker")

	hitRegistry := make(map[string]bool)

	glg.Info("Retrieving my orders for the update")
	//! Need to iterate through existing orders and update them
	wgUpdate := sync.WaitGroup{}
	updateFunctor := func(cfg SimplePairMarketMakerConf, makerOrder mm2_data_structure.MakerOrderContent, combination string) {
		defer wgUpdate.Done()
		updateOrderFromCfg(cfg, makerOrder)
	}
	if resp, err := mm2_tools_generics.MyOrders(); resp != nil {
		for _, curMakerOrder := range resp.Result.MakerOrders {
			combination := curMakerOrder.Base + "/" + curMakerOrder.Rel
			if val, ok := gSimpleMarketMakerRegistry[combination]; ok && val.Enable {
				wgUpdate.Add(1)
				go updateFunctor(val, curMakerOrder, combination)
				hitRegistry[combination] = true
			}
		}
	} else {
		glg.Errorf("err on my_orders: %v", err)
	}
	wgUpdate.Wait()
	glg.Info("Orders updated")

	glg.Info("Iterating over order that have not been updated and that need a creation")
	//! If i didn't visit one of the supported coin i need to create an order
	wgCreate := sync.WaitGroup{}
	creatorFunctor := func(cfg SimplePairMarketMakerConf, combination string) {
		defer wgCreate.Done()
		glg.Infof("Need to create order for pair: [%s]", combination)
		createOrderFromConf(cfg)
	}

	for curCombination, cfg := range gSimpleMarketMakerRegistry {
		if _, ok := hitRegistry[curCombination]; !ok && cfg.Enable {
			wgCreate.Add(1)
			go creatorFunctor(cfg, curCombination)
		}
	}
	wgCreate.Wait()
	glg.Info("Orders created")
}

func startSimpleMarketMakerBotService() {
	for {
		select {
		case <-gQuitMarketMakerBot:
			glg.Info("Simple Market Maker Bot service successfully stopped")
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

func StopSimpleMarketMakerBotService() error {
	if !constants.GSimpleMarketMakerBotStopping && constants.GSimpleMarketMakerBotRunning {
		constants.GSimpleMarketMakerBotStopping = true
		//! Also need to cancel all existing orders (Could use by UUID meanwhile this time)
		glg.Info("Stopping Simple Market Maker Bot Service - may take up to 30 seconds")
		cancelPendingOrders()
		go func() {
			gQuitMarketMakerBot <- struct{}{}
		}()
		return nil
	} else {
		fmt.Println("Simple market maker is still shutting down or not running")
		return errors.New("simple market maker is still shutting down or not running")
	}
}

func cancelPendingOrders() {
	if resp, err := mm2_tools_generics.CancelAllOrders("all", []string{}); resp != nil {
		glg.Info("Successfully cancelled all pending orders")
	} else {
		glg.Warnf("Couldn't cancel all pending orders: %v", err)
	}
}

func StartSimpleMarketMakerBot(path string, appName string) error {
	if constants.GMM2Running {
		if constants.GSimpleMarketMakerBotRunning {
			fmt.Println("Simple Market Maker bot is already running (or being stopped) - skipping")
			return errors.New("simple Market Maker bot is already running (or being stopped) - skipping")
		} else {
			if resp := NewMarketMakerConfFromFile(path); resp {
				cancelPendingOrders()
				glg.Infof("Starting simple market maker bot with %d coin(s)", len(gSimpleMarketMakerRegistry))
				constants.GSimpleMarketMakerBotRunning = true
				gQuitMarketMakerBot = make(chan struct{})
				go startSimpleMarketMakerBotService()
				return nil
			} else {
				fmt.Println("Couldn't start simple market maker without valid conf")
				return errors.New("couldn't start simple market maker without valid conf")
			}
		}
	} else {
		fmt.Println("MM2 need to be started before starting the simple market maker bot")
		return errors.New("mm2 need to be started before starting the simple market maker bot")
	}
}
