package main

import (
	"flag"
	"github.com/kpango/glg"
	"mm2_client/config"
	"mm2_client/log"
	"mm2_client/mm2_tools_server"
	"path/filepath"
)

func main() {
	flag.Parse()
	args := flag.Args()
	appName := "standard"
	if len(args) == 1 {
		appName = args[0]
	}
	infolog := log.InitLogger(filepath.Join(config.GetDesktopPath(appName), "logs"), glg.BOTH, "mm2.tools.server")
	defer infolog.Close()
	if appName == "standard" {
		_ = glg.Info("Logger initialized for app: AtomicDEX")
	} else {
		_ = glg.Infof("Logger initialized for app: %s", appName)
	}
	mm2_tools_server.LaunchServer(appName)
}
