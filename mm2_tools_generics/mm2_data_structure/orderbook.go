package mm2_data_structure

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"mm2_client/helpers"
	"os"
	"runtime"
	"strconv"
)

type OrderbookRequest struct {
	Userpass string `json:"userpass"`
	Method   string `json:"method"`
	Base     string `json:"base"`
	Rel      string `json:"rel"`
}

type OrderbookContent struct {
	Coin          string          `json:"coin"`
	Address       string          `json:"address"`
	Price         string          `json:"price"`
	PriceRat      [][]interface{} `json:"price_rat"`
	PriceFraction struct {
		Numer string `json:"numer"`
		Denom string `json:"denom"`
	} `json:"price_fraction"`
	Maxvolume         string          `json:"maxvolume"`
	MaxVolumeRat      [][]interface{} `json:"max_volume_rat"`
	MaxVolumeFraction struct {
		Numer string `json:"numer"`
		Denom string `json:"denom"`
	} `json:"max_volume_fraction"`
	MinVolume         string          `json:"min_volume"`
	MinVolumeRat      [][]interface{} `json:"min_volume_rat"`
	MinVolumeFraction struct {
		Numer string `json:"numer"`
		Denom string `json:"denom"`
	} `json:"min_volume_fraction"`
	Pubkey                string `json:"pubkey"`
	Age                   int    `json:"age"`
	Zcredits              int    `json:"zcredits"`
	Uuid                  string `json:"uuid"`
	IsMine                bool   `json:"is_mine"`
	BaseMaxVolume         string `json:"base_max_volume"`
	BaseMaxVolumeFraction struct {
		Numer string `json:"numer"`
		Denom string `json:"denom"`
	} `json:"base_max_volume_fraction"`
	BaseMaxVolumeRat      [][]interface{} `json:"base_max_volume_rat"`
	BaseMinVolume         string          `json:"base_min_volume"`
	BaseMinVolumeFraction struct {
		Numer string `json:"numer"`
		Denom string `json:"denom"`
	} `json:"base_min_volume_fraction"`
	BaseMinVolumeRat     [][]interface{} `json:"base_min_volume_rat"`
	RelMaxVolume         string          `json:"rel_max_volume"`
	RelMaxVolumeFraction struct {
		Numer string `json:"numer"`
		Denom string `json:"denom"`
	} `json:"rel_max_volume_fraction"`
	RelMaxVolumeRat      [][]interface{} `json:"rel_max_volume_rat"`
	RelMinVolume         string          `json:"rel_min_volume"`
	RelMinVolumeFraction struct {
		Numer string `json:"numer"`
		Denom string `json:"denom"`
	} `json:"rel_min_volume_fraction"`
	RelMinVolumeRat           [][]interface{} `json:"rel_min_volume_rat"`
	BaseMaxVolumeAggr         string          `json:"base_max_volume_aggr"`
	BaseMaxVolumeAggrFraction struct {
		Numer string `json:"numer"`
		Denom string `json:"denom"`
	} `json:"base_max_volume_aggr_fraction"`
	BaseMaxVolumeAggrRat     [][]interface{} `json:"base_max_volume_aggr_rat"`
	RelMaxVolumeAggr         string          `json:"rel_max_volume_aggr"`
	RelMaxVolumeAggrFraction struct {
		Numer string `json:"numer"`
		Denom string `json:"denom"`
	} `json:"rel_max_volume_aggr_fraction"`
	RelMaxVolumeAggrRat [][]interface{} `json:"rel_max_volume_aggr_rat"`
}

type OrderbookAnswer struct {
	Askdepth                 int                `json:"askdepth"`
	Asks                     []OrderbookContent `json:"asks"`
	Base                     string             `json:"base"`
	Biddepth                 int                `json:"biddepth"`
	Bids                     []OrderbookContent `json:"bids"`
	Netid                    int                `json:"netid"`
	Numasks                  int                `json:"numasks"`
	Numbids                  int                `json:"numbids"`
	Rel                      string             `json:"rel"`
	Timestamp                int                `json:"timestamp"`
	TotalAsksBaseVol         string             `json:"total_asks_base_vol"`
	TotalAsksBaseVolFraction struct {
		Numer string `json:"numer"`
		Denom string `json:"denom"`
	} `json:"total_asks_base_vol_fraction"`
	TotalAsksBaseVolRat     [][]interface{} `json:"total_asks_base_vol_rat"`
	TotalAsksRelVol         string          `json:"total_asks_rel_vol"`
	TotalAsksRelVolFraction struct {
		Numer string `json:"numer"`
		Denom string `json:"denom"`
	} `json:"total_asks_rel_vol_fraction"`
	TotalAsksRelVolRat       [][]interface{} `json:"total_asks_rel_vol_rat"`
	TotalBidsBaseVol         string          `json:"total_bids_base_vol"`
	TotalBidsBaseVolFraction struct {
		Numer string `json:"numer"`
		Denom string `json:"denom"`
	} `json:"total_bids_base_vol_fraction"`
	TotalBidsBaseVolRat     [][]interface{} `json:"total_bids_base_vol_rat"`
	TotalBidsRelVol         string          `json:"total_bids_rel_vol"`
	TotalBidsRelVolFraction struct {
		Numer string `json:"numer"`
		Denom string `json:"denom"`
	} `json:"total_bids_rel_vol_fraction"`
	TotalBidsRelVolRat [][]interface{} `json:"total_bids_rel_vol_rat"`
}

func NewOrderbookRequest(base string, rel string) *OrderbookRequest {
	genReq := NewGenericRequest("orderbook")
	req := &OrderbookRequest{Userpass: genReq.Userpass, Method: genReq.Method, Base: base, Rel: rel}
	return req
}

func (req *OrderbookRequest) ToJson() string {
	b, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}

func renderTable(contents []OrderbookContent, base string, rel string, depth int, size int, isAsks bool) {
	var data [][]string

	for _, cur := range contents {
		var out = []string{helpers.ResizeNb(cur.Price), helpers.ResizeNb(cur.BaseMaxVolume), helpers.ResizeNb(cur.RelMaxVolume), helpers.TransformBool(cur.IsMine)}
		data = append(data, out)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"Price (" + rel + ")", "Qty (" + base + ")", "Total (" + rel + ")", "Is Mine"})
	if runtime.GOARCH != "wasm" {
		if isAsks {
			table.SetHeaderColor(
				tablewriter.Colors{tablewriter.FgHiRedColor, tablewriter.Bold, tablewriter.BgBlackColor},
				tablewriter.Colors{tablewriter.FgHiRedColor, tablewriter.Bold, tablewriter.BgBlackColor},
				tablewriter.Colors{tablewriter.FgHiRedColor, tablewriter.Bold, tablewriter.BgBlackColor},
				tablewriter.Colors{tablewriter.FgHiRedColor, tablewriter.Bold, tablewriter.BgBlackColor})
		} else {
			table.SetHeaderColor(
				tablewriter.Colors{tablewriter.FgHiGreenColor, tablewriter.Bold, tablewriter.BgBlackColor},
				tablewriter.Colors{tablewriter.FgHiGreenColor, tablewriter.Bold, tablewriter.BgBlackColor},
				tablewriter.Colors{tablewriter.FgHiGreenColor, tablewriter.Bold, tablewriter.BgBlackColor},
				tablewriter.Colors{tablewriter.FgHiGreenColor, tablewriter.Bold, tablewriter.BgBlackColor})
		}
	}
	table.SetFooter([]string{"", "Depth: " + strconv.Itoa(depth), "NbOrders: " + strconv.Itoa(size), ""})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}

func (answer *OrderbookAnswer) ToTable(base string, rel string) {
	renderTable(answer.Asks, base, rel, answer.Askdepth, answer.Numasks, true)
	fmt.Println()
	renderTable(answer.Bids, base, rel, answer.Biddepth, answer.Numbids, false)
}
