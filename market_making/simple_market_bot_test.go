package market_making

import (
	"github.com/stretchr/testify/suite"
	"mm2_client/helpers"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"testing"
)

type Content struct {
	BaseAmount string
	RelAmount  string
}

func NewSwapContentInfos(values []Content) *mm2_data_structure.MyRecentSwapsAnswer {
	resp := &mm2_data_structure.MyRecentSwapsAnswer{}
	respSuccess := mm2_data_structure.MyRecentSwapsAnswerSuccess{}
	var swaps []mm2_data_structure.SwapContent
	for _, cur := range values {
		swaps = append(swaps, mm2_data_structure.SwapContent{MyInfo: mm2_data_structure.MyInfoContents{MyAmount: cur.BaseAmount, OtherAmount: cur.RelAmount}})
	}
	respSuccess.Swaps = swaps
	resp.Result = respSuccess
	return resp
}

type CalculateThreshHoldFromLastTradesTestSuite struct {
	suite.Suite
}

// This is an example test that will always succeed
func (suite *CalculateThreshHoldFromLastTradesTestSuite) TestSingleTradeReversedSide() {
	suite.Equal(true, true)
	cfg := SimplePairMarketMakerConf{Base: "FIRO", Rel: "KMD", Spread: "1.015", Max: true, MinVolume: nil, Enable: true,
		PriceElapsedValidity: nil, CheckLastBidirectionalTradeThreshHold: helpers.BoolAddr(true)}
	baseResp := &mm2_data_structure.MyRecentSwapsAnswer{}
	relResp := NewSwapContentInfos([]Content{{BaseAmount: "219.4709", RelAmount: "29.99999"}})

	expectedPrice := calculateThreshHoldFromLastTrades(cfg, "7.14455729", baseResp, relResp, "7.14455729")
	suite.Equal(expectedPrice, "7.42543460")
}

func (suite *CalculateThreshHoldFromLastTradesTestSuite) TestSingleTradeReversedSideForcePrice() {
	suite.Equal(true, true)
	cfg := SimplePairMarketMakerConf{Base: "FIRO", Rel: "KMD", Spread: "1.015", Max: true, MinVolume: nil, Enable: true,
		PriceElapsedValidity: nil, CheckLastBidirectionalTradeThreshHold: helpers.BoolAddr(true)}
	baseResp := &mm2_data_structure.MyRecentSwapsAnswer{}
	relResp := NewSwapContentInfos([]Content{{BaseAmount: "219.4709", RelAmount: "29.99999"}})

	expectedPrice := calculateThreshHoldFromLastTrades(cfg, "7.6", baseResp, relResp, "7.6")
	suite.Equal(expectedPrice, "7.6")
}

func (suite *CalculateThreshHoldFromLastTradesTestSuite) TestMultipleTradeReversedSide() {
	suite.Equal(true, true)
	cfg := SimplePairMarketMakerConf{Base: "FIRO", Rel: "KMD", Spread: "1.015", Max: true, MinVolume: nil, Enable: true,
		PriceElapsedValidity: nil, CheckLastBidirectionalTradeThreshHold: helpers.BoolAddr(true)}
	baseResp := &mm2_data_structure.MyRecentSwapsAnswer{}
	relResp := NewSwapContentInfos([]Content{{BaseAmount: "219.4709", RelAmount: "29.99999"}, {BaseAmount: "222.762", RelAmount: "29.99999"}})

	expectedPrice := calculateThreshHoldFromLastTrades(cfg, "7.14455729", baseResp, relResp, "7.14455729")
	suite.Equal(expectedPrice, "7.48110906")
}

func (suite *CalculateThreshHoldFromLastTradesTestSuite) TestMultipleTradeReversedSideForcePrice() {
	suite.Equal(true, true)
	cfg := SimplePairMarketMakerConf{Base: "FIRO", Rel: "KMD", Spread: "1.015", Max: true, MinVolume: nil, Enable: true,
		PriceElapsedValidity: nil, CheckLastBidirectionalTradeThreshHold: helpers.BoolAddr(true)}
	baseResp := &mm2_data_structure.MyRecentSwapsAnswer{}
	relResp := NewSwapContentInfos([]Content{{BaseAmount: "219.4709", RelAmount: "29.99999"}, {BaseAmount: "222.762", RelAmount: "29.99999"}})

	expectedPrice := calculateThreshHoldFromLastTrades(cfg, "7.54455729", baseResp, relResp, "7.54455729")
	suite.Equal(expectedPrice, "7.54455729")
}

func (suite *CalculateThreshHoldFromLastTradesTestSuite) TestZeroTradeReversedSide() {
	suite.Equal(true, true)
	cfg := SimplePairMarketMakerConf{Base: "FIRO", Rel: "KMD", Spread: "1.015", Max: true, MinVolume: nil, Enable: true,
		PriceElapsedValidity: nil, CheckLastBidirectionalTradeThreshHold: helpers.BoolAddr(true)}
	baseResp := &mm2_data_structure.MyRecentSwapsAnswer{}
	relResp := &mm2_data_structure.MyRecentSwapsAnswer{}

	expectedPrice := calculateThreshHoldFromLastTrades(cfg, "7.14455729", baseResp, relResp, "7.14455729")
	suite.Equal(expectedPrice, "7.14455729")
}

func (suite *CalculateThreshHoldFromLastTradesTestSuite) TestSingleTradeBaseSide() {
	suite.Equal(true, true)
	cfg := SimplePairMarketMakerConf{Base: "FIRO", Rel: "KMD", Spread: "1.015", Max: true, MinVolume: nil, Enable: true,
		PriceElapsedValidity: nil, CheckLastBidirectionalTradeThreshHold: helpers.BoolAddr(true)}
	baseResp := NewSwapContentInfos([]Content{{BaseAmount: "29.99997438", RelAmount: "222.76277576"}})
	relResp := &mm2_data_structure.MyRecentSwapsAnswer{}

	expectedPrice := calculateThreshHoldFromLastTrades(cfg, "7.14455729", baseResp, relResp, "7.14455729")
	suite.Equal(expectedPrice, "7.53681368")
}

func (suite *CalculateThreshHoldFromLastTradesTestSuite) TestSingleTradeBaseSideForPrice() {
	suite.Equal(true, true)
	cfg := SimplePairMarketMakerConf{Base: "FIRO", Rel: "KMD", Spread: "1.015", Max: true, MinVolume: nil, Enable: true,
		PriceElapsedValidity: nil, CheckLastBidirectionalTradeThreshHold: helpers.BoolAddr(true)}
	baseResp := NewSwapContentInfos([]Content{{BaseAmount: "29.99997438", RelAmount: "222.76277576"}})
	relResp := &mm2_data_structure.MyRecentSwapsAnswer{}

	expectedPrice := calculateThreshHoldFromLastTrades(cfg, "7.6", baseResp, relResp, "7.6")
	suite.Equal(expectedPrice, "7.6")
}

func TestCalculateThreshHoldFromLastTradesTestSuite(t *testing.T) {
	suite.Run(t, new(CalculateThreshHoldFromLastTradesTestSuite))
}
