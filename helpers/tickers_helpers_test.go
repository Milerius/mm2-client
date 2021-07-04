package helpers

import (
	"fmt"
	"testing"
)

func TestRetrieveMainTicker(t *testing.T) {
	fmt.Println(RetrieveMainTicker("KMD"))
	fmt.Println(RetrieveMainTicker("KMD-BEP20"))
}
