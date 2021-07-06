package main

import (
	"fmt"
	"mm2_client/config"
	"mm2_client/http"
	"sort"
)

func main() {
	config.ParseDesktopRegistry(http.GetLastDesktopVersion())
	keys := make([]string, 0, len(config.GCFGRegistry))

	fmt.Println("var subCommandsEnable = []prompt.Suggest{")
	for key, _ := range config.GCFGRegistry {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		toPrint := fmt.Sprintf("{Text: %q, Description: \"Enable %s\"},", key, key)
		fmt.Println(toPrint)
	}

	fmt.Println("}")
}
