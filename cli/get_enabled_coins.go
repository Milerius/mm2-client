package cli

import (
	"mm2_client/http"
)

func GetEnabledCoins() {
	resp := http.GetEnabledCoins()
	if resp != nil {
		resp.ToTable()
	}
}
