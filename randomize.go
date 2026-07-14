package main

import (
	"math/rand/v2"
)

func randomizePosition(grid [][]coords) coords {
	randomRowIndex := rand.N(len(grid))
	randomRow := grid[randomRowIndex]

	randomColumnIndex := rand.N(len(randomRow))
	randomValue := randomRow[randomColumnIndex]
	
	return randomValue
}