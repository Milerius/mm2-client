package cli

import (
	"fmt"
	"mm2_client/http"
)

func CancelOrder(uuid string) {
	if resp := http.CancelOrder(uuid); resp != nil {
		fmt.Println(resp.Result)
	}
}
