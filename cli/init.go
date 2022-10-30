package cli

import (
	"fmt"
	"github.com/kyokomi/emoji/v2"
	"mm2_client/config"
	"mm2_client/helpers"
	mm2http "mm2_client/http"
	"os"
)

const targetCoinsUrl = "https://raw.githubusercontent.com/KomodoPlatform/coins/master/coins"

func downloadCoinsFile(filePath string) {
	_, _ = emoji.Printf("Downloading coins file %s :arrows_counterclockwise:\n", targetCoinsUrl)
	err := helpers.DownloadFile(filePath, targetCoinsUrl, true)
	if err != nil {
		_, _ = emoji.Printf("Error when Downloading %s: %v", targetCoinsUrl, err)
		os.Exit(1)
	}
}

func processCoinsFile() {
	targetDir := helpers.GetWorkingDir() + "/mm2"
	if !helpers.CreateDirIfNotExist(targetDir) {
		return
	}
	targetPath := targetDir + "/coins.json"
	if !helpers.FileExists(targetPath) {
		helpers.PrintCheck("Checking if coins file is present:", false)
		downloadCoinsFile(targetPath)
	} else {
		helpers.PrintCheck("Checking if coins file is present:", true)
	}
}

func downloadLastMM2(targetPath string, zipTarget string, targetDir string) {
	_, downloadURL, err := mm2http.GetUrlLastMM2()
	if err != nil {
		_, _ = emoji.Printf("Error when Getting the Last URL Infos: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Download url is: %s\n", downloadURL)
	err = helpers.DownloadFile(zipTarget, downloadURL, false)
	if err != nil {
		_, _ = emoji.Printf("Error when Downloading %s: %v", downloadURL, err)
		os.Exit(1)
	}
	uz := helpers.NewUnzip()
	_, _ = emoji.Printf("Extracting: %s :outbox_tray:\n", zipTarget)
	files, err := uz.Extract(zipTarget, targetDir)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
	_, _ = emoji.Printf("Successfully extracted %d files, extracted: %v :white_check_mark:\n", len(files), files)
	_ = os.Remove(zipTarget)
}

func processMM2Release() {
	targetDir := helpers.GetWorkingDir() + "/mm2"
	targetPath := targetDir + "/mm2"
	zipTarget := targetDir + "/mm2.zip"
	if !helpers.FileExists(targetPath) {
		helpers.PrintCheck("Checking if mm2 binary is present:", false)
		downloadLastMM2(targetPath, zipTarget, targetDir)
	} else {
		helpers.PrintCheck("Checking if mm2 binary is present:", true)
	}
}

func processMM2Json() {
	targetDir := helpers.GetWorkingDir() + "/mm2"
	targetPath := targetDir + "/MM2.json"
	if !helpers.FileExists(targetPath) {
		helpers.PrintCheck("Checking if MM2.json is present and configured", false)
		mm2CFG := config.NewMM2Config()
		freshCfg := mm2CFG.ToJson()
		f, err := os.Create(targetPath)
		if err != nil {
			fmt.Printf("Err: %v", err)
			os.Exit(1)
		}
		defer f.Close()
		_, err = f.WriteString(freshCfg)
		if err != nil {
			fmt.Printf("Err: %v", err)
			os.Exit(1)
		}
		helpers.PrintCheck("Successfully generated a MM2.json template configuration", true)
		CheckMM2Configuration(mm2CFG)
		CheckMM2DB(mm2CFG)
		CheckMM2SeedNode(mm2CFG)
	} else {
		CheckMM2Configuration(config.NewMM2ConfigFromFile(targetPath))
	}
}

func processVersionConfiguration() {
	version := mm2http.GetLastDesktopVersion()
	targetDir := helpers.GetWorkingDir() + "/mm2"
	targetPath := targetDir + "/coins_config.json"
	targetUrl := "https://raw.githubusercontent.com/KomodoPlatform/coins/master/utils/coins_config.json"
	if !helpers.FileExists(targetPath) {
		helpers.PrintCheck("Checking if desktop coins json is present", false)
		emoji.Printf("Downloading now %s :recycle:\n", targetUrl)
		helpers.DownloadFile(targetPath, targetUrl, true)
	} else {
		helpers.PrintCheck("Checking if desktop coins json is present", true)
	}
}

func processMarketMakerConf() {
	targetDir := helpers.GetWorkingDir() + "/mm2"
	targetPath := targetDir + "/mm2_market_maker.json"
	if !helpers.FileExists(targetPath) {
		helpers.PrintCheck("Checking if mm2_market_maker.json is present and configured", false)
		marketMakerCfg := config.NewMarketMakerTemplateConfig()
		freshCfg := marketMakerCfg.ToJson()
		f, err := os.Create(targetPath)
		if err != nil {
			fmt.Printf("Err: %v", err)
			os.Exit(1)
		}
		defer f.Close()
		_, err = f.WriteString(freshCfg)
		if err != nil {
			fmt.Printf("Err: %v", err)
			os.Exit(1)
		}
		helpers.PrintCheck("Successfully generated a mm2_market_maker.json template configuration", true)
	} else {
		helpers.PrintCheck("Checking if market maker conf json is present", true)
	}
}

func InitMM2() {
	processCoinsFile()
	processMM2Release()
	processMM2Json()
	processVersionConfiguration()
	processMarketMakerConf()
}
