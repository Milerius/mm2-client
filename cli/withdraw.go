package cli

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"mm2_client/config"
	"mm2_client/http"
)

func PostWithdraw(answer *http.WithdrawAnswer) {
	prompt := promptui.Select{
		Label: "Do you want to broadcast your transaction ?",
		Items: []string{"Yes", "No"},
	}
	_, result, _ := prompt.Run()
	if result == "Yes" {
		Broadcast(answer.TxHex)
	}
}

func Withdraw(coin string, amount string, address string, fees []string) {
	if val, ok := config.GCFGRegistry[coin]; ok {
		var resp *http.WithdrawAnswer = nil
		if len(fees) > 0 {
			switch val.Type {
			case "ERC-20", "BEP-20", "QRC-20":
				if len(fees) == 3 {
					resp = http.Withdraw(coin, amount, address, fees, val.Type)
				} else {
					ShowCommandHelp("withdraw")
				}
			case "UTXO", "Smart Chain":
				if len(fees) == 2 {
					resp = http.Withdraw(coin, amount, address, fees, val.Type)
				} else {
					ShowCommandHelp("withdraw")
				}
			}
		} else {
			resp = http.Withdraw(coin, amount, address, []string{}, "")
		}
		if resp != nil {
			resp.ToTable()
			PostWithdraw(resp)
		}
	} else {
		fmt.Printf("%s is not present in the cfg - skipping\n", coin)
	}
}
