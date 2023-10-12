package external_services

import (
	"mm2_client/helpers"
)

func RetrieveUSDValIfSupported(coin string, expirePriceValidity int) (string, string, string) {
	//! Binance
	val, date, _, provider := BinanceRetrieveUSDValIfSupported(coin)

	elapsed := helpers.DateToTimeElapsed(date)
	expirePriceValidityF := float64(expirePriceValidity)

	//! Forex
	if val == "0" || (expirePriceValidity > 0 && elapsed > expirePriceValidityF) {
		val, date, provider = ForexRetrieveUSDValIfSupported(coin)
		if val != "0" {
			return val, date, provider
		}
	}

	//! Nomics
	if val == "0" || (expirePriceValidity > 0 && elapsed > expirePriceValidityF) {
		val, date, provider = NomicsRetrieveUSDValIfSupported(coin)
		elapsed = helpers.DateToTimeElapsed(date)
	}

	//! Gecko
	if val == "0" || (expirePriceValidity > 0 && elapsed > expirePriceValidityF) {
		val, date, provider = CoingeckoRetrieveUSDValIfSupported(coin)
		elapsed = helpers.DateToTimeElapsed(date)
	}

	//! Paprika
	if val == "0" || (expirePriceValidity > 0 && elapsed > expirePriceValidityF) {
		val, date, provider = CoinpaprikaRetrieveUSDValIfSupported(coin)
		if val == "0" {
			val, date, provider = CoinpaprikaRetrieveUSDValIfSupported(helpers.RetrieveMainTicker(coin))
		}
		elapsed = helpers.DateToTimeElapsed(date)
	}

	//! LCW
	if val == "0" || (expirePriceValidity > 0 && elapsed > expirePriceValidityF) {
		val, date, provider = LcwRetrieveUSDValIfSupported(coin)
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

	//! Nomics
	if val == "0" {
		val, calculated, date, provider = NomicsRetrieveCEXRatesFromPair(base, rel)
	}

	//! LWC
	if val == "0" {
		val, calculated, date, provider = LcwRetrieveCEXRatesFromPair(base, rel)
	}

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
	if volume == "0" {
		volume, date, provider = LcwGetTotalVolume(coin)
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
		change24h, date, provider = NomicsGetChange24h(coin)
	}

	if change24h == "0" {
		change24h, date, provider = CoingeckoGetChange24h(coin)
	}

	if change24h == "0" {
		change24h, date, provider = CoinpaprikaGetChange24h(coin)
	}

	if change24h == "0" {
		change24h, date, provider = LcwGetChange24h(coin)
	}

	if change24h != "0" {
		return change24h, date, provider
	} else {
		return change24h, date, "unknown"
	}
}
