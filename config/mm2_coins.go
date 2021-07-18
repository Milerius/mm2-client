package config

import (
	"encoding/json"
	"errors"
	"github.com/kpango/glg"
	"io"
	"io/ioutil"
	"mm2_client/constants"
	"mm2_client/helpers"
	"strconv"
)

type TProtocolData struct {
	Platform        string `json:"platform"`
	ContractAddress string `json:"contract_address"`
}

type TAddressFormat struct {
	Format  string `json:"format"`
	Network string `json:"network"`
}

type MM2CFG struct {
	Coin                  string  `json:"coin"`
	Name                  string  `json:"name,omitempty"`
	Fname                 string  `json:"fname,omitempty"`
	Rpcport               int     `json:"rpcport,omitempty"`
	Mm2                   int     `json:"mm2,omitempty"`
	ChainId               int     `json:"chain_id,omitempty"`
	AvgBlocktime          float64 `json:"avg_blocktime,omitempty"`
	RequiredConfirmations int     `json:"required_confirmations,omitempty"`
	Protocol              struct {
		Type         string         `json:"type"`
		ProtocolData *TProtocolData `json:"protocol_data,omitempty"`
	} `json:"protocol"`
	Decimals             int             `json:"decimals,omitempty"`
	Pubtype              int             `json:"pubtype,omitempty"`
	P2Shtype             int             `json:"p2shtype,omitempty"`
	Wiftype              int             `json:"wiftype,omitempty"`
	Txfee                int             `json:"txfee,omitempty"`
	Segwit               bool            `json:"segwit,omitempty"`
	Bech32Hrp            string          `json:"bech32_hrp,omitempty"`
	Confpath             string          `json:"confpath,omitempty"`
	Asset                string          `json:"asset,omitempty"`
	Txversion            int             `json:"txversion,omitempty"`
	Overwintered         int             `json:"overwintered,omitempty"`
	RequiresNotarization bool            `json:"requires_notarization,omitempty"`
	Dust                 int             `json:"dust,omitempty"`
	EstimateFeeBlocks    int             `json:"estimate_fee_blocks,omitempty"`
	ForkId               string          `json:"fork_id,omitempty"`
	AddressFormat        *TAddressFormat `json:"address_format,omitempty"`
	IsPoS                int             `json:"isPoS,omitempty"`
	MatureConfirmations  int             `json:"mature_confirmations,omitempty"`
	EstimateFeeMode      string          `json:"estimate_fee_mode,omitempty"`
	Taddr                int             `json:"taddr,omitempty"`
	ForceMinRelayFee     bool            `json:"force_min_relay_fee,omitempty"`
	P2P                  int             `json:"p2p,omitempty"`
	Magic                string          `json:"magic,omitempty"`
	NSPV                 string          `json:"nSPV,omitempty"`
	VersionGroupId       string          `json:"version_group_id,omitempty"`
	ConsensusBranchId    string          `json:"consensus_branch_id,omitempty"`
	GuiCoin              string          `json:"gui_coin,omitempty"`
}

var GMM2CFGRegistry = make(map[string]*MM2CFG)
var GMM2CFGArray []*MM2CFG

func ParseMM2CFGRegistry() {
	file, _ := ioutil.ReadFile(constants.GMM2CoinsPath)
	_ = json.Unmarshal([]byte(file), &GMM2CFGArray)
	for _, v := range GMM2CFGArray {
		GMM2CFGRegistry[v.Coin] = v
	}
	helpers.PrintCheck("Successfully load MM2 cfg with "+strconv.Itoa(len(GMM2CFGRegistry))+" coins", true)
}

func ParseMM2CFGFromUrl(url string) error {
	resp, err := helpers.CrossGet(url)
	if err != nil {
		glg.Errorf("err cross-get")
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	decodeErr := json.NewDecoder(resp.Body).Decode(&GMM2CFGArray)
	if decodeErr != nil {
		return decodeErr
	}

	if len(GMM2CFGArray) > 0 {
		return nil
	}
	return errors.New("unknown error")
}

func ParseMM2CFGRegistryFromFile(path string) bool {
	file, _ := ioutil.ReadFile(path)
	err := json.Unmarshal(file, &GMM2CFGArray)
	if err != nil {
		return false
	}
	for _, v := range GMM2CFGArray {
		GMM2CFGRegistry[v.Coin] = v
	}
	return true
}

func RetrieveContractsInfo(coin string) string {
	if val, ok := GMM2CFGRegistry[coin]; ok {
		return val.Protocol.ProtocolData.ContractAddress
	}
	return ""
}
