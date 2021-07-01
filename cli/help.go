package cli

import "fmt"

const initHelp = `The init command allow you to bootstrap mm2 by downloading all the requirements
usage: init`

func ShowGlobalHelp() {
	fmt.Println("TODO: global help")
}

func ShowCommandHelp(command string) {
	switch command {
	case "init":
		fmt.Println(initHelp)
	}
}
