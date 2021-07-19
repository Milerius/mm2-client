package config

import (
	"encoding/json"
	"github.com/kpango/glg"
	"mm2_client/config/wasm_storage"
)

func UpdateWasm() {
	glg.Infof("Updating cfg in wasm")
	file, _ := json.MarshalIndent(GCFGRegistry, "", " ")
	cfgStr := string(file)
	wasm_storage.Store("wasm_coins_cfg", cfgStr)
}
