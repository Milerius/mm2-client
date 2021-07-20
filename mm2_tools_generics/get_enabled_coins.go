package mm2_tools_generics

import (
	"fmt"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"mm2_client/mm2_tools_generics/mm2_http_request"
	"mm2_client/mm2_tools_generics/mm2_wasm_request"
	"runtime"
)

func GetEnabledCoins() (*mm2_data_structure.GetEnabledCoinsAnswer, error) {
	if runtime.GOARCH == "wasm" {
		return mm2_wasm_request.GetEnabledCoins()
	} else {
		return mm2_http_request.GetEnabledCoins()
	}
}

func GetEnabledCoinsCLI() {
	resp, err := GetEnabledCoins()
	if resp != nil && len(resp.Result) > 0 {
		resp.ToTable()
	} else {
		fmt.Println(err)
	}
}
