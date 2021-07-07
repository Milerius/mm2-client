package cli

import (
	"mm2_client/http"
)

func processMyRecentSwaps(limit string, pageNumber string, args []string) *http.MyRecentSwapsAnswer {
	baseCoin := ""
	relCoin := ""
	from := ""
	to := ""
	if len(args) >= 1 {
		baseCoin = args[0]
	} else if len(args) >= 2 {
		relCoin = args[1]
	} else if len(args) >= 3 {
		from = args[2]
	} else if len(args) >= 4 {
		to = args[3]
	}
	return http.ProcessMyRecentSwaps(limit, pageNumber, baseCoin, relCoin, from, to)
}

func MyRecentSwaps(limit string, pageNumber string, args []string) {
	if resp := processMyRecentSwaps(limit, pageNumber, args); resp != nil {
		resp.ToTable()
	}
}
