package market_making

import (
	"github.com/stretchr/testify/suite"
	"mm2_client/helpers"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"testing"
	"time"
)

type Content struct {
	BaseAmount string
	RelAmount  string
}

func NewSwapContentInfos(values []Content) *mm2_data_structure.MyRecentSwapsAnswer {
	resp := &mm2_data_structure.MyRecentSwapsAnswer{}
	respSuccess := mm2_data_structure.MyRecentSwapsAnswerSuccess{}
	var swaps []mm2_data_structure.SwapContent
	var events []mm2_data_structure.SwapEventContent
	eventData := mm2_data_structure.SwapEvent{Type: "Finished"}
	curEvent := mm2_data_structure.SwapEventContent{Timestamp: time.Now().Unix(), Event: eventData}
	events = append(events, curEvent)
	for _, cur := range values {
		swaps = append(swaps, mm2_data_structure.SwapContent{MyInfo: mm2_data_structure.MyInfoContents{MyAmount: cur.BaseAmount, OtherAmount: cur.RelAmount}, Events: events})
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
	suite.Equal(expectedPrice, "7.31569911")
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
	relResp := NewSwapContentInfos([]Content{{BaseAmount: "219.4709", RelAmount: "19.99999"}, {BaseAmount: "222.762", RelAmount: "29.99999"}})

	expectedPrice := calculateThreshHoldFromLastTrades(cfg, "7.14455729", baseResp, relResp, "7.14455729")
	suite.Equal(expectedPrice, "8.84466154")
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
	suite.Equal(expectedPrice, "7.42543220")
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

func (suite *CalculateThreshHoldFromLastTradesTestSuite) TestMultipleTradeBaseSide() {
	suite.Equal(true, true)
	cfg := SimplePairMarketMakerConf{Base: "FIRO", Rel: "KMD", Spread: "1.015", Max: true, MinVolume: nil, Enable: true,
		PriceElapsedValidity: nil, CheckLastBidirectionalTradeThreshHold: helpers.BoolAddr(true)}
	baseResp := NewSwapContentInfos([]Content{{BaseAmount: "29.99997438", RelAmount: "222.76277576"}, {BaseAmount: "14.99998719", RelAmount: "105.38138788"}})
	relResp := &mm2_data_structure.MyRecentSwapsAnswer{}
	expectedPrice := calculateThreshHoldFromLastTrades(cfg, "7.14455729", baseResp, relResp, "7.14455729")
	suite.Equal(expectedPrice, "7.29209875")
}

func (suite *CalculateThreshHoldFromLastTradesTestSuite) TestMultipleTradeBaseSideForcePrice() {
	suite.Equal(true, true)
	cfg := SimplePairMarketMakerConf{Base: "FIRO", Rel: "KMD", Spread: "1.015", Max: true, MinVolume: nil, Enable: true,
		PriceElapsedValidity: nil, CheckLastBidirectionalTradeThreshHold: helpers.BoolAddr(true)}
	baseResp := NewSwapContentInfos([]Content{{BaseAmount: "29.99997438", RelAmount: "222.76277576"}, {BaseAmount: "29.99997438", RelAmount: "190.76277576"}})
	relResp := &mm2_data_structure.MyRecentSwapsAnswer{}
	expectedPrice := calculateThreshHoldFromLastTrades(cfg, "7.14455729", baseResp, relResp, "7.14455729")
	suite.Equal(expectedPrice, "7.14455729")
}

func (suite *CalculateThreshHoldFromLastTradesTestSuite) TestMultipleTradeBothSide() {
	suite.Equal(true, true)
	cfg := SimplePairMarketMakerConf{Base: "LTC", Rel: "KMD", Spread: "1.013", Max: true, MinVolume: nil, Enable: true,
		PriceElapsedValidity: nil, CheckLastBidirectionalTradeThreshHold: helpers.BoolAddr(true)}
	baseResp := NewSwapContentInfos([]Content{{BaseAmount: "4.102174", RelAmount: "752.8892"}, {BaseAmount: "1.719676", RelAmount: "316.4945"}, {BaseAmount: "0.052971", RelAmount: "10.24148"}, {BaseAmount: "0.133952", RelAmount: "25.98481"}})
	relResp := NewSwapContentInfos([]Content{
		{BaseAmount: "27.18649", RelAmount: "0.161065"},
		{BaseAmount: "13.17659", RelAmount: "0.077728"},
		{BaseAmount: "37.00451", RelAmount: "0.208161"},
		{BaseAmount: "743.0625", RelAmount: "4.266546"},
		{BaseAmount: "2819.370", RelAmount: "15.68015"},
		{BaseAmount: "2633.345", RelAmount: "14.61151"},
		{BaseAmount: "559.7940", RelAmount: "3.006166"}})
	expectedPrice := calculateThreshHoldFromLastTrades(cfg, "174.1375", baseResp, relResp, "174.1375")
	suite.Equal(expectedPrice, "181.87995125")
}

func TestCalculateThreshHoldFromLastTradesTestSuite(t *testing.T) {
	suite.Run(t, new(CalculateThreshHoldFromLastTradesTestSuite))
}
