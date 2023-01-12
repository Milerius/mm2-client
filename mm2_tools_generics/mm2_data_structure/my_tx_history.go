package mm2_data_structure

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"mm2_client/config"
	"mm2_client/external_services"
	"mm2_client/helpers"
	"mm2_client/mm2_tools_generics/common"
	"os"
	"strconv"
	"sync"
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

type MyTxHistoryRequest struct {
	Userpass string `json:"userpass"`
	Method   string `json:"method"`
	Coin     string `json:"coin"`
	Limit    int    `json:"limit"`
	//FromId     string `json:"from_id,omitempty"`
	PageNumber int  `json:"page_number,omitempty"`
	Max        bool `json:"max,omitempty"`
}

func NewMyTxHistoryRequest(coin string, defaultNbTx int, defaultPage int, max bool) *MyTxHistoryRequest {
	genReq := NewGenericRequest("my_tx_history")
	req := &MyTxHistoryRequest{Userpass: genReq.Userpass, Method: genReq.Method, Coin: coin, Limit: defaultNbTx, PageNumber: defaultPage, Max: max}
	return req
}

func (req *MyTxHistoryRequest) ToJson() string {
	b, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}

func (answer *MyTxHistoryAnswer) ToTable(coinReq string, page int, tx int, withOriginalFiatValue bool, max bool, custom bool) {
	var data [][]string
	cfg, cfgExist := config.GCFGRegistry[coinReq]

	functor := func(timestamp int64, geckoID string, wg *sync.WaitGroup) {
		defer wg.Done()
		common.HandleGeckoPrice(timestamp, geckoID)
	}

	if withOriginalFiatValue {
		var wg sync.WaitGroup
		visited := make(map[string]bool)
		for _, curAnswer := range answer.Result.Transactions {
			if cfgExist {
				if !common.ExistInGeckoRegistry(curAnswer.Timestamp, cfg.CoingeckoID) {
					key := cfg.CoingeckoID + "-" + common.TimestampToGeckoDate(curAnswer.Timestamp)
					if _, ok := visited[key]; !ok {
						//fmt.Printf("key %s don't exist processing\n", key)
						wg.Add(1)
						go functor(curAnswer.Timestamp, cfg.CoingeckoID, &wg)
						visited[key] = true
					}
				}
			}
		}
		wg.Wait()
	}

	for _, curAnswer := range answer.Result.Transactions {
		if curAnswer.Coin != "" {
			val := "0"
			if !withOriginalFiatValue {
				val, _, _ = external_services.RetrieveUSDValIfSupported(curAnswer.Coin, 0)
				if val != "0" {
					val = helpers.BigFloatMultiply(curAnswer.MyBalanceChange, val, 2)
				}
			} else {
				if cfgExist {
					val = helpers.BigFloatMultiply(curAnswer.MyBalanceChange, common.GetFromRegistry(curAnswer.Timestamp, cfg.CoingeckoID), 2)
				}
			}

			totalFee := curAnswer.FeeDetails.Amount
			feeCoin := curAnswer.Coin
			if custom {
				totalFee = curAnswer.FeeDetails.TotalFee
				feeCoin = curAnswer.FeeDetails.Coin
			}

			txUrl := ""
			if cfgExist {
				if cfg.ExplorerTxURL != "" {
					txUrl = cfg.ExplorerURL + cfg.ExplorerTxURL + curAnswer.TxHash
				} else {
					txUrl = cfg.ExplorerURL + "tx/" + curAnswer.TxHash
				}
			}

			cur := []string{curAnswer.From[0], curAnswer.To[0], curAnswer.MyBalanceChange + " (" + val + "$)", totalFee + " " + feeCoin, helpers.GetDateFromTimestamp(curAnswer.Timestamp, true), txUrl}
			data = append(data, cur)
		}
	}

	helpers.SortDoubleSliceByDate(data, 4, false)

	table := tablewriter.NewWriter(os.Stdout)
	if !custom && !max {
		table.SetFooter([]string{"", "", "Current Page", strconv.Itoa(page), "Nb Pages", strconv.Itoa(answer.Result.TotalPages)}) // Add Footer
	}
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"From", "To", "Balance Change", "Fee", "Date", "TxUrl"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}
