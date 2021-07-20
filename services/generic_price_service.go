package services

func RetrieveUSDValIfSupported(coin string) (string, string, string) {
	//! Binance
	val, date, _, provider := BinanceRetrieveUSDValIfSupported(coin)

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

func RetrieveVolume24h(coin string) (string, string, string) {
	volume, date, provider := CoingeckoGetTotalVolume(coin)
	if volume == "0" {
		volume, date, provider = CoinpaprikaTotalVolume(coin)
	}
	if volume != "0" {
		return volume, date, provider
	} else {
		return volume, date, "unknown"
	}
}

func RetrieveSparkline7D(coin string) (*[]float64, string, string) {
	sparklineData, date, provider := CoingeckoGetSparkline7D(coin)
	if sparklineData == nil {
		return sparklineData, date, "unknown"
	}
	return sparklineData, date, provider
}

func RetrievePercentChange24h(coin string) (string, string, string) {
	_, date, change24h, provider := BinanceRetrieveUSDValIfSupported(coin)
	if change24h == "0" {
		change24h, date, provider = CoingeckoGetChange24h(coin)
	}
	if change24h == "0" {
		change24h, date, provider = CoinpaprikaGetChange24h(coin)
	}
	if change24h != "0" {
		return change24h, date, provider
	} else {
		return change24h, date, "unknown"
	}
}
