package mm2_data_structure

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"mm2_client/external_services"
	"mm2_client/helpers"
	"os"
	"strconv"
)

type GenericEnableAnswer struct {
	Coin                  string `json:"coin"`
	Address               string `json:"address"`
	Balance               string `json:"balance"`
	RequiredConfirmations int    `json:"required_confirmations"`
	RequiresNotarization  bool   `json:"requires_notarization"`
	UnspendableBalance    string `json:"unspendable_balance"`
	Result                string `json:"result"`
	Error                 string `json:"error,omitempty"`
}

func (answer *GenericEnableAnswer) ToTable() {
	if answer.Coin != "" {
		val, _, provider := external_services.RetrieveUSDValIfSupported(answer.Coin)
		if val != "0" {
			val = helpers.BigFloatMultiply(answer.Balance, val, 2)
		}

		data := [][]string{
			{answer.Coin, answer.Address, answer.Balance, val, strconv.Itoa(answer.RequiredConfirmations), strconv.FormatBool(answer.RequiresNotarization), answer.UnspendableBalance, answer.Result, provider},
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoWrapText(false)
		table.SetHeader([]string{"Coin", "Address", "Balance", "Balance (USD)", "Confirmations", "Notarization", "Unspendable", "Status", "Price Provider"})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		table.SetCenterSeparator("|")
		table.AppendBulk(data) // Add Bulk Data
		table.Render()
	}
}

func ToTableGenericEnableAnswers(answers []GenericEnableAnswer) {
	var data [][]string

	for _, answer := range answers {
		if answer.Coin != "" {
			val, _, provider := external_services.RetrieveUSDValIfSupported(answer.Coin)
			if val != "0" {
				val = helpers.BigFloatMultiply(answer.Balance, val, 2)
			}

			cur := []string{answer.Coin, answer.Address, answer.Balance, val, strconv.Itoa(answer.RequiredConfirmations),
				strconv.FormatBool(answer.RequiresNotarization), answer.UnspendableBalance, answer.Result, provider}
			data = append(data, cur)
		} else {
			fmt.Println("Error: " + answer.Error)
		}
	}

	helpers.SortDoubleSlice(data, 3, false)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"Coin", "Address", "Balance", "Balance (USD)", "Confirmations", "Notarization", "Unspendable", "Status", "Price Provider"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}
