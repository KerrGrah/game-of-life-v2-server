package main

import "testing"

func benchmarkGetLivingNeighbourCount(cell Cell, b *testing.B) {
	for n := 0; n < b.N; n++ {
		getLivingNeighbourCount(testBoards[3], cell)
	}
}

func BenchmarkGetLivingNeighbourCountCorner(b *testing.B) {
	benchmarkGetLivingNeighbourCount(Cell{0, 0}, b)
}
func BenchmarkGetLivingNeighbourCountXEdge(b *testing.B) {
	benchmarkGetLivingNeighbourCount(Cell{0, 10}, b)
}
func BenchmarkGetLivingNeighbourCountYEdge(b *testing.B) {
	benchmarkGetLivingNeighbourCount(Cell{10, 0}, b)
}
func BenchmarkGetLivingNeighbourCountEdgeClose(b *testing.B) {
	benchmarkGetLivingNeighbourCount(Cell{10, 10}, b)
}
func BenchmarkGetLivingNeighbourCountEdgeFar(b *testing.B) {
	benchmarkGetLivingNeighbourCount(Cell{40, 40}, b)
}

func BenchmarkCopyArray(b *testing.B) {
	for n := 0; n < b.N; n++ {
		copyArray(testBoards[3][0])
	}
}

func BenchmarkDeepCopyCellTable(b *testing.B) {
	for n := 0; n < b.N; n++ {
		deepCopyCellTable(testBoards[3])
	}
}

func BenchmarkTurn(b *testing.B) {
	for n := 0; n < b.N; n++ {
		turn(testBoards[3])
	}
}

func benchmarkFind(el int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		find(testBoards[3][6], el)
	}
}

func BenchmarkFindStart(b *testing.B)        { benchmarkFind(0, b) }
func BenchmarkFindEnd(b *testing.B)          { benchmarkFind(79, b) }
func BenchmarkFindMissingSmall(b *testing.B) { benchmarkFind(1, b) }
func BenchmarkFindMissingBig(b *testing.B)   { benchmarkFind(100, b) }

func BenchmarkGetPotentialLiveCells(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getPotentialLiveCells(testBoards[3])
	}
}

func BenchmarkTrimCellTable(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		trimCellTable(testBoards[3], 70, 30)
	}
}

func BenchmarkTrimArray(b *testing.B) {
	testList := []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 2, 3, 7, 9, 13, 15, 16, 17, 19, 23, 24, 26, 27, 30, 34, 36, 37, 38, 39, 40, 42, 43, 44, 46, 47, 52, 55, 56, 59, 61, 62, 63, 64, 66, 70, 73, 76, 78, 79}
	for n := 0; n < b.N; n++ {
		trimArray(testList, 70)
	}
}
