package main

//func getNextMaxLengths(liveCells CellTable) []int {
//	// limit new board size to current live plus one for each edge
//	xLength := len(liveCells) + 2
//	yLength := getLongestLen(liveCells) + 2
//
//	return []int{xLength, yLength}
//}

// used to get the longest Y slice length when creating new board
//func getLongestLen(board CellTable) (length int) {
//	for _, col := range board {
//		if len(col) > length {
//			length = len(col)
//		}
//	}
//	return
//}

// turns live to full dimension boll board
//func fillBinBoard(emptyBoard [][]bool, liveCells CellTable) [][]bool {
//	for xKey := range emptyBoard {
//		for yKey := range emptyBoard[0] {
//			emptyBoard[xKey][yKey] = find(liveCells[xKey], yKey)
//		}
//	}
//	return emptyBoard
//}

//func BenchmarkFilterZeroToLimit(b *testing.B) {
//	testList := []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 2, 3, 7, 9, 13, 15, 16, 17, 19, 23, 24, 26, 27, 30, 34, 36, 37, 38, 39, 40, 42, 43, 44, 46, 47, 52, 55, 56, 59, 61, 62, 63, 64, 66, 70, 73, 76, 78, 79}
//	for n := 0; n < b.N; n++ {
//		filterZeroToLimit(testList, 70)
//	}
//}
//
//func filterZeroToLimit(list []int, value int) []int {
//	result := []int{}
//
//	for _, v := range list {
//		if v > -1 && v < value {
//			result = append(result, v)
//		}
//	}
//	return result
//}
