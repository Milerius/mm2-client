package config

import (
	"encoding/json"
	"io/ioutil"
	"mm2_client/constants"
	"mm2_client/helpers"
	"strconv"
)

type DesktopCFG struct {
	Coin          string `json:"coin"`
	Name          string `json:"name"`
	CoinpaprikaID string `json:"coinpaprika_id"`
	CoingeckoID   string `json:"coingecko_id"`
	Electrum      []struct {
		URL                     string `json:"url"`
		Protocol                string `json:"protocol"`
		DisableCertVerification bool   `json:"disable_cert_verification"`
	} `json:"electrum,omitempty"`
	Nodes              []string `json:"nodes,omitempty"`
	ExplorerURL        []string `json:"explorer_url"`
	ExplorerTxURL      string   `json:"explorer_tx_url,omitempty"`
	ExplorerAddressURL string   `json:"explorer_address_url,omitempty"`
	Type               string   `json:"type"`
	Active             bool     `json:"active"`
	CurrentlyEnabled   bool     `json:"currently_enabled"`
	IsClaimable        bool     `json:"is_claimable,omitempty"`
	WalletOnly         bool     `json:"wallet_only,omitempty"`
	IsTestNet          bool     `json:"is_testnet,omitempty"`
}

const (
	GasStation                            = "https://ethgasstation.info/json/ethgasAPI.json"
	ErcSwapContractAddress                = "0x24ABE4c71FC658C91313b6552cd40cD808b3Ea80"
	ErcTestnetSwapContractAddress         = "0x6b5A52217006B965BB190864D62dc3d270F7AaFD"
	BnbTestnetSwapContractAddress         = "0xcCD17C913aD7b772755Ad4F0BDFF7B34C6339150"
	BnbSwapContractAddress                = "0xeDc5b89Fe1f0382F9E4316069971D90a0951DB31"
	ErcFallbackSwapContractAddress        = "0x8500AFc0bc5214728082163326C2FF0C73f4a871"
	ErcTestnetFallbackSwapContractAddress = "0x7Bc1bBDD6A0a722fC9bffC49c921B685ECB84b94"
	BnbFallbackSwapContractAddress        = BnbSwapContractAddress
	BnbTestnetFallbackSwapContractAddress = BnbTestnetSwapContractAddress
)

var GCFGRegistry = make(map[string]*DesktopCFG)

func ParseDesktopRegistry(version string) {
	var desktopCoinsPath = constants.GMM2Dir + "/" + version + "-coins.json"
	file, _ := ioutil.ReadFile(desktopCoinsPath)
	_ = json.Unmarshal([]byte(file), &GCFGRegistry)
	helpers.PrintCheck("Successfully load desktop cfg with "+strconv.Itoa(len(GCFGRegistry))+" coins", true)
}

func (cfg *DesktopCFG) RetrieveContracts() (string, string) {
	switch cfg.Type {
	case "BEP-20":
		if cfg.IsTestNet {
			return BnbTestnetSwapContractAddress, BnbTestnetFallbackSwapContractAddress
		} else {
			return BnbSwapContractAddress, BnbFallbackSwapContractAddress
		}
	case "ERC-20":
		if cfg.IsTestNet {
			return ErcTestnetSwapContractAddress, ErcTestnetFallbackSwapContractAddress
		} else {
			return ErcSwapContractAddress, ErcFallbackSwapContractAddress
		}
	default:
		return "", ""
	}
}
