package cli

import (
	"fmt"
	"mm2_client/config"
	"mm2_client/helpers"
	"os"
	"os/exec"
)

func StartMM2() {
	if !helpers.FileExists(gMM2BinPath) {
		fmt.Println("MM2 need to be configured, please run the init command")
	} else {
		CheckMM2Configuration(config.NewMM2ConfigFromFile(gMM2ConfPath))
		fmt.Println("Starting mm2")
		cmd := exec.Command(gMM2BinPath)
		cmd.Env = os.Environ()
		cmd.Env = append(cmd.Env, "MM_COINS_PATH="+gMM2CoinsPath)
		cmd.Env = append(cmd.Env, "MM_CONF_PATH="+gMM2ConfPath)
		cmd.Env = append(cmd.Env, "MM_LOG=/tmp/mm2.log")
		err := cmd.Start()
		if err != nil {
			fmt.Printf("cmd.Start failed: %v\n", err)
		} else {
			err := cmd.Process.Release()
			if err != nil {
				fmt.Printf("cmd.Release failed: %v\n", err)
			} else {
				helpers.PrintCheck("MM2 Successfully started", true)
			}
		}
	}
}
