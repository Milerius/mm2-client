package mm2_tools_generics

import (
	"fmt"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"mm2_client/mm2_tools_generics/mm2_http_request"
	"mm2_client/mm2_tools_generics/mm2_wasm_request"
	"runtime"
)

func StopSimpleMarketMakerBot() (*mm2_data_structure.StopSimpleMarketMakerAnswer, error) {
	if runtime.GOARCH == "wasm" {
		return mm2_wasm_request.StopSimpleMarketMakerBot()
	} else {
		return mm2_http_request.StopSimpleMarketMakerBot()
	}
}

func StopSimpleMarketMakerBotCLI() {
	if resp, err := StopSimpleMarketMakerBot(); resp != nil {
		fmt.Printf("%s\n", resp.Result.Result)
	} else {
		fmt.Println(err)
	}
}
