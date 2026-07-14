package main

import (
	"bufio"
	"fmt"
	"strings"
)

func getInputs(r *bufio.Reader, prompt string) (string, error) {
	fmt.Println(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}