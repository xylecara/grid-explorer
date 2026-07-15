package main

import "fmt"

func main() {
	grid := setGrid()
	treasureCoord := randomizePosition(grid)

	player := newPlayer(grid)

	fmt.Printf(
		"Treasure is at (%v, %v)\n",
		treasureCoord.X, treasureCoord.Y,
	)

	player.playerControls(grid, treasureCoord)
}
