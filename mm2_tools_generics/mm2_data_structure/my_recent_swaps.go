package mm2_data_structure

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"mm2_client/external_services"
	"mm2_client/helpers"
	"os"
	"strconv"
)

type MyRecentSwapsRequest struct {
	Userpass      string `json:"userpass"`
	Method        string `json:"method"`
	Limit         int    `json:"limit"`
	PageNumber    int    `json:"page_number"`
	MyCoin        string `json:"my_coin,omitempty"`
	OtherCoin     string `json:"other_coin,omitempty"`
	FromTimestamp int64  `json:"from_timestamp,omitempty"`
	ToTimestamp   int64  `json:"to_timestamp,omitempty"`
}

type MyInfoContents struct {
	MyAmount    string `json:"my_amount"`
	MyCoin      string `json:"my_coin"`
	OtherAmount string `json:"other_amount"`
	OtherCoin   string `json:"other_coin"`
	StartedAt   int64  `json:"started_at"`
}

type SwapEventData struct {
	Error                     string `json:"error,omitempty"`
	LockDuration              int    `json:"lock_duration,omitempty"`
	MakerAmount               string `json:"maker_amount,omitempty"`
	MakerCoin                 string `json:"maker_coin,omitempty"`
	MakerCoinStartBlock       int    `json:"maker_coin_start_block,omitempty"`
	MakerPaymentConfirmations int    `json:"maker_payment_confirmations,omitempty"`
	MakerPaymentRequiresNota  bool   `json:"maker_payment_requires_nota,omitempty"`
	MakerPaymentLock          int    `json:"maker_payment_lock,omitempty"`
	MyPersistentPub           string `json:"my_persistent_pub,omitempty"`
	Secret                    string `json:"secret,omitempty"`
	StartedAt                 int    `json:"started_at,omitempty"`
	Taker                     string `json:"taker,omitempty"`
	TakerAmount               string `json:"taker_amount,omitempty"`
	TakerCoin                 string `json:"taker_coin,omitempty"`
	TakerCoinStartBlock       int    `json:"taker_coin_start_block,omitempty"`
	TakerPaymentConfirmations int    `json:"taker_payment_confirmations,omitempty"`
	TakerPaymentRequiresNota  bool   `json:"taker_payment_requires_nota,omitempty"`
	Uuid                      string `json:"uuid,omitempty"`
	TakerPaymentLocktime      int    `json:"taker_payment_locktime,omitempty"`
	TakerPubkey               string `json:"taker_pubkey,omitempty"`
	TxHash                    string `json:"tx_hash,omitempty"`
	TxHex                     string `json:"tx_hex,omitempty"`
	Maker                     string `json:"maker,omitempty"`
	MakerPaymentWait          int    `json:"maker_payment_wait,omitempty"`
	TakerPaymentLock          int    `json:"taker_payment_lock,omitempty"`
	MakerPaymentLocktime      int    `json:"maker_payment_locktime,omitempty"`
	MakerPubkey               string `json:"maker_pubkey,omitempty"`
	SecretHash                string `json:"secret_hash,omitempty"`
	Transaction               struct {
		TxHash string `json:"tx_hash"`
		TxHex  string `json:"tx_hex"`
	} `json:"transaction,omitempty"`
}

type SwapEvent struct {
	Data SwapEventData `json:"data,omitempty"`
	Type string        `json:"type"`
}

type SwapEventContent struct {
	Event     SwapEvent `json:"event"`
	Timestamp int64     `json:"timestamp"`
}

type SwapContent struct {
	ErrorEvents   []string           `json:"error_events"`
	Events        []SwapEventContent `json:"events"`
	MyInfo        MyInfoContents     `json:"my_info"`
	MakerCoin     string             `json:"maker_coin"`
	MakerAmount   string             `json:"maker_amount"`
	TakerCoin     string             `json:"taker_coin"`
	TakerAmount   string             `json:"taker_amount"`
	Gui           string             `json:"gui,omitempty"`
	MmVersion     string             `json:"mm_version"`
	SuccessEvents []string           `json:"success_events"`
	Type          string             `json:"type"`
	Uuid          string             `json:"uuid"`
	MyOrderUuid   string             `json:"my_order_uuid,omitempty"`
}

type MyRecentSwapsAnswerSuccess struct {
	FromUuid     string        `json:"from_uuid"`
	Limit        int           `json:"limit"`
	Skipped      int           `json:"skipped"`
	Total        int           `json:"total"`
	FoundRecords int           `json:"found_records"`
	PageNumber   int           `json:"page_number,omitempty"`
	TotalPages   int           `json:"total_pages"`
	Swaps        []SwapContent `json:"swaps"`
}

type MyRecentSwapsAnswer struct {
	Result MyRecentSwapsAnswerSuccess `json:"result"`
}

func (receiver *SwapContent) GetLastStatus() string {
	if len(receiver.Events) > 0 {
		lastEvent := receiver.Events[len(receiver.Events)-1].Event.Type
		if lastEvent == "Finished" {
			for _, cur := range receiver.Events {
				if cur.Event.Data.Error != "" {
					lastEvent = cur.Event.Type
				}
			}
		}
		return lastEvent
	} else {
		return "Unknown"
	}
}

func (receiver *SwapContent) ToMessage() string {
	str := fmt.Sprintf("%s: %s (%s) vs %s (%s), status: %s", receiver.Uuid, receiver.MyInfo.MyCoin, receiver.MyInfo.MyAmount, receiver.MyInfo.OtherCoin, receiver.MyInfo.OtherAmount, receiver.GetLastStatus())
	return str
}

func NewMyRecentSwapsRequest(limit string, pageNumber string, baseCoin string, relCoin string, from string, to string) *MyRecentSwapsRequest {
	genReq := NewGenericRequest("my_recent_swaps")
	limitNb, err := strconv.Atoi(limit)
	if err != nil {
		limitNb = 50
	}
	pageNumberNb, pageNumberErr := strconv.Atoi(pageNumber)
	if pageNumberErr != nil {
		pageNumberNb = 50
	}
	req := &MyRecentSwapsRequest{Userpass: genReq.Userpass, Method: genReq.Method, Limit: limitNb, PageNumber: pageNumberNb}
	if len(baseCoin) > 0 {
		req.MyCoin = baseCoin
	}
	if len(relCoin) > 0 {
		req.OtherCoin = relCoin
	}
	if len(from) > 0 {
		req.FromTimestamp = helpers.SimpleDateToTimestamp(from)
	}
	if len(to) > 0 {
		req.ToTimestamp = helpers.SimpleDateToTimestamp(to)
	}
	return req
}

func (req *MyRecentSwapsRequest) ToJson() string {
	b, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}

func (answer *MyRecentSwapsAnswer) ToTable() {
	var data [][]string
	baseTotalValue := "0"
	relTotalValue := "0"

	for _, cur := range answer.Result.Swaps {
		baseVal, _, _ := external_services.RetrieveUSDValIfSupported(cur.MyInfo.MyCoin, 0)
		if baseVal != "0" {
			baseVal = helpers.BigFloatMultiply(cur.MyInfo.MyAmount, baseVal, 2)
		}
		relVal, _, _ := external_services.RetrieveUSDValIfSupported(cur.MyInfo.OtherCoin, 0)
		if relVal != "0" {
			relVal = helpers.BigFloatMultiply(cur.MyInfo.OtherAmount, relVal, 2)
		}
		out := []string{cur.MyInfo.MyCoin, helpers.ResizeNb(cur.MyInfo.MyAmount) + " (" + baseVal + " $)", "",
			helpers.ResizeNb(cur.MyInfo.OtherAmount) + " (" + relVal + " $)", cur.MyInfo.OtherCoin,
			helpers.GetDateFromTimestamp(cur.MyInfo.StartedAt, true), cur.Uuid, cur.GetLastStatus()}
		data = append(data, out)
		baseTotalValue = helpers.BigFloatAdd(baseTotalValue, baseVal, 2)
		relTotalValue = helpers.BigFloatAdd(relTotalValue, relVal, 2)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"MyCoin", "MyAmount", "", "OtherAmount", "OtherCoin", "Date", "UUID", "Status"})
	table.SetFooter([]string{
		"Page: " + strconv.Itoa(answer.Result.PageNumber),
		"Total USD (base): " + baseTotalValue + " $",
		"Nb swaps (DB): " + strconv.Itoa(answer.Result.Total),
		"Total USD (rel): " + relTotalValue + " $",
		"Limit: " + strconv.Itoa(answer.Result.Limit),
		"Skipped: " + strconv.Itoa(answer.Result.Skipped), "Total pages: " + strconv.Itoa(answer.Result.TotalPages), ""})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}

func (answer *MyRecentSwapsAnswer) ToMap() map[string]SwapContent {
	out := make(map[string]SwapContent)
	for _, cur := range answer.Result.Swaps {
		out[cur.Uuid] = cur
	}
	return out
}
