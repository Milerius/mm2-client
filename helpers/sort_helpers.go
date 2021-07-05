package helpers

import (
	"sort"
	"strconv"
)

func SortDoubleSlice(data [][]string, idx int, ascending bool) {
	sort.Slice(data[:], func(i int, j int) bool {
		if left, err := strconv.ParseFloat(data[i][idx], 64); err == nil {
			if right, rightErr := strconv.ParseFloat(data[j][idx], 64); rightErr == nil {
				if !ascending {
					return left > right
				} else {
					return left < right
				}
			}
		}
		return false
	})
}

func SortDoubleSliceByDate(data [][]string, idx int, ascending bool) {
	sort.Slice(data[:], func(i int, j int) bool {
		left := DateToTimestamp(data[i][idx], true)
		right := DateToTimestamp(data[j][idx], true)
		if !ascending {
			return left > right
		} else {
			return left < right
		}
	})
}
