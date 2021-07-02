package http

import (
	"encoding/json"
	"fmt"
	"github.com/kyokomi/emoji/v2"
	"net/http"
)

const desktopTargetUrl = "https://raw.githubusercontent.com/KomodoPlatform/atomicDEX-Desktop/dev/vcpkg.json"

type VcpkgAnswer struct {
	Name          string        `json:"name"`
	VersionString string        `json:"version-string"`
	Dependencies  []interface{} `json:"dependencies"`
}

func fetchLastVersion() (*VcpkgAnswer, error) {
	resp, err := http.Get(desktopTargetUrl)
	if err != nil {
		fmt.Printf("Error occured: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()
	var cResp *VcpkgAnswer = new(VcpkgAnswer)
	if err := json.NewDecoder(resp.Body).Decode(cResp); err != nil {
		fmt.Printf("Error occured: %v\n", err)
		return nil, err
	}
	_, _ = emoji.Println("Downloaded information about the last vcpkg desktop version: :white_check_mark:")
	return cResp, nil
}

func GetLastDesktopVersion() string {
	resp, err := fetchLastVersion()
	if err != nil {
		return ""
	}
	return resp.VersionString
}
