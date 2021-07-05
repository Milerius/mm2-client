package cli

import (
	"github.com/manifoldco/promptui"
	"mm2_client/http"
)

func postKmdRewardsInfo(resp *http.KMDRewardsInfoAnswer) {
	if resp.ToTable() {
		prompt := promptui.Select{
			Label: "Do you want to claim your rewards ?",
			Items: []string{"Yes", "No"},
		}
		_, result, _ := prompt.Run()
		if result == "Yes" {
			if respBalance := http.MyBalance("KMD"); respBalance != nil {
				Send("KMD", "max", respBalance.Address, []string{})
			}
		}
	}
}

func KmdRewardsInfo() {
	if resp := http.KmdRewardsInfo(); resp != nil {

		postKmdRewardsInfo(resp)
	}
}
