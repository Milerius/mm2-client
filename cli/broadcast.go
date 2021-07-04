package cli

import (
	"fmt"
	"mm2_client/http"
)

func Broadcast(coin string, txHex string) {
	if resp := http.Broadcast(coin, txHex); resp != nil {
		fmt.Println(resp.TxUrl)
	}
}
