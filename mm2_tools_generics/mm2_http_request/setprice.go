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

func NewSetPriceRequest(base string, rel string, price string, volume *string, max *bool, cancelPrevious bool, minVolume *string,
	baseConfs *int, baseNota *bool, relConfs *int, relNota *bool) *mm2_data_structure.SetPriceRequest {
	genReq := mm2_data_structure.NewGenericRequest("setprice")
	req := &mm2_data_structure.SetPriceRequest{Userpass: genReq.Userpass, Method: genReq.Method, Base: base, Rel: rel, Price: price, CancelPrevious: cancelPrevious}
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

func SetPrice(base string, rel string, price string, volume *string, max *bool, cancelPrevious bool, minVolume *string,
	baseConfs *int, baseNota *bool, relConfs *int, relNota *bool) (*mm2_data_structure.SetPriceAnswer, error) {
	req := NewSetPriceRequest(base, rel, price, volume, max, cancelPrevious, minVolume, baseConfs, baseNota, relConfs, relNota).ToJson()
	resp, err := http.Post(mm2_data_structure.GMM2Endpoint, "application/json", bytes.NewBuffer([]byte(req)))
	if err != nil {
		glg.Errorf("Err: %v", err)
		return nil, err
	}
	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		var answer = &mm2_data_structure.SetPriceAnswer{}
		decodeErr := json.NewDecoder(resp.Body).Decode(answer)
		if decodeErr != nil {
			glg.Errorf("Err: %v", err)
			return nil, decodeErr
		}
		return answer, nil
	} else {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		errStr := fmt.Sprintf("Err: %s\n", bodyBytes)
		return nil, errors.New(errStr)
	}
}
