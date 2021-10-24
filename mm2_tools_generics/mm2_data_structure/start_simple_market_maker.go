package mm2_data_structure

import (
	"mm2_client/config"
)

type StartSimpleMarketMakerRequest struct {
	Mmrpc    *string                              `json:"mmrpc,omitempty"`
	Userpass string                               `json:"userpass"`
	Method   string                               `json:"method"`
	Params   *config.StartSimpleMarketMakerParams `json:"params"`
}

func NewStartSimpleMarketMakerRequest() *StartSimpleMarketMakerRequest {
	genReq := NewGenericRequestV2("start_simple_market_maker_bot")
	out := &StartSimpleMarketMakerRequest{Mmrpc: genReq.MMRpc, Userpass: genReq.Userpass, Params: config.GSimpleMarketMakerConf}
	return out
}
