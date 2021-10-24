package mm2_data_structure

import (
	"encoding/json"
	"fmt"
)

type StopSimpleMarketMakerRequest struct {
	Mmrpc    *string `json:"mmrpc,omitempty"`
	Userpass string  `json:"userpass"`
	Method   string  `json:"method"`
	Params   *string `json:"params"`
}

type StopSimpleMarketMakerAnswerSuccess struct {
	Result string `json:"result"`
}

type StopSimpleMarketMakerAnswer struct {
	Mmrpc  string                              `json:"mmrpc"`
	Result *StopSimpleMarketMakerAnswerSuccess `json:"result,omitempty"`
	Id     int                                 `json:"id"`
	Error  string                              `json:"error,omitempty"`
}

func (req *StopSimpleMarketMakerRequest) ToJson() string {
	b, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}

func NewStopSimpleMarketMakerRequest() *StopSimpleMarketMakerRequest {
	genReq := NewGenericRequestV2("stop_simple_market_maker_bot")
	out := &StopSimpleMarketMakerRequest{Mmrpc: genReq.MMRpc, Userpass: genReq.Userpass, Method: genReq.Method, Params: nil}
	return out
}
