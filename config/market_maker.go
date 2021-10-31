package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kpango/glg"
	"io"
	"io/ioutil"
	"mm2_client/constants"
	"mm2_client/helpers"
	"strconv"
)

type SimplePairMarketMakerConf struct {
	Base                                  string   `json:"base"`
	Rel                                   string   `json:"rel"`
	Max                                   bool     `json:"max,omitempty"`
	BalancePercent                        string   `json:"balance_percent,omitempty"`
	MinVolume                             *string  `json:"min_volume,omitempty"`
	Spread                                string   `json:"spread"`
	BaseConfs                             int      `json:"base_confs"`
	BaseNota                              bool     `json:"base_nota"`
	RelConfs                              int      `json:"rel_confs"`
	RelNota                               bool     `json:"rel_nota"`
	Enable                                bool     `json:"enable"`
	PriceElapsedValidity                  *float64 `json:"price_elapsed_validity,omitempty"`
	CheckLastBidirectionalTradeThreshHold *bool    `json:"check_last_bidirectional_trade_thresh_hold,omitempty"`
}

type StartSimpleMarketMakerParams struct {
	PriceUrl       string                               `json:"price_url,omitempty"`
	TelegramApiKey string                               `json:"telegram_api_key,omitempty"`
	TelegramChatId string                               `json:"telegram_chat_id,omitempty"`
	Cfg            map[string]SimplePairMarketMakerConf `json:"cfg"`
}

var GSimpleMarketMakerConf *StartSimpleMarketMakerParams

func init() {
	GSimpleMarketMakerConf = &StartSimpleMarketMakerParams{}
	GSimpleMarketMakerConf.Cfg = make(map[string]SimplePairMarketMakerConf)
}

func (cfg *StartSimpleMarketMakerParams) ToJson() string {
	b, err := json.MarshalIndent(cfg, "", "    ")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}

func NewMarketMakerTemplateConfig() StartSimpleMarketMakerParams {
	cfg := make(map[string]SimplePairMarketMakerConf)
	minVol := "0.25"
	priceElapsedValidity := 30.0
	cfg["KMD/LTC"] = SimplePairMarketMakerConf{
		Base:                                  "KMD",
		Rel:                                   "LTC",
		Max:                                   true,
		BalancePercent:                        "",
		MinVolume:                             &minVol,
		Spread:                                "1.015",
		BaseConfs:                             1,
		BaseNota:                              false,
		RelConfs:                              1,
		RelNota:                               false,
		Enable:                                true,
		PriceElapsedValidity:                  &priceElapsedValidity,
		CheckLastBidirectionalTradeThreshHold: helpers.BoolAddr(true),
	}
	out := StartSimpleMarketMakerParams{
		PriceUrl:       "http://price.cipig.net:1313/api/v2/tickers?expire_at=600",
		TelegramApiKey: "",
		TelegramChatId: "",
		Cfg:            cfg,
	}
	return out
}

func ParseMarketMakerConf() {
	if constants.GMarketMakerCfgLoaded {
		_ = glg.Infof("Simple Market Maker conf already loaded")
		return
	}
	var desktopCoinsPath = constants.GMM2Dir + "/" + "mm2_market_maker.json"
	file, _ := ioutil.ReadFile(desktopCoinsPath)
	err := json.Unmarshal([]byte(file), &GSimpleMarketMakerConf)
	if err != nil {
		_ = glg.Errorf("Cannot parse cfg: %v", err)
	}
	if len(GSimpleMarketMakerConf.Cfg) > 0 {
		constants.GMarketMakerCfgLoaded = true
	}
	helpers.PrintCheck("Successfully load MarketMakerConf with: "+strconv.Itoa(len(GSimpleMarketMakerConf.Cfg))+" pairs.", true)
}

func ParseMarketMakerConfFromUrl(url string) error {
	if constants.GMarketMakerCfgLoaded {
		return nil
	}

	resp, err := helpers.CrossGet(url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	decodeErr := json.NewDecoder(resp.Body).Decode(&GSimpleMarketMakerConf)
	if decodeErr != nil {
		return decodeErr
	}

	if len(GSimpleMarketMakerConf.Cfg) > 0 {
		constants.GMarketMakerCfgLoaded = true
		return nil
	}
	return errors.New("empty cfg")
}
