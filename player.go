package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type player struct {
	coordinates coords
	health      int
}

func newPlayer(grid [][]coords) player {
	position := randomizePosition(grid)

	return player{
		coordinates: position,
		health:      3,
	}
}

func (p *player) movementControls(grid [][]coords) {
	reader := bufio.NewReader(os.Stdin)
	var gridLimit coords

	for x := range grid {
		for y := range grid[x] {
			gridLimit = grid[x][y]
		}
	}

	for {
		input, err := getInputs(reader, "Move around(Up, Down, Left, Right, Exit)")
		if err != nil {
			fmt.Println("Invalid input! Try again.")
		}

		inputLowercased := strings.ToLower(input)
		switch inputLowercased{
		case "up":
			if p.coordinates.Y < gridLimit.Y {
				p.coordinates.Y++
			    fmt.Printf("Player is at (%v, %v)\n", p.coordinates.X, p.coordinates.Y)
			} else {
				fmt.Println("You hit a wall")
			}
			continue

		case "down":
			if p.coordinates.Y > 0 {
				p.coordinates.Y--
			    fmt.Printf("Player is at (%v, %v)\n", p.coordinates.X, p.coordinates.Y)
			} else {
				fmt.Println("You hit a wall")
			}
			continue
		
		case "right":
			if p.coordinates.X < gridLimit.X {
				p.coordinates.X++
			    fmt.Printf("Player is at (%v, %v)\n", p.coordinates.X, p.coordinates.Y)
			} else {
				fmt.Println("You hit a wall")
			}
			continue
		case "left":
			if p.coordinates.X > 0 {
				p.coordinates.X--
			    fmt.Printf("Player is at (%v, %v)\n", p.coordinates.X, p.coordinates.Y)
			} else {
				fmt.Println("You hit a wall")
			}
			continue
		
		case "exit":
			fmt.Println("Exiting...")
		default:
			fmt.Println("Invalid input! Try again.\n(Up, Down, Left, Right, Exit)")
			continue
		}
		
		break
	}
}

