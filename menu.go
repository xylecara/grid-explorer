package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func menu() (play bool) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to Grid Explorer! by xyleware2026")
	fmt.Println("In this game you have to find a treasure! Check the wind for warmth or the ground for rumbles!")

	for {
		input, err := getInput(reader, "What will you do?\n(Play, Exit)")
		if err != nil {
			fmt.Println("Invalid input! Try again.")
		}

		inputLowercased := strings.ToLower(input)
		switch inputLowercased {
		case "play":
			fmt.Println("Loading...")
			play = true
		case "exit":
			play = false
		default:
			fmt.Println("Invalid input, Try again!(Play, Exit)")
			continue
		}

		break
	}

	return play
}
