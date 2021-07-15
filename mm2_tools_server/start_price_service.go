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
	DesktopCfgPath string `json:"desktop_cfg_path"`
}

func StartPriceService(ctx *fasthttp.RequestCtx) {
	if constants.GPricesServicesRunning {
		ctx.SetStatusCode(200)
		ctx.SetBodyString("Price service already running - skipping")
	} else {
		out := &StartPriceRequest{}
		r := bytes.NewReader(ctx.PostBody())
		err := json.NewDecoder(r).Decode(out)
		if err != nil {
			_ = glg.Errorf("%v", err)
			ctx.SetStatusCode(http.StatusBadRequest)
			ctx.SetBodyString(err.Error())
			return
		}
		if res := config.ParseDesktopRegistryFromFile(out.DesktopCfgPath); res {
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
