package main

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/kpango/glg"
	"mm2_client/config"
	"mm2_client/config/wasm_storage"
	"mm2_client/constants"
	"mm2_client/mm2_tools_generics"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"mm2_client/mm2_tools_generics/mm2_wasm_request"
	"mm2_client/services"
	"net/url"
	"strconv"
	"syscall/js"
)

func startPriceService() js.Func {
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if constants.GPricesServicesRunning {
			_ = glg.Warn("Price service already running - skipping")
			return nil
		}
		if len(config.GCFGRegistry) == 0 {
			_ = glg.Warn("Desktop cfg need to be loaded first before running the price service")
			return nil
		}
		services.LaunchServices()
		return nil
	})
	return jsfunc
}

func loadCoinsCfgFromUrl() js.Func {
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		_ = glg.Info("load_coins_cfg_from_url called")
		if len(args) != 1 {
			usage := "invalid nb args - usage: load_coins_cfg_from_url(\"my_url_to_coins_cfg\")"
			_ = glg.Error(usage)
			result := map[string]interface{}{
				"error": usage,
			}
			return result
		}
		inputUrl := args[0].String()
		_, err := url.ParseRequestURI(inputUrl)
		if err != nil {
			errStr := fmt.Sprintf("invalid url: %v\n", err)
			_ = glg.Errorf("%s", errStr)
			result := map[string]interface{}{
				"error": errStr,
			}
			return result
		}
		_ = glg.Infof("url is: %s", inputUrl)
		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			resolve := args[0]
			reject := args[1]
			go func() {
				err = config.ParseMM2CFGFromUrl(inputUrl)
				if err != nil {
					errStr := fmt.Sprintf("error when parsing cfg: %v\n", err)
					rejectErr := errors.New(errStr)
					errorConstructor := js.Global().Get("Error")
					errorObject := errorConstructor.New(rejectErr.Error())
					reject.Invoke(errorObject)
					_ = glg.Errorf("%s", errStr)
				} else {
					_ = glg.Infof("cfg successfully parsed: %d", len(config.GCFGRegistry))
					resolve.Invoke(map[string]interface{}{
						"message": "cfg successfully parsed",
						"error":   nil,
					})
				}
			}()
			return nil
		})
		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
	return jsfunc
}

func loadDesktopCfgFromUrl() js.Func {
	storageHandler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		resolve := args[0]
		reject := args[1]
		resp := wasm_storage.Retrieve("wasm_coins_cfg")
		go func() {
			if config.ParseDesktopRegistryFromString(resp) {
				resolve.Invoke(map[string]interface{}{
					"message": "desktop cfg successfully parsed",
					"len":     strconv.Itoa(len(config.GCFGRegistry)),
					"error":   nil,
				})
			} else {
				rejectErr := errors.New("error when parsing cfg")
				errorConstructor := js.Global().Get("Error")
				errorObject := errorConstructor.New(rejectErr.Error())
				reject.Invoke(errorObject)
			}
		}()
		return nil
	})
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if resp := wasm_storage.Retrieve("wasm_coins_cfg"); resp != "" {
			return js.Global().Get("Promise").New(storageHandler)
		} else {
			_ = glg.Info("load_desktop_cfg_from_url called")
			if len(args) != 1 {
				usage := "invalid nb args - usage: load_desktop_cfg_from_url(\"my_url_to_desktop_cfg\")"
				_ = glg.Error(usage)
				result := map[string]interface{}{
					"error": usage,
				}
				return result
			}
			inputUrl := args[0].String()
			_, err := url.ParseRequestURI(inputUrl)
			if err != nil {
				errStr := fmt.Sprintf("invalid url: %v\n", err)
				_ = glg.Errorf("%s", errStr)
				result := map[string]interface{}{
					"error": errStr,
				}
				return result
			}
			_ = glg.Infof("url is: %s", inputUrl)
			handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
				resolve := args[0]
				reject := args[1]
				go func() {
					err = config.ParseDesktopRegistryFromUrl(inputUrl)
					if err != nil {
						errStr := fmt.Sprintf("error when parsing cfg: %v\n", err)
						rejectErr := errors.New(errStr)
						errorConstructor := js.Global().Get("Error")
						errorObject := errorConstructor.New(rejectErr.Error())
						reject.Invoke(errorObject)
						_ = glg.Errorf("%s", errStr)
					} else {
						_ = glg.Infof("cfg successfully parsed: %d", len(config.GCFGRegistry))
						resolve.Invoke(map[string]interface{}{
							"message": "cfg successfully parsed",
							"error":   nil,
						})
						if resp == "" {
							config.UpdateWasm()
						}
					}
				}()
				return nil
			})
			promiseConstructor := js.Global().Get("Promise")
			return promiseConstructor.New(handler)
		}
	})
	return jsfunc
}

func getTickerInfos() js.Func {
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if !constants.GPricesServicesRunning {
			_ = glg.Warn("Price service need to run to call this function")
			return nil
		}
		if !constants.GDesktopCfgLoaded {
			_ = glg.Warn("Desktop cfg need to be loaded to continue")
			return nil
		}
		if len(args) != 1 {
			usage := "invalid nb args"
			_ = glg.Error(usage)
			result := map[string]interface{}{
				"error": usage,
			}
			return result
		}
		resp := mm2_tools_generics.GetTickerInfos(args[0].String())
		return resp.ToWeb()
	})
	return jsfunc
}

func getAllTickerInfos() js.Func {
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if !constants.GPricesServicesRunning {
			_ = glg.Warn("Price service need to run to call this function")
			return nil
		}
		if !constants.GDesktopCfgLoaded {
			_ = glg.Warn("Desktop cfg need to be loaded to continue")
			return nil
		}
		var out = make(map[string]interface{})
		for _, cur := range config.GCFGRegistry {
			resp := mm2_tools_generics.GetTickerInfos(cur.Coin).ToWeb()
			out[cur.Coin] = resp
		}
		return out
	})
	return jsfunc
}

func StartMM2() js.Func {
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 2 {
			usage := "invalid nb args - usage: start_mm2(userpass, passphrase)"
			_ = glg.Error(usage)
			result := map[string]interface{}{
				"error": usage,
			}
			return result
		}
		var out []string
		if len(args) > 2 {
			raw := valueToBytes(args[2])
			extraArgs := []string{}
			buf := bytes.NewBuffer(raw)
			if buf != nil {
				gob.NewDecoder(buf).Decode(&extraArgs)
				for _, cur := range extraArgs {
					out = append(out, cur)
				}
			} else {
				glg.Error("err decode")
			}
		}
		return config.NewMM2ConfigWasm(args[0].String(), args[1].String(), out)
	})
	return jsfunc
}

func bootstrap() js.Func {
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		userpass := "wasmtest"
		passphrase := "hardcoded_password"
		if len(args) >= 2 {
			userpass = args[0].String()
			passphrase = args[1].String()
		}
		mm2_data_structure.GRuntimeUserpass = userpass
		go func() {
			val, errVal := mm2_wasm_request.Await(js.Global().Call("init_wasm"))
			if val != nil {
				glg.Infof("done from the promise")
				parseVal, _ := mm2_wasm_request.Await(js.Global().Call("load_desktop_cfg_from_url", "http://localhost:8080/static/assets/wasm.coins.json"))
				if parseVal != nil {
					parseMM2Val, _ := mm2_wasm_request.Await(js.Global().Call("load_coins_cfg_from_url", "http://localhost:8080/static/assets/coins.json"))
					if parseMM2Val != nil {
						var startVal []js.Value
						if len(args) <= 2 {
							startVal, _ = mm2_wasm_request.Await(js.Global().Call("run_mm2", js.Global().Call("start_mm2", userpass, passphrase)))
						} else {
							var anotherSlice []string
							for _, cur := range args[2:] {
								anotherSlice = append(anotherSlice, cur.String())
							}
							buf := &bytes.Buffer{}
							gob.NewEncoder(buf).Encode(anotherSlice)
							extraArgs := bytesToValue(buf.Bytes())
							startVal, _ = mm2_wasm_request.Await(js.Global().Call("run_mm2", js.Global().Call("start_mm2", userpass, passphrase, extraArgs)))
						}
						constants.GMM2Running = true
						if startVal != nil {
							js.Global().Call("enable_active_coins")
							glg.Info("Bootstrap done !")
						}
					}
				}
			} else {
				glg.Errorf("bad from the promise: %v", errVal)
			}
		}()
		return "done"
	})
	return jsfunc
}

func main() {
	mm2_data_structure.GRuntimeUserpass = "wasmtest"
	glg.Get().SetMode(glg.STD)
	_ = glg.Info("Hello from webassembly - Slyris tools running")

	//! Internal
	js.Global().Set("load_desktop_cfg_from_url", loadDesktopCfgFromUrl())
	js.Global().Set("load_coins_cfg_from_url", loadCoinsCfgFromUrl())
	js.Global().Set("start_mm2", StartMM2())
	js.Global().Set("enable_active_coins", enableActiveCoins())

	//! Price API
	js.Global().Set("get_ticker_infos", getTickerInfos())
	js.Global().Set("get_all_ticker_infos", getAllTickerInfos())
	js.Global().Set("start_price_service", startPriceService())

	//! CLI API
	js.Global().Set("bootstrap", bootstrap())
	js.Global().Set("my_balance", myBalance())
	js.Global().Set("balance_all", myBalanceAll())
	js.Global().Set("enable", enable())
	js.Global().Set("disable_coin", disableCoin())
	js.Global().Set("disable_enabled_coins", disableEnabledCoins())
	js.Global().Set("disable_zero_balance", disableZeroBalance())
	js.Global().Set("get_enabled_coins", getEnabledCoins())
	js.Global().Set("my_orders", myOrders())
	js.Global().Set("start_simple_market_maker_bot", startSimpleMarketMakerBot())
	js.Global().Set("stop_simple_market_maker_bot", stopSimpleMarketMakerBot())
	js.Global().Set("kmd_rewards_infos", kmdRewardsInfos())
	js.Global().Set("withdraw", withdraw())
	js.Global().Set("send", send())
	js.Global().Set("broadcast", broadcast())
	js.Global().Set("my_tx_history", myTxHistory())
	js.Global().Set("orderbook", orderbook())

	<-make(chan bool)
}
