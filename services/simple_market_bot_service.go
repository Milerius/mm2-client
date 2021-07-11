package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mm2_client/constants"
	"os"
	"path/filepath"
	"time"
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

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

var gSimpleMarketMakerRegistry = make(map[string]SimplePairMarketMakerConf)
var gQuitMarketMakerBot chan bool

func init() {
	_ = os.MkdirAll(filepath.Join(constants.GetAppDataPath(), "logs"), os.ModePerm)
	file, err := os.OpenFile(constants.GetAppDataPath()+"/logs/simple.market.maker.logs", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func NewMarketMakerConfFromFile(targetPath string) bool {
	file, _ := ioutil.ReadFile(targetPath)
	err := json.Unmarshal(file, &gSimpleMarketMakerRegistry)
	if err != nil {
		fmt.Println("Couldn't parse cfg file - aborting")
		return false
	}
	return true
}

func marketMakerProcess() {
	InfoLogger.Println("process market maker")
	for key, value := range gSimpleMarketMakerRegistry {
		if value.Enable {
			InfoLogger.Printf("Treating %s\n", key)
		}
	}
}

func StartSimpleMarketMakerBotService() {
	for {
		select {
		case <-gQuitMarketMakerBot:
			InfoLogger.Println("Simple Market Maker Bot service successfully stopped")
			constants.GSimpleMarketMakerBotRunning = false
			close(gQuitMarketMakerBot)
			return
		default:
			marketMakerProcess()
			time.Sleep(30 * time.Second)
		}
	}
}

func StopSimpleMarketMakerBotService() {
	InfoLogger.Println("Stopping Simple Market Maker Bot Service - may take up to 30 seconds")
	go func() {
		gQuitMarketMakerBot <- true
	}()
}

func StartSimpleMarketMakerBot() {
	if constants.GMM2Running {
		if constants.GSimpleMarketMakerBotRunning {
			fmt.Println("Simple Market Maker bot is already running (or being stopped) - skipping")
		} else {
			if resp := NewMarketMakerConfFromFile(constants.GSimpleMarketMakerConf); resp {
				InfoLogger.Printf("Starting simple market maker bot with %d coin(s)\n", len(gSimpleMarketMakerRegistry))
				constants.GSimpleMarketMakerBotRunning = true
				gQuitMarketMakerBot = make(chan bool)
				go StartSimpleMarketMakerBotService()
			} else {
				fmt.Println("Couldn't start simple market maker without valid conf")
			}
		}
	} else {
		fmt.Println("MM2 need to be started before starting the simple market maker bot")
	}
}
