package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"net/http"
	"os"
)

type GetEnabledCoinsAnswer struct {
	Result []struct {
		Address string `json:"address"`
		Ticker  string `json:"ticker"`
	} `json:"result"`
}

func (receiver *GetEnabledCoinsAnswer) ToSlice() []string {
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

func GetEnabledCoins() *GetEnabledCoinsAnswer {
	req := NewGenericRequest("get_enabled_coins").ToJson()
	resp, err := http.Post(GMM2Endpoint, "application/json", bytes.NewBuffer([]byte(req)))
	if err != nil {
		fmt.Printf("Err: %v\n", err)
		return nil
	}
	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		res := &GetEnabledCoinsAnswer{}
		decodeErr := json.NewDecoder(resp.Body).Decode(res)
		if decodeErr != nil {
			fmt.Printf("Err: %v\n", err)
			return nil
		}
		return res
	}
	return nil
}
