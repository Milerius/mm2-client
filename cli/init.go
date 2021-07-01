package cli

import (
	"github.com/kyokomi/emoji/v2"
	"mm2_client/helpers"
	"os"
)

const targetCoinsUrl = "https://raw.githubusercontent.com/KomodoPlatform/coins/master/coins"

func downloadCoinsFile(filePath string) {
	_, _ = emoji.Printf("Downloading coins file %s :arrows_counterclockwise:\n", targetCoinsUrl)
	err := helpers.DownloadFile(filePath, targetCoinsUrl)
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

func InitMM2() {
	processCoinsFile()
}
