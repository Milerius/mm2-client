package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mm2_client/constants"
)

type SimplePairMarketMakerConf struct {
	Base           string  `json:"base"`
	Rel            string  `json:"rel"`
	Max            bool    `json:"max,omitempty"`
	BalancePercent int     `json:"balance_percent,omitempty"`
	MinVolume      int     `json:"min_volume"`
	Spread         float64 `json:"spread"`
	BaseConfs      int     `json:"base_confs"`
	BaseNota       bool    `json:"base_nota"`
	RelConfs       int     `json:"rel_confs"`
	RelNota        bool    `json:"rel_nota"`
	Enable         bool    `json:"enable"`
}

var gSimpleMarketMakerRegistry = make(map[string]SimplePairMarketMakerConf)

func NewMarketMakerConfFromFile(targetPath string) bool {
	file, _ := ioutil.ReadFile(targetPath)
	err := json.Unmarshal(file, &gSimpleMarketMakerRegistry)
	if err != nil {
		fmt.Println("Couldn't parse cfg file - aborting")
		return false
	}
	return true
}

func StartSimplePairMarketConf() {
	if resp := NewMarketMakerConfFromFile(constants.GSimpleMarketMakerConf); resp {
		fmt.Printf("Starting simple market maker bot with %d\n", len(gSimpleMarketMakerRegistry))
	} else {
		fmt.Println("Couldn't start simple market maker without valid conf")
	}
}
