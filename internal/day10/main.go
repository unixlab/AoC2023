// Package day10 is our package for the 10th AoC day
package day10

import (
	"fmt"
	"strings"
)

// Grid is our own string array struct to build function on
type Grid [][]string

// FindS returns the start point in our grid
func (g Grid) FindS() Cord {
	for y, yData := range g {
		for x, xData := range yData {
			if xData == "S" {
				return Cord{X: x, Y: y}
			}
		}
	}
	return Cord{X: -1, Y: -1}
}

func topPossible(char string) bool {
	switch char {
	case "S":
		return true
	case "|":
		return true
	case "J":
		return true
	case "L":
		return true
	}
	return false
}

func rightPossible(char string) bool {
	switch char {
	case "S":
		return true
	case "-":
		return true
	case "L":
		return true
	case "F":
		return true
	}
	return false
}

func bottomPossible(char string) bool {
	switch char {
	case "S":
		return true
	case "|":
		return true
	case "F":
		return true
	case "7":
		return true
	}
	return false
}

func leftPossible(char string) bool {
	switch char {
	case "S":
		return true
	case "-":
		return true
	case "J":
		return true
	case "7":
		return true
	}
	return false
}

// FindNextBestPath returns the next possible path
func (g Grid) FindNextBestPath(c Cord) Cord {
	// top
	if c.Y-1 >= 0 && topPossible(g[c.Y][c.X]) {
		switch g[c.Y-1][c.X] {
		case "|":
			return Cord{Y: c.Y - 1, X: c.X}
		case "7":
			return Cord{Y: c.Y - 1, X: c.X}
		case "F":
			return Cord{Y: c.Y - 1, X: c.X}
		case "S":
			return Cord{Y: c.Y - 1, X: c.X}
		}
	}
	// right
	if c.X+1 < len(g[0]) && rightPossible(g[c.Y][c.X]) {
		switch g[c.Y][c.X+1] {
		case "-":
			return Cord{Y: c.Y, X: c.X + 1}
		case "7":
			return Cord{Y: c.Y, X: c.X + 1}
		case "J":
			return Cord{Y: c.Y, X: c.X + 1}
		case "S":
			return Cord{Y: c.Y, X: c.X + 1}
		}
	}
	// bottom
	if c.Y+1 < len(g) && bottomPossible(g[c.Y][c.X]) {
		switch g[c.Y+1][c.X] {
		case "|":
			return Cord{Y: c.Y + 1, X: c.X}
		case "L":
			return Cord{Y: c.Y + 1, X: c.X}
		case "J":
			return Cord{Y: c.Y + 1, X: c.X}
		case "S":
			return Cord{Y: c.Y + 1, X: c.X}
		}
	}
	// left
	if c.X-1 >= 0 && leftPossible(g[c.Y][c.X]) {
		switch g[c.Y][c.X-1] {
		case "-":
			return Cord{Y: c.Y, X: c.X - 1}
		case "L":
			return Cord{Y: c.Y, X: c.X - 1}
		case "F":
			return Cord{Y: c.Y, X: c.X - 1}
		case "S":
			return Cord{Y: c.Y, X: c.X - 1}
		}
	}
	return Cord{X: -1, Y: -1}
}

// Cord is our struct representing a coordinate
type Cord struct {
	Y int
	X int
}

func readGrids(input []string) []Grid {
	var grids []Grid
	var grid Grid
	for _, line := range input {
		if line == "" {
			grids = append(grids, grid)
			grid = [][]string{}
			continue
		}
		grid = append(grid, strings.Split(line, ""))
	}
	return grids
}

// RunPart1 is for the first star of the day
func RunPart1(input []string) int {
	stepsNeeded := 0
	grids := readGrids(input)
	for _, grid := range grids {
		steps := 0
		s := grid.FindS()
		next := grid.FindNextBestPath(s)
		for s != next {
			steps++
			prev := next
			next = grid.FindNextBestPath(next)
			grid[prev.Y][prev.X] = "x"
		}
		stepsNeeded = (steps + 1) / 2
	}
	return stepsNeeded
}

// RunPart2 is for the second star of the day
func RunPart2(input []string) int {
	for _, line := range input {
		fmt.Println(line)
	}
	return 0
}
