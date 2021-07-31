package cli

import (
	"fmt"
	"mm2_client/config"
	"mm2_client/constants"
	"mm2_client/external_services"
	"mm2_client/helpers"
	"mm2_client/http"
	"os"
	"os/exec"
	"path"
)

func StartMM2(withServices bool) {
	if !helpers.FileExists(constants.GMM2BinPath) {
		fmt.Println("MM2 need to be configured, please run the init command")
	} else {
		CheckMM2Configuration(config.NewMM2ConfigFromFile(constants.GMM2ConfPath))
		fmt.Println("Starting mm2")
		cmd := exec.Command(constants.GMM2BinPath)
		cmd.Env = os.Environ()
		cmd.Env = append(cmd.Env, "MM_COINS_PATH="+constants.GMM2CoinsPath)
		cmd.Env = append(cmd.Env, "MM_CONF_PATH="+constants.GMM2ConfPath)
		cmd.Env = append(cmd.Env, "MM_LOG=/tmp/mm2.log")
		err := cmd.Start()
		if err != nil {
			fmt.Printf("cmd.Start failed: %v\n", err)
		} else {
			err := cmd.Process.Release()
			if err != nil {
				fmt.Printf("cmd.Release failed: %v\n", err)
			} else {
				config.ParseDesktopRegistry(http.GetLastDesktopVersion())
				config.ParseMM2CFGRegistry()
				if withServices {
					external_services.LaunchServices("file", path.Join(constants.GMM2Dir, "notify_service.json"))
				} else {
					fmt.Println("Start MM2 without extra services")
				}
				helpers.PrintCheck("MM2 successfully started", true)
				constants.GMM2Running = true
			}
		}
	}
}
