package main

import (
	"fmt"
)

type coords struct {
	X int
	Y int
}

func main() {
	rows, columns := 5, 3

	grid := make([][]coords, columns)
	for i := range grid {
		grid[i] = make([]coords, rows)
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
