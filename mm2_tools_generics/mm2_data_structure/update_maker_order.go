package mm2_data_structure

import (
	"encoding/json"
	"fmt"
)

type UpdateMakerOrderRequest struct {
	Userpass    string  `json:"userpass"`
	Method      string  `json:"method"`
	Uuid        string  `json:"uuid"`
	NewPrice    *string `json:"new_price"`
	VolumeDelta *string `json:"volume_delta,omitempty"`
	Max         *bool   `json:"max,omitempty"`
	MinVolume   *string `json:"min_volume,omitempty"`
	BaseConfs   *int    `json:"base_confs,omitempty"`
	BaseNota    *bool   `json:"base_nota,omitempty"`
	RelConfs    *int    `json:"rel_confs,omitempty"`
	RelNota     *bool   `json:"rel_nota,omitempty"`
}

type UpdateMakerOrderAnswer struct {
	Result struct {
		Base          string          `json:"base"`
		Rel           string          `json:"rel"`
		MaxBaseVol    string          `json:"max_base_vol"`
		MaxBaseVolRat [][]interface{} `json:"max_base_vol_rat"`
		MinBaseVol    interface{}     `json:"min_base_vol"`
		CreatedAt     int64           `json:"created_at"`
		UpdatedAt     int64           `json:"updated_at"`
		Matches       struct {
		} `json:"matches"`
		Price        string          `json:"price"`
		PriceRat     [][]interface{} `json:"price_rat"`
		StartedSwaps []interface{}   `json:"started_swaps"`
		Uuid         string          `json:"uuid"`
		ConfSettings struct {
			BaseConfs int  `json:"base_confs"`
			BaseNota  bool `json:"base_nota"`
			RelConfs  int  `json:"rel_confs"`
			RelNota   bool `json:"rel_nota"`
		} `json:"conf_settings"`
	} `json:"result"`
}

func NewUpdateMakerRequest(uuid string, newPrice *string, volumeDelta *string, max *bool, minVolume *string,
	baseConfs *int, baseNota *bool, relConfs *int, relNota *bool) *UpdateMakerOrderRequest {
	genReq := NewGenericRequest("update_maker_order" +
		"")
	req := &UpdateMakerOrderRequest{Userpass: genReq.Userpass, Method: genReq.Method, Uuid: uuid}
	if newPrice != nil {
		req.NewPrice = newPrice
	}
	if volumeDelta != nil {
		req.VolumeDelta = volumeDelta
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

func (req *UpdateMakerOrderRequest) ToJson() string {
	b, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}
