package mm2_tools_server

import (
	"bytes"
	"encoding/json"
	"github.com/kpango/glg"
	"github.com/valyala/fasthttp"
	"mm2_client/config"
	"mm2_client/constants"
	"mm2_client/external_services"
	"net/http"
)

type StartNotifyRequest struct {
	NotifierCfgPath *string `json:"notifier_cfg_path,omitempty"`
	NotifierRawCfg  *string `json:"notifier_raw_cfg,omitempty"`
}

func StartNotifyService(ctx *fasthttp.RequestCtx) {
	if constants.GMessageServiceRunning {
		ctx.SetStatusCode(200)
		ctx.SetBodyString("notify service already running - skipping")
	} else {
		out := &StartNotifyRequest{}
		r := bytes.NewReader(ctx.PostBody())
		err := json.NewDecoder(r).Decode(out)
		if err != nil || (out.NotifierRawCfg == nil && out.NotifierCfgPath == nil) {
			_ = glg.Errorf("%v", err)
			ctx.SetStatusCode(http.StatusBadRequest)
			ctx.SetBodyString(err.Error())
			return
		}
		res := false
		kind := ""
		target := ""
		if out.NotifierRawCfg != nil {
			res = config.ParseNotifyCfgFromString(*out.NotifierRawCfg)
			kind = "file"
			target = *out.NotifierRawCfg
		} else if out.NotifierCfgPath != nil {
			res = config.ParseNotifyCfgFromFile(*out.NotifierCfgPath)
			kind = "file"
			target = *out.NotifierCfgPath
		}
		if res {
			_ = glg.Info("Launch message service")
			external_services.LaunchMessagesService(kind, target)
			ctx.SetStatusCode(200)
			ctx.SetBodyString("Message service launched")
		} else {
			ctx.SetStatusCode(http.StatusInternalServerError)
			ctx.SetBodyString("Couldn't parse notify cfg")
		}
	}
}
