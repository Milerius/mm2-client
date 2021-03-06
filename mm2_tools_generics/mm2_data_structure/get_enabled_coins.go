package mm2_data_structure

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

type GetEnabledCoinsAnswer struct {
	Result []struct {
		Address string `json:"address"`
		Ticker  string `json:"ticker"`
	} `json:"result"`
}

func (receiver *GetEnabledCoinsAnswer) Contains(ticker string) bool {
	for _, v := range receiver.Result {
		if v.Ticker == ticker {
			return true
		}
	}

	return false
}

func (receiver *GetEnabledCoinsAnswer) ToSlice() []string {
	var out []string
	for _, cur := range receiver.Result {
		out = append(out, cur.Ticker)
	}
	return out
}

func (receiver *GetEnabledCoinsAnswer) ToSliceEmptyBalance() []string {
	var out []string
	for _, cur := range receiver.Result {
		out = append(out, cur.Ticker)
	}
	return out
}

func (receiver *GetEnabledCoinsAnswer) ToTable() {
	var data [][]string

	for _, answer := range receiver.Result {
		cur := []string{answer.Ticker, answer.Address}
		data = append(data, cur)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"Coin", "Address"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}
