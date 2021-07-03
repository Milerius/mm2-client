package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"mm2_client/config"
	"net/http"
	"os"
	"strconv"
)

type EnableRequest struct {
	Coin                 string   `json:"coin"`
	FallbackSwapContract string   `json:"fallback_swap_contract"`
	Method               string   `json:"method"`
	SwapContractAddress  string   `json:"swap_contract_address"`
	TxHistory            bool     `json:"tx_history"`
	Urls                 []string `json:"urls"`
	Userpass             string   `json:"userpass"`
}

type EnableAnswer struct {
	Coin                  string `json:"coin"`
	Address               string `json:"address"`
	Balance               string `json:"balance"`
	RequiredConfirmations int    `json:"required_confirmations"`
	RequiresNotarization  bool   `json:"requires_notarization"`
	UnspendableBalance    string `json:"unspendable_balance"`
	Result                string `json:"result"`
}

func newEnableRequest(cfg *config.DesktopCFG) *EnableRequest {
	genReq := NewGenericRequest("enable")
	req := &EnableRequest{Userpass: genReq.Userpass, Method: genReq.Method}
	//cfg := config.GCFGRegistry[coin]
	req.Coin = cfg.Coin
	req.TxHistory = true
	req.Urls = cfg.Nodes
	req.SwapContractAddress, req.FallbackSwapContract = cfg.RetrieveContracts()
	return req
}

func (req *EnableRequest) toJson() string {
	b, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}

func (answer *EnableAnswer) toTable() {
	data := [][]string{
		{answer.Coin, answer.Address, answer.Balance, strconv.Itoa(answer.RequiredConfirmations), strconv.FormatBool(answer.RequiresNotarization), answer.UnspendableBalance, answer.Result},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"Coin", "Address", "Balance", "Confirmations", "Notarization", "Unspendable", "Status"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}

func Enable(coin string) bool {
	if val, ok := config.GCFGRegistry[coin]; ok {
		req := newEnableRequest(val).toJson()
		resp, err := http.Post(GMM2Endpoint, "application/json", bytes.NewBuffer([]byte(req)))
		if err != nil {
			fmt.Printf("Err: %v\n", err)
			return false
		}
		if resp.StatusCode == http.StatusOK {
			defer resp.Body.Close()
			var answer = &EnableAnswer{}
			decodeErr := json.NewDecoder(resp.Body).Decode(answer)
			if decodeErr != nil {
				fmt.Printf("Err: %v\n", err)
				return false
			}
			answer.toTable()
			return answer.Result == "success"
		} else {
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			fmt.Printf("Err: %s\n", bodyBytes)
		}
	} else {
		fmt.Printf("coin: %s doesn't exist or is not present in the desktop configuration\n", coin)
		return false
	}
	return false
}
