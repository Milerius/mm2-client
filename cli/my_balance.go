package cli

import (
	"fmt"
	"mm2_client/http"
)

func MyBalance(coin string) {
	resp := http.MyBalance(coin)
	if resp != nil {
		resp.ToTable()
	}
}

func MyBalanceMultipleCoins(coins []string) {
	fmt.Println(coins)
}
