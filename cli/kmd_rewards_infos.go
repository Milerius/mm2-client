package cli

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"mm2_client/mm2_tools_generics"
)

func postKmdRewardsInfo() {
	prompt := promptui.Select{
		Label: "Do you want to claim your rewards ?",
		Items: []string{"Yes", "No"},
	}
	_, result, _ := prompt.Run()
	if result == "Yes" {
		if respBalance, err := mm2_tools_generics.MyBalance("KMD"); respBalance != nil {
			mm2_tools_generics.Send("KMD", "max", respBalance.Address, []string{})
		} else {
			fmt.Println(err)
		}
	}
}
