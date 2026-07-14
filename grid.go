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

func formatGrid() string {
	var gridFormat strings.Builder

	grid := setGrid()

	for i := range grid {
		for j := range grid[i] {
			fmt.Fprintf(&gridFormat, "(%v, %v)", grid[i][j].X, grid[i][j].Y)
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

		break
	}

	return columns
}
