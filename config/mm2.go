package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mm2_client/helpers"
	"os"
)

type MM2Config struct {
	Dbdir       string `json:"dbdir"`
	Gui         string `json:"gui"`
	Netid       int    `json:"netid"`
	Passphrase  string `json:"passphrase"`
	RPCPassword string `json:"rpc_password"`
	Userhome    string `json:"userhome"`
}

func NewMM2Config() *MM2Config {
	return &MM2Config{
		Dbdir:       os.Getenv("HOME") + "/atomicdex_cli/mm2/db",
		Gui:         "AtomicDEX Client CLI",
		Netid:       7777,
		Passphrase:  "your_seed",
		RPCPassword: "your_rpc_password",
		Userhome:    os.Getenv("HOME")}
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
