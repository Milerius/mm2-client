package cli

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"mm2_client/mm2_tools_generics"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
)

func PostWithdraw(answer *mm2_data_structure.WithdrawAnswer) {
	if answer != nil {
		if answer.Result != nil {
			answer.ToTable()
			prompt := promptui.Select{
				Label: "Do you want to broadcast your transaction ?",
				Items: []string{"Yes", "No"},
			}
			_, result, _ := prompt.Run()
			if result == "Yes" {
				mm2_tools_generics.BroadcastCLI(answer.Result.Coin, answer.Result.TxHex)
			}
		} else {
			fmt.Println(answer.Error)
		}
	}
}
