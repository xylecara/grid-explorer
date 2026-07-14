package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coords struct {
	X int
	Y int
}

func setGrid() [][]coords {
	rows := getRows()
	columns := getColumns()

	grid := make([][]coords, rows)
	for i := range grid {
		grid[i] = make([]coords, columns)
	}

	for x := range grid {
		for y := range grid[x] {
			grid[x][y] = coords{X: x, Y: y}
		}
	}

	return grid
}

func formatGrid(grid [][]coords) string {
	var gridFormat strings.Builder

	for x := range grid {
		for y := range grid[x] {
			fmt.Fprintf(&gridFormat, "(%v, %v)", grid[x][y].X, grid[x][y].Y)
		}
		gridFormat.WriteString("\n")
	}

	return gridFormat.String()
}

func getRows() int {
	reader := bufio.NewReader(os.Stdin)
	var rows int
	var rowsString string
	var err error

	for {
		rowsString, err = getInputs(reader, "How many rows?")
		if err != nil {
			fmt.Println("Invalid input! Try again.")
			continue
		}

		rows, err = strconv.Atoi(rowsString)
		if err != nil {
			fmt.Println("Invalid input! Only integers are allowed, try again.")
			continue
		}

		if rows <= 0 {
			fmt.Println("Invalid input! You cant use 0 or less as \nan input, try agian.")
			continue
		}

		break
	}

	return rows
}

func getColumns() int {
	reader := bufio.NewReader(os.Stdin)
	var columns int
	var columnsString string
	var err error

	for {
		columnsString, err = getInputs(reader, "How many columns?")
		if err != nil {
			fmt.Println("Invalid input! Try again.")
			continue
		}

		columns, err = strconv.Atoi(columnsString)
		if err != nil {
			fmt.Println("Invalid input! Only integers are allowed, try again.")
			continue
		}

		if columns <= 0 {
			fmt.Println("Invalid input! You cant use 0 or less as \nan input, try agian.")
			continue
		}

		break
	}

	return columns
}
