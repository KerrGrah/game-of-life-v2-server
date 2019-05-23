package main

import (
	"reflect"
	"testing"
)

var testBoards = []CellTable{
	{
		0:  []int{1, 2, 3, 4},
		1:  []int{1, 2, 3, 4},
		2:  []int{1, 2},
		10: []int{0, 4},
	},
	{
		0:  []int{1, 3},
		1:  []int{2, 4},
		2:  []int{1, 3},
		10: []int{2, 4},
	},
	{
		0: []int{1, 4, 7, 10},
		1: []int{3, 6, 9, 12},
		2: []int{5, 8, 11, 14},
		3: []int{7, 10, 13, 16},
	},
}

func TestGetLivingNeighbourCount(t *testing.T) {
	type CellAndExpected struct {
		Cell
		Expected int
	}
	// max return is 3 though up to 8 possible
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

		if len(board) != test.height || len(board[0]) != test.width {
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
