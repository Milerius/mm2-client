package cli

import (
	"fmt"
	"os"
	"strings"
)

func Executor(fullCommand string) {
	fullCommand = strings.TrimSpace(fullCommand)
	command := strings.Split(fullCommand, " ")
	switch command[0] {
	case "init":
		InitMM2()
	case "help":
		if len(command) == 1 {
			ShowGlobalHelp()
		} else if len(command) > 1 {
			ShowCommandHelp(command[1])
		}
	case "start":
		StartMM2()
	case "stop":
		StopMM2()
	case "enable":
		if len(command) == 1 {
			ShowCommandHelp(command[0])
		} else if len(command) == 2 {
			Enable(command[1])
		} else {
			EnableMultipleCoins(command[1:])
		}
	case "disable_coin":
		if len(command) == 1 {
			ShowCommandHelp(command[0])
		} else if len(command) == 2 {
			DisableCoin(command[1])
		} else {
			DisableCoins(command[1:])
		}
	case "get_enabled_coins":
		if len(command) > 1 {
			ShowCommandHelp("get_enabled_coins")
		} else {
			GetEnabledCoins()
		}
	case "exit":
		fmt.Println("Bye")
		StopMM2()
		os.Exit(0)
	}
	return
}
