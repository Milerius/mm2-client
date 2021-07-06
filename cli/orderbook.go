package cli

import "mm2_client/http"

func Orderbook(base string, rel string) {
	if resp := http.Orderbook(base, rel); resp != nil {
		resp.ToTable(base, rel)
	}
}
