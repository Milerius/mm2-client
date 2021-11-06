package external_services

import (
	"github.com/kpango/glg"
	"mm2_client/constants"
	"os"
)

func LaunchPriceServices() {
	glg.Info("Starting binance websocket service")
	go StartBinanceWebsocketService()
	if os.Getenv("NOMICS_API_KEY") != "" {
		glg.Info("Starting nomics price service")
		go StartNomicsService()
	}
	glg.Info("Starting coingecko price service")
	go StartCoingeckoService()
	glg.Info("Starting coinpaprika price service")
	go StartCoinpaprikaService()
	constants.GPricesServicesRunning = true
}

func LaunchMessagesService(kind string, target string) {
	go StartNotifierMessagesService(kind, target)
}

func LaunchServices(kind string, target string) {
	glg.Info("Launching price services")
	LaunchPriceServices()
	glg.Info("Launching extra services")
	LaunchMessagesService(kind, target)
}
