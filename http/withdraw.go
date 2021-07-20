package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"mm2_client/config"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"net/http"
	"os"
	"strconv"
)

type Fee struct {
	Type     string `json:"type,omitempty"`
	Amount   string `json:"amount,omitempty"`
	GasPrice string `json:"gas_price,omitempty"`
	Gas      int    `json:"gas,omitempty"`
	GasLimit int    `json:"gas_limit,omitempty"`
}

type WithdrawRequest struct {
	Method   string `json:"method"`
	Userpass string `json:"userpass"`
	Coin     string `json:"coin"`
	To       string `json:"to"`
	Amount   string `json:"amount,omitempty"`
	Max      bool   `json:"max,omitempty"`
	Fee      *Fee   `json:"fee,omitempty"`
}

type WithdrawAnswer struct {
	BlockHeight int    `json:"block_height"`
	Coin        string `json:"coin"`
	FeeDetails  *struct {
		Type     string `json:"type,omitempty"`
		Coin     string `json:"coin,omitempty"`
		Amount   string `json:"amount,omitempty"`    //< UTXO, Smart Chain
		Gas      int    `json:"gas,omitempty"`       //< QRC, ERC, BEP
		GasPrice string `json:"gas_price,omitempty"` //< ERC, BEP
		TotalFee string `json:"total_fee,omitempty"` //< ERC, QRC, BEP
		MinerFee string `json:"miner_fee,omitempty"` //< QRC
		GasLimit int    `json:"gas_limit,omitempty"` //< QRC
	} `json:"fee_details,omitempty"`
	From            []string `json:"from"`
	MyBalanceChange string   `json:"my_balance_change"`
	ReceivedByMe    string   `json:"received_by_me"`
	SpentByMe       string   `json:"spent_by_me"`
	To              []string `json:"to"`
	TotalAmount     string   `json:"total_amount"`
	TxHash          string   `json:"tx_hash"`
	TxHex           string   `json:"tx_hex"`
	KmdRewards      *struct {
		Amount      string `json:"amount"`
		ClaimedByMy bool   `json:"claimed_by_my"`
	} `json:"kmd_rewards,omitempty"`
	Error string `json:"error,omitempty"`
}

func NewWithdrawRequest(coin string, amount string, address string, fees []string, coinType string) *WithdrawRequest {
	genReq := mm2_data_structure.NewGenericRequest("withdraw")
	req := &WithdrawRequest{Userpass: genReq.Userpass, Method: genReq.Method, Coin: coin, To: address}
	if amount == "max" {
		req.Max = true
	} else {
		req.Amount = amount
	}
	if len(fees) > 0 {
		req.Fee = &Fee{}
		switch coinType {
		case "ERC-20", "BEP-20":
			req.Fee.Type = "EthGas"
			req.Fee.GasPrice = fees[1]
			req.Fee.Gas, _ = strconv.Atoi(fees[2])
		case "QRC-20":
			req.Fee.Type = "Qrc20Gas"
			req.Fee.GasPrice = fees[1]
			req.Fee.GasLimit, _ = strconv.Atoi(fees[2])
		case "UTXO", "Smart Chain":
			switch fees[0] {
			case "utxo_fixed":
				req.Fee.Type = "UtxoFixed"
			case "utxo_per_kbyte":
				req.Fee.Type = "UtxoPerKbyte"
			}
			req.Fee.Amount = fees[1]
		}
	}
	return req
}

func (req *WithdrawRequest) ToJson() string {
	b, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}

func (receiver *WithdrawAnswer) RetrieveTotalFee() string {
	if receiver.FeeDetails.Amount != "" {
		return receiver.FeeDetails.Amount
	} else {
		return receiver.FeeDetails.TotalFee
	}
}

func (receiver *WithdrawAnswer) ToTable() {
	data := [][]string{
		{receiver.From[0], receiver.To[0], receiver.TotalAmount, receiver.MyBalanceChange, receiver.RetrieveTotalFee()},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	headers := []string{"From", "To", "Amount", "Balance Change", "Fee"}
	if receiver.KmdRewards != nil {
		headers = append(headers, "KMD Rewards")
		data[0] = append(data[0], receiver.KmdRewards.Amount)
	}
	table.SetHeader(headers)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
	fmt.Printf("\ntx_hex: %s\n", receiver.TxHex)
}

func Withdraw(coin string, amount string, address string, fees []string, coinType string) *WithdrawAnswer {
	//NewWithdrawRequest(coin, address, amount, fees, coinType)
	if _, ok := config.GCFGRegistry[coin]; ok {
		req := NewWithdrawRequest(coin, amount, address, fees, coinType).ToJson()
		//fmt.Println(req)
		resp, err := http.Post(mm2_data_structure.GMM2Endpoint, "application/json", bytes.NewBuffer([]byte(req)))
		if err != nil {
			fmt.Printf("Err: %v\n", err)
			return nil
		}
		if resp.StatusCode == http.StatusOK {
			defer resp.Body.Close()
			var answer = &WithdrawAnswer{}
			decodeErr := json.NewDecoder(resp.Body).Decode(answer)
			if decodeErr != nil {
				fmt.Printf("Err: %v\n", err)
				return nil
			}
			return answer
		} else {
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			fmt.Printf("Err: %s\n", bodyBytes)
		}
	} else {
		fmt.Printf("coin: %s doesn't exist or is not present in the desktop configuration\n", coin)
		return nil
	}
	return nil
}
