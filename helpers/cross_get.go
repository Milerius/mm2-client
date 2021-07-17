package helpers

import (
	"github.com/kpango/glg"
	"net/http"
	"runtime"
)

func CrossGet(url string) (resp *http.Response, err error) {
	if runtime.GOARCH != "wasm" {
		return http.Get(url)
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		_ = glg.Errorf("cross_get err: %v", err)
		return nil, err
	}
	req.Header.Add("js.fetch:mode", "cors")
	return http.DefaultClient.Do(req)
}
