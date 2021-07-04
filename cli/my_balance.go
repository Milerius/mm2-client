package cli

import "fmt"

func MyBalance(coin string) {
	fmt.Println(coin)
}

func MyBalanceMultipleCoins(coins []string) {
	fmt.Println(coins)
}
