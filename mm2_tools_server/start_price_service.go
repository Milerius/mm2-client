package mm2_tools_server

import (
	"bytes"
	"encoding/json"
	"github.com/kpango/glg"
	"github.com/valyala/fasthttp"
	"mm2_client/config"
	"mm2_client/constants"
	"mm2_client/services"
	"net/http"
)

type StartPriceRequest struct {
	DesktopCfgPath *string `json:"desktop_cfg_path,omitempty"`
	DesktopRawCfg  *string `json:"desktop_raw_cfg,omitempty"`
}

func StartPriceService(ctx *fasthttp.RequestCtx) {
	if constants.GPricesServicesRunning {
		ctx.SetStatusCode(200)
		ctx.SetBodyString("Price service already running - skipping")
	} else {
		out := &StartPriceRequest{}
		r := bytes.NewReader(ctx.PostBody())
		err := json.NewDecoder(r).Decode(out)
		if err != nil || (out.DesktopRawCfg == nil && out.DesktopCfgPath == nil) {
			_ = glg.Errorf("%v", err)
			ctx.SetStatusCode(http.StatusBadRequest)
			ctx.SetBodyString(err.Error())
			return
		}
		res := false
		if out.DesktopRawCfg != nil {
			res = config.ParseDesktopRegistryFromString(*out.DesktopRawCfg)
		} else if out.DesktopCfgPath != nil {
			res = config.ParseDesktopRegistryFromFile(*out.DesktopCfgPath)
		}
		if res {
			_ = glg.Info("Launch price services")
			services.LaunchServices()
			ctx.SetStatusCode(200)
			ctx.SetBodyString("Price service launched")
		} else {
			ctx.SetStatusCode(http.StatusInternalServerError)
			ctx.SetBodyString("Couldn't parse desktop cfg")
		}
	}
}
