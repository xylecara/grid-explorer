package main

import "fmt"

func main() {
	grid := setGrid()
	gridString := formatGrid(grid)
	randomCoord := randomizeTreasure(grid)

	fmt.Println(gridString)
	fmt.Println(randomCoord)
}
