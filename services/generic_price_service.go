package services

func RetrieveUSDValIfSupported(coin string) (string, string) {
	val, date := BinanceRetrieveUSDValIfSupported(coin)
	//! Later add coingecko/paprika
	return val, date
}

func RetrieveCEXRatesFromPair(base string, rel string) (string, bool, string) {
	val, calculated, date := BinanceRetrieveCEXRatesFromPair(base, rel)
	//! Later add coingecko / paprika
	return val, calculated, date
}
