package helpers

import (
	"github.com/soniah/evaler"
)

func BigFloatMultiply(first string, second string, prec int) string {
	result, err := evaler.Eval(first + "*" + second)
	if err != nil {
		return "0"
	}
	return result.FloatString(prec)
}
