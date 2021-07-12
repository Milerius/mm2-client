package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kyokomi/emoji/v2"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"mm2_client/helpers"
	"mm2_client/services"
	"net/http"
	"os"
)

type KMDRewardsInfoAnswer struct {
	Result []struct {
		AccrueStopAt   int64 `json:"accrue_stop_at"`
		AccrueStartAt  int64 `json:"accrue_start_at"`
		AccruedRewards struct {
			Accrued          string `json:"Accrued,omitempty"`
			NotAccruedReason string `json:"NotAccruedReason,omitempty"`
		} `json:"accrued_rewards"`
		Amount     string `json:"amount"`
		Height     int    `json:"height"`
		InputIndex int    `json:"input_index"`
		Locktime   int64  `json:"locktime"`
		TxHash     string `json:"tx_hash"`
	} `json:"result"`
}

func (answer *KMDRewardsInfoAnswer) ToTable() bool {
	var data [][]string
	valid := false
	for _, cur := range answer.Result {
		val, _, provider := services.RetrieveUSDValIfSupported("KMD")
		accrued := cur.AccruedRewards.Accrued
		if val != "0" {
			if cur.AccruedRewards.Accrued != "" {
				val = helpers.BigFloatMultiply(cur.AccruedRewards.Accrued, val, 2)
				valid = true
			} else {
				accrued = emoji.Sprintf(":x: %s", cur.AccruedRewards.NotAccruedReason)
				//accrued = cur.AccruedRewards.NotAccruedReason
				val = "0"
			}
		}
		toInsert := []string{
			cur.Amount, accrued, val,
			helpers.GetDateFromTimestamp(cur.AccrueStartAt, true),
			helpers.GetDateFromTimestamp(cur.AccrueStopAt, true), provider}
		data = append(data, toInsert)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	headers := []string{"Amount", "Accrued", "Accrued (USD)", "Start at", "Stop At", "Price Provider"}
	table.SetHeader(headers)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
	return valid
}

func KmdRewardsInfo() *KMDRewardsInfoAnswer {
	answerEnabled := GetEnabledCoins()
	if answerEnabled != nil && answerEnabled.Contains("KMD") {
		req := NewGenericRequest("kmd_rewards_info").ToJson()
		resp, err := http.Post(GMM2Endpoint, "application/json", bytes.NewBuffer([]byte(req)))
		if err != nil {
			fmt.Printf("Err: %v\n", err)
			return nil
		}
		if resp.StatusCode == http.StatusOK {
			defer resp.Body.Close()
			var answer = &KMDRewardsInfoAnswer{}
			decodeErr := json.NewDecoder(resp.Body).Decode(answer)
			if decodeErr != nil {
				fmt.Printf("Err: %v\n", err)
				return nil
			}
			return answer
		} else {
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			fmt.Printf("Err: %s\n", bodyBytes)
			return nil
		}
	} else {
		fmt.Println("KMD need to be enabled in order to call KmdRewardsInfo")
	}
	return nil
}
