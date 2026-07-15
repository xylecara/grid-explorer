package main

import "fmt"

func main() {
	grid := setGrid()
	gridString := formatGrid(grid)
	treasureCoord := randomizePosition(grid)

	player := newPlayer(grid)

	fmt.Println(gridString)
	fmt.Printf(
		"Player is at (%v, %v) while treasure is at (%v, %v)\n",
		player.coordinates.X, player.coordinates.Y, treasureCoord.X, treasureCoord.Y,
	)

	player.playerControls(grid, treasureCoord)
}
