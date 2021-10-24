package config

import (
	"encoding/json"
	"errors"
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
	PriceUrl          string                               `json:"price_url,omitempty"`
	TelegramApiKey    string                               `json:"telegram_api_key,omitempty"`
	TelegramApiChatId string                               `json:"telegram_api_chat_id,omitempty"`
	Cfg               map[string]SimplePairMarketMakerConf `json:"cfg"`
}

var GSimpleMarketMakerConf *StartSimpleMarketMakerParams

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
	return errors.New("unknown error")
}
