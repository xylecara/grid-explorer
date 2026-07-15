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

func (p *player) playerControls(grid [][]coords, treasure coords, play bool) bool {
	reader := bufio.NewReader(os.Stdin)
	var gridLimit coords

	for x := range grid {
		for y := range grid[x] {
			gridLimit = grid[x][y]
		}
	}

	for p.health > 0 {
		input, err := getInput(reader, "Do something \n(Up, Down, Left, Right, Dig, Coords, Map, Check, Exit)")
		if err != nil {
			fmt.Println("Invalid input! Try again.")
		}

		inputLowercased := strings.ToLower(input)
		switch inputLowercased {
		case "up":
			if p.coordinates.Y < gridLimit.Y {
				p.coordinates.Y++
			} else {
				fmt.Println("You hit a wall")
			}
			continue

		case "down":
			if p.coordinates.Y > 0 {
				p.coordinates.Y--
			} else {
				fmt.Println("You hit a wall")
			}
			continue

		case "right":
			if p.coordinates.X < gridLimit.X {
				p.coordinates.X++
			} else {
				fmt.Println("You hit a wall")
			}
			continue

		case "left":
			if p.coordinates.X > 0 {
				p.coordinates.X--
			} else {
				fmt.Println("You hit a wall")
			}
			continue

		case "dig":
			fmt.Println("Digging...")
			if p.isTreasureUnder(treasure) {
				fmt.Println("You Win!")
				play = playAgain(play)
			} else {
				fmt.Println("KABOOM! That was a landmine!")
				p.health--
				if p.health <= 0 {
					fmt.Println("You died...")
					play = playAgain(play)
				}
				continue
			}

		case "show":
			fmt.Printf("You are at (%v, %v)\n", p.coordinates.X, p.coordinates.Y)
			continue

		case "map":
			fmt.Println("Here is the map:\n" + formatGrid(grid))
			continue

		case "check":
			p.isNearTreasure(treasure)
			continue
		case "exit":
			fmt.Println("Exiting...")
			play = playAgain(play)

		default:
			fmt.Println("Invalid input! Not an available command.")
			continue
		}

		break
	}

	return play
}

func (p *player) isTreasureUnder(treasure coords) bool {
	return p.coordinates == treasure
}

func (p *player) isNearTreasure(teasure coords) {
	if p.coordinates == teasure {
		fmt.Println("The ground rumbles...")
	} else if p.coordinates.X == teasure.X {
		fmt.Println("The air in the east and west feels warm...")
	} else if p.coordinates.Y == teasure.Y {
		fmt.Println("The air in the north and south feels warm...")
	} else {
		fmt.Println("Hmmm, nothing out of the ordinary... Odd...")
	}
}

func playAgain(play bool) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		input, err := getInput(reader, "Would you like to try again?(Yes/No)")
		if err != nil {
			fmt.Println("Invalid input!, try again.")
		}

		inputLowercased := strings.ToLower(input)
		switch inputLowercased {
		case "yes":
			play = true

		case "no":
			play = false

		default:
			fmt.Println("Invalid input!, try again.")
			continue
		}

		break
	}

	return play
}
