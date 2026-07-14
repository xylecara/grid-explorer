package main

import (
	"fmt"
	"math/rand/v2"
)

func randomizeTreasure(grid [][]coords) string {
	randomRowIndex := rand.N(len(grid))
	randomRow := grid[randomRowIndex]

	randomColumnIndex := rand.N(len(randomRow))
	randomValue := randomRow[randomColumnIndex]
	
	return fmt.Sprintf("The random value is (%v, %v)", randomValue.X, randomValue.Y) 	
}