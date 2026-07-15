package main

import "fmt"

func main() {
	play := menu()
	for play {
		grid := setGrid()

		treasureCoord := randomizePosition(grid)
		player := newPlayer(grid)
		play = player.playerControls(grid, treasureCoord, play)
	} 
    fmt.Println("Bye!\nExiting...")
}
