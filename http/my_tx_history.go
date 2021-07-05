package http

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"mm2_client/config"
	"mm2_client/helpers"
	"mm2_client/services"
	"net/http"
	"os"
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
			Timestamp       int64    `json:"timestamp"`
			To              []string `json:"to"`
			TotalAmount     string   `json:"total_amount"`
			TxHash          string   `json:"tx_hash"`
			TxHex           string   `json:"tx_hex"`
		} `json:"transactions"`
	} `json:"result"`
	CoinType string
}

func (answer *MyTxHistoryAnswer) ToTable(page int, tx int, withOriginalFiatValue bool, max bool, custom bool) {
	var data [][]string

	for _, curAnswer := range answer.Result.Transactions {
		if curAnswer.Coin != "" {
			val := "0"
			if !withOriginalFiatValue {
				val = services.RetrieveUSDValIfSupported(curAnswer.Coin)
				if val != "0" {
					val = helpers.BigFloatMultiply(curAnswer.MyBalanceChange, val, 2)
				}
			}

			totalFee := curAnswer.FeeDetails.Amount
			feeCoin := curAnswer.Coin
			if custom {
				totalFee = curAnswer.FeeDetails.TotalFee
				feeCoin = curAnswer.FeeDetails.Coin
			}

			txUrl := ""
			coin := curAnswer.Coin
			if (answer.CoinType == "ERC20" || answer.CoinType == "BEP20") &&
				(curAnswer.Coin != "BNB" && curAnswer.Coin != "BNBT" && curAnswer.Coin != "ETH" && curAnswer.Coin != "ETHR") {
				coin = coin + "-" + answer.CoinType
			}

			if cfg, ok := config.GCFGRegistry[coin]; ok {
				if cfg.ExplorerTxURL != "" {
					txUrl = cfg.ExplorerURL[0] + cfg.ExplorerTxURL + curAnswer.TxHash
				} else {
					txUrl = cfg.ExplorerURL[0] + "tx/" + curAnswer.TxHash
				}
			}

			cur := []string{curAnswer.From[0], curAnswer.To[0], curAnswer.MyBalanceChange + " (" + val + "$)", totalFee + " " + feeCoin, helpers.GetDateFromTimestamp(curAnswer.Timestamp, true), txUrl}
			data = append(data, cur)
		}
	}

	helpers.SortDoubleSliceByDate(data, 4, false)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"From", "To", "Balance Change", "Fee", "Date", "TxUrl"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}

const customTxEndpoint = "https://komodo.live:3334/api/"

func MyTxHistory(coin string, defaultNbTx int, defaultPage int, withFiatValue bool, isMax bool) *MyTxHistoryAnswer {
	fmt.Printf("%s %d %d %t %t\n", coin, defaultNbTx, defaultPage, withFiatValue, isMax)
	return nil
}

func CustomMyTxHistory(coin string, defaultNbTx int, defaultPage int, withFiatValue bool, isMax bool, contract string,
	query string, address string, coinType string) *MyTxHistoryAnswer {
	endpoint := customTxEndpoint
	if contract != "" {
		endpoint = endpoint + "v2/" + query + "/" + contract + "/" + address
	} else {
		endpoint = endpoint + "v1/" + query + "/" + address
	}
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
	cResp.CoinType = coinType
	return cResp
}
