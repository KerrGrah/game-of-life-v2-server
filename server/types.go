package main

// 				    X     Y
type CellTable map[int][]int

type Cell struct {
	X, Y int
}

type GameSetup struct {
	Width    int
	Height   int
	Density  float32
	Speed    float64
	Initiate bool
}
