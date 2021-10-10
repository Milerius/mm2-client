package mm2_tools_server

import "github.com/valyala/fasthttp"

func Ping(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(200)
	ctx.SetBodyString("Ok")
}
