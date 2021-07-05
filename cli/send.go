package cli

import "fmt"

func Send(coin string, amount string, address string, fees []string) {
	if resp := Withdraw(coin, amount, address, fees); resp != nil {
		resp.ToTable()
		fmt.Println()
		Broadcast(resp.Coin, resp.TxHex)
	}
}
