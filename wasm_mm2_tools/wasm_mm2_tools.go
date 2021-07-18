package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kpango/glg"
	"mm2_client/config"
	"mm2_client/constants"
	"mm2_client/http"
	"mm2_client/mm2_tools_generics"
	"mm2_client/services"
	"net/url"
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
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
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
				}
			}()
			return nil
		})
		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
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
		return config.NewMM2ConfigWasm()
	})
	return jsfunc
}

func enableActiveCoins() js.Func {
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		coins := config.RetrieveActiveCoins()
		var outBatch []interface{}
		for _, v := range coins {
			if val, ok := config.GCFGRegistry[v]; ok {
				switch val.Type {
				case "BEP-20", "ERC-20":
					req := http.NewEnableRequest(val)
					//fmt.Println(req)
					outBatch = append(outBatch, req)
					if !val.Active {
						val.Active = true
						config.GCFGRegistry[v] = val

					}
				case "UTXO", "QRC-20", "Smart Chain":
					req := http.NewElectrumRequest(val)
					//fmt.Println(req.ToJson())
					outBatch = append(outBatch, req)
					if !val.Active {
						val.Active = true
						config.GCFGRegistry[v] = val
					}
				default:
					glg.Warnf("Not supported yet")
				}
			} else {
				glg.Warnf("coin %s doesn't exist - skipping", v)
			}
		}

		p, _ := json.Marshal(outBatch)
		rawReq := string(p)
		glg.Infof("req: %s", rawReq)
		go func() {
			val, errVal := await(js.Global().Call("rpc_request", rawReq))
			if errVal != nil {
				glg.Info("not ok")
			}
			glg.Infof("ok %v", val)
		}()
		return "done"
	})
	return jsfunc
}

func Bootstrap() js.Func {
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		go func() {
			val, errVal := await(js.Global().Call("init_wasm"))
			if val != nil {
				glg.Infof("done from the promise")
				parseVal, _ := await(js.Global().Call("load_desktop_cfg_from_url", "http://localhost:8080/static/assets/wasm.coins.json"))
				if parseVal != nil {
					parseMM2Val, _ := await(js.Global().Call("load_coins_cfg_from_url", "http://localhost:8080/static/assets/coins.json"))
					if parseMM2Val != nil {
						startVal, _ := await(js.Global().Call("run_mm2", js.Global().Call("start_mm2")))
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
	http.GRuntimeUserpass = "wasmtest"
	glg.Get().SetMode(glg.STD)
	_ = glg.Info("Hello from webassembly")
	js.Global().Set("load_desktop_cfg_from_url", loadDesktopCfgFromUrl())
	js.Global().Set("load_coins_cfg_from_url", loadCoinsCfgFromUrl())
	js.Global().Set("get_ticker_infos", getTickerInfos())
	js.Global().Set("get_all_ticker_infos", getAllTickerInfos())
	js.Global().Set("start_price_service", startPriceService())
	js.Global().Set("start_mm2", StartMM2())
	js.Global().Set("enable_active_coins", enableActiveCoins())
	js.Global().Set("bootstrap", Bootstrap())
	//js.Global().Set("load_desktop_cfg_from_string", startPriceService())
	//js.Global().Set("load_desktop_cfg_from_file", startPriceService())
	<-make(chan bool)
}
