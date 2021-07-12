package services

func RetrieveUSDValIfSupported(coin string) (string, string) {
	val, date := BinanceRetrieveUSDValIfSupported(coin)
	if val == "0" {
		val, date = CoingeckoRetrieveUSDValIfSupported(coin)
	}
	return val, date
}

func RetrieveCEXRatesFromPair(base string, rel string) (string, bool, string) {
	val, calculated, date := BinanceRetrieveCEXRatesFromPair(base, rel)
	if val == "0" {
		val, calculated, date = CoingeckoRetrieveCEXRatesFromPair(base, rel)
	}
	//! Later add coingecko / paprika
	return val, calculated, date
}
