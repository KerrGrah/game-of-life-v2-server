package main

import (
	"reflect"
	"testing"
)

func TestGetLivingNeighbourCount(t *testing.T) {
	type CellAndExpected struct {
		Cell
		Expected int
	}
	// return is up to 3 as is maximum interesting value (though up to 8 possible)
	var testList = []CellAndExpected{
		{Cell{10, 2}, 0},
		{Cell{1, 4}, 3},
		{Cell{1, 2}, 4},
	}

	for _, test := range testList {
		count := getLivingNeighbourCount(testBoards[0], test.Cell)
		if count != test.Expected {
			t.Errorf("Count was incorrect - input %v got: %d, want: %d.", test.Cell, count, test.Expected)
		}
	}
}

func TestGetIsLiving(t *testing.T) {
	type ArgsAndExpected struct {
		Count    int
		IsAlive  bool
		Expected bool
	}
	testList := []ArgsAndExpected{
		{0, false, false},
		{0, true, false},
		{1, false, false},
		{1, true, false},
		{2, false, false},
		{2, true, true},
		{3, false, true},
		{3, true, true},
		{4, false, false},
		{4, true, false},
	}

	for _, test := range testList {
		isLiving := getIsLiving(test.Count, test.IsAlive)
		if isLiving != test.Expected {
			t.Errorf("Is living result for count %d was incorrect, got: %t, want: %t.", test.Count, isLiving, test.Expected)
		}
	}
}

func TestDeepCopyCellTable(t *testing.T) {
	testBoardCopy := deepCopyCellTable(testBoards[0])

	delete(testBoardCopy, 2)

	if _, ok := testBoards[0][2]; !ok {
		t.Errorf("Copy map entry failed")
	}

	testBoardCopy[0][0] = 99

	if testBoards[0][0][0] == 99 {
		t.Errorf("Copy deepe slice failed - %d", testBoardCopy[0][0])
	}
}

func TestGetNeighbours(t *testing.T) {
	testCell := Cell{2, 2}
	expected := CellTable{
		1: []int{1, 2, 3},
		2: []int{1, 3},
		3: []int{1, 2, 3},
	}
	neighbours := getNeighbours(testCell)

	if !reflect.DeepEqual(neighbours, expected) {
		t.Errorf("Neighbours not correct - received %v", neighbours)
	}
}

func TestGetPotentialLiveCells(t *testing.T) {
	var testBoard = CellTable{
		0:  []int{1, 2, 3, 4},
		1:  []int{1, 2, 3, 4},
		2:  []int{1, 2},
		10: []int{0, 4},
	}
	expected := CellTable{
		-1: []int{0, 1, 2, 3, 4, 5},
		0:  []int{0, 1, 2, 3, 4, 5},
		1:  []int{0, 1, 2, 3, 4, 5},
		2:  []int{0, 1, 2, 3, 4, 5},
		3:  []int{0, 1, 2, 3},
		9:  []int{-1, 0, 1, 3, 4, 5},
		10: []int{-1, 0, 1, 3, 4, 5},
		11: []int{-1, 0, 1, 3, 4, 5},
	}
	potentialLiving := getPotentialLiveCells(testBoard)

	if !reflect.DeepEqual(potentialLiving, expected) {
		t.Errorf("Potential living not correct - received %v", potentialLiving)
	}
}

func TestMergeWithoutDuplicates(t *testing.T) {
	type Args struct {
		arr1, arr2 []int
	}
	type TestInputAndExpected struct {
		args     Args
		expected []int
	}

	testList := []TestInputAndExpected{
		{Args{[]int{0, 0}, []int{1}}, []int{0, 1}},
		{Args{[]int{0}, []int{1, 1}}, []int{0, 1}},
		{Args{[]int{0, 0}, []int{1, 1}}, []int{0, 1}},
		{Args{[]int{0, 0, 1}, []int{1, 0, -1, 0}}, []int{-1, 0, 1}},
		{Args{[]int{1, 2, 3, 4}, []int{3, 4, 5, 6}}, []int{1, 2, 3, 4, 5, 6}},
	}
	for _, test := range testList {
		set := mergeWithoutDuplicates(test.args.arr1, test.args.arr2)

		if !reflect.DeepEqual(set, test.expected) {
			t.Errorf("Merge slices without duplicates not correct - input %v, %v and  received %v", test.args.arr1, test.args.arr2, set)
		}
	}
}

func TestGetNextLivingCells(t *testing.T) {
	type Args struct {
		potential, liveCells CellTable
	}
	type TestInputAndExpected struct {
		args     Args
		expected CellTable
	}

	testList := []TestInputAndExpected{
		{Args{getPotentialLiveCells(testBoards[0]), testBoards[0]}, CellTable{
			-1: []int{2, 3},
			0:  []int{1, 4},
			1:  []int{0, 4},
			2:  []int{1},
		}},

		{Args{getPotentialLiveCells(testBoards[1]), testBoards[1]}, CellTable{
			0: []int{2, 3},
			1: []int{1, 4},
			2: []int{2, 3},
		}},

		{Args{getPotentialLiveCells(testBoards[2]), testBoards[2]}, CellTable{
			1: []int{4, 5, 6, 7, 8, 9, 10, 11},
			2: []int{6, 7, 8, 9, 10, 11, 12, 13},
		}},
	}

	for _, test := range testList {
		nextLiving := getNextLivingCells(test.args.potential, test.args.liveCells)

		if !reflect.DeepEqual(nextLiving, test.expected) {
			t.Errorf("Get next living cells not correct - received %v", nextLiving)
		}
	}
}

func TestMakeEmptyBoard(t *testing.T) {
	type TestInput struct {
		width, height int
	}

	testList := []TestInput{{width: 10, height: 12}, {width: 1, height: 102}, {width: 100, height: 100}}

	for _, test := range testList {
		board := makeEmptyBoard(test.width, test.height)

		if len(board) != test.width || len(board[0]) != test.height {
			t.Errorf("Make empty board not correct - expected %d / %d -  received %d / %d", test.width, test.height, len(board), len(board[0]))
		}
	}
}

func TestRandPopulate(t *testing.T) {
	var (
		density  float32 = 0.3
		width            = 1000
		height           = 1000
		expected         = 300
		lim              = 3
		avg              = 0
	)

	populated := randPopulate(width, height, density)

	for _, row := range populated {
		avg += len(row)
	}
	avg /= len(populated)

	if avg > expected+lim || avg < expected-lim {
		t.Errorf("Populate random board likely not correct - with density of %f - received avg length of %d expected %d", density, avg, expected)
	}
}

func TestFind(t *testing.T) {
	type ArgsAndExpected struct {
		List     []int
		El       int
		Expected bool
	}
	testList := []ArgsAndExpected{
		{[]int{1, 2, 3, 4, 5, 6}, 1, true},
		{[]int{1, 2, 3, 4, 5, 6}, 7, false},
		{[]int{1, 2, 3, 4, 5, 6}, 0, false},
		{[]int{1, 2, 3, 4, 5, 6}, 5, true},
		{[]int{0, 2, 4, 5, 6, 8, 9, 10, 11, 12, 17, 23, 25, 29, 31, 32, 35, 37, 38, 39, 43, 49, 52, 53, 54, 55, 58, 59, 60, 62, 63, 72, 75, 77, 78}, 78, true},
		{[]int{0, 2, 4, 5, 6, 8, 9, 10, 11, 12, 17, 23, 25, 29, 31, 32, 35, 37, 38, 39, 43, 49, 52, 53, 54, 55, 58, 59, 60, 62, 63, 72, 75, 77, 78}, 100, false},
	}

	for _, test := range testList {
		result := find(test.List, test.El)
		if result != test.Expected {
			t.Errorf("Find el in list result was incorrect - find %v in %v got: %v, want: %v.", test.List, test.El, result, test.Expected)
		}
	}
}

func TestFilter(t *testing.T) {
	type ArgsAndExpected struct {
		List     []int
		limit    int
		expected []int
	}
	testList := []ArgsAndExpected{
		{[]int{}, 1, []int{}},
		{[]int{1, 2, 3, 4, 5, 6}, 1, []int{}},
		{[]int{1, 2, 3, 4, 5, 6}, 2, []int{1}},
		{[]int{1, 2, 3, 4, 5, 6}, 3, []int{1, 2}},
		{[]int{1, 2, 3, 4, 5, 6}, 4, []int{1, 2, 3}},
		{[]int{1, 2, 3, 5, 6}, 5, []int{1, 2, 3}},
		{[]int{1, 2, 3, 5, 6}, 6, []int{1, 2, 3, 5}},
		{[]int{1, 2, 3, 5, 6}, 9, []int{1, 2, 3, 5, 6}},
		{[]int{-2, -1, 0, 1, 2, 3, 4, 5, 6}, 3, []int{0, 1, 2}},
		{[]int{-2, 0, 1, 2, 3, 4, 5, 6}, 3, []int{0, 1, 2}},
		{[]int{-2, 1, 2, 3, 4, 5, 6}, 3, []int{1, 2}},
	}

	for _, test := range testList {
		result := filter(test.List, test.limit)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Filter arrray was incorrect - from %v limit %d got: %v, wanted: %v.", test.List, test.limit, result, test.expected)
		}
	}
}

func TestTrimCellTable(t *testing.T) {

	type Args struct {
		table         CellTable
		width, height int
	}
	type TestInputAndExpected struct {
		args     Args
		expected CellTable
	}

	testList := []TestInputAndExpected{
		{Args{testBoards[0], 1, 2}, CellTable{
			0: []int{1},
		}},
		{Args{testBoards[0], 2, 2}, CellTable{
			0: []int{1},
			1: []int{1},
		}},
		{Args{testBoards[0], 3, 3}, CellTable{
			0: []int{1, 2},
			1: []int{1, 2},
			2: []int{1, 2},
		}},
		{Args{testBoards[4], 3, 6}, CellTable{
			0: []int{1, 4},
			1: []int{3},
		}},
		{Args{testBoards[4], 5, 11}, CellTable{
			0: []int{1, 4, 7, 10},
			1: []int{3, 6, 9},
			3: []int{7, 10},
		}},
	}
	for _, test := range testList {
		result := trimCellTable(test.args.table, test.args.width, test.args.height)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Table trim not correct - input %v, %d / %d,  and received %v, expected %v", test.args.table, test.args.height, test.args.width, result, test.expected)
		}
	}
}

func TestTrimArray(t *testing.T) {

	type Args struct {
		arr   []int
		limit int
	}
	type TestInputAndExpected struct {
		args     Args
		expected []int
	}

	testList := []TestInputAndExpected{
		{Args{[]int{}, 0}, []int{}},
		{Args{[]int{}, 100}, []int{}},
		{Args{[]int{-4, 0}, 100}, []int{0}},
		{Args{[]int{0, 1, 2, 3}, 1}, []int{0}},
		{Args{[]int{3, 6, 9, 12}, 11}, []int{3, 6, 9}},
		{Args{[]int{-3, -2}, 100}, []int{}},
		{Args{[]int{-3, 0, 1, 2, 3}, 1}, []int{0}},
		{Args{[]int{-3, 0, 1, 2, 3}, 2}, []int{0, 1}},
		{Args{[]int{-3, -2, -1, 0, 1, 2, 3}, 8}, []int{0, 1, 2, 3}},
		{Args{[]int{-3, -2, -1, 0, 1, 2, 3, 6, 7}, 5}, []int{0, 1, 2, 3}},
		{Args{[]int{-3, -2, -1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 8}, []int{1, 2, 3, 4, 5, 6, 7}},
	}
	for _, test := range testList {
		result := trimArray(test.args.arr, test.args.limit)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Array trim not correct - input %v, height %v and received %v, expected %v", test.args.arr, test.args.limit, result, test.expected)
		}
	}
}
