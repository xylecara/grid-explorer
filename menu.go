package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func menu() (play bool) {
	reader := bufio.NewReader(os.Stdin)
	

	fmt.Println("Welcome to Grid Explorer!\nby xyleware2026")
	fmt.Println("In this game, (0, 0) is at the top left corner of the map!\n keep that in mind if you want to get the treasre!")

	for {
		input, err := getInputs(reader, "What will you do?\n(Play, Exit)")
		if err != nil {
			fmt.Println("Invalid input! Try again.")
		}

		inputLowercased := strings.ToLower(input)
		switch inputLowercased {
		case "play":
			fmt.Println("Loading...")
			play = true
		case "exit":
			fmt.Println("Bye!\nExiting...")
			play = false
		default:
			fmt.Println("Invalid input, Try again!(Play, Exit)")
			continue
		}

		break
	}

	return play
}