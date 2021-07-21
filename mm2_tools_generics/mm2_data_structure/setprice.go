package mm2_data_structure

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"mm2_client/helpers"
	"os"
	"strconv"
)

type SetPriceRequest struct {
	Userpass       string  `json:"userpass"`
	Method         string  `json:"method"`
	Base           string  `json:"base"`
	Rel            string  `json:"rel"`
	Price          string  `json:"price"`
	CancelPrevious bool    `json:"cancel_previous"`
	Volume         *string `json:"volume,omitempty"`
	Max            *bool   `json:"max,omitempty"`
	MinVolume      *string `json:"min_volume,omitempty"`
	BaseConfs      *int    `json:"base_confs,omitempty"`
	BaseNota       *bool   `json:"base_nota,omitempty"`
	RelConfs       *int    `json:"rel_confs,omitempty"`
	RelNota        *bool   `json:"rel_nota,omitempty"`
}

type SetPriceAnswer struct {
	Result struct {
		Base          string          `json:"base"`
		Rel           string          `json:"rel"`
		MaxBaseVol    string          `json:"max_base_vol"`
		MaxBaseVolRat [][]interface{} `json:"max_base_vol_rat"`
		MinBaseVol    string          `json:"min_base_vol"`
		CreatedAt     int64           `json:"created_at"`
		Matches       struct {
		} `json:"matches"`
		Price        string          `json:"price"`
		PriceRat     [][]interface{} `json:"price_rat"`
		StartedSwaps []string        `json:"started_swaps"`
		Uuid         string          `json:"uuid"`
		ConfSettings struct {
			BaseConfs int  `json:"base_confs"`
			BaseNota  bool `json:"base_nota"`
			RelConfs  int  `json:"rel_confs"`
			RelNota   bool `json:"rel_nota"`
		} `json:"conf_settings"`
	} `json:"result"`
}

func NewSetPriceRequest(base string, rel string, price string, volume *string, max *bool, cancelPrevious bool, minVolume *string,
	baseConfs *int, baseNota *bool, relConfs *int, relNota *bool) *SetPriceRequest {
	genReq := NewGenericRequest("setprice")
	req := &SetPriceRequest{Userpass: genReq.Userpass, Method: genReq.Method, Base: base, Rel: rel, Price: price, CancelPrevious: cancelPrevious}
	if volume != nil {
		req.Volume = volume
	}
	if max != nil {
		req.Max = max
	}
	if minVolume != nil {
		req.MinVolume = minVolume
	}
	if baseConfs != nil {
		req.BaseConfs = baseConfs
	}
	if baseNota != nil {
		req.BaseNota = baseNota
	}
	if relConfs != nil {
		req.RelConfs = relConfs
	}
	if relNota != nil {
		req.RelNota = relNota
	}
	return req
}

func (req *SetPriceRequest) ToJson() string {
	b, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}

func (answer *SetPriceAnswer) ToTable() {

	baseNota := strconv.FormatBool(answer.Result.ConfSettings.BaseNota)
	relNota := strconv.FormatBool(answer.Result.ConfSettings.RelNota)
	data := [][]string{

		{answer.Result.Base, answer.Result.MinBaseVol, answer.Result.MaxBaseVol,
			answer.Result.Price, strconv.Itoa(answer.Result.ConfSettings.BaseConfs), baseNota, "", answer.Result.Rel,
			helpers.BigFloatMultiply(answer.Result.MaxBaseVol, answer.Result.Price, 8),
			strconv.Itoa(answer.Result.ConfSettings.RelConfs), relNota, answer.Result.Uuid},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"Base", "Base Min Vol", "Base Amount", "Base Price", "Base Confs", "Base Nota", " ", "Rel", "Rel Amount", "Rel confs", "Rel nota", "UUID"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}
