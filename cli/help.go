package cli

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
)

const (
	initHelp             = `The init command allow you to bootstrap mm2 by downloading all the requirements`
	initUsage            = `init`
	startHelp            = `The start command allow you to start MM2 into a detached process`
	startUsage           = `start`
	stopHelp             = `The stop command allow you to stop MM2`
	stopUsage            = `stop`
	exitHelp             = `Quit the mm2-client CLI, doesn't shutdown mm2`
	exitUsage            = `exit`
	enableHelp           = `Enable the specified coin(s) within MM2`
	enableUsage          = `enable <coin_1> <coin_2> ...`
	disableCoinHelp      = `Disable the specified coin(s) within MM2`
	disableCoinUsage     = `disable_coin <coin_1> <coin_2> ...`
	getEnabledCoinsHelp  = `List the enabled coins`
	getEnabledCoinsUsage = `get_enabled_coins`
)

func ShowGlobalHelp() {
	data := [][]string{
		{"init", "", initHelp, initUsage},
		{"exit", "", exitHelp, exitUsage},
		{"start", "", startHelp, startUsage},
		{"stop", "", stopHelp, stopUsage},
		{"enable", "<coin_1> <coin_2> ...", enableHelp, enableUsage},
		{"disable_coin", "<coin_1> <coin_2> ...", disableCoinHelp, disableCoinUsage},
		{"get_enabled_coins", "", getEnabledCoinsHelp, getEnabledCoinsUsage},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"Command", "Args", "Description", "Usage"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}

func ShowCommandHelp(command string) {
	switch command {
	case "init":
		fmt.Println(initHelp)
		fmt.Printf("usage: %s\n", initUsage)
	case "exit":
		fmt.Println(exitHelp)
		fmt.Printf("usage: %s\n", exitUsage)
	case "start":
		fmt.Println(startHelp)
		fmt.Printf("usage: %s\n", startUsage)
	case "stop":
		fmt.Println(stopHelp)
		fmt.Printf("usage: %s\n", stopUsage)
	case "enable":
		fmt.Println(enableHelp)
		fmt.Printf("usage: %s\n", enableUsage)
	case "disable_coin":
		fmt.Println(disableCoinHelp)
		fmt.Printf("usage: %s\n", disableCoinUsage)
	case "get_enabled_coins":
		fmt.Println(getEnabledCoinsHelp)
		fmt.Printf("usage: %s\n", getEnabledCoinsUsage)
	default:
		fmt.Printf("Command %s not found\n", command)
	}
}
