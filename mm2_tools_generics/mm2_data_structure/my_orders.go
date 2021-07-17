package mm2_data_structure

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"mm2_client/helpers"
	"os"
)

type MakerMatchesContent struct {
	Connect struct {
		DestPubKey     string `json:"dest_pub_key"`
		MakerOrderUuid string `json:"maker_order_uuid"`
		Method         string `json:"method"`
		SenderPubkey   string `json:"sender_pubkey"`
		TakerOrderUuid string `json:"taker_order_uuid"`
	} `json:"connect"`
	Connected struct {
		DestPubKey     string `json:"dest_pub_key"`
		MakerOrderUuid string `json:"maker_order_uuid"`
		Method         string `json:"method"`
		SenderPubkey   string `json:"sender_pubkey"`
		TakerOrderUuid string `json:"taker_order_uuid"`
	} `json:"connected"`
	LastUpdated int64 `json:"last_updated"`
	Request     struct {
		Action       string `json:"action"`
		Base         string `json:"base"`
		BaseAmount   string `json:"base_amount"`
		DestPubKey   string `json:"dest_pub_key"`
		Method       string `json:"method"`
		Rel          string `json:"rel"`
		RelAmount    string `json:"rel_amount"`
		SenderPubkey string `json:"sender_pubkey"`
		Uuid         string `json:"uuid"`
	} `json:"request"`
	Reserved struct {
		Base           string `json:"base"`
		BaseAmount     string `json:"base_amount"`
		DestPubKey     string `json:"dest_pub_key"`
		MakerOrderUuid string `json:"maker_order_uuid"`
		Method         string `json:"method"`
		Rel            string `json:"rel"`
		RelAmount      string `json:"rel_amount"`
		SenderPubkey   string `json:"sender_pubkey"`
		TakerOrderUuid string `json:"taker_order_uuid"`
	} `json:"reserved"`
}

type MakerOrderContent struct {
	AvailableAmount string                         `json:"available_amount"`
	Base            string                         `json:"base"`
	Cancellable     bool                           `json:"cancellable"`
	CreatedAt       int64                          `json:"created_at"`
	Matches         map[string]MakerMatchesContent `json:"matches"`
	MaxBaseVol      string                         `json:"max_base_vol"`
	MaxBaseVolRat   [][]interface{}                `json:"max_base_vol_rat"`
	MinBaseVol      string                         `json:"min_base_vol"`
	MinBaseVolRat   [][]interface{}                `json:"min_base_vol_rat"`
	Price           string                         `json:"price"`
	PriceRat        [][]interface{}                `json:"price_rat"`
	Rel             string                         `json:"rel"`
	StartedSwaps    []string                       `json:"started_swaps"`
	Uuid            string                         `json:"uuid"`
}

type TakerMatchesContent struct {
	Connect struct {
		DestPubKey     string `json:"dest_pub_key"`
		MakerOrderUuid string `json:"maker_order_uuid"`
		Method         string `json:"method"`
		SenderPubkey   string `json:"sender_pubkey"`
		TakerOrderUuid string `json:"taker_order_uuid"`
	} `json:"connect"`
	Connected struct {
		DestPubKey     string `json:"dest_pub_key"`
		MakerOrderUuid string `json:"maker_order_uuid"`
		Method         string `json:"method"`
		SenderPubkey   string `json:"sender_pubkey"`
		TakerOrderUuid string `json:"taker_order_uuid"`
	} `json:"connected"`
	LastUpdated int64 `json:"last_updated"`
	Reserved    struct {
		Base           string `json:"base"`
		BaseAmount     string `json:"base_amount"`
		DestPubKey     string `json:"dest_pub_key"`
		MakerOrderUuid string `json:"maker_order_uuid"`
		Method         string `json:"method"`
		Rel            string `json:"rel"`
		RelAmount      string `json:"rel_amount"`
		SenderPubkey   string `json:"sender_pubkey"`
		TakerOrderUuid string `json:"taker_order_uuid"`
	} `json:"reserved"`
}

type TakerOrderContent struct {
	Cancellable bool                           `json:"cancellable"`
	CreatedAt   int64                          `json:"created_at"`
	Matches     map[string]TakerMatchesContent `json:"matches"`
	Request     struct {
		Action        string          `json:"action"`
		Base          string          `json:"base"`
		BaseAmount    string          `json:"base_amount"`
		BaseAmountRat [][]interface{} `json:"base_amount_rat"`
		DestPubKey    string          `json:"dest_pub_key"`
		Method        string          `json:"method"`
		Rel           string          `json:"rel"`
		RelAmount     string          `json:"rel_amount"`
		RelAmountRat  [][]interface{} `json:"rel_amount_rat"`
		SenderPubkey  string          `json:"sender_pubkey"`
		Uuid          string          `json:"uuid"`
		MatchBy       struct {
			Type string `json:"type"`
		} `json:"match_by"`
	} `json:"request"`
	OrderType struct {
		Type string `json:"type"`
	} `json:"order_type"`
}

type MyOrdersAnswer struct {
	Result struct {
		MakerOrders map[string]MakerOrderContent `json:"maker_orders,omitempty"`
		TakerOrders map[string]TakerOrderContent `json:"taker_orders,omitempty"`
	} `json:"result"`
}

func renderTableTakerOrders(takerOrders map[string]TakerOrderContent) {
	var data [][]string

	for _, cur := range takerOrders {
		var out = []string{cur.Request.Base, cur.Request.BaseAmount, "", cur.Request.RelAmount, cur.Request.Rel, cur.Request.Uuid}
		data = append(data, out)
	}

	if len(data) > 0 {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoWrapText(false)
		table.SetHeader([]string{"Base", "Base Amount", "", "Rel Amount", "Rel", "UUID"})
		table.SetFooter([]string{"", "", "", "", "", "TakerOrders"})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		table.SetCenterSeparator("|")
		table.AppendBulk(data) // Add Bulk Data
		table.Render()
		fmt.Println()
	} else {
		//goland:noinspection GoPrintFunctions
		fmt.Println("No taker orders\n")
	}
}

func renderTableMakerOrders(makerOrders map[string]MakerOrderContent) {
	var data [][]string

	for _, cur := range makerOrders {
		var out = []string{cur.Base, cur.MinBaseVol, cur.AvailableAmount, "", helpers.BigFloatMultiply(cur.AvailableAmount, cur.Price, 8), cur.Rel, cur.Price, cur.Uuid}
		data = append(data, out)
	}

	if len(data) > 0 {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoWrapText(false)
		table.SetHeader([]string{"Base", "Base MinVol", "Base Amount", "", "Rel Amount", "Rel", "Price", "UUID"})
		table.SetFooter([]string{"", "", "", "", "", "", "", "MakerOrders"})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		table.SetCenterSeparator("|")
		table.AppendBulk(data) // Add Bulk Data
		table.Render()
	} else {
		fmt.Println("No maker orders")
	}
}

func (answer *MyOrdersAnswer) ToTable() {
	renderTableTakerOrders(answer.Result.TakerOrders)
	renderTableMakerOrders(answer.Result.MakerOrders)
}
