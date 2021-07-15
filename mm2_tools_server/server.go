package mm2_tools_server

import (
	"fmt"
	"github.com/kpango/glg"
	"github.com/ulule/limiter/v3"
	mfasthttp "github.com/ulule/limiter/v3/drivers/middleware/fasthttp"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	"github.com/valyala/fasthttp"
	"log"
	"runtime"
)

var gAppName = ""

func LaunchServer(appName string) {
	if runtime.GOOS == "ios" {
		glg.Get().SetMode(glg.STD)
		glg.Info("Launch MM2 Tools Server from ios")
	}

	gAppName = appName
	router := InitRooter()
	rate, err := limiter.NewRateFromFormatted("30-M")
	if err != nil {
		log.Fatal(err)
		return
	}

	store := memory.NewStore()

	// Create a fasthttp middleware.
	middleware := mfasthttp.NewMiddleware(limiter.New(store, rate, limiter.WithTrustForwardHeader(true)))

	glg.Fatal(fasthttp.ListenAndServe(":"+fmt.Sprintf("%d", 1313), middleware.Handle(router.Handler)))
}
