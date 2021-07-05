package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MyTxHistoryAnswer struct {
	Result struct {
		Skipped      int         `json:"skipped"`
		Limit        int         `json:"limit"`
		Total        int         `json:"total"`
		CurrentBlock int         `json:"current_block,omitempty"`
		PageNumber   interface{} `json:"page_number,omitempty"`
		SyncStatus   struct {
			State string `json:"state,omitempty"`
		} `json:"sync_status,omitempty"`
		TotalPages   int `json:"total_pages"`
		Transactions []struct {
			BlockHeight   int    `json:"block_height"`
			Coin          string `json:"coin"`
			Confirmations int    `json:"confirmations"`
			FeeDetails    struct {
				Coin     string `json:"coin"`
				Gas      int    `json:"gas,omitempty"`
				GasPrice string `json:"gas_price,omitempty"`
				Amount   string `json:"amount,omitempty"`
				TotalFee string `json:"total_fee,omitempty"`
			} `json:"fee_details"`
			From            []string `json:"from"`
			InternalId      string   `json:"internal_id"`
			MyBalanceChange string   `json:"my_balance_change"`
			ReceivedByMe    string   `json:"received_by_me"`
			SpentByMe       string   `json:"spent_by_me"`
			Timestamp       int      `json:"timestamp"`
			To              []string `json:"to"`
			TotalAmount     string   `json:"total_amount"`
			TxHash          string   `json:"tx_hash"`
			TxHex           string   `json:"tx_hex"`
		} `json:"transactions"`
	} `json:"result"`
}

const customTxEndpoint = "https://komodo.live:3334/api/"

func MyTxHistory(coin string, defaultNbTx int, defaultPage int, withFiatValue bool, isMax bool) *MyTxHistoryAnswer {
	fmt.Printf("%s %d %d %t %t\n", coin, defaultNbTx, defaultPage, withFiatValue, isMax)
	return nil
}

func CustomMyTxHistory(coin string, defaultNbTx int, defaultPage int, withFiatValue bool, isMax bool, contract string,
	query string, address string) *MyTxHistoryAnswer {
	endpoint := customTxEndpoint
	if contract != "" {
		endpoint = endpoint + "v2/" + query + "/" + contract + "/" + address
	} else {
		endpoint = endpoint + "v1/" + query + "/" + address
	}
	fmt.Printf("%s %d %d %t %t %s\n", coin, defaultNbTx, defaultPage, withFiatValue, isMax, endpoint)
	resp, err := http.Get(endpoint)
	if err != nil {
		fmt.Printf("Error occured: %v\n", err)
		return nil
	}
	defer resp.Body.Close()
	var cResp = new(MyTxHistoryAnswer)
	if decodeErr := json.NewDecoder(resp.Body).Decode(cResp); decodeErr != nil {
		fmt.Printf("Error occured: %v\n", decodeErr)
		return nil
	}

	return cResp
}
