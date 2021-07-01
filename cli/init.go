package cli

import (
	"github.com/kyokomi/emoji/v2"
	helpers "mm2_client/helpers"
)

const targetCoinsUrl = "https://raw.githubusercontent.com/KomodoPlatform/coins/master/coins"

func downloadCoinsFile() {
	_, _ = emoji.Println("Downloading coins file :arrows_counterclockwise:")
}

func processCoinsFile() {
	targetDir := helpers.GetWorkingDir() + "/mm2"
	if !helpers.CreateDirIfNotExist(targetDir) {
		return
	}
	targetPath := targetDir + "/coins.json"
	if !helpers.FileExists(targetPath) {
		helpers.PrintCheck("Checking if coins file is present:", false)
		downloadCoinsFile()
	} else {
		helpers.PrintCheck("Checking if coins file is present:", true)
	}
}

func InitMM2() {
	processCoinsFile()
}
