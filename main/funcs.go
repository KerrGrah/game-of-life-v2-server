package main

import (
	"math/rand"
)

func getRandBool(density float32) bool {
	// eg density = 0.7 => 70% this will be true, 30% false
	return rand.Float32() < density
}

func turn(liveCells CellTable) CellTable {
	potentialLiveCells := getPotentialLiveCells(liveCells)
	nextLiving := getNextLivingCells(potentialLiveCells, liveCells)
	return nextLiving
}

func makeEmptyBoard(height, width int) CellTable {
	board := CellTable{}

	for xKey := range make([]int, width) {
		board[xKey] = make([]int, height)
	}

	return board
}

func randPopulate(width, height int, density float32) CellTable {
	board := CellTable{}

	for xKey := range make([]int, width) {
		for yKey := range make([]int, height) {
			if getRandBool(density) {
				board[xKey] = append(board[xKey], yKey)
			}
		}
	}
	return board
}


//for every live cell get live neighbour count - break at 4, when overcrowding occurs
func getLivingNeighbourCount(liveCells CellTable, cell Cell) (count int) {
	x, y := cell.X, cell.Y
	for i := x - 1; i <= x+1; i++ {
		if col, ok := liveCells[i]; ok {
			for j := y - 1; j <= y+1; j++ {
				// skip arg "cell"
				if i == x && j == y {
					continue
				}
				if find(col, j) {
					count++

					if count == 4 {
						return
					}
				}
			}
		}
	}
	return
}

//Births: Each dead cell adjacent to exactly three live neighbors will become live in the next generation.
//Death by isolation: Each live cell with one or fewer live neighbors will die in the next generation.
//Death by overcrowding: Each live cell with four or more live neighbors will die in the next generation.
//Survival: Each live cell with either two or three live neighbors will remain alive for the next generation.
// will limit to 0 <= n <= 4
func getIsLiving(neighbourCount int, isCurrentlyAlive bool) bool {
	return neighbourCount == 3 ||
		neighbourCount == 2 && isCurrentlyAlive
}

func getNeighbours(cell Cell) CellTable {
	x, y := cell.X, cell.Y
	result := CellTable{}

	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			// skip arg "cell"
			if i == x && j == y {
				continue
			}
			result[i] = append(result[i], j)
		}
	}
	return result

}

// TODO lots of redundant action - eg removing dups on every addition
// TODO consider using structs in  y slice, x: [{y: liveNeighbours}, ...]
// TODO will be more performant after a density threshold, eg liveCells length < original length / 3, otherwise just use full dimensions
func getPotentialLiveCells(liveCells CellTable) CellTable {
	// keep already living
	result := deepCopyCellTable(liveCells)

	for xKey := range liveCells {
		for _, yVal := range liveCells[xKey] {
			neighbours := getNeighbours(Cell{xKey, yVal})

			for neighboursX, neighboursY := range neighbours {
				if resultY, ok := result[neighboursX]; ok {
					result[neighboursX] = mergeWithoutDuplicates(resultY, neighboursY)
				} else {
					result[neighboursX] = neighboursY
				}
			}
		}
	}
	return result
}


func getNextLivingCells(potentialLiveCells CellTable, liveCells CellTable) CellTable {
	nextLiveCells := CellTable{}

	for xKey := range potentialLiveCells {
		for _, yVal := range potentialLiveCells[xKey] {
			isCurrentlyAlive := find(liveCells[xKey], yVal)
			livingNeighbours := getLivingNeighbourCount(liveCells, Cell{xKey, yVal})
			willBeAlive := getIsLiving(livingNeighbours, isCurrentlyAlive)

			if willBeAlive {
				if nextY, ok := nextLiveCells[xKey]; ok {
					nextLiveCells[xKey] = append(nextY, yVal)
				} else {
					nextLiveCells[xKey] = []int{yVal}
				}
			}
		}
	}
	return nextLiveCells
}

