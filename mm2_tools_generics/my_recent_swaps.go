package mm2_tools_generics

import (
	"fmt"
	"mm2_client/mm2_tools_generics/mm2_data_structure"
	"mm2_client/mm2_tools_generics/mm2_http_request"
	"mm2_client/mm2_tools_generics/mm2_wasm_request"
	"runtime"
)

func MyRecentSwaps(limit string, pageNumber string, baseCoin string, relCoin string, from string, to string) (*mm2_data_structure.MyRecentSwapsAnswer, error) {
	if runtime.GOARCH == "wasm" {
		return mm2_wasm_request.MyRecentSwaps(limit, pageNumber, baseCoin, relCoin, from, to)
	} else {
		return mm2_http_request.MyRecentSwaps(limit, pageNumber, baseCoin, relCoin, from, to)
	}
}

func processMyRecentSwaps(limit string, pageNumber string, args []string) (*mm2_data_structure.MyRecentSwapsAnswer, error) {
	baseCoin := ""
	relCoin := ""
	from := ""
	to := ""
	if len(args) >= 1 {
		baseCoin = args[0]
	}
	if len(args) >= 2 {
		relCoin = args[1]
	}
	if len(args) >= 3 {
		from = args[2]
	}
	if len(args) >= 4 {
		to = args[3]
	}
	return MyRecentSwaps(limit, pageNumber, baseCoin, relCoin, from, to)
}

func MyRecentSwapsCLI(limit string, pageNumber string, args []string) {
	if resp, err := processMyRecentSwaps(limit, pageNumber, args); resp != nil {
		resp.ToTable()
	} else {
		fmt.Println(err)
	}
}
