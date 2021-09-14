package mm2_tools_server

import (
	"fmt"
	"github.com/kpango/glg"
	"github.com/ulule/limiter/v3"
	mfasthttp "github.com/ulule/limiter/v3/drivers/middleware/fasthttp"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	"github.com/valyala/fasthttp"
	"runtime"
)

var gAppName = ""

func LaunchServer(appName string, onlyPriceService bool) {
	if runtime.GOOS == "ios" {
		glg.Get().SetMode(glg.STD)
		glg.Info("Launch MM2 Tools Server from ios")
	}

	if runtime.GOOS == "android" {
		glg.Get().SetMode(glg.STD)
		glg.Info("Launch MM2 Tools Server from android")
	}

	gAppName = appName
	router := InitRooter(onlyPriceService)
	rate, err := limiter.NewRateFromFormatted("30-M")
	if err != nil {
		glg.Fatalf("error on limiter: %v", err)
		return
	}

	store := memory.NewStore()
	glg.Info("Memory store created")

	// Create a fasthttp middleware.
	middleware := mfasthttp.NewMiddleware(limiter.New(store, rate, limiter.WithTrustForwardHeader(true)))
	glg.Info("Middleware created")

        glg.Fatal(fasthttp.ListenAndServeTLS(":"+fmt.Sprintf("%d", 1313), "/etc/letsencrypt/live/prices.komodo.live/fullchain.pem", "/etc/letsencrypt/live/prices.komodo.live/privkey.pem", middleware.Handle(router.Handler)))
}
