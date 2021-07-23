package mm2_tools_generics

import "fmt"

func Send(coin string, amount string, address string, fees []string) {
	if resp, err := WithdrawCLI(coin, amount, address, fees); resp != nil {
		if resp.Result != nil {
			resp.ToTable()
			fmt.Println()
			_, broadcastErr := Broadcast(resp.Result.Coin, resp.Result.TxHex)
			if broadcastErr != nil {
				fmt.Println(broadcastErr)
			}
		} else {
			fmt.Println(resp.Error)
		}
	} else {
		fmt.Println(err)
	}
}
