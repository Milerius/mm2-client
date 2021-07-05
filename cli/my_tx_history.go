package cli

import (
	"fmt"
	"mm2_client/config"
	"mm2_client/http"
	"strconv"
)

func processTxHistory(coin string, args []string) *http.MyTxHistoryAnswer {
	if val, ok := config.GCFGRegistry[coin]; ok {
		defaultNbTx := 50
		defaultPage := 1
		withFiatValue := false
		isMax := false
		if len(args) >= 1 {
			if args[0] == "max" {
				isMax = true
			} else {
				defaultNbTx, _ = strconv.Atoi(args[0])
			}
		}
		if len(args) >= 2 {
			defaultPage, _ = strconv.Atoi(args[1])
		}
		if len(args) >= 3 {
			withFiatValue, _ = strconv.ParseBool(args[2])
		}
		switch val.Type {
		case "BEP-20":
			contract := ""
			toQuery := "bnb_tx_history"
			if val.Coin != "BNB" && val.Coin != "BNBT" {
				contract = config.RetrieveContractsInfo(val.Coin)
				toQuery = "bep_tx_history"
			}
			if resp := http.MyBalance(val.Coin); resp != nil {
				return http.CustomMyTxHistory(coin, defaultNbTx, defaultPage, withFiatValue, isMax, contract, toQuery, resp.Address)
			}
		case "ERC-20":
			contract := ""
			toQuery := "eth_tx_history"
			if val.Coin != "ETH" && val.Coin != "ETHR" {
				contract = config.RetrieveContractsInfo(val.Coin)
				toQuery = "erc_tx_history"
			}
			if resp := http.MyBalance(val.Coin); resp != nil {
				return http.CustomMyTxHistory(coin, defaultNbTx, defaultPage, withFiatValue, isMax, contract, toQuery, resp.Address)
			}
		default:
			return http.MyTxHistory(coin, defaultNbTx, defaultPage, withFiatValue, isMax)
		}
	}
	return nil
}

// MyTxHistory /**
// eg MyTxHistory("KMD") //< shortcut for the last 50 transactions with the fiat price of now
// eg MyTxHistory("KMD", "50", "1") //< return 50 last transaction page 1
// eg MyTxHistory("KMD", "50", "1", "true") //< return 50 last transaction page 1 with fiat price at the time of the tx
// eg MyTxHistory("KMD", "50", "1", "false") //< return 50 last transaction page 1 with fiat price of now
// eg MyTxHistory("KMD", "max") //< return all transactions
// eg MyTxHistory("KMD", "50", 2) //< return 50 last transactions page 2
func MyTxHistory(coin string, args []string) {
	if resp := processTxHistory(coin, args); resp != nil {
		fmt.Println("success")
	}
}
