package config

import (
	"encoding/json"
	"fmt"
	"github.com/kpango/glg"
	"io/ioutil"
	"mm2_client/helpers"
	"os"
	"strconv"
)

type MM2Config struct {
	Dbdir       *string    `json:"dbdir,omitempty"`
	Gui         string     `json:"gui"`
	MM2         *int       `json:"mm2,omitempty"`
	Netid       int        `json:"netid"`
	Passphrase  string     `json:"passphrase"`
	RPCPassword string     `json:"rpc_password"`
	Userhome    *string    `json:"userhome,omitempty"`
	IMASeed     bool       `json:"im_a_seed,omitempty"`
	Coins       *[]*MM2CFG `json:"coins,omitempty"`
	SeedNodes   *[]string  `json:"seednodes,omitempty"`
}

func NewMM2ConfigFromFile(targetPath string) *MM2Config {
	cfg := &MM2Config{}
	file, _ := ioutil.ReadFile(targetPath)
	_ = json.Unmarshal([]byte(file), cfg)
	return cfg
}

func NewMM2Config() *MM2Config {
	dbdir := os.Getenv("HOME") + "/atomicdex_cli/mm2/db"
	userhome := os.Getenv("HOME")
	return &MM2Config{
		Dbdir:       &dbdir,
		Gui:         "AtomicDEX Client CLI",
		Netid:       7777,
		Passphrase:  "your_seed",
		RPCPassword: "your_rpc_password",
		Userhome:    &userhome}
}

func NewMM2ConfigWasm(userpass string, passphrase string, extraArgs []string) string {
	mm2 := 1
	cfg := &MM2Config{
		MM2:         &mm2,
		Gui:         "AtomicDEX Webassembly",
		Netid:       7777,
		Passphrase:  passphrase,
		RPCPassword: userpass,
		Coins:       &GMM2CFGArray}
	if len(extraArgs) > 0 {
		value, err := strconv.Atoi(extraArgs[0])
		if err == nil {
			cfg.Netid = value
		} else {
			glg.Errorf("err atoi: %v", err)
		}
		if len(extraArgs) > 1 {
			var seednodes = extraArgs[1:]
			cfg.SeedNodes = &seednodes
		}
	}
	cfgCopy := *cfg
	cfgCopy.Coins = nil
	cfgCopy.Passphrase = "*******"
	cfgCopy.RPCPassword = "*******"
	glg.Infof("extraArgs: %s", extraArgs)
	glg.Infof("cfg: %s", cfgCopy.ToJson())
	return cfg.ToJson()
}

func (cfg *MM2Config) ToJson() string {
	b, err := json.MarshalIndent(cfg, "", "    ")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}

func (cfg *MM2Config) WriteToFile() {
	targetDir := helpers.GetWorkingDir() + "/mm2"
	targetPath := targetDir + "/MM2.json"
	if helpers.FileExists(targetPath) {
		_ = os.Remove(targetPath)
	}
	_ = ioutil.WriteFile(targetPath, []byte(cfg.ToJson()), 0644)
}
