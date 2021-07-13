package config

import (
	"log"
	"os"
	"testing"
)

func TestGetDesktopPath(t *testing.T) {
	folderInfo, err := os.Stat(GetDesktopPath("standard"))
	if os.IsNotExist(err) {
		t.Errorf("Failed, expecting directory to exist: %v", err)
	}
	log.Println(folderInfo)
}
