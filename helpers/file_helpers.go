package helpers

import (
	"fmt"
	"os"
)

func FileExists(filepath string) bool {
	if fileInfo, err := os.Stat(filepath); err != nil {
		if os.IsNotExist(err) {
			return false
		} else if os.IsPermission(err) {
			return false
		}
	} else {
		return !fileInfo.IsDir()
	}
	return true
}

func GetWorkingDir() string {
	getwd, err := os.Getwd()
	if err != nil {
		return ""
	}
	return getwd
}

func CreateDirIfNotExist(targetDir string) bool {
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		err := os.Mkdir(targetDir, os.FileMode(0700))
		if err != nil {
			fmt.Printf("Dir error: %v", err)
			return false
		} else {
			PrintCheck("Creating directory "+targetDir+":", true)
		}
	}
	return true
}
