package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mm2_client/helpers"
	"os"
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

func NewMM2ConfigWasm() string {
	mm2 := 1
	cfg := &MM2Config{
		MM2:         &mm2,
		Gui:         "AtomicDEX Client CLI",
		Netid:       7777,
		Passphrase:  "hardcoded_password",
		RPCPassword: "wasmtest",
		Coins:       &GMM2CFGArray}
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
