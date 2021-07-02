package cli

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
)

const initHelp = `The init command allow you to bootstrap mm2 by downloading all the requirements`

const initUsage = `init`

const exitHelp = `Quit the mm2-client CLI, doesn't shutdown mm2'`
const exitUsage = `exit`

func ShowGlobalHelp() {
	data := [][]string{
		{"init", "", initHelp, initUsage},
		{"exit", "", exitHelp, exitUsage},
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
		fmt.Printf("usage %s\n", initUsage)
	}
}
