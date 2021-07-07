package cli

import (
	"fmt"
	"mm2_client/config"
	"mm2_client/http"
	"os"
	"strings"
)

func Executor(fullCommand string) {
	fullCommand = strings.TrimSpace(fullCommand)
	command := strings.Split(fullCommand, " ")
	switch command[0] {
	case "init":
		InitMM2()
	case "help":
		if len(command) == 1 {
			ShowGlobalHelp()
		} else if len(command) > 1 {
			ShowCommandHelp(command[1])
		}
	case "start":
		StartMM2()
	case "stop":
		StopMM2()
	case "enable":
		if len(command) == 1 {
			ShowCommandHelp(command[0])
		} else if len(command) == 2 {
			Enable(command[1])
		} else {
			EnableMultipleCoins(command[1:])
		}
	case "enable_active_coins":
		EnableMultipleCoins(config.RetrieveActiveCoins())
	case "enable_all_coins":
		EnableMultipleCoins(config.RetrieveAllCoins())
	case "disable_coin":
		if len(command) == 1 {
			ShowCommandHelp(command[0])
		} else if len(command) == 2 {
			DisableCoin(command[1])
		} else {
			DisableCoins(command[1:])
		}
	case "my_balance":
		if len(command) == 1 {
			ShowCommandHelp(command[0])
		} else if len(command) == 2 {
			MyBalance(command[1])
		} else {
			MyBalanceMultipleCoins(command[1:])
		}
	case "balance_all":
		MyBalanceMultipleCoins(config.RetrieveActiveCoins())
	case "kmd_rewards_info":
		KmdRewardsInfo()
	case "disable_enabled_coins":
		DisableCoins(http.GetEnabledCoins().ToSlice())
	case "disable_zero_balance":
		DisableCoins(http.ToSliceEmptyBalance(MyBalanceMultipleCoinsSilent(http.GetEnabledCoins().ToSlice()), true))
	case "orderbook":
		if len(command) != 3 {
			ShowCommandHelp("orderbook")
		} else {
			Orderbook(command[1], command[2])
		}
	case "my_tx_history":
		if len(command) == 1 {
			ShowCommandHelp("my_tx_history")
		} else {
			MyTxHistory(command[1], command[2:])
		}
	case "my_recent_swaps":
		if len(command) == 1 {
			MyRecentSwaps("50", "1", []string{})
		} else if len(command) == 2 {
			MyRecentSwaps(command[1], "1", []string{})
		} else if len(command) == 3 {
			MyRecentSwaps(command[1], command[2], []string{})
		} else {
			MyRecentSwaps(command[1], command[2], command[3:])
		}
	case "get_enabled_coins":
		if len(command) > 1 {
			ShowCommandHelp("get_enabled_coins")
		} else {
			GetEnabledCoins()
		}
	case "withdraw":
		if len(command) < 4 {
			ShowCommandHelp("withdraw")
		} else {
			PostWithdraw(Withdraw(command[1], command[2], command[3], command[4:]))
		}
	case "send":
		if len(command) < 4 {
			ShowCommandHelp("send")
		} else {
			Send(command[1], command[2], command[3], command[4:])
		}
	case "broadcast":
		if len(command) != 3 {
			ShowCommandHelp("broadcast")
		} else {
			Broadcast(command[1], command[2])
		}
	case "exit":
		fmt.Println("Quitting the application - trying to shutdown MM2")
		StopMM2()
		os.Exit(0)
	}
	return
}
