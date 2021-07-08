package cli

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
)

const (
	initHelp                      = `The init command allow you to bootstrap mm2 by downloading all the requirements`
	initUsage                     = `init`
	startHelp                     = `The start command allow you to start MM2 into a detached process`
	startUsage                    = `start`
	stopHelp                      = `The stop command allow you to stop MM2`
	stopUsage                     = `stop`
	exitHelp                      = `Quit the mm2-client CLI, doesn't shutdown mm2`
	exitUsage                     = `exit`
	enableHelp                    = `Enable the specified coin(s) within MM2`
	enableUsage                   = `enable <coin_1> <coin_2> ...`
	enableActiveCoinsHelp         = `Enable the active coins from the cfg within MM2`
	enableActiveCoinsUsage        = `enable_active_coins`
	enableAllCoinsHelp            = `Enable all the coins from the cfg within MM2`
	enableAllCoinsUsage           = `enable_all_coins`
	disableCoinHelp               = `Disable the specified coin(s) within MM2`
	disableCoinUsage              = `disable_coin <coin_1> <coin_2> ...`
	disableEnabledCoinHelp        = `Disable the enabled coin(s) from MM2`
	disableEnabledCoinUsage       = `disable_enabled_coins`
	disableZeroBalanceHelp        = `Disable all coins that have 0 balance`
	disableZeroBalanceUsage       = `disable_zero_balance`
	getEnabledCoinsHelp           = `List the enabled coins`
	getEnabledCoinsUsage          = `get_enabled_coins`
	myBalanceHelp                 = `Show the balance of the specified coin(s)`
	myBalanceUsage                = `my_balance <coin_1> <coin_2> ...`
	balanceAllHelp                = `Show the balance of the active coin(s)`
	balanceAllUsage               = `balance_all`
	kmdRewardsInfoHelp            = `Show the Komodo rewards information`
	kmdRewardsInfoUsage           = `kmd_rewards_info`
	BroadcastHelp                 = `Broadcast a transaction to the network`
	BroadcastUsage                = `broadcast <coin> <tx_hex>`
	OrderbookHelp                 = `Show the orderbook for the given pair`
	OrderbookUsage                = `orderbook <base> <rel>`
	SendHelp                      = `withdraw + broadcast equivalent`
	SendUsage                     = `same as withdraw`
	GetBinanceSupportedPairsHelp  = `Show supported pair for binance auto market maker bot`
	GetBinanceSupportedPairsUsage = `get_binance_supported_pairs`
	MyTxHistoryHelp               = `Show the tx history of the given coin`
	MyTxHistoryUsage              = `my_tx_history <coin> (limit|max) (page_number) (true|false)
eg: my_tx_history KMD
eg: my_tx_history KMD 50
eg: my_tx_history KMD 50 2
eg: my_tx_history KMD 50 1 true
eg: my_tx_history KMD max
eg: my_tx_history KMD max 1 true`
	MyRecentSwapsHelp  = `Show the swaps history`
	MyRecentSwapsUsage = `my_recent_swaps <limit> <page_number> (base_coin) (rel_coin) (from) (to)
eg: my_recent_swaps
eg: my_recent_swaps 50
eg: my_recent_swaps 50 1
eg: my_recent_swaps 50 1 KMD
eg: my_recent_swaps 50 1 KMD LTC
eg: my_recent_swaps 50 1 KMD LTC 01-02-2021
eg: my_recent_swaps 50 1 KMD LTC 01-02-2021 01-03-2021`
	withdrawHelp  = `Prepare a transaction to send`
	withdrawUsage = `withdraw <coin> amount|max <address> fees...
eg: withdraw KMD 1 RWaZ8yDea2j5peA6J5ftC1huPywxK66X2s
eg: withdraw KMD max RWaZ8yDea2j5peA6J5ftC1huPywxK66X2s
eg: withdraw KMD 1 RWaZ8yDea2j5peA6J5ftC1huPywxK66X2s utxo_fixed 0.1
eg: withdraw KMD 1 RWaZ8yDea2j5peA6J5ftC1huPywxK66X2s utxo_per_kbyte 1
eg: withdraw ETH 1 0x6beb7d81b03a785a79d5d9d31a896934eaac7cc0 eth_gas 3.5 55000
eg: withdraw QC 1 qHmJ3KA6ZAjR9wGjpFASn4gtUSeFAqdZgs qrc_gas 40 250000`
)

func ShowGlobalHelp() {
	data := [][]string{
		{"init", "", initHelp, initUsage},
		{"exit", "", exitHelp, exitUsage},
		{"start", "", startHelp, startUsage},
		{"stop", "", stopHelp, stopUsage},
		{"enable", "<coin_1> <coin_2> ...", enableHelp, enableUsage},
		{"enable_active_coins", "", enableActiveCoinsHelp, enableActiveCoinsUsage},
		{"enable_all_coins", "", enableAllCoinsUsage, enableAllCoinsUsage},
		{"disable_coin", "<coin_1> <coin_2> ...", disableCoinHelp, disableCoinUsage},
		{"disable_enabled_coin", "", disableEnabledCoinHelp, disableEnabledCoinUsage},
		{"disable_zero_balance", "", disableEnabledCoinHelp, disableEnabledCoinUsage},
		{"get_enabled_coins", "", getEnabledCoinsHelp, getEnabledCoinsUsage},
		{"my_balance", "<coin_1> <coin_2> ...", myBalanceHelp, myBalanceUsage},
		{"balance_all", "", balanceAllHelp, balanceAllUsage},
		{"kmd_rewards_info", "", kmdRewardsInfoHelp, kmdRewardsInfoUsage},
		{"broadcast", "<coin> <tx_hex>", BroadcastHelp, BroadcastUsage},
		{"withdraw", "", withdrawHelp, withdrawUsage},
		{"my_tx_history", "", MyTxHistoryHelp, MyTxHistoryUsage},
		{"my_recent_swaps", "", MyRecentSwapsHelp, MyRecentSwapsUsage},
		{"send", "", SendHelp, SendUsage},
		{"orderbook", "<base> <rel>", OrderbookHelp, OrderbookUsage},
		{"get_binance_supported_pairs", "", GetBinanceSupportedPairsHelp, GetBinanceSupportedPairsUsage},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"Command", "Args", "Description", "Usage"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}

func ShowCommandHelp(command string) {
	switch command {
	case "init":
		fmt.Println(initHelp)
		fmt.Printf("usage: %s\n", initUsage)
	case "exit":
		fmt.Println(exitHelp)
		fmt.Printf("usage: %s\n", exitUsage)
	case "start":
		fmt.Println(startHelp)
		fmt.Printf("usage: %s\n", startUsage)
	case "stop":
		fmt.Println(stopHelp)
		fmt.Printf("usage: %s\n", stopUsage)
	case "enable":
		fmt.Println(enableHelp)
		fmt.Printf("usage: %s\n", enableUsage)
	case "enable_active_coins":
		fmt.Println(enableActiveCoinsHelp)
		fmt.Printf("usage: %s\n", enableActiveCoinsUsage)
	case "enable_all_coins":
		fmt.Println(enableAllCoinsHelp)
		fmt.Printf("usage: %s\n", enableAllCoinsUsage)
	case "disable_coin":
		fmt.Println(disableCoinHelp)
		fmt.Printf("usage: %s\n", disableCoinUsage)
	case "disable_enabled_coin":
		fmt.Println(disableEnabledCoinHelp)
		fmt.Printf("usage: %s\n", disableEnabledCoinUsage)
	case "disable_zero_balance":
		fmt.Println(disableZeroBalanceHelp)
		fmt.Printf("usage: %s\n", disableZeroBalanceUsage)
	case "get_enabled_coins":
		fmt.Println(getEnabledCoinsHelp)
		fmt.Printf("usage: %s\n", getEnabledCoinsUsage)
	case "my_balance":
		fmt.Println(myBalanceHelp)
		fmt.Printf("usage: %s\n", myBalanceUsage)
	case "balance_all":
		fmt.Println(balanceAllHelp)
		fmt.Printf("usage: %s\n", balanceAllUsage)
	case "kmd_rewards_info":
		fmt.Println(kmdRewardsInfoHelp)
		fmt.Printf("usage: %s\n", kmdRewardsInfoUsage)
	case "withdraw":
		fmt.Println(withdrawHelp)
		fmt.Printf("usage: %s\n", withdrawUsage)
	case "broadcast":
		fmt.Println(BroadcastHelp)
		fmt.Printf("usage: %s\n", BroadcastUsage)
	case "send":
		fmt.Println(SendHelp)
		fmt.Printf("usage: %s\n", SendUsage)
	case "my_tx_history":
		fmt.Println(MyTxHistoryHelp)
		fmt.Printf("usage: %s\n", MyTxHistoryUsage)
	case "my_recent_swaps":
		fmt.Println(MyRecentSwapsHelp)
		fmt.Printf("usage: %s\n", MyRecentSwapsUsage)
	case "orderbook":
		fmt.Println(OrderbookHelp)
		fmt.Printf("usage: %s\n", OrderbookUsage)
	case "get_binance_supported_pairs":
		fmt.Println(GetBinanceSupportedPairsHelp)
		fmt.Printf("usage: %s\n", GetBinanceSupportedPairsUsage)
	default:
		fmt.Printf("Command %s not found\n", command)
	}
}
