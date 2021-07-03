package cli

import (
	"mm2_client/helpers"
	"mm2_client/http"
)

func StopMM2() {
	if http.Stop() {
		helpers.PrintCheck("MM2 successfully stopped", true)
	}
}
