package mm2_data_structure

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"mm2_client/helpers"
	"mm2_client/services"
	"os"
)

type MyBalanceRequest struct {
	Userpass string `json:"userpass"`
	Method   string `json:"method"`
	Coin     string `json:"coin"`
}

type MyBalanceAnswer struct {
	Address            string `json:"address"`
	Balance            string `json:"balance"`
	UnspendableBalance string `json:"unspendable_balance"`
	Coin               string `json:"coin"`
}

func (req *MyBalanceRequest) ToJson() string {
	b, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}

func (answer *MyBalanceAnswer) ToTable() {
	if answer.Coin != "" {
		val, _, provider := services.RetrieveUSDValIfSupported(answer.Coin)
		if val != "0" {
			val = helpers.BigFloatMultiply(answer.Balance, val, 2)
		}

		data := [][]string{
			{answer.Coin, answer.Address, answer.Balance, val, answer.UnspendableBalance, provider},
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoWrapText(false)
		table.SetHeader([]string{"Coin", "Address", "Balance", "Balance (USD)", "Unspendable", "Price Provider"})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		table.SetCenterSeparator("|")
		table.AppendBulk(data) // Add Bulk Data
		table.Render()
	}
}
