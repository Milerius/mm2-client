package cli

import (
	"github.com/c-bata/go-prompt"
	"mm2_client/config"
	"strings"
)

var commands = []prompt.Suggest{
	{Text: "exit", Description: "Quit the application"},
	{Text: "init", Description: "Init MM2 Dependencies, Download/Setup"},
	{Text: "help", Description: "Show the global help"},
	{Text: "start", Description: "Start MM2 into a detached process"},
	{Text: "stop", Description: "Stop MM2"},
	{Text: "enable", Description: "Enable the specified coin(s) in MM2"},
	{Text: "enable_active_coins", Description: "Enable the active coin(s) from cfg"},
	{Text: "enable_all_coins", Description: "Enable all coin(s) from cfg"},
	{Text: "get_enabled_coins", Description: "List the enabled coins"},
	{Text: "disable_coin", Description: "Disable the specified coin(s)"},
	{Text: "disable_enabled_coins", Description: "Disable the enabled coin(s)"},
	{Text: "disable_zero_balance", Description: "Disable all coins that have 0 balance"},
	{Text: "my_balance", Description: "Show the balance of the specified coin(s)"},
	{Text: "balance_all", Description: "Show the balance of all the active coin(s)"},
	{Text: "kmd_rewards_info", Description: "Show the Komodo rewards information"},
	{Text: "withdraw", Description: "Prepare a transaction to send an asset to another address"},
	{Text: "broadcast", Description: "Send a transaction to the network"},
	{Text: "send", Description: "withdraw + broadcast equivalent"},
	{Text: "my_tx_history", Description: "Show the tx history of the specified coin"},
	{Text: "my_recent_swaps", Description: "Show the swaps history"},
	{Text: "my_orders", Description: "Show the active orders"},
	{Text: "cancel_order", Description: "Cancel the given order"},
	{Text: "orderbook", Description: "Show the orderbook of the given pair"},
	{Text: "start_simple_market_maker_bot", Description: "Start the simple market maker bot"},
	{Text: "stop_simple_market_maker_bot", Description: "Stop the simple market maker bot"},
	{Text: "setprice", Description: "The setprice method places an order on the orderbook, and it relies on this node acting as a maker, also called a Bob node."},
	{Text: "get_binance_supported_pairs", Description: "Show a table of binance supported pairs with average and real calculation"},
}

var subCommandsHelp = []prompt.Suggest{
	{Text: "exit", Description: "Shows help of the help command"},
	{Text: "init", Description: "Shows help of the init command"},
	{Text: "start", Description: "Shows help of the start command"},
	{Text: "stop", Description: "Shows help of the stop command"},
	{Text: "enable", Description: "Shows help of the enable command"},
	{Text: "enable_active_coins", Description: "Shows help of the enable_active_coins command"},
	{Text: "enable_all_coins", Description: "Shows help of the enable_all_coins command"},
	{Text: "get_enabled_coins", Description: "Shows help of the get_enabled_coins command"},
	{Text: "disable_coin", Description: "Shows help of the disable_coin command"},
	{Text: "disable_enabled_coins", Description: "Shows help of the disable_enabled_coins command"},
	{Text: "my_balance", Description: "Show the help of the my_balance command"},
	{Text: "balance_all", Description: "Show the help of the balance_all command"},
	{Text: "kmd_rewards_info", Description: "Show the help of the kmd_rewards_info"},
	{Text: "withdraw", Description: "Show the help of the withdraw command"},
	{Text: "broadcast", Description: "Show the help of the broadcast command"},
	{Text: "send", Description: "Show the help of the send command"},
	{Text: "my_tx_history", Description: "Show the help of the my_tx_history command"},
	{Text: "my_recent_swaps", Description: "Show the help of the my_recent_swaps command"},
	{Text: "my_orders", Description: "Show the help of the my_orders command"},
	{Text: "setprice", Description: "Show the help of the setprice command"},
	{Text: "cancel_order", Description: "Show the help of the cancel_order command"},
	{Text: "orderbook", Description: "Show the help of the orderbook command"},
	{Text: "get_binance_supported_pairs", Description: "Show the help of the get_binance_supported_pairs command"},
	{Text: "start_simple_market_maker_bot", Description: "Show the help of the start_simple_market_maker_bot command"},
	{Text: "stop_simple_market_maker_bot", Description: "Show the help of the stop_simple_market_maker_bot command"},
}

var subCommandsEnable = []prompt.Suggest{
	{Text: "1INCH-BEP20", Description: "Enable 1INCH-BEP20"},
	{Text: "1INCH-ERC20", Description: "Enable 1INCH-ERC20"},
	{Text: "AAVE-BEP20", Description: "Enable AAVE-BEP20"},
	{Text: "AAVE-ERC20", Description: "Enable AAVE-ERC20"},
	{Text: "ABY", Description: "Enable ABY"},
	{Text: "ADA-BEP20", Description: "Enable ADA-BEP20"},
	{Text: "ADX-BEP20", Description: "Enable ADX-BEP20"},
	{Text: "ADX-ERC20", Description: "Enable ADX-ERC20"},
	{Text: "AGI", Description: "Enable AGI"},
	{Text: "ANKR-BEP20", Description: "Enable ANKR-BEP20"},
	{Text: "ANKR-ERC20", Description: "Enable ANKR-ERC20"},
	{Text: "ANT", Description: "Enable ANT"},
	{Text: "ARPA", Description: "Enable ARPA"},
	{Text: "ATOM-BEP20", Description: "Enable ATOM-BEP20"},
	{Text: "AUR", Description: "Enable AUR"},
	{Text: "AVA-BEP20", Description: "Enable AVA-BEP20"},
	{Text: "AVAX-BEP20", Description: "Enable AVAX-BEP20"},
	{Text: "AWC", Description: "Enable AWC"},
	{Text: "AXE", Description: "Enable AXE"},
	{Text: "BABYDOGE", Description: "Enable BABYDOGE"},
	{Text: "BAL-BEP20", Description: "Enable BAL-BEP20"},
	{Text: "BAL-ERC20", Description: "Enable BAL-ERC20"},
	{Text: "BAND-BEP20", Description: "Enable BAND-BEP20"},
	{Text: "BAND-ERC20", Description: "Enable BAND-ERC20"},
	{Text: "BAT-BEP20", Description: "Enable BAT-BEP20"},
	{Text: "BAT-ERC20", Description: "Enable BAT-ERC20"},
	{Text: "BCH", Description: "Enable BCH"},
	{Text: "BEST", Description: "Enable BEST"},
	{Text: "BET", Description: "Enable BET"},
	{Text: "BIDR-BEP20", Description: "Enable BIDR-BEP20"},
	{Text: "BLK", Description: "Enable BLK"},
	{Text: "BNB", Description: "Enable BNB"},
	{Text: "BNBT", Description: "Enable BNBT"},
	{Text: "BNT-BEP20", Description: "Enable BNT-BEP20"},
	{Text: "BNT-ERC20", Description: "Enable BNT-ERC20"},
	{Text: "BOTS", Description: "Enable BOTS"},
	{Text: "BSTY", Description: "Enable BSTY"},
	{Text: "BTC", Description: "Enable BTC"},
	{Text: "BTC-BEP20", Description: "Enable BTC-BEP20"},
	{Text: "BTCH", Description: "Enable BTCH"},
	{Text: "BTCZ", Description: "Enable BTCZ"},
	{Text: "BTE", Description: "Enable BTE"},
	{Text: "BTT-BEP20", Description: "Enable BTT-BEP20"},
	{Text: "BTU", Description: "Enable BTU"},
	{Text: "BUSD-BEP20", Description: "Enable BUSD-BEP20"},
	{Text: "BUSD-ERC20", Description: "Enable BUSD-ERC20"},
	{Text: "CAKE", Description: "Enable CAKE"},
	{Text: "CCL", Description: "Enable CCL"},
	{Text: "CDN", Description: "Enable CDN"},
	{Text: "CEL", Description: "Enable CEL"},
	{Text: "CENNZ", Description: "Enable CENNZ"},
	{Text: "CHIPS", Description: "Enable CHIPS"},
	{Text: "CHSB", Description: "Enable CHSB"},
	{Text: "CHZ", Description: "Enable CHZ"},
	{Text: "CIPHS", Description: "Enable CIPHS"},
	{Text: "CLC", Description: "Enable CLC"},
	{Text: "COMP-BEP20", Description: "Enable COMP-BEP20"},
	{Text: "COMP-ERC20", Description: "Enable COMP-ERC20"},
	{Text: "COQUI", Description: "Enable COQUI"},
	{Text: "CRO", Description: "Enable CRO"},
	{Text: "CRV", Description: "Enable CRV"},
	{Text: "CRYPTO", Description: "Enable CRYPTO"},
	{Text: "CVC", Description: "Enable CVC"},
	{Text: "CVT", Description: "Enable CVT"},
	{Text: "DAI-BEP20", Description: "Enable DAI-BEP20"},
	{Text: "DAI-ERC20", Description: "Enable DAI-ERC20"},
	{Text: "DASH", Description: "Enable DASH"},
	{Text: "DEX", Description: "Enable DEX"},
	{Text: "DGB", Description: "Enable DGB"},
	{Text: "DGC", Description: "Enable DGC"},
	{Text: "DIA", Description: "Enable DIA"},
	{Text: "DIMI", Description: "Enable DIMI"},
	{Text: "DODO-BEP20", Description: "Enable DODO-BEP20"},
	{Text: "DODO-ERC20", Description: "Enable DODO-ERC20"},
	{Text: "DOGE", Description: "Enable DOGE"},
	{Text: "DOGE-BEP20", Description: "Enable DOGE-BEP20"},
	{Text: "DOT-BEP20", Description: "Enable DOT-BEP20"},
	{Text: "DP", Description: "Enable DP"},
	{Text: "DX", Description: "Enable DX"},
	{Text: "ECA", Description: "Enable ECA"},
	{Text: "EFL", Description: "Enable EFL"},
	{Text: "EGLD-BEP20", Description: "Enable EGLD-BEP20"},
	{Text: "ELF-BEP20", Description: "Enable ELF-BEP20"},
	{Text: "ELF-ERC20", Description: "Enable ELF-ERC20"},
	{Text: "EMC2", Description: "Enable EMC2"},
	{Text: "ENJ", Description: "Enable ENJ"},
	{Text: "EOS-BEP20", Description: "Enable EOS-BEP20"},
	{Text: "ETC-BEP20", Description: "Enable ETC-BEP20"},
	{Text: "ETH", Description: "Enable ETH"},
	{Text: "ETH-BEP20", Description: "Enable ETH-BEP20"},
	{Text: "ETHR", Description: "Enable ETHR"},
	{Text: "EURS", Description: "Enable EURS"},
	{Text: "FIL-BEP20", Description: "Enable FIL-BEP20"},
	{Text: "FIRO", Description: "Enable FIRO"},
	{Text: "FIRO-BEP20", Description: "Enable FIRO-BEP20"},
	{Text: "FJC", Description: "Enable FJC"},
	{Text: "FTC", Description: "Enable FTC"},
	{Text: "FTM-BEP20", Description: "Enable FTM-BEP20"},
	{Text: "FTM-ERC20", Description: "Enable FTM-ERC20"},
	{Text: "GLEEC", Description: "Enable GLEEC"},
	{Text: "GLEEC-OLD", Description: "Enable GLEEC-OLD"},
	{Text: "GNO", Description: "Enable GNO"},
	{Text: "GRS", Description: "Enable GRS"},
	{Text: "HEX", Description: "Enable HEX"},
	{Text: "HLC", Description: "Enable HLC"},
	{Text: "HODL", Description: "Enable HODL"},
	{Text: "HOT", Description: "Enable HOT"},
	{Text: "HPY", Description: "Enable HPY"},
	{Text: "HT", Description: "Enable HT"},
	{Text: "HUSD", Description: "Enable HUSD"},
	{Text: "IL8P", Description: "Enable IL8P"},
	{Text: "ILN", Description: "Enable ILN"},
	{Text: "INJ-BEP20", Description: "Enable INJ-BEP20"},
	{Text: "INJ-ERC20", Description: "Enable INJ-ERC20"},
	{Text: "INK", Description: "Enable INK"},
	{Text: "IOTA-BEP20", Description: "Enable IOTA-BEP20"},
	{Text: "IOTX-BEP20", Description: "Enable IOTX-BEP20"},
	{Text: "JRT-ERC20", Description: "Enable JRT-ERC20"},
	{Text: "JSTR", Description: "Enable JSTR"},
	{Text: "JUMBLR", Description: "Enable JUMBLR"},
	{Text: "KMD", Description: "Enable KMD"},
	{Text: "KMD-BEP20", Description: "Enable KMD-BEP20"},
	{Text: "KNC", Description: "Enable KNC"},
	{Text: "KOIN", Description: "Enable KOIN"},
	{Text: "LABS", Description: "Enable LABS"},
	{Text: "LBC", Description: "Enable LBC"},
	{Text: "LCC", Description: "Enable LCC"},
	{Text: "LEO", Description: "Enable LEO"},
	{Text: "LINK-BEP20", Description: "Enable LINK-BEP20"},
	{Text: "LINK-ERC20", Description: "Enable LINK-ERC20"},
	{Text: "LRC", Description: "Enable LRC"},
	{Text: "LSTR", Description: "Enable LSTR"},
	{Text: "LTC", Description: "Enable LTC"},
	{Text: "LYNX", Description: "Enable LYNX"},
	{Text: "MANA", Description: "Enable MANA"},
	{Text: "MATIC-BEP20", Description: "Enable MATIC-BEP20"},
	{Text: "MATIC-ERC20", Description: "Enable MATIC-ERC20"},
	{Text: "MCL", Description: "Enable MCL"},
	{Text: "MESH", Description: "Enable MESH"},
	{Text: "MGW", Description: "Enable MGW"},
	{Text: "MKR-BEP20", Description: "Enable MKR-BEP20"},
	{Text: "MKR-ERC20", Description: "Enable MKR-ERC20"},
	{Text: "MLN", Description: "Enable MLN"},
	{Text: "MM-ERC20", Description: "Enable MM-ERC20"},
	{Text: "MONA", Description: "Enable MONA"},
	{Text: "MORTY", Description: "Enable MORTY"},
	{Text: "MSHARK", Description: "Enable MSHARK"},
	{Text: "NAV", Description: "Enable NAV"},
	{Text: "NEAR-BEP20", Description: "Enable NEAR-BEP20"},
	{Text: "NMC", Description: "Enable NMC"},
	{Text: "NVC", Description: "Enable NVC"},
	{Text: "OC", Description: "Enable OC"},
	{Text: "OCEAN-BEP20", Description: "Enable OCEAN-BEP20"},
	{Text: "OCEAN-ERC20", Description: "Enable OCEAN-ERC20"},
	{Text: "OKB", Description: "Enable OKB"},
	{Text: "ONT-BEP20", Description: "Enable ONT-BEP20"},
	{Text: "OOT", Description: "Enable OOT"},
	{Text: "PANGEA", Description: "Enable PANGEA"},
	{Text: "PAX-BEP20", Description: "Enable PAX-BEP20"},
	{Text: "PAX-ERC20", Description: "Enable PAX-ERC20"},
	{Text: "PAXG-BEP20", Description: "Enable PAXG-BEP20"},
	{Text: "PAXG-ERC20", Description: "Enable PAXG-ERC20"},
	{Text: "PNK", Description: "Enable PNK"},
	{Text: "POWR", Description: "Enable POWR"},
	{Text: "PUT", Description: "Enable PUT"},
	{Text: "QBT", Description: "Enable QBT"},
	{Text: "QC", Description: "Enable QC"},
	{Text: "QI", Description: "Enable QI"},
	{Text: "QIAIR", Description: "Enable QIAIR"},
	{Text: "QKC-BEP20", Description: "Enable QKC-BEP20"},
	{Text: "QKC-ERC20", Description: "Enable QKC-ERC20"},
	{Text: "QNT", Description: "Enable QNT"},
	{Text: "QRC20", Description: "Enable QRC20"},
	{Text: "QTUM", Description: "Enable QTUM"},
	{Text: "REN", Description: "Enable REN"},
	{Text: "REP", Description: "Enable REP"},
	{Text: "REV", Description: "Enable REV"},
	{Text: "REVS", Description: "Enable REVS"},
	{Text: "RICK", Description: "Enable RICK"},
	{Text: "RLC", Description: "Enable RLC"},
	{Text: "RSR", Description: "Enable RSR"},
	{Text: "RVN", Description: "Enable RVN"},
	{Text: "S4F", Description: "Enable S4F"},
	{Text: "SCA", Description: "Enable SCA"},
	{Text: "SFUSD", Description: "Enable SFUSD"},
	{Text: "SHR", Description: "Enable SHR"},
	{Text: "SKL", Description: "Enable SKL"},
	{Text: "SMTF", Description: "Enable SMTF"},
	{Text: "SNT", Description: "Enable SNT"},
	{Text: "SNX-BEP20", Description: "Enable SNX-BEP20"},
	{Text: "SNX-ERC20", Description: "Enable SNX-ERC20"},
	{Text: "SOULJA", Description: "Enable SOULJA"},
	{Text: "SPACE", Description: "Enable SPACE"},
	{Text: "SPC", Description: "Enable SPC"},
	{Text: "SRM", Description: "Enable SRM"},
	{Text: "STFIRO", Description: "Enable STFIRO"},
	{Text: "STORJ", Description: "Enable STORJ"},
	{Text: "SUPERNET", Description: "Enable SUPERNET"},
	{Text: "SUSHI-BEP20", Description: "Enable SUSHI-BEP20"},
	{Text: "SUSHI-ERC20", Description: "Enable SUSHI-ERC20"},
	{Text: "SXP-BEP20", Description: "Enable SXP-BEP20"},
	{Text: "SXP-ERC20", Description: "Enable SXP-ERC20"},
	{Text: "SYS", Description: "Enable SYS"},
	{Text: "THC", Description: "Enable THC"},
	{Text: "TMTG", Description: "Enable TMTG"},
	{Text: "TRAC", Description: "Enable TRAC"},
	{Text: "TRC", Description: "Enable TRC"},
	{Text: "TRX-BEP20", Description: "Enable TRX-BEP20"},
	{Text: "TRYB-BEP20", Description: "Enable TRYB-BEP20"},
	{Text: "TRYB-ERC20", Description: "Enable TRYB-ERC20"},
	{Text: "TSL", Description: "Enable TSL"},
	{Text: "TTT", Description: "Enable TTT"},
	{Text: "TUSD-BEP20", Description: "Enable TUSD-BEP20"},
	{Text: "TUSD-ERC20", Description: "Enable TUSD-ERC20"},
	{Text: "UBT", Description: "Enable UBT"},
	{Text: "UIS", Description: "Enable UIS"},
	{Text: "UMA", Description: "Enable UMA"},
	{Text: "UNI-BEP20", Description: "Enable UNI-BEP20"},
	{Text: "UNI-ERC20", Description: "Enable UNI-ERC20"},
	{Text: "UNO", Description: "Enable UNO"},
	{Text: "UOS", Description: "Enable UOS"},
	{Text: "UQC", Description: "Enable UQC"},
	{Text: "USDC-BEP20", Description: "Enable USDC-BEP20"},
	{Text: "USDC-ERC20", Description: "Enable USDC-ERC20"},
	{Text: "USDT-BEP20", Description: "Enable USDT-BEP20"},
	{Text: "USDT-ERC20", Description: "Enable USDT-ERC20"},
	{Text: "UTK", Description: "Enable UTK"},
	{Text: "VAL", Description: "Enable VAL"},
	{Text: "VGX", Description: "Enable VGX"},
	{Text: "VITE-BEP20", Description: "Enable VITE-BEP20"},
	{Text: "VRA", Description: "Enable VRA"},
	{Text: "VRM", Description: "Enable VRM"},
	{Text: "VRSC", Description: "Enable VRSC"},
	{Text: "WBTC", Description: "Enable WBTC"},
	{Text: "WCN", Description: "Enable WCN"},
	{Text: "WLC", Description: "Enable WLC"},
	{Text: "WSB", Description: "Enable WSB"},
	{Text: "XLM-BEP20", Description: "Enable XLM-BEP20"},
	{Text: "XMY", Description: "Enable XMY"},
	{Text: "XOR", Description: "Enable XOR"},
	{Text: "XPM", Description: "Enable XPM"},
	{Text: "XRP-BEP20", Description: "Enable XRP-BEP20"},
	{Text: "XTZ-BEP20", Description: "Enable XTZ-BEP20"},
	{Text: "XVC", Description: "Enable XVC"},
	{Text: "XVC-BEP20", Description: "Enable XVC-BEP20"},
	{Text: "XVS", Description: "Enable XVS"},
	{Text: "YFI-BEP20", Description: "Enable YFI-BEP20"},
	{Text: "YFI-ERC20", Description: "Enable YFI-ERC20"},
	{Text: "YFII-BEP20", Description: "Enable YFII-BEP20"},
	{Text: "YFII-ERC20", Description: "Enable YFII-ERC20"},
	{Text: "ZEC", Description: "Enable ZEC"},
	{Text: "ZER", Description: "Enable ZER"},
	{Text: "ZET", Description: "Enable ZET"},
	{Text: "ZIL-BEP20", Description: "Enable ZIL-BEP20"},
	{Text: "ZILLA", Description: "Enable ZILLA"},
	{Text: "ZRX", Description: "Enable ZRX"},
	{Text: "tBTC-TEST", Description: "Enable tBTC-TEST"},
	{Text: "tQTUM", Description: "Enable tQTUM"},
}

type Completer struct {
}

func NewCompleter() (*Completer, error) {
	return &Completer{}, nil
}

func (c *Completer) argumentsCompleter(args []string) []prompt.Suggest {
	if len(args) <= 1 {
		return prompt.FilterContains(commands, args[0], true)
	}
	first := args[0]
	switch first {
	case "my_orders":
		cur := args[len(args)-1]
		if len(args) == 2 {
			var subCommandsMyOrders = []prompt.Suggest{
				{Text: "true", Description: "My orders with fees (trade_preimage) (can be slow)"},
				{Text: "false", Description: "My orders without fees (faster)"},
			}
			return prompt.FilterContains(subCommandsMyOrders, cur, true)
		}
	case "start":
		cur := args[len(args)-1]
		if len(args) == 2 {
			var subCommandsStart = []prompt.Suggest{
				{Text: "true", Description: "Start MM2 with extra services (price, data)"},
				{Text: "false", Description: "Start MM2 without extra services (price, data)"},
			}
			return prompt.FilterContains(subCommandsStart, cur, true)
		}
	case "setprice":
		cur := args[len(args)-1]
		if len(args) == 2 || len(args) == 3 {
			return prompt.FilterHasPrefix(subCommandsEnable, cur, true)
		}
		if len(args) == 4 {
			var subCommandsSetPriceFirst = []prompt.Suggest{
				{Text: "1", Description: "Set the price per unit for (" + args[1] + ")"},
			}
			return prompt.FilterContains(subCommandsSetPriceFirst, cur, true)
		}
		if len(args) == 5 {
			var subCommandsSetPriceSecond = []prompt.Suggest{
				{Text: "1", Description: "Set the volume of (" + args[1] + ") that you want to sell"},
				{Text: "max", Description: "Use the max balance of (" + args[1] + ")"},
			}
			return prompt.FilterContains(subCommandsSetPriceSecond, cur, true)
		}
	case "orderbook":
		cur := args[len(args)-1]
		if len(args) == 2 || len(args) == 3 {
			return prompt.FilterHasPrefix(subCommandsEnable, cur, true)
		}
	case "help":
		second := args[1]
		if len(args) == 2 {
			return prompt.FilterHasPrefix(subCommandsHelp, second, true)
		}
	case "enable", "disable_coin", "my_balance":
		cur := args[len(args)-1]
		if len(args) >= 2 {
			return prompt.FilterHasPrefix(subCommandsEnable, cur, true)
		}
	case "broadcast":
		cur := args[len(args)-1]
		if len(args) == 2 {
			return prompt.FilterHasPrefix(subCommandsEnable, cur, true)
		}
	case "my_recent_swaps":
		cur := args[len(args)-1]
		if len(args) == 2 {
			var subCommandsRecentSwapsSecond = []prompt.Suggest{
				{Text: "50", Description: "Show the N last swaps"},
			}
			return prompt.FilterContains(subCommandsRecentSwapsSecond, cur, true)
		}
		if len(args) == 3 {
			var subCommandsRecentSwapsThird = []prompt.Suggest{
				{Text: "1", Description: "Cursor the swap history on this page"},
			}
			return prompt.FilterContains(subCommandsRecentSwapsThird, cur, true)
		}
		if len(args) == 4 || len(args) == 5 {
			return prompt.FilterHasPrefix(subCommandsEnable, cur, true)
		}
		if len(args) == 6 || len(args) == 7 {
			var subCommandsRecentSwapsFifth = []prompt.Suggest{
				{Text: "01-02-2021", Description: "Choose the date from/to"},
			}
			return prompt.FilterContains(subCommandsRecentSwapsFifth, cur, true)
		}
	case "my_tx_history":
		cur := args[len(args)-1]
		if len(args) == 2 {
			return prompt.FilterHasPrefix(subCommandsEnable, cur, true)
		}
		if len(args) == 3 {
			var subCommandsTxSecond = []prompt.Suggest{
				{Text: "max", Description: "Show the full history"},
				{Text: "50", Description: "Show the N last transactions"},
			}
			return prompt.FilterContains(subCommandsTxSecond, cur, true)
		}
		if len(args) == 4 {
			if args[2] != "max" {
				var subCommandsTxThird = []prompt.Suggest{
					{Text: "1", Description: "Specify the page of history"},
				}
				return prompt.FilterContains(subCommandsTxThird, cur, true)
			}
		}
		if len(args) == 5 {
			var subCommandsTxFourth = []prompt.Suggest{
				{Text: "true", Description: "If you want fiat value at the time of the tx"},
				{Text: "false", Description: "If you want fiat value at the time of today"},
			}
			return prompt.FilterContains(subCommandsTxFourth, cur, true)
		}
	case "withdraw", "send":
		cur := args[len(args)-1]
		if len(args) == 2 {
			return prompt.FilterHasPrefix(subCommandsEnable, cur, true)
		}
		if len(args) == 3 {
			var subCommandsWithdrawFirst = []prompt.Suggest{
				{Text: "1", Description: "Use withdraw with a manual amount"},
				{Text: "max", Description: "Use withdraw with max balance"},
			}
			return prompt.FilterContains(subCommandsWithdrawFirst, cur, true)
		}
		if len(args) == 4 {
			var subCommandsWithdrawSecond = []prompt.Suggest{
				{Text: "<address>", Description: "Address where you want to send " + args[1]},
			}
			return prompt.FilterContains(subCommandsWithdrawSecond, cur, true)
		}
		if len(args) == 5 {
			var out []prompt.Suggest
			if len(config.GCFGRegistry) > 0 {
				if val, ok := config.GCFGRegistry[args[1]]; ok {
					switch val.Type {
					case "BEP-20", "ERC-20":
						out = append(out, prompt.Suggest{Text: "eth_gas", Description: "(optional) if you want to specify eth_gas (also work for BEP-20)"})
					case "QRC-20":
						out = append(out, prompt.Suggest{Text: "qrc_gas", Description: "(optional) if you want to specify qrc_gas"})
					case "UTXO", "Smart Chain":
						out = append(out, prompt.Suggest{Text: "utxo_fixed", Description: "(optional) if you want to specify utxo amount"})
						out = append(out, prompt.Suggest{Text: "utxo_per_kbyte", Description: "(optional) if you want to specify utxo per kbyte"})
					}
				}
			}
			return prompt.FilterContains(out, cur, true)
		}
		if len(args) == 6 {
			var out []prompt.Suggest
			if len(config.GCFGRegistry) > 0 {
				if val, ok := config.GCFGRegistry[args[1]]; ok {
					switch val.Type {
					case "BEP-20", "ERC-20", "QRC-20":
						out = append(out, prompt.Suggest{Text: "<gas_price>", Description: "specify gas_price for " + args[4]})
					case "UTXO", "Smart Chain":
						out = append(out, prompt.Suggest{Text: "<amount>", Description: "specify the utxo amount for " + args[4]})
					}
				}
			}
			return prompt.FilterContains(out, cur, true)
		}
		if len(args) == 7 {
			var out []prompt.Suggest
			if len(config.GCFGRegistry) > 0 {
				if val, ok := config.GCFGRegistry[args[1]]; ok {
					switch val.Type {
					case "BEP-20", "ERC-20":
						out = append(out, prompt.Suggest{Text: "<gas>", Description: "specify gas " + args[4]})
					case "QRC-20":
						out = append(out, prompt.Suggest{Text: "<gas_limit>", Description: "specify the gas limit for " + args[4]})
					}
				}
			}
			return prompt.FilterContains(out, cur, true)
		}
	}
	return []prompt.Suggest{}
}

func (c *Completer) Complete(d prompt.Document) []prompt.Suggest {
	if d.TextBeforeCursor() == "" {
		return []prompt.Suggest{}
	}
	args := strings.Split(d.TextBeforeCursor(), " ")
	return c.argumentsCompleter(args)
}
