package mm2_http_request

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kpango/glg"
	"io/ioutil"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"net/http"
)

func NewUpdateMakerRequest(uuid string, newPrice *string, volumeDelta *string, max *bool, minVolume *string,
	baseConfs *int, baseNota *bool, relConfs *int, relNota *bool) *mm2_data_structure.UpdateMakerOrderRequest {
	genReq := mm2_data_structure.NewGenericRequest("update_maker_order" +
		"")
	req := &mm2_data_structure.UpdateMakerOrderRequest{Userpass: genReq.Userpass, Method: genReq.Method, Uuid: uuid}
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

func UpdateMakerOrder(uuid string, newPrice *string, volumeDelta *string, max *bool, minVolume *string,
	baseConfs *int, baseNota *bool, relConfs *int, relNota *bool) (*mm2_data_structure.UpdateMakerOrderAnswer, error) {
	req := NewUpdateMakerRequest(uuid, newPrice, volumeDelta, max, minVolume, baseConfs, baseNota, relConfs, relNota).ToJson()
	resp, err := http.Post(mm2_data_structure.GMM2Endpoint, "application/json", bytes.NewBuffer([]byte(req)))
	if err != nil {
		glg.Errorf("Err: %v", err)
		return nil, err
	}
	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		var answer = &mm2_data_structure.UpdateMakerOrderAnswer{}
		decodeErr := json.NewDecoder(resp.Body).Decode(answer)
		if decodeErr != nil {
			glg.Errorf("Err: %v", decodeErr)
			return nil, decodeErr
		}
		return answer, nil
	} else {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		errStr := fmt.Sprintf("Err: %s\n", bodyBytes)
		return nil, errors.New(errStr)
	}
}
