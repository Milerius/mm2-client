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

// Returns only unique values from a list of strings
func UniqueStrings(input []string) []string {
	u := make([]string, 0, len(input))
	m := make(map[string]bool)
	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}

func ChunkStringList(slice []string, chunkSize int) [][]string {
	var chunks [][]string
	for {
		if len(slice) == 0 {
			break
		}
		if len(slice) < chunkSize {
			chunkSize = len(slice)
		}
		chunks = append(chunks, slice[0:chunkSize])
		slice = slice[chunkSize:]
	}
	return chunks
}
