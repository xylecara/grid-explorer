package main

import (
	"fmt"
)

type coords struct {
	X int
	Y int
}

func main() {
	rows, columns := 10, 10 

	grid := make([][]coords, rows)
	for i := range grid {
		grid[i] = make([]coords, columns)
	}

	for x := range grid {
		for y := range grid[x] {
			grid[x][y] = coords{X: x, Y: y}
		}
	}

	for i := range grid {
		for j := range grid[i] {
			fmt.Printf("(%v, %v)", grid[i][j].X, grid[i][j].Y)
		}
		fmt.Println()
	}
}
