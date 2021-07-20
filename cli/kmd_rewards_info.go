package cli

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"mm2_client/mm2_tools_generics/mm2_http_request"
)

func postKmdRewardsInfo(resp *mm2_http_request.KMDRewardsInfoAnswer) {
	if resp.ToTable() {
		prompt := promptui.Select{
			Label: "Do you want to claim your rewards ?",
			Items: []string{"Yes", "No"},
		}
		_, result, _ := prompt.Run()
		if result == "Yes" {
			if respBalance, err := mm2_http_request.MyBalance("KMD"); respBalance != nil {
				Send("KMD", "max", respBalance.Address, []string{})
			} else {
				fmt.Println(err)
			}
		}
	}
}

func KmdRewardsInfo() {
	if resp := mm2_http_request.KmdRewardsInfo(); resp != nil {

		postKmdRewardsInfo(resp)
	}
}
