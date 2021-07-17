package cli

import (
	"fmt"
	"mm2_client/mm2_tools_generics/mm2_http_request"
)

func MyOrders() {
	if resp, err := mm2_http_request.MyOrders(); resp != nil {
		resp.ToTable()
	} else {
		fmt.Println(err)
	}
}
