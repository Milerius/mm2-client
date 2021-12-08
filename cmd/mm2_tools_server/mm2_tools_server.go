package main

import (
	"flag"
	"github.com/kpango/glg"
	"mm2_client/config"
	"mm2_client/external_services"
	"mm2_client/log"
	"mm2_client/mm2_tools_server"
	"path/filepath"
	"strings"
)

func main() {
	onlyPriceService := false
	flag.BoolVar(&onlyPriceService, "only_price_service", false, "-only_price_service=true")
	flag.Parse()
	args := flag.Args()
	appName := "standard"
	if len(args) == 1 {
		appName = args[0]
	}
	infolog := log.InitLogger(filepath.Join(config.GetDesktopPath(appName), "logs"), glg.BOTH, "mm2.tools.server")
	if onlyPriceService {
		glg.Info("only price service is true")
	}
	defer infolog.Close()
	if appName == "standard" {
		_ = glg.Info("Logger initialized for app: AtomicDEX")
	} else {
		_ = glg.Infof("Logger initialized for app: %s", appName)
	}
	matches, _ := filepath.Glob("*.json")
	for _, curMatch := range matches {
		if strings.Contains(curMatch, "coins.json") {
			res := config.ParseDesktopRegistryFromFile(curMatch)
			if res {
				glg.Info("starting price service from within the server")
				external_services.LaunchPriceServices()
			}
		}
	}
	mm2_tools_server.LaunchServer(appName, onlyPriceService)
}
