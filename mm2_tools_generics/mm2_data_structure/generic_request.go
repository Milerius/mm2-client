package mm2_data_structure

import (
	"encoding/json"
	"fmt"
	"mm2_client/config"
	"mm2_client/constants"
)

type MM2GenericRequest struct {
	Method   string  `json:"method"`
	Userpass string  `json:"userpass"`
	MMRpc    *string `json:"mmrpc,omitempty"`
}

var GRuntimeUserpass = ""

const GMM2Endpoint = "http://127.0.0.1:7783"

func NewGenericRequest(method string) *MM2GenericRequest {
	if GRuntimeUserpass == "" {
		GRuntimeUserpass = config.NewMM2ConfigFromFile(constants.GMM2ConfPath).RPCPassword
	}
	return &MM2GenericRequest{Method: method, Userpass: GRuntimeUserpass}
}

func NewGenericRequestV2(method string) *MM2GenericRequest {
	resp := NewGenericRequest(method)
	mmrpc := "2.0"
	resp.MMRpc = &mmrpc
	return resp
}

func (req MM2GenericRequest) ToJson() string {
	b, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}
