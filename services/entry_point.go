package services

import (
	"mm2_client/helpers"
)

func LaunchServices() {
	helpers.PrintCheck("Starting binance websocket service", true)
	go StartBinanceWebsocketService()
}
