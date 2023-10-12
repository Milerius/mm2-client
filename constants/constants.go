package constants

import (
	"mm2_client/helpers"
	"os"
	"runtime"
	"time"
)

var (
	GMM2Dir                             = helpers.GetWorkingDir() + "/mm2"
	GMM2BinPath                         = GMM2Dir + "/mm2"
	GMM2ConfPath                        = GMM2Dir + "/MM2.json"
	GMM2CoinsPath                       = GMM2Dir + "/coins.json"
	GSimpleMarketMakerConf              = GMM2Dir + "/simple_market_bot.json"
	GMM2Running                         = false
	GSimpleMarketMakerBotRunning        = false
	GSimpleMarketMakerBotStopping       = false
	GMyRecentSwapsUpdaterServiceRunning = false
	GPricesServicesRunning              = false
	GMessageServiceRunning              = false
	GDesktopCfgLoaded                   = false
	GMarketMakerCfgLoaded               = false
	GNotifyCfgLoaded                    = false
	GPricesLoopTime                     = time.Duration(60) * time.Second
	GNomicsApiKey                       = os.Getenv("NOMICS_API_KEY")
	GLcwApiKey                          = os.Getenv("LCW_API_KEY")
	GPaprikaApiKey                      = os.Getenv("PAPRIKA_API_KEY")
	GGeckoApiKey                        = os.Getenv("GECKO_API_KEY")
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
