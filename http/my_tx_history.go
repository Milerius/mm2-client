package http

import "fmt"

func MyTxHistory(coin string, defaultNbTx int, defaultPage int, withFiatValue bool, isMax bool) {
	fmt.Printf("%s %d %d %t %t\n", coin, defaultNbTx, defaultPage, withFiatValue, isMax)
}

func CustomMyTxHistory(coin string, defaultNbTx int, defaultPage int, withFiatValue bool, isMax bool) {
	fmt.Printf("%s %d %d %t %t\n", coin, defaultNbTx, defaultPage, withFiatValue, isMax)
}
