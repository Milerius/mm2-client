package main

import (
	"fmt"
	"mm2_client/config"
	"mm2_client/http"
)

func main() {
	config.ParseDesktopRegistry(http.GetLastDesktopVersion())
	fmt.Println("var subCommandsEnable = []prompt.Suggest{")
	for key, _ := range config.GCFGRegistry {
		toPrint := fmt.Sprintf("{Text: %q, Description: \"Enable %s\"},", key, key)
		fmt.Println(toPrint)
	}
	fmt.Println("}")
}
