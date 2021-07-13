package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mm2_client/constants"
	"mm2_client/helpers"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

type ElectrumData struct {
	URL                     string  `json:"url"`
	Protocol                *string `json:"protocol,omitempty"`
	DisableCertVerification *bool   `json:"disable_cert_verification,omitempty"`
}

type DesktopCFG struct {
	Coin               string         `json:"coin"`
	Name               string         `json:"name"`
	CoinpaprikaID      string         `json:"coinpaprika_id"`
	CoingeckoID        string         `json:"coingecko_id"`
	Electrum           []ElectrumData `json:"electrum,omitempty"`
	Nodes              []string       `json:"nodes,omitempty"`
	ExplorerURL        []string       `json:"explorer_url"`
	ExplorerTxURL      string         `json:"explorer_tx_url,omitempty"`
	ExplorerAddressURL string         `json:"explorer_address_url,omitempty"`
	Type               string         `json:"type"`
	Active             bool           `json:"active"`
	CurrentlyEnabled   bool           `json:"currently_enabled"`
	IsClaimable        bool           `json:"is_claimable,omitempty"`
	WalletOnly         bool           `json:"wallet_only,omitempty"`
	IsTestNet          bool           `json:"is_testnet,omitempty"`
}

const (
	GasStation                            = "https://ethgasstation.info/json/ethgasAPI.json"
	ErcSwapContractAddress                = "0x24ABE4c71FC658C91313b6552cd40cD808b3Ea80"
	ErcTestnetSwapContractAddress         = "0x6b5A52217006B965BB190864D62dc3d270F7AaFD"
	BnbTestnetSwapContractAddress         = "0xcCD17C913aD7b772755Ad4F0BDFF7B34C6339150"
	BnbSwapContractAddress                = "0xeDc5b89Fe1f0382F9E4316069971D90a0951DB31"
	ErcFallbackSwapContractAddress        = "0x8500AFc0bc5214728082163326C2FF0C73f4a871"
	ErcTestnetFallbackSwapContractAddress = "0x7Bc1bBDD6A0a722fC9bffC49c921B685ECB84b94"
	QrcTestnetSwapContractAddress         = "0xba8b71f3544b93e2f681f996da519a98ace0107a"
	QrcSwapContractAddress                = "0x2f754733acd6d753731c00fee32cb484551cc15d"
	BnbFallbackSwapContractAddress        = BnbSwapContractAddress
	BnbTestnetFallbackSwapContractAddress = BnbTestnetSwapContractAddress
	QrcTestnetFallbackSwapContractAddress = QrcTestnetSwapContractAddress
	QrcFallbackSwapContractAddress        = QrcSwapContractAddress
)

var GCFGRegistry = make(map[string]*DesktopCFG)

func GetDesktopDB() string {
	switch runtime.GOOS {
	case "linux":
		return os.Getenv("HOME") + "/atomic_qt/mm2/DB"
	case "darwin":
		return os.Getenv("HOME") + "/Library/Application Support/AtomicDex Desktop/mm2/DB"
	case "windows":
		return os.Getenv("APPDATA") + "/atomic_qt/mm2/DB"
	default:
		return os.Getenv("HOME") + "atomic_qt/mm2/DB"
	}
}

func GetDesktopPath(appName string) string {
	if appName == "standard" {
		switch runtime.GOOS {
		case "linux":
			return filepath.Join(os.Getenv("HOME"), "atomic_qt")
		case "darwin":
			return filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "AtomicDex Desktop")
		case "windows":
			return filepath.Join(os.Getenv("APPDATA"), "atomic_qt")
		default:
			return filepath.Join(os.Getenv("HOME"), "atomic_qt")
		}
	} else {
		switch runtime.GOOS {
		case "linux":
			return filepath.Join(os.Getenv("HOME"), appName)
		case "darwin":
			return filepath.Join(os.Getenv("HOME"), "Library", "Application Support", appName)
		case "windows":
			return filepath.Join(os.Getenv("APPDATA"), appName)
		default:
			return filepath.Join(os.Getenv("HOME"), appName)
		}
	}
}

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
	case "QRC-20":
		if cfg.IsTestNet {
			return QrcTestnetSwapContractAddress, QrcTestnetFallbackSwapContractAddress
		} else {
			return QrcSwapContractAddress, QrcFallbackSwapContractAddress
		}
	default:
		return "", ""
	}
}

func (cfg *DesktopCFG) RetrieveElectrums() []ElectrumData {
	switch cfg.Type {
	case "QRC-20":
		if cfg.IsTestNet {
			if val, ok := GCFGRegistry["tQTUM"]; ok {
				return val.Electrum
			}
		} else {
			if val, ok := GCFGRegistry["QTUM"]; ok {
				return val.Electrum
			}
		}
	default:
		return cfg.Electrum
	}
	return cfg.Electrum
}

func Update(version string) {
	//fmt.Println("Updating cfg")
	var desktopCoinsPath = constants.GMM2Dir + "/" + version + "-coins.json"
	e := os.Remove(desktopCoinsPath)
	if e != nil {
		fmt.Printf("Err: %v", e)
	} else {
		file, _ := json.MarshalIndent(GCFGRegistry, "", " ")
		err := ioutil.WriteFile(desktopCoinsPath, file, 0644)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func RetrieveActiveCoins() []string {
	var out []string
	for _, value := range GCFGRegistry {
		if value.Active {
			out = append(out, value.Coin)
		}
	}
	return out
}

func RetrieveAllCoins() []string {
	var out []string
	for _, value := range GCFGRegistry {
		out = append(out, value.Coin)
	}
	return out
}
