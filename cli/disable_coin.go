package cli

import (
	"mm2_client/helpers"
	"mm2_client/http"
)

func DisableCoin(coin string) {
	resp := http.DisableCoin(coin)
	if resp != nil {
		helpers.PrintCheck(coin+" successfully disabled", true)
	}
}

func DisableCoins(coins []string) {

}
