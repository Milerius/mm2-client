package cli

import (
	"fmt"
	"mm2_client/mm2_tools_generics/mm2_http_request"
)

func Broadcast(coin string, txHex string) {
	if resp, err := mm2_http_request.Broadcast(coin, txHex); resp != nil {
		fmt.Println(resp.TxUrl)
	} else {
		fmt.Println(err)
	}
}
