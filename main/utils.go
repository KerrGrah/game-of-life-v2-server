package main

import (
	"sort"
)

func copyArray(arr []int) []int {
	return append(make([]int, 0, len(arr)), arr...)
}

func deepCopyCellTable(table CellTable) (result CellTable) {
	result = CellTable{}
	for x, yList := range table {
		result[x] = copyArray(yList)
	}
	return
}

func find(list []int, el int) bool {
	for _, v := range list {
		if v < el {
			continue
		}
		if v > el {
			return false
		}
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

func findMinIndex(list []int) int {
	for i, v := range list {
		if v >= 0 {
			return i
		}
	}
	return 0
}

func findMaxIndex(list []int, limit int) int {
	lastI := len(list) - 1

	for i, v := range list {
		if v >= limit {
			return i
		}
		if i == lastI {
			if v < 0 {
				return 0
			}
		}
	}
	return len(list)
}

func trimArray(arr []int, limit int) []int {
	if len(arr) == 0 {
		return arr
	}

	minI := findMinIndex(arr)
	maxI := findMaxIndex(arr, limit)

	result := make([]int, len(arr))
	copy(result, arr)

	return result[minI:maxI]
}

func trimCellTable(table CellTable, width, height int) (result CellTable) {
	result = CellTable{}
	for x := range make([]int, width) {
		trimmed := trimArray(table[x], height)
		if len(trimmed) != 0 {
			result[x] = trimmed
		}
	}
	return
}
