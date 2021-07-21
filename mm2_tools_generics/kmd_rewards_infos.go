package mm2_tools_generics

import (
	"fmt"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"mm2_client/mm2_tools_generics/mm2_http_request"
	"mm2_client/mm2_tools_generics/mm2_wasm_request"
	"runtime"
)

func KmdRewardsInfo() (*mm2_data_structure.KMDRewardsInfoAnswer, error) {
	if runtime.GOARCH == "wasm" {
		return mm2_wasm_request.KmdRewardsInfo()
	} else {
		return mm2_http_request.KmdRewardsInfo()
	}
}

func KmdRewardsInfoCLI() bool {
	if resp, err := KmdRewardsInfo(); resp != nil {
		return resp.ToTable()
	} else {
		fmt.Println(err)
		return false
	}
}
