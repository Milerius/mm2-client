package http

import (
	"encoding/json"
	"fmt"
	"mm2_client/config"
	"mm2_client/constants"
)

type MM2GenericRequest struct {
	Method   string `json:"method"`
	Userpass string `json:"userpass"`
}

var gRuntimeUserpass = ""

func NewGenericRequest(method string) *MM2GenericRequest {
	if gRuntimeUserpass == "" {
		gRuntimeUserpass = config.NewMM2ConfigFromFile(constants.GMM2ConfPath).RPCPassword
	}
	return &MM2GenericRequest{Method: method, Userpass: gRuntimeUserpass}
}

func (req MM2GenericRequest) ToJson() string {
	b, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}
