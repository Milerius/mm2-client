package mm2_data_structure

import (
	"encoding/json"
	"fmt"
	"mm2_client/helpers"
)

type TradePreimageParams struct {
	Base       string  `json:"base"`
	Rel        string  `json:"rel"`
	Price      string  `json:"price"`
	Volume     *string `json:"volume,omitempty"`
	Max        *bool   `json:"max,omitempty"`
	SwapMethod string  `json:"swap_method"`
}

type TradePreimageRequest struct {
	Mmrpc    *string             `json:"mmrpc,omitempty"`
	Userpass string              `json:"userpass"`
	Method   string              `json:"method"`
	Params   TradePreimageParams `json:"params"`
}

type TBaseCoinFee struct {
	Amount         string `json:"amount"`
	AmountFraction struct {
		Denom string `json:"denom"`
		Numer string `json:"numer"`
	} `json:"amount_fraction"`
	AmountRat          [][]interface{} `json:"amount_rat"`
	Coin               string          `json:"coin"`
	PaidFromTradingVol bool            `json:"paid_from_trading_vol"`
}

type TRelCoinFee struct {
	Coin           string `json:"coin"`
	Amount         string `json:"amount"`
	AmountFraction struct {
		Numer string `json:"numer"`
		Denom string `json:"denom"`
	} `json:"amount_fraction"`
	AmountRat          [][]interface{} `json:"amount_rat"`
	PaidFromTradingVol bool            `json:"paid_from_trading_vol"`
}

type TotalFeeContent struct {
	Coin           string `json:"coin"`
	Amount         string `json:"amount"`
	AmountFraction struct {
		Numer string `json:"numer"`
		Denom string `json:"denom"`
	} `json:"amount_fraction"`
	AmountRat               [][]interface{} `json:"amount_rat"`
	RequiredBalance         string          `json:"required_balance"`
	RequiredBalanceFraction struct {
		Numer string `json:"numer"`
		Denom string `json:"denom"`
	} `json:"required_balance_fraction"`
	RequiredBalanceRat [][]interface{} `json:"required_balance_rat"`
}

type TTakerFee struct {
	Amount         string `json:"amount"`
	AmountFraction struct {
		Denom string `json:"denom"`
		Numer string `json:"numer"`
	} `json:"amount_fraction"`
	AmountRat          [][]interface{} `json:"amount_rat"`
	Coin               string          `json:"coin"`
	PaidFromTradingVol bool            `json:"paid_from_trading_vol"`
}

type TFeeToSendTakerFee struct {
	Amount         string `json:"amount"`
	AmountFraction struct {
		Denom string `json:"denom"`
		Numer string `json:"numer"`
	} `json:"amount_fraction"`
	AmountRat          [][]interface{} `json:"amount_rat"`
	Coin               string          `json:"coin"`
	PaidFromTradingVol bool            `json:"paid_from_trading_vol"`
}

type TradePreimageAnswerSuccess struct {
	BaseCoinFee       TBaseCoinFee        `json:"base_coin_fee"`
	RelCoinFee        TRelCoinFee         `json:"rel_coin_fee"`
	TakerFee          *TTakerFee          `json:"taker_fee"`
	FeeToSendTakerFee *TFeeToSendTakerFee `json:"fee_to_send_taker_fee"`
	TotalFees         []TotalFeeContent   `json:"total_fees"`
}

type TradePreimageAnswer struct {
	Mmrpc  string                      `json:"mmrpc"`
	Result *TradePreimageAnswerSuccess `json:"result,omitempty"`
	Id     int                         `json:"id"`
	Error  string                      `json:"error,omitempty"`
}

func (req *TradePreimageRequest) ToJson() string {
	b, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}

func NewTradePreimageRequest(base string, rel string, price string, method string, volumeOrMax string) *TradePreimageRequest {
	var max *bool = nil
	var volume *string = nil
	if volumeOrMax == "max" {
		max = helpers.BoolAddr(true)
	} else {
		volume = &volumeOrMax
	}
	genReq := NewGenericRequestV2("trade_preimage")
	params := TradePreimageParams{Base: base, Rel: rel, Price: price, SwapMethod: method}
	if max != nil {
		params.Max = max
	}
	if volume != nil {
		params.Volume = volume
	}
	req := &TradePreimageRequest{Userpass: genReq.Userpass, Method: genReq.Method, Params: params, Mmrpc: genReq.MMRpc}
	return req
}
