package cli

import (
	"fmt"
	"mm2_client/mm2_tools_generics/mm2_http_request"
)

func CancelOrder(uuid string) {
	if resp, cancelErr := mm2_http_request.CancelOrder(uuid); resp != nil {
		fmt.Println(resp.Result)
	} else {
		fmt.Println(cancelErr)
	}
}
