package constants

import (
	"mm2_client/helpers"
	"os"
	"runtime"
)

var (
	GMM2Dir                       = helpers.GetWorkingDir() + "/mm2"
	GMM2BinPath                   = GMM2Dir + "/mm2"
	GMM2ConfPath                  = GMM2Dir + "/MM2.json"
	GMM2CoinsPath                 = GMM2Dir + "/coins.json"
	GSimpleMarketMakerConf        = GMM2Dir + "/simple_market_bot.json"
	GMM2Running                   = false
	GSimpleMarketMakerBotRunning  = false
	GSimpleMarketMakerBotStopping = false
	GPricesServicesRunning        = false
	GDesktopCfgLoaded             = false
)

func GetAppDataPath() string {
	switch runtime.GOOS {
	case "linux":
		return os.Getenv("HOME") + "/.atomicdex_cli"
	case "darwin":
		return os.Getenv("HOME") + "/Library/Application Support/atomicdex_cli"
	case "windows":
		return os.Getenv("APPDATA") + "/atomicdex_cli"
	default:
		return os.Getenv("HOME") + "/atomicdex_cli"
	}
}
