package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kpango/glg"
	"io"
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
	WSURL                   *string `json:"ws_url,omitempty"`
	Protocol                *string `json:"protocol,omitempty"`
	DisableCertVerification *bool   `json:"disable_cert_verification,omitempty"`
}

type NodesData struct {
	URL                     string  `json:"url"`
	GuiAuth                 *bool   `json:"gui_auth,omitempty"`
}

type DesktopCFG struct {
	Coin               string         `json:"coin"`
	Name               string         `json:"name"`
	AliasTicker        *string        `json:"alias_ticker,omitempty"`
	CoinpaprikaID      string         `json:"coinpaprika_id"`
	CoingeckoID        string         `json:"coingecko_id"`
	NomicsId           *string        `json:"nomics_id,omitempty"`
	ForexId            *string        `json:"forex_id,omitempty"`
	Electrum           []ElectrumData `json:"electrum,omitempty"`
	Nodes              []NodesData    `json:"nodes,omitempty"`
	ExplorerURL        string         `json:"explorer_url"`
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
	GasStationErc20                        = "https://ethgasstation.info/json/ethgasAPI.json"
	GasStationMatic                        = "https://gasstation-mainnet.matic.network"
	ArbitrumContractAddress                = "0x9130b257d37a52e52f21054c4da3450c72f595ce"
	ArbitrumFallbackContractAddress        = ArbitrumContractAddress
	ArbitrumTestnetContractAddress         = "0x9130b257d37a52e52f21054c4da3450c72f595ce"
	ArbitrumTestnetFallbackContractAddress = ArbitrumTestnetContractAddress
	AvaxContractAddress                    = "0x9130b257d37a52e52f21054c4da3450c72f595ce"
	AvaxFallbackContractAddress            = AvaxContractAddress
	AvaxTestnetContractAddress             = "0x9130b257d37a52e52f21054c4da3450c72f595ce"
	AvaxTestnetFallbackContractAddress     = AvaxTestnetContractAddress
	BnbTestnetSwapContractAddress          = "0xcCD17C913aD7b772755Ad4F0BDFF7B34C6339150"
	BnbSwapContractAddress                 = "0xeDc5b89Fe1f0382F9E4316069971D90a0951DB31"
	BnbFallbackSwapContractAddress         = BnbSwapContractAddress
	BnbTestnetFallbackSwapContractAddress  = BnbTestnetSwapContractAddress
	ErcSwapContractAddress                 = "0x24ABE4c71FC658C91313b6552cd40cD808b3Ea80"
	ErcTestnetSwapContractAddress          = "0x6b5A52217006B965BB190864D62dc3d270F7AaFD"
	ErcFallbackSwapContractAddress         = "0x8500AFc0bc5214728082163326C2FF0C73f4a871"
	ErcTestnetFallbackSwapContractAddress  = "0x7Bc1bBDD6A0a722fC9bffC49c921B685ECB84b94"
	EtcContractAddress                     = "0x6d9ce4BD298DE38bAfEFD15f5C6f5c95313B1d94"
	EtcFallbackContractAddress             = EtcContractAddress
	EtcTestnetContractAddress              = "0x6d9ce4BD298DE38bAfEFD15f5C6f5c95313B1d94"
	EtcTestnetFallbackContractAddress      = EtcTestnetContractAddress
	FtmContractAddress                     = "0x9130b257d37a52e52f21054c4da3450c72f595ce"
	FtmFallbackContractAddress             = FtmContractAddress
	FtmTestnetContractAddress              = "0x9130b257d37a52e52f21054c4da3450c72f595ce"
	FtmTestnetFallbackContractAddress      = FtmTestnetContractAddress
	HcoContractAddress                     = "0x9130b257d37a52e52f21054c4da3450c72f595ce"
	HcoFallbackContractAddress             = HcoContractAddress
	HcoTestnetContractAddress              = "0x9130b257d37a52e52f21054c4da3450c72f595ce"
	HcoTestnetFallbackContractAddress      = HcoTestnetContractAddress
	KrcContractAddress                     = "0x9130b257d37a52e52f21054c4da3450c72f595ce"
	KrcFallbackContractAddress             = KrcContractAddress
	KrcTestnetContractAddress              = "0x9130b257d37a52e52f21054c4da3450c72f595ce"
	KrcTestnetFallbackContractAddress      = KrcTestnetContractAddress
	MaticContractAddress                   = "0x9130b257d37a52e52f21054c4da3450c72f595ce"
	MaticFallbackContractAddress           = MaticContractAddress
	MaticTestnetContractAddress            = "0x73c1Dd989218c3A154C71Fc08Eb55A24Bd2B3A10"
	MaticTestnetFallbackContractAddress    = MaticTestnetContractAddress
	MovrContractAddress                    = "0x9130b257d37a52e52f21054c4da3450c72f595ce"
	MovrFallbackContractAddress            = MovrContractAddress
	MovrTestnetContractAddress             = "0x9130b257d37a52e52f21054c4da3450c72f595ce"
	MovrTestnetFallbackContractAddress     = MovrTestnetContractAddress
	GlmrContractAddress                    = "0x6d9ce4BD298DE38bAfEFD15f5C6f5c95313B1d94"
	GlmrFallbackContractAddress            = MovrContractAddress
	GlmrTestnetContractAddress             = "0x6d9ce4BD298DE38bAfEFD15f5C6f5c95313B1d94"
	GlmrTestnetFallbackContractAddress     = MovrTestnetContractAddress
	OneContractAddress                     = "0x9130b257d37a52e52f21054c4da3450c72f595ce"
	OneFallbackContractAddress             = OneContractAddress
	OneTestnetContractAddress              = "0x9130b257d37a52e52f21054c4da3450c72f595ce"
	OneTestnetFallbackContractAddress      = OneTestnetContractAddress
	OptimismContractAddress                = "0x9130b257d37a52e52f21054c4da3450c72f595ce"
	OptimismFallbackContractAddress        = OptimismContractAddress
	OptimismTestnetContractAddress         = "0x9130b257d37a52e52f21054c4da3450c72f595ce"
	OptimismTestnetFallbackContractAddress = OptimismTestnetContractAddress
	QrcTestnetSwapContractAddress          = "0xba8b71f3544b93e2f681f996da519a98ace0107a"
	QrcSwapContractAddress                 = "0x2f754733acd6d753731c00fee32cb484551cc15d"
	QrcTestnetFallbackSwapContractAddress  = QrcTestnetSwapContractAddress
	QrcFallbackSwapContractAddress         = QrcSwapContractAddress
	SmartBchContractAddress                = "0x25bF2AAB8749AD2e4360b3e0B738f3Cd700C4D68"
	SmartBchFallbackContractAddress        = SmartBchContractAddress
	SmartBchTestnetContractAddress         = "0x25bF2AAB8749AD2e4360b3e0B738f3Cd700C4D68"
	SmartBchTestnetFallbackContractAddress = SmartBchTestnetContractAddress
	UbiqContractAddress                    = "0x9130b257d37a52e52f21054c4da3450c72f595ce"
	UbiqFallbackContractAddress            = UbiqContractAddress
	UbiqTestnetContractAddress             = "0x9130b257d37a52e52f21054c4da3450c72f595ce"
	UbiqTestnetFallbackContractAddress     = UbiqTestnetContractAddress
)

var GCFGRegistry = make(map[string]*DesktopCFG)

var GasStationMaticDecimals = 9
var GasStationErc20Decimals = 8

func GetDesktopDB() *string {
	result := os.Getenv("HOME") + "atomic_qt/mm2/DB"
	switch runtime.GOOS {
	case "linux":
		result = os.Getenv("HOME") + "/atomic_qt/mm2/DB"
	case "darwin":
		result = os.Getenv("HOME") + "/Library/Application Support/AtomicDex Desktop/mm2/DB"
	case "windows":
		result = os.Getenv("APPDATA") + "/atomic_qt/mm2/DB"
	default:
		result = os.Getenv("HOME") + "atomic_qt/mm2/DB"
	}
	return &result
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
	var desktopCoinsPath = constants.GMM2Dir + "/coins_config.json"
	file, _ := ioutil.ReadFile(desktopCoinsPath)
	err := json.Unmarshal([]byte(file), &GCFGRegistry)
	if err != nil {
		_ = glg.Errorf("Cannot parse cfg: %v", err)
	}
	helpers.PrintCheck("Successfully load desktop cfg with "+strconv.Itoa(len(GCFGRegistry))+" coins", true)
}

func ParseDesktopRegistryFromFile(path string) bool {
	if constants.GDesktopCfgLoaded {
		return true
	}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		_ = glg.Errorf("err when parsing: %v", err)
		return false
	}
	unmarshalErr := json.Unmarshal([]byte(file), &GCFGRegistry)
	if unmarshalErr != nil {
		_ = glg.Errorf("err when unmarshaling: %v", unmarshalErr)
		return false
	}
	if len(GCFGRegistry) > 0 {
		constants.GDesktopCfgLoaded = true
		return true
	}
	return false
}

func ParseDesktopRegistryFromString(cfg string) bool {
	if constants.GDesktopCfgLoaded {
		return true
	}
	_ = json.Unmarshal([]byte(cfg), &GCFGRegistry)
	if len(GCFGRegistry) > 0 {
		constants.GDesktopCfgLoaded = true
		return true
	}
	return false
}

func ParseDesktopRegistryFromUrl(url string) error {
	if constants.GDesktopCfgLoaded {
		return nil
	}

	resp, err := helpers.CrossGet(url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	decodeErr := json.NewDecoder(resp.Body).Decode(&GCFGRegistry)
	if decodeErr != nil {
		return decodeErr
	}

	if len(GCFGRegistry) > 0 {
		constants.GDesktopCfgLoaded = true
		return nil
	}
	return errors.New("unknown error")
}

func (cfg *DesktopCFG) RetrieveGasStationUrl() string {
	switch cfg.Type {
	case "ERC-20":
		return GasStationErc20
	case "Matic":
		return GasStationMatic
	}
	return ""
}

func (cfg *DesktopCFG) RetrieveGasStationDecimals() *int {
	switch cfg.Type {
	case "ERC-20":
		return &GasStationErc20Decimals
	case "Matic":
		return &GasStationMaticDecimals
	}
	return nil
}

func (cfg *DesktopCFG) RetrieveContracts() (string, string) {
	switch cfg.Type {
	case "Arbitrum":
		if cfg.IsTestNet {
			return ArbitrumTestnetContractAddress, ArbitrumTestnetFallbackContractAddress
		} else {
			return ArbitrumContractAddress, ArbitrumFallbackContractAddress
		}
	case "AVX-20":
		if cfg.IsTestNet {
			return AvaxTestnetContractAddress, AvaxTestnetFallbackContractAddress
		} else {
			return AvaxContractAddress, AvaxFallbackContractAddress
		}
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
	case "Ethereum Classic":
		if cfg.IsTestNet {
			return EtcTestnetContractAddress, EtcTestnetFallbackContractAddress
		} else {
			return EtcContractAddress, EtcFallbackContractAddress
		}
	case "SmartBCH":
		if cfg.IsTestNet {
			return SmartBchTestnetContractAddress, SmartBchTestnetFallbackContractAddress
		} else {
			return SmartBchContractAddress, SmartBchFallbackContractAddress
		}
	case "FTM-20":
		if cfg.IsTestNet {
			return FtmTestnetContractAddress, FtmTestnetFallbackContractAddress
		} else {
			return FtmContractAddress, FtmFallbackContractAddress
		}
	case "HecoChain":
		if cfg.IsTestNet {
			return HcoTestnetContractAddress, HcoTestnetFallbackContractAddress
		} else {
			return HcoContractAddress, HcoFallbackContractAddress
		}
	case "HRC-20":
		if cfg.IsTestNet {
			return OneTestnetContractAddress, OneTestnetFallbackContractAddress
		} else {
			return OneContractAddress, OneFallbackContractAddress
		}
	case "KRC-20":
		if cfg.IsTestNet {
			return KrcTestnetContractAddress, KrcTestnetFallbackContractAddress
		} else {
			return KrcContractAddress, KrcFallbackContractAddress
		}
	case "Matic":
		if cfg.IsTestNet {
			return MaticTestnetContractAddress, MaticTestnetFallbackContractAddress
		} else {
			return MaticContractAddress, MaticFallbackContractAddress
		}
	case "Moonriver":
		if cfg.IsTestNet {
			return MovrTestnetContractAddress, MovrTestnetFallbackContractAddress
		} else {
			return MovrContractAddress, MovrFallbackContractAddress
		}
	case "Moonbeam":
		if cfg.IsTestNet {
			return GlmrTestnetContractAddress, GlmrTestnetFallbackContractAddress
		} else {
			return GlmrContractAddress, GlmrFallbackContractAddress
		}
	case "Optimism":
		if cfg.IsTestNet {
			return OptimismTestnetContractAddress, OptimismTestnetFallbackContractAddress
		} else {
			return OptimismContractAddress, OptimismFallbackContractAddress
		}
	case "QRC-20":
		if cfg.IsTestNet {
			return QrcTestnetSwapContractAddress, QrcTestnetFallbackSwapContractAddress
		} else {
			return QrcSwapContractAddress, QrcFallbackSwapContractAddress
		}
	case "Ubiq":
		if cfg.IsTestNet {
			return UbiqTestnetContractAddress, UbiqTestnetFallbackContractAddress
		} else {
			return UbiqContractAddress, UbiqFallbackContractAddress
		}
	default:
		return "", ""
	}
}

func (cfg *DesktopCFG) RetrieveElectrums() []ElectrumData {
	functorElectrum := func(in []ElectrumData) []ElectrumData {
		if runtime.GOARCH != "wasm" {
			return in
		}
		var out []ElectrumData
		for _, cur := range in {
			curOut := ElectrumData{URL: cur.URL, WSURL: nil, DisableCertVerification: cur.DisableCertVerification, Protocol: cur.Protocol}
			if runtime.GOARCH == "wasm" {
				if cur.WSURL != nil {
					curOut.URL = *cur.WSURL
					protocol := "WSS"
					curOut.Protocol = &protocol
				} else {
					continue
				}
			}
			out = append(out, curOut)
		}
		return out
	}
	switch cfg.Type {
	case "QRC-20":
		if cfg.IsTestNet {
			if val, ok := GCFGRegistry["tQTUM"]; ok {
				return functorElectrum(val.Electrum)
			}
		} else {
			if val, ok := GCFGRegistry["QTUM"]; ok {
				return functorElectrum(val.Electrum)
			}
		}
	default:
		return functorElectrum(cfg.Electrum)
	}
	return functorElectrum(cfg.Electrum)
}

func Update(version string) {
	//fmt.Println("Updating cfg")
	_ = glg.Infof("Updating cfg")
	if runtime.GOARCH != "wasm" {
		var desktopCoinsPath = constants.GMM2Dir + "/coins_config.json"
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
	} else {
		UpdateWasm()
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
