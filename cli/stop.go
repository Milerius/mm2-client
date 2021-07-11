package cli

import (
	"mm2_client/constants"
	"mm2_client/helpers"
	"mm2_client/http"
)

func StopMM2() {
	if http.Stop() {
		helpers.PrintCheck("MM2 successfully stopped", true)
		constants.GMM2Running = false
	}
}
