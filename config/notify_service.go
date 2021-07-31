package config

import (
	"encoding/json"
	"github.com/kpango/glg"
	"io/ioutil"
	"mm2_client/constants"
)

type TelegramApiConfig struct {
	TelegramApiToken string `json:"telegram_api_token"`
	TelegramReceiver int64  `json:"telegram_receiver"`
}

type NotifierConfig struct {
	Telegram              *TelegramApiConfig `json:"telegram"`
	MyRecentSwapsNotifier *bool              `json:"my_recent_swaps_notifier"`
}

var GNotifyCFG *NotifierConfig = nil

func ParseNotifyCfgFromFile(path string) bool {
	if constants.GNotifyCfgLoaded {
		return true
	}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		glg.Errorf("err when parsing: %v", err)
		return false
	}
	unmarshalErr := json.Unmarshal([]byte(file), &GNotifyCFG)
	if unmarshalErr != nil {
		glg.Errorf("err when unmarshaling: %v", unmarshalErr)
		return false
	}
	if GNotifyCFG != nil {
		constants.GNotifyCfgLoaded = true
		return true
	}
	return false
}

func ParseNotifyCfgFromString(cfg string) bool {
	if constants.GNotifyCfgLoaded {
		return true
	}
	_ = json.Unmarshal([]byte(cfg), &GNotifyCFG)
	if GNotifyCFG != nil {
		constants.GNotifyCfgLoaded = true
		return true
	}
	return false
}
