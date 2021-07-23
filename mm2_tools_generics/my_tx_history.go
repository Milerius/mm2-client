package mm2_tools_generics

import (
	"errors"
	"fmt"
	"mm2_client/config"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"mm2_client/mm2_tools_generics/mm2_http_request"
	"mm2_client/mm2_tools_generics/mm2_wasm_request"
	"runtime"
	"strconv"
)

func MyTxHistory(coin string, defaultNbTx int, defaultPage int,
	withFiatValue bool, isMax bool) (*mm2_data_structure.MyTxHistoryAnswer, error) {
	if runtime.GOARCH == "wasm" {
		return mm2_wasm_request.MyTxHistory(coin, defaultNbTx, defaultPage, withFiatValue, isMax)
	} else {
		return mm2_http_request.MyTxHistory(coin, defaultNbTx, defaultPage, withFiatValue, isMax)
	}
}

func processTxHistory(coin string, args []string) (*mm2_data_structure.MyTxHistoryAnswer, int, int, bool, bool, bool, error) {
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
			if resp, err := MyBalance(val.Coin); resp != nil {
				answer, errAnswer := mm2_http_request.CustomMyTxHistory(coin, defaultNbTx, defaultPage, withFiatValue,
					isMax, contract, toQuery, resp.Address, "BEP20")
				return answer, defaultPage, defaultNbTx, withFiatValue, isMax, true, errAnswer
			} else {
				fmt.Println(err)
			}
		case "ERC-20":
			contract := ""
			toQuery := "eth_tx_history"
			if val.Coin != "ETH" && val.Coin != "ETHR" {
				contract = config.RetrieveContractsInfo(val.Coin)
				toQuery = "erc_tx_history"
			}
			if resp, err := MyBalance(val.Coin); resp != nil {
				answer, answerErr := mm2_http_request.CustomMyTxHistory(coin, defaultNbTx, defaultPage, withFiatValue, isMax, contract, toQuery, resp.Address, "ERC20")
				return answer, defaultPage, defaultNbTx, withFiatValue, isMax, true, answerErr
			} else {
				fmt.Println(err)
			}
		default:
			answer, answerErr := MyTxHistory(coin, defaultNbTx, defaultPage, withFiatValue, isMax)
			return answer, defaultPage, defaultNbTx, withFiatValue, isMax, false, answerErr
		}
	}
	return nil, 0, 0, false, false, false, errors.New("unknown error")
}

// MyTxHistoryCLI /**
// eg MyTxHistoryCLI("KMD") //< shortcut for the last 50 transactions with the fiat price of now
// eg MyTxHistoryCLI("KMD", "50", "1") //< return 50 last transaction page 1
// eg MyTxHistoryCLI("KMD", "50", "1", "true") //< return 50 last transaction page 1 with fiat price at the time of the tx
// eg MyTxHistoryCLI("KMD", "50", "1", "false") //< return 50 last transaction page 1 with fiat price of now
// eg MyTxHistoryCLI("KMD", "max") //< return all transactions
// eg MyTxHistoryCLI("KMD", "50", "2") //< return 50 last transactions page 2
func MyTxHistoryCLI(coin string, args []string) {
	if resp, page, nbTx, withFiatValue, isMax, isCustom, err := processTxHistory(coin, args); resp != nil {
		resp.ToTable(coin, page, nbTx, withFiatValue, isMax, isCustom)
	} else {
		fmt.Println(err)
	}
}
