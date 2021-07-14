package services

import (
	"github.com/kpango/glg"
	"mm2_client/constants"
)

func LaunchServices() {
	glg.Info("Starting binance websocket service")
	go StartBinanceWebsocketService()
	glg.Info("Starting coingecko price service")
	go StartCoingeckoService()
	glg.Info("Starting coinpaprika price service")
	go StartCoinpaprikaService()
	constants.GPricesServicesRunning = true
}
