package helpers

import (
	"sort"
	"strconv"
)

func SortDoubleSlice(data [][]string, idx int) {
	sort.Slice(data[:], func(i int, j int) bool {
		if left, err := strconv.ParseFloat(data[i][idx], 64); err == nil {
			if right, rightErr := strconv.ParseFloat(data[j][idx], 64); rightErr == nil {
				return left > right
			}
		}
		return false
	})
}
