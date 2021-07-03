package cli

import (
	"fmt"
	"mm2_client/config"
	"mm2_client/http"
)

func Enable(coin string) {
	if val, ok := config.GCFGRegistry[coin]; ok {
		switch val.Type {
		case "BEP-20", "ERC-20":
			http.Enable(coin)
		default:
			fmt.Println("Not supported yet")
		}
	} else {
		fmt.Printf("Cannot enable %s, did you start MM2 with the command <start> ?\n", coin)
	}
}

/*func EnableMultipleCoins(coins string[]) {

}*/
