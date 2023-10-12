package external_services

import (
	"github.com/kpango/glg"
	"mm2_client/constants"
)

func LaunchPriceServices() {
	glg.Info("Starting binance websocket service")
	go StartBinanceWebsocketService()
	if constants.GNomicsApiKey != "" {
		glg.Info("Starting nomics price service")
		go StartNomicsService()
	}
	glg.Info("Start forex service")
	go StartForexService()
	glg.Info("Starting coingecko price service")
	go StartCoingeckoService()
	glg.Info("Starting coinpaprika price service")
	go StartCoinpaprikaService()
	if constants.GLcwApiKey != "" {
		glg.Info("Start LiveCoinWatch service")
		go StartLcwService()
	}
	constants.GPricesServicesRunning = true
}

func LaunchServices(kind string, target string) {
	glg.Info("Launching price services")
	LaunchPriceServices()
}
