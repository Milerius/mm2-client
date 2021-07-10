package cli

import (
	"mm2_client/http"
)

func MyOrders() {
	if resp := http.MyOrders(); resp != nil {
		resp.ToTable()
	}
}
