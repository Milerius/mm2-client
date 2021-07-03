package cli

import (
	"fmt"
	"mm2_client/http"
)

func GetEnabledCoins() {
	resp := http.GetEnabledCoins()
	if resp != nil && len(resp.Result) > 0 {
		resp.ToTable()
	} else {
		fmt.Println("No coins enabled")
	}
}
