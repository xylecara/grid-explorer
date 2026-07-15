package main

import "os"

func main() {
	play := menu()
	if play {
		grid := setGrid()

		treasureCoord := randomizePosition(grid)
		player := newPlayer(grid)
		player.playerControls(grid, treasureCoord)
	} else {
		os.Exit(0)
	}
}
