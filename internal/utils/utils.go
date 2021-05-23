package utils

import (
	"errors"
)

// Split slice into batches.
// Batch size must be greater than zero.
// If the slice equals nil then result is nil.
func SplitSlice(slice []int, batchSize int) ([][]int, error) {
	if slice == nil {
		return nil, nil
	}

	if batchSize <= 0 {
		return nil, errors.New("Batch size must be greater than zero")
	}

	var fullBatchCount = len(slice) / batchSize
	var totalBatchCount = fullBatchCount

	if len(slice)%batchSize != 0 {
		totalBatchCount++
	}

	result := make([][]int, 0, totalBatchCount)
	i := 0

	for fullBatchCount > 0 {
		result = append(result, slice[i:i+batchSize])
		i += batchSize
		fullBatchCount--
	}

	if len(slice)%batchSize != 0 {
		result = append(result, slice[i:])
	}

	return result, nil
}

// Swap map keys and values.
// If incoming map is nil then result is nil.
// Incoming map must not contain duplicate values.
func InvertMap(m map[string]int) map[int]string {
	if m == nil {
		return nil
	}

	result := make(map[int]string)

	for key, value := range m {
		_, exists := result[value]

		if exists {
			panic("Doublicate key")
		}

		result[value] = key
	}

	return result
}

// Filter a incoming list  by forbidden characters..
// If the incoming list is nil then result is nil.
// If the incoming filter is nill, empty then result is list.
func FilterList(list []rune, filter []rune) []rune {
	if list == nil || filter == nil {
		return list
	}

	set := map[rune]struct{}{}

	for _, ch := range filter {
		set[ch] = struct{}{}
	}

	result := make([]rune, 0, len(list))

	for _, ch := range list {
		_, exists := set[ch]

		if !exists {
			result = append(result, ch)
		}
	}

	return result
}
