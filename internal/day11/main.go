// Package day11 is our package for the 11th AoC day
package day11

import (
	"math"
	"strings"

	"gonum.org/v1/gonum/stat/combin"
)

// Cord is our struct representing a coordinate
type Cord struct {
	Y int
	X int
}

// AbsDiffInt calculates the absolut difference between to integers and returns an integer
func AbsDiffInt(a, b int) int {
	return int(math.Abs(math.Abs(float64(a)) - math.Abs(float64(b))))
}

// Run is for both stars
func Run(input []string, increase int) int {
	// We need to subtract one, because the initial line is already there
	increase--
	var grid [][]string
	var galaxies []Cord
	var rowsToAdd []int
	var colsToAdd []int
	for y, line := range input {
		splitData := strings.Split(line, "")
		grid = append(grid, splitData)
		for x, char := range splitData {
			if char == "#" {
				galaxies = append(galaxies, Cord{Y: y, X: x})
			}
		}
		if strings.Count(line, "#") == 0 {
			rowsToAdd = append(rowsToAdd, y)
		}
	}
	for x := 0; x < len(grid[0]); x++ {
		var row strings.Builder
		for y := 0; y < len(grid); y++ {
			row.WriteString(grid[y][x])
		}
		if strings.Count(row.String(), "#") == 0 {
			colsToAdd = append(colsToAdd, x)
		}
	}
	for offset, row := range rowsToAdd {
		for i := 0; i < len(galaxies); i++ {
			if galaxies[i].Y > row+(offset*increase) {
				galaxies[i].Y += increase
			}
		}
	}
	for offset, col := range colsToAdd {
		for i := 0; i < len(galaxies); i++ {
			if galaxies[i].X > col+(offset*increase) {
				galaxies[i].X += increase
			}
		}
	}
	combs := combin.Combinations(len(galaxies), 2)
	sum := 0
	for _, comb := range combs {
		sum += AbsDiffInt(galaxies[comb[0]].X, galaxies[comb[1]].X)
		sum += AbsDiffInt(galaxies[comb[0]].Y, galaxies[comb[1]].Y)
	}
	return sum
}
