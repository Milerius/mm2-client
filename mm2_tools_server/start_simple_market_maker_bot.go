package mm2_tools_server

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/kpango/glg"
	"github.com/valyala/fasthttp"
	"mm2_client/config"
	"mm2_client/constants"
	mm2_http "mm2_client/http"
	"mm2_client/market_making"
	"mm2_client/services"
	"net/http"
)

type StartSimpleMarketMakerRequest struct {
	DesktopCfgPath     string `json:"desktop_cfg_path"`
	Mm2CoinsCfgPath    string `json:"mm2_coins_cfg_path"`
	MarketMakerCfgPath string `json:"market_maker_cfg_path"`
	Mm2Userpass        string `json:"mm2_userpass"`
}

func internalStartSimpleMarketMakerBot(out *StartSimpleMarketMakerRequest) error {
	//! We assume mm2 is already running
	constants.GMM2Running = true
	if out.Mm2Userpass != "" {
		mm2_http.GRuntimeUserpass = out.Mm2Userpass
	} else {
		return errors.New("mm2 userpass cannot be empty")
	}
	if res := config.ParseDesktopRegistryFromFile(out.DesktopCfgPath); res {
		_ = glg.Infof("Desktop cfg successfully parsed, nb coins: %d", len(config.GCFGRegistry))
		res = config.ParseMM2CFGRegistryFromFile(out.Mm2CoinsCfgPath)
		if res {
			//! Launch price services
			services.LaunchServices()

			//! Launch the bot afterwards
			return market_making.StartSimpleMarketMakerBot(out.MarketMakerCfgPath, gAppName)
		} else {
			return errors.New("couldn't parse MM2 coins file")
		}
	} else {
		return errors.New("couldn't parse desktop cfg")
	}
}

func StartSimpleMarketMakerBot(ctx *fasthttp.RequestCtx) {
	out := &StartSimpleMarketMakerRequest{}
	r := bytes.NewReader(ctx.PostBody())
	err := json.NewDecoder(r).Decode(out)
	if err != nil {
		_ = glg.Errorf("%v", err)
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.SetBodyString(err.Error())
		return
	}
	err = internalStartSimpleMarketMakerBot(out)
	if err == nil {
		ctx.SetStatusCode(200)
		ctx.SetBodyString("Successfully started")
		//ctx.SetContentType("application/json")
	} else {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		_ = glg.Errorf("Error during initialization: %v", err)
		ctx.SetBodyString(err.Error())
	}
}
