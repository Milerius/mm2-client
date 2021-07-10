package helpers

import (
	"github.com/soniah/evaler"
	"strconv"
)

func BigFloatMultiply(first string, second string, prec int) string {
	result, err := evaler.Eval(first + "*" + second)
	if err != nil {
		return "0"
	}
	return result.FloatString(prec)
}

func BigFloatAdd(first string, second string, prec int) string {
	result, err := evaler.Eval(first + "+" + second)
	if err != nil {
		return "0"
	}
	return result.FloatString(prec)
}

func BigFloatDivide(first string, second string, prec int) string {
	result, err := evaler.Eval(first + "/" + second)
	if err != nil {
		return "0"
	}
	return result.FloatString(prec)
}

func ResizeNb(nb string) string {
	if len(nb) >= 8 {
		return nb[0:8]
	} else {
		return nb
	}
}

func AsFloat(input string) float64 {
	val, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0
	}
	return val
}
