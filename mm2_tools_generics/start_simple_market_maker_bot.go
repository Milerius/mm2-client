package mm2_tools_generics

import (
	"fmt"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"mm2_client/mm2_tools_generics/mm2_http_request"
	"mm2_client/mm2_tools_generics/mm2_wasm_request"
	"runtime"
)

func StartSimpleMarketMakerBot() (*mm2_data_structure.StartSimpleMarketMakerAnswer, error) {
	if runtime.GOARCH == "wasm" {
		return mm2_wasm_request.StartSimpleMarketMakerBot()
	} else {
		return mm2_http_request.StartSimpleMarketMakerBot()
	}
}

func StartSimpleMarketMakerBotCLI() {
	if resp, err := StartSimpleMarketMakerBot(); resp != nil {
		fmt.Printf("%v", resp)
	} else {
		fmt.Println(err)
	}
}
