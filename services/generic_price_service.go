package services

func RetrieveUSDValIfSupported(coin string) (string, string, string) {
	//! Binance
	val, date, provider := BinanceRetrieveUSDValIfSupported(coin)

	//! Gecko
	if val == "0" {
		val, date, provider = CoingeckoRetrieveUSDValIfSupported(coin)
	}

	//! Paprika
	if val == "0" {
		val, date, provider = CoinpaprikaRetrieveUSDValIfSupported(coin)
	}

	//! Verification
	if val != "0" {
		return val, date, provider
	} else {
		return val, date, "unknown"
	}
}

func RetrieveCEXRatesFromPair(base string, rel string) (string, bool, string, string) {
	//! Binance
	val, calculated, date, provider := BinanceRetrieveCEXRatesFromPair(base, rel)

	//! Gecko
	if val == "0" {
		val, calculated, date, provider = CoingeckoRetrieveCEXRatesFromPair(base, rel)
	}

	//! Paprika
	if val == "0" {
		val, calculated, date, provider = CoinpaprikaRetrieveCEXRatesFromPair(base, rel)
	}

	//! Verification
	if val != "0" {
		return val, calculated, date, provider
	} else {
		return val, calculated, date, "unknown"
	}
}
