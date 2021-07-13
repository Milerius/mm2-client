package services

func LaunchServices() {
	infoLogger.Println("Starting binance websocket service")
	go StartBinanceWebsocketService()
	infoLogger.Println("Starting coingecko price service")
	go StartCoingeckoService()
	infoLogger.Println("Starting coinpaprika price service")
	go StartCoinpaprikaService()
}
