package cli

import (
	"fmt"
	"mm2_client/config"
	"mm2_client/constants"
	"mm2_client/market_making"
	"mm2_client/mm2_tools_generics"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"mm2_client/services"
	"os"
	"strconv"
	"strings"
)

func Executor(fullCommand string) {
	fullCommand = strings.TrimSpace(fullCommand)
	command := strings.Split(fullCommand, " ")
	switch command[0] {
	case "setprice":
		if len(command) != 5 {
			mm2_tools_generics.ShowCommandHelp(command[0])
		} else {
			SetPrice(command[1], command[2], command[3], command[4])
		}
	case "cancel_order":
		if len(command) != 2 {
			mm2_tools_generics.ShowCommandHelp(command[0])
		} else {
			mm2_tools_generics.CancelOrderCLI(command[1])
		}
	case "init":
		InitMM2()
	case "help":
		if len(command) == 1 {
			mm2_tools_generics.ShowGlobalHelp()
		} else if len(command) > 1 {
			mm2_tools_generics.ShowCommandHelp(command[1])
		}
	case "start":
		if len(command) == 1 {
			StartMM2(true)
		} else if len(command) == 2 {
			withServices, err := strconv.ParseBool(command[1])
			if err != nil {
				withServices = false
				fmt.Printf("%v - mm2 starting without services\n", err)
			}
			StartMM2(withServices)
		}
	case "stop":
		StopMM2()
	case "enable":
		if len(command) == 1 {
			mm2_tools_generics.ShowCommandHelp(command[0])
		} else if len(command) == 2 {
			mm2_tools_generics.EnableCLI(command[1])
		} else {
			mm2_tools_generics.EnableMultipleCoins(command[1:])
		}
	case "enable_active_coins":
		mm2_tools_generics.EnableMultipleCoins(config.RetrieveActiveCoins())
	case "enable_all_coins":
		mm2_tools_generics.EnableMultipleCoins(config.RetrieveAllCoins())
	case "disable_coin":
		if len(command) == 1 {
			mm2_tools_generics.ShowCommandHelp(command[0])
		} else if len(command) == 2 {
			mm2_tools_generics.DisableCoinCLI(command[1])
		} else {
			mm2_tools_generics.DisableCoins(command[1:])
		}
	case "my_balance":
		if len(command) == 1 {
			mm2_tools_generics.ShowCommandHelp(command[0])
		} else if len(command) == 2 {
			mm2_tools_generics.MyBalanceCLI(command[1])
		} else {
			mm2_tools_generics.MyBalanceMultipleCoinsCLI(command[1:])
		}
	case "balance_all":
		mm2_tools_generics.MyBalanceMultipleCoinsCLI(config.RetrieveActiveCoins())
	case "kmd_rewards_info":
		if mm2_tools_generics.KmdRewardsInfoCLI() {
			postKmdRewardsInfo()
		}
	case "disable_enabled_coins":
		val, _ := mm2_tools_generics.GetEnabledCoins()
		mm2_tools_generics.DisableCoins(val.ToSlice())
	case "disable_zero_balance":
		val, _ := mm2_tools_generics.GetEnabledCoins()
		mm2_tools_generics.DisableCoins(mm2_data_structure.ToSliceEmptyBalance(mm2_tools_generics.MyBalanceMultipleCoinsSilent(val.ToSlice()), true))
	case "orderbook":
		if len(command) != 3 {
			mm2_tools_generics.ShowCommandHelp("orderbook")
		} else {
			mm2_tools_generics.OrderbookCLI(command[1], command[2])
		}
	case "my_tx_history":
		if len(command) == 1 {
			mm2_tools_generics.ShowCommandHelp("my_tx_history")
		} else {
			mm2_tools_generics.MyTxHistoryCLI(command[1], command[2:])
		}
	case "my_orders":
		mm2_tools_generics.MyOrdersCLI()
	case "my_recent_swaps":
		if len(command) == 1 {
			mm2_tools_generics.MyRecentSwapsCLI("50", "1", []string{})
		} else if len(command) == 2 {
			mm2_tools_generics.MyRecentSwapsCLI(command[1], "1", []string{})
		} else if len(command) == 3 {
			mm2_tools_generics.MyRecentSwapsCLI(command[1], command[2], []string{})
		} else {
			mm2_tools_generics.MyRecentSwapsCLI(command[1], command[2], command[3:])
		}
	case "get_enabled_coins":
		if len(command) > 1 {
			mm2_tools_generics.ShowCommandHelp("get_enabled_coins")
		} else {
			mm2_tools_generics.GetEnabledCoinsCLI()
		}
	case "withdraw":
		if len(command) < 4 {
			mm2_tools_generics.ShowCommandHelp("withdraw")
		} else {
			resp, err := mm2_tools_generics.WithdrawCLI(command[1], command[2], command[3], command[4:])
			if err != nil {
				PostWithdraw(resp)
			} else {
				fmt.Println(err)
			}
		}
	case "send":
		if len(command) < 4 {
			mm2_tools_generics.ShowCommandHelp("send")
		} else {
			mm2_tools_generics.Send(command[1], command[2], command[3], command[4:])
		}
	case "broadcast":
		if len(command) != 3 {
			mm2_tools_generics.ShowCommandHelp("broadcast")
		} else {
			mm2_tools_generics.BroadcastCLI(command[1], command[2])
		}
	case "get_binance_supported_pairs":
		if len(command) == 1 {
			services.GetBinanceSupportedPairs("")
		} else if len(command) == 2 {
			services.GetBinanceSupportedPairs(command[1])
		}
	case "start_simple_market_maker_bot":
		_ = market_making.StartSimpleMarketMakerBot(constants.GSimpleMarketMakerConf, "file")
	case "stop_simple_market_maker_bot":
		_ = market_making.StopSimpleMarketMakerBotService()
	case "exit":
		fmt.Println("Quitting the application - trying to shutdown MM2")
		_ = market_making.StopSimpleMarketMakerBotService()
		StopMM2()
		os.Exit(0)
	}
	return
}
