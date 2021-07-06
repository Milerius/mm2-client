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
	{Text: "orderbook", Description: "Show the orderbook of the given pair"},
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
}

var subCommandsEnable = []prompt.Suggest{
	{Text: "SOULJA", Description: "Enable SOULJA"},
	{Text: "DX", Description: "Enable DX"},
	{Text: "LRC", Description: "Enable LRC"},
	{Text: "VAL", Description: "Enable VAL"},
	{Text: "XVC-BEP20", Description: "Enable XVC-BEP20"},
	{Text: "POWR", Description: "Enable POWR"},
	{Text: "BIDR-BEP20", Description: "Enable BIDR-BEP20"},
	{Text: "tBTC-TEST", Description: "Enable tBTC-TEST"},
	{Text: "TSL", Description: "Enable TSL"},
	{Text: "CEL", Description: "Enable CEL"},
	{Text: "ARPA", Description: "Enable ARPA"},
	{Text: "MSHARK", Description: "Enable MSHARK"},
	{Text: "IL8P", Description: "Enable IL8P"},
	{Text: "VRA", Description: "Enable VRA"},
	{Text: "PUT", Description: "Enable PUT"},
	{Text: "BAND-BEP20", Description: "Enable BAND-BEP20"},
	{Text: "BAT-BEP20", Description: "Enable BAT-BEP20"},
	{Text: "MKR-BEP20", Description: "Enable MKR-BEP20"},
	{Text: "CENNZ", Description: "Enable CENNZ"},
	{Text: "PNK", Description: "Enable PNK"},
	{Text: "BAND-ERC20", Description: "Enable BAND-ERC20"},
	{Text: "EOS-BEP20", Description: "Enable EOS-BEP20"},
	{Text: "RVN", Description: "Enable RVN"},
	{Text: "SKL", Description: "Enable SKL"},
	{Text: "TRAC", Description: "Enable TRAC"},
	{Text: "ANKR-BEP20", Description: "Enable ANKR-BEP20"},
	{Text: "CCL", Description: "Enable CCL"},
	{Text: "FJC", Description: "Enable FJC"},
	{Text: "HLC", Description: "Enable HLC"},
	{Text: "SNT", Description: "Enable SNT"},
	{Text: "1INCH-ERC20", Description: "Enable 1INCH-ERC20"},
	{Text: "AVAX-BEP20", Description: "Enable AVAX-BEP20"},
	{Text: "COQUI", Description: "Enable COQUI"},
	{Text: "UBT", Description: "Enable UBT"},
	{Text: "UMA", Description: "Enable UMA"},
	{Text: "BNBT", Description: "Enable BNBT"},
	{Text: "JUMBLR", Description: "Enable JUMBLR"},
	{Text: "UOS", Description: "Enable UOS"},
	{Text: "BAL-BEP20", Description: "Enable BAL-BEP20"},
	{Text: "SNX-ERC20", Description: "Enable SNX-ERC20"},
	{Text: "SRM", Description: "Enable SRM"},
	{Text: "AXE", Description: "Enable AXE"},
	{Text: "CDN", Description: "Enable CDN"},
	{Text: "ECA", Description: "Enable ECA"},
	{Text: "AWC", Description: "Enable AWC"},
	{Text: "MCL", Description: "Enable MCL"},
	{Text: "XLM-BEP20", Description: "Enable XLM-BEP20"},
	{Text: "CHZ", Description: "Enable CHZ"},
	{Text: "MANA", Description: "Enable MANA"},
	{Text: "QNT", Description: "Enable QNT"},
	{Text: "VGX", Description: "Enable VGX"},
	{Text: "BET", Description: "Enable BET"},
	{Text: "SUSHI-ERC20", Description: "Enable SUSHI-ERC20"},
	{Text: "MLN", Description: "Enable MLN"},
	{Text: "XRP-BEP20", Description: "Enable XRP-BEP20"},
	{Text: "BAL-ERC20", Description: "Enable BAL-ERC20"},
	{Text: "DGC", Description: "Enable DGC"},
	{Text: "LYNX", Description: "Enable LYNX"},
	{Text: "OCEAN-BEP20", Description: "Enable OCEAN-BEP20"},
	{Text: "QIAIR", Description: "Enable QIAIR"},
	{Text: "NVC", Description: "Enable NVC"},
	{Text: "BTCH", Description: "Enable BTCH"},
	{Text: "BOTS", Description: "Enable BOTS"},
	{Text: "PANGEA", Description: "Enable PANGEA"},
	{Text: "INK", Description: "Enable INK"},
	{Text: "VRM", Description: "Enable VRM"},
	{Text: "ZILLA", Description: "Enable ZILLA"},
	{Text: "BTU", Description: "Enable BTU"},
	{Text: "FIL-BEP20", Description: "Enable FIL-BEP20"},
	{Text: "JSTR", Description: "Enable JSTR"},
	{Text: "LTC", Description: "Enable LTC"},
	{Text: "VOTE2021", Description: "Enable VOTE2021"},
	{Text: "XVS", Description: "Enable XVS"},
	{Text: "SUSHI-BEP20", Description: "Enable SUSHI-BEP20"},
	{Text: "ILN", Description: "Enable ILN"},
	{Text: "BEST", Description: "Enable BEST"},
	{Text: "CVC", Description: "Enable CVC"},
	{Text: "USDT-ERC20", Description: "Enable USDT-ERC20"},
	{Text: "AAVE-ERC20", Description: "Enable AAVE-ERC20"},
	{Text: "IOTA-BEP20", Description: "Enable IOTA-BEP20"},
	{Text: "RICK", Description: "Enable RICK"},
	{Text: "WLC", Description: "Enable WLC"},
	{Text: "tQTUM", Description: "Enable tQTUM"},
	{Text: "KNC", Description: "Enable KNC"},
	{Text: "DODO-ERC20", Description: "Enable DODO-ERC20"},
	{Text: "EMC2", Description: "Enable EMC2"},
	{Text: "TRYB-ERC20", Description: "Enable TRYB-ERC20"},
	{Text: "BTC-BEP20", Description: "Enable BTC-BEP20"},
	{Text: "BUSD-ERC20", Description: "Enable BUSD-ERC20"},
	{Text: "DOGE", Description: "Enable DOGE"},
	{Text: "EGLD-BEP20", Description: "Enable EGLD-BEP20"},
	{Text: "DIA", Description: "Enable DIA"},
	{Text: "GNO", Description: "Enable GNO"},
	{Text: "DAI-ERC20", Description: "Enable DAI-ERC20"},
	{Text: "UNO", Description: "Enable UNO"},
	{Text: "USDC-ERC20", Description: "Enable USDC-ERC20"},
	{Text: "KMD-BEP20", Description: "Enable KMD-BEP20"},
	{Text: "MORTY", Description: "Enable MORTY"},
	{Text: "XMY", Description: "Enable XMY"},
	{Text: "ETH-BEP20", Description: "Enable ETH-BEP20"},
	{Text: "CHSB", Description: "Enable CHSB"},
	{Text: "ADX-BEP20", Description: "Enable ADX-BEP20"},
	{Text: "DIMI", Description: "Enable DIMI"},
	{Text: "DP", Description: "Enable DP"},
	{Text: "ONT-BEP20", Description: "Enable ONT-BEP20"},
	{Text: "CRO", Description: "Enable CRO"},
	{Text: "HEX", Description: "Enable HEX"},
	{Text: "EURS", Description: "Enable EURS"},
	{Text: "GLEEC", Description: "Enable GLEEC"},
	{Text: "KOIN", Description: "Enable KOIN"},
	{Text: "PAXG-ERC20", Description: "Enable PAXG-ERC20"},
	{Text: "LINK-BEP20", Description: "Enable LINK-BEP20"},
	{Text: "TRYB-BEP20", Description: "Enable TRYB-BEP20"},
	{Text: "ENJ", Description: "Enable ENJ"},
	{Text: "ADX-ERC20", Description: "Enable ADX-ERC20"},
	{Text: "WSB", Description: "Enable WSB"},
	{Text: "OCEAN-ERC20", Description: "Enable OCEAN-ERC20"},
	{Text: "PAXG-BEP20", Description: "Enable PAXG-BEP20"},
	{Text: "TUSD-ERC20", Description: "Enable TUSD-ERC20"},
	{Text: "TTT", Description: "Enable TTT"},
	{Text: "UNI-ERC20", Description: "Enable UNI-ERC20"},
	{Text: "ELF-BEP20", Description: "Enable ELF-BEP20"},
	{Text: "ETH", Description: "Enable ETH"},
	{Text: "1INCH-BEP20", Description: "Enable 1INCH-BEP20"},
	{Text: "NEAR-BEP20", Description: "Enable NEAR-BEP20"},
	{Text: "HT", Description: "Enable HT"},
	{Text: "ATOM-BEP20", Description: "Enable ATOM-BEP20"},
	{Text: "AUR", Description: "Enable AUR"},
	{Text: "YFII-BEP20", Description: "Enable YFII-BEP20"},
	{Text: "HPY", Description: "Enable HPY"},
	{Text: "OKB", Description: "Enable OKB"},
	{Text: "UQC", Description: "Enable UQC"},
	{Text: "ABY", Description: "Enable ABY"},
	{Text: "NMC", Description: "Enable NMC"},
	{Text: "QKC-BEP20", Description: "Enable QKC-BEP20"},
	{Text: "LSTR", Description: "Enable LSTR"},
	{Text: "BNB", Description: "Enable BNB"},
	{Text: "QTUM", Description: "Enable QTUM"},
	{Text: "RLC", Description: "Enable RLC"},
	{Text: "STORJ", Description: "Enable STORJ"},
	{Text: "LCC", Description: "Enable LCC"},
	{Text: "S4F", Description: "Enable S4F"},
	{Text: "WBTC", Description: "Enable WBTC"},
	{Text: "CVT", Description: "Enable CVT"},
	{Text: "FET", Description: "Enable FET"},
	{Text: "FTM-BEP20", Description: "Enable FTM-BEP20"},
	{Text: "BAT-ERC20", Description: "Enable BAT-ERC20"},
	{Text: "SMTF", Description: "Enable SMTF"},
	{Text: "WCN", Description: "Enable WCN"},
	{Text: "ZER", Description: "Enable ZER"},
	{Text: "SPC", Description: "Enable SPC"},
	{Text: "BTC", Description: "Enable BTC"},
	{Text: "EFL", Description: "Enable EFL"},
	{Text: "SXP-BEP20", Description: "Enable SXP-BEP20"},
	{Text: "DOT-BEP20", Description: "Enable DOT-BEP20"},
	{Text: "PAX-BEP20", Description: "Enable PAX-BEP20"},
	{Text: "THC", Description: "Enable THC"},
	{Text: "CRYPTO", Description: "Enable CRYPTO"},
	{Text: "IOTX-BEP20", Description: "Enable IOTX-BEP20"},
	{Text: "ZEC", Description: "Enable ZEC"},
	{Text: "HOT", Description: "Enable HOT"},
	{Text: "REP", Description: "Enable REP"},
	{Text: "BUSD-BEP20", Description: "Enable BUSD-BEP20"},
	{Text: "JRT-ERC20", Description: "Enable JRT-ERC20"},
	{Text: "KMD", Description: "Enable KMD"},
	{Text: "XVC", Description: "Enable XVC"},
	{Text: "SFUSD", Description: "Enable SFUSD"},
	{Text: "REN", Description: "Enable REN"},
	{Text: "BTCZ", Description: "Enable BTCZ"},
	{Text: "CAKE", Description: "Enable CAKE"},
	{Text: "DODO-BEP20", Description: "Enable DODO-BEP20"},
	{Text: "PAX-ERC20", Description: "Enable PAX-ERC20"},
	{Text: "SPACE", Description: "Enable SPACE"},
	{Text: "CRV", Description: "Enable CRV"},
	{Text: "COMP-ERC20", Description: "Enable COMP-ERC20"},
	{Text: "ANT", Description: "Enable ANT"},
	{Text: "DGB", Description: "Enable DGB"},
	{Text: "GLEEC-OLD", Description: "Enable GLEEC-OLD"},
	{Text: "CHIPS", Description: "Enable CHIPS"},
	{Text: "YFI-BEP20", Description: "Enable YFI-BEP20"},
	{Text: "QBT", Description: "Enable QBT"},
	{Text: "UNI-BEP20", Description: "Enable UNI-BEP20"},
	{Text: "FTC", Description: "Enable FTC"},
	{Text: "ETHR", Description: "Enable ETHR"},
	{Text: "HUSD", Description: "Enable HUSD"},
	{Text: "BNT-ERC20", Description: "Enable BNT-ERC20"},
	{Text: "BCH", Description: "Enable BCH"},
	{Text: "TRC", Description: "Enable TRC"},
	{Text: "ANKR-ERC20", Description: "Enable ANKR-ERC20"},
	{Text: "LEO", Description: "Enable LEO"},
	{Text: "ZRX", Description: "Enable ZRX"},
	{Text: "BABYDOGE", Description: "Enable BABYDOGE"},
	{Text: "SNX-BEP20", Description: "Enable SNX-BEP20"},
	{Text: "UTK", Description: "Enable UTK"},
	{Text: "AAVE-BEP20", Description: "Enable AAVE-BEP20"},
	{Text: "QI", Description: "Enable QI"},
	{Text: "SCA", Description: "Enable SCA"},
	{Text: "FTM-ERC20", Description: "Enable FTM-ERC20"},
	{Text: "LINK-ERC20", Description: "Enable LINK-ERC20"},
	{Text: "REVS", Description: "Enable REVS"},
	{Text: "OC", Description: "Enable OC"},
	{Text: "TMTG", Description: "Enable TMTG"},
	{Text: "FIRO-BEP20", Description: "Enable FIRO-BEP20"},
	{Text: "YFII-ERC20", Description: "Enable YFII-ERC20"},
	{Text: "DASH", Description: "Enable DASH"},
	{Text: "ELF-ERC20", Description: "Enable ELF-ERC20"},
	{Text: "ETC-BEP20", Description: "Enable ETC-BEP20"},
	{Text: "DEX", Description: "Enable DEX"},
	{Text: "CIPHS", Description: "Enable CIPHS"},
	{Text: "YFI-ERC20", Description: "Enable YFI-ERC20"},
	{Text: "ZET", Description: "Enable ZET"},
	{Text: "RSR", Description: "Enable RSR"},
	{Text: "BLK", Description: "Enable BLK"},
	{Text: "GRS", Description: "Enable GRS"},
	{Text: "SYS", Description: "Enable SYS"},
	{Text: "VRSC", Description: "Enable VRSC"},
	{Text: "ZIL-BEP20", Description: "Enable ZIL-BEP20"},
	{Text: "QC", Description: "Enable QC"},
	{Text: "COMP-BEP20", Description: "Enable COMP-BEP20"},
	{Text: "MONA", Description: "Enable MONA"},
	{Text: "NAV", Description: "Enable NAV"},
	{Text: "MATIC-ERC20", Description: "Enable MATIC-ERC20"},
	{Text: "UIS", Description: "Enable UIS"},
	{Text: "XOR", Description: "Enable XOR"},
	{Text: "TRX-BEP20", Description: "Enable TRX-BEP20"},
	{Text: "BSTY", Description: "Enable BSTY"},
	{Text: "DAI-BEP20", Description: "Enable DAI-BEP20"},
	{Text: "MATIC-BEP20", Description: "Enable MATIC-BEP20"},
	{Text: "USDC-BEP20", Description: "Enable USDC-BEP20"},
	{Text: "SXP-ERC20", Description: "Enable SXP-ERC20"},
	{Text: "MKR-ERC20", Description: "Enable MKR-ERC20"},
	{Text: "USDT-BEP20", Description: "Enable USDT-BEP20"},
	{Text: "QRC20", Description: "Enable QRC20"},
	{Text: "BTT-BEP20", Description: "Enable BTT-BEP20"},
	{Text: "DOGE-BEP20", Description: "Enable DOGE-BEP20"},
	{Text: "STFIRO", Description: "Enable STFIRO"},
	{Text: "OOT", Description: "Enable OOT"},
	{Text: "XPM", Description: "Enable XPM"},
	{Text: "BNT-BEP20", Description: "Enable BNT-BEP20"},
	{Text: "FIRO", Description: "Enable FIRO"},
	{Text: "QKC-ERC20", Description: "Enable QKC-ERC20"},
	{Text: "HODL", Description: "Enable HODL"},
	{Text: "LABS", Description: "Enable LABS"},
	{Text: "MESH", Description: "Enable MESH"},
	{Text: "XTZ-BEP20", Description: "Enable XTZ-BEP20"},
	{Text: "SHR", Description: "Enable SHR"},
	{Text: "CLC", Description: "Enable CLC"},
	{Text: "MGW", Description: "Enable MGW"},
	{Text: "SUPERNET", Description: "Enable SUPERNET"},
	{Text: "ADA-BEP20", Description: "Enable ADA-BEP20"},
	{Text: "BTE", Description: "Enable BTE"},
	{Text: "TUSD-BEP20", Description: "Enable TUSD-BEP20"},
	{Text: "AGI", Description: "Enable AGI"},
	{Text: "REV", Description: "Enable REV"},
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
