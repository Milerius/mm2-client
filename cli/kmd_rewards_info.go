package cli

import (
	"mm2_client/http"
)

func KmdRewardsInfo() {
	if resp := http.KmdRewardsInfo(); resp != nil {
		resp.ToTable()
	}
}
