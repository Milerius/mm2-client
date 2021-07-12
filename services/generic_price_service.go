package services

func RetrieveUSDValIfSupported(coin string) (string, string, string) {
	val, date, provider := BinanceRetrieveUSDValIfSupported(coin)
	if val == "0" {
		val, date, provider = CoingeckoRetrieveUSDValIfSupported(coin)
	}
	if val != "0" {
		return val, date, provider
	} else {
		return val, date, "unknown"
	}
}

func RetrieveCEXRatesFromPair(base string, rel string) (string, bool, string, string) {
	val, calculated, date, provider := BinanceRetrieveCEXRatesFromPair(base, rel)
	if val == "0" {
		val, calculated, date, provider = CoingeckoRetrieveCEXRatesFromPair(base, rel)
	}

	if val != "0" {
		return val, calculated, date, provider
	} else {
		return val, calculated, date, "unknown"
	}
}
