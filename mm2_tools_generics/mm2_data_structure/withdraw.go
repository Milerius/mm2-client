package mm2_data_structure

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
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

type WithdrawRequestParams struct {
	Coin   string `json:"coin"`
	To     string `json:"to"`
	Amount string `json:"amount,omitempty"`
	Max    bool   `json:"max,omitempty"`
	Fee    *Fee   `json:"fee,omitempty"`
}

type WithdrawRequest struct {
	Method                string                `json:"method"`
	Userpass              string                `json:"userpass"`
	MMRpc                 *string               `json:"mmrpc,omitempty"`
	WithdrawRequestParams WithdrawRequestParams `json:"params"`
}

type WithdrawAnswerSuccess struct {
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
}

type WithdrawAnswer struct {
	Mmrpc  string                 `json:"mmrpc"`
	Error  string                 `json:"error,omitempty"`
	Id     int                    `json:"id"`
	Result *WithdrawAnswerSuccess `json:"result,omitempty"`
}

func NewWithdrawRequest(coin string, amount string, address string, fees []string, coinType string) *WithdrawRequest {
	genReq := NewGenericRequestV2("withdraw")
	params := WithdrawRequestParams{Coin: coin, To: address}
	req := &WithdrawRequest{Userpass: genReq.Userpass, Method: genReq.Method, WithdrawRequestParams: params, MMRpc: genReq.MMRpc}
	if amount == "max" {
		req.WithdrawRequestParams.Max = true
	} else {
		req.WithdrawRequestParams.Amount = amount
	}
	if len(fees) > 0 {
		req.WithdrawRequestParams.Fee = &Fee{}
		switch coinType {
		case "ERC-20", "BEP-20":
			req.WithdrawRequestParams.Fee.Type = "EthGas"
			req.WithdrawRequestParams.Fee.GasPrice = fees[1]
			req.WithdrawRequestParams.Fee.Gas, _ = strconv.Atoi(fees[2])
		case "QRC-20":
			req.WithdrawRequestParams.Fee.Type = "Qrc20Gas"
			req.WithdrawRequestParams.Fee.GasPrice = fees[1]
			req.WithdrawRequestParams.Fee.GasLimit, _ = strconv.Atoi(fees[2])
		case "UTXO", "Smart Chain":
			switch fees[0] {
			case "utxo_fixed":
				req.WithdrawRequestParams.Fee.Type = "UtxoFixed"
			case "utxo_per_kbyte":
				req.WithdrawRequestParams.Fee.Type = "UtxoPerKbyte"
			}
			req.WithdrawRequestParams.Fee.Amount = fees[1]
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
	if receiver.Result != nil {
		if receiver.Result.FeeDetails.Amount != "" {
			return receiver.Result.FeeDetails.Amount
		} else {
			return receiver.Result.FeeDetails.TotalFee
		}
	}
	return "0"
}

func (receiver *WithdrawAnswer) ToTable() {
	data := [][]string{
		{receiver.Result.From[0], receiver.Result.To[0], receiver.Result.TotalAmount, receiver.Result.MyBalanceChange, receiver.RetrieveTotalFee()},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	headers := []string{"From", "To", "Amount", "Balance Change", "Fee"}
	if receiver.Result.KmdRewards != nil {
		headers = append(headers, "KMD Rewards")
		data[0] = append(data[0], receiver.Result.KmdRewards.Amount)
	}
	table.SetHeader(headers)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
	fmt.Printf("\ntx_hex: %s\n", receiver.Result.TxHex)
}
