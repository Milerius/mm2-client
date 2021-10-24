package mm2_data_structure

import (
	"encoding/json"
	"fmt"
	"mm2_client/config"
)

type StartSimpleMarketMakerRequest struct {
	Mmrpc    *string                              `json:"mmrpc,omitempty"`
	Userpass string                               `json:"userpass"`
	Method   string                               `json:"method"`
	Params   *config.StartSimpleMarketMakerParams `json:"params"`
}

type StartSimpleMarketMakerAnswerSuccess struct {
	Result string `json:"result"`
}

type StartSimpleMarketMakerAnswer struct {
	Mmrpc  string                               `json:"mmrpc"`
	Result *StartSimpleMarketMakerAnswerSuccess `json:"result,omitempty"`
	Id     int                                  `json:"id"`
	Error  string                               `json:"error,omitempty"`
}

func (req *StartSimpleMarketMakerRequest) ToJson() string {
	b, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}

func NewStartSimpleMarketMakerRequest() *StartSimpleMarketMakerRequest {
	genReq := NewGenericRequestV2("start_simple_market_maker_bot")
	out := &StartSimpleMarketMakerRequest{Mmrpc: genReq.MMRpc, Userpass: genReq.Userpass, Method: genReq.Method, Params: config.GSimpleMarketMakerConf}
	return out
}
