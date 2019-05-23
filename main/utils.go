package main

import "sort"

func copyArray(arr []int) []int {
	return append(arr[:0:0], arr...)
}

func deepCopyCellTable(table CellTable) (result CellTable) {
	result = CellTable{}
	for x, yList := range table {
		result[x] = copyArray(yList)
	}
	return
}

// TODO if list is sorted could return false when v > el
func find(list []int, el int) bool {
	for _, v := range list {
		if v == el {
			return true
		}
	}
	return false
}

func mergeWithoutDuplicates(a, b []int) []int {
	merged := append(a, b...)
	sort.Ints(merged)

	result := []int{}

	for i, v := range merged {
		if i == len(merged)-1 || v < merged[i+1] {
			result = append(result, v)
		}
	}
	return result
}

