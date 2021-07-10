package services

func RetrieveUSDValIfSupported(coin string) (string, string) {
	val, date := BinanceRetrieveUSDValIfSupported(coin)
	//! Later add coingecko/paprika
	return val, date
}
