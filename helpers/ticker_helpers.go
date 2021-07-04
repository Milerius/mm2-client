package helpers

import "strings"

func RetrieveMainTicker(ticker string) string {
	if strings.Contains(ticker, "-") {
		return ticker[0:strings.Index(ticker, "-")]
	}
	return ticker
}
