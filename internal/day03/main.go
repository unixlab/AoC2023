// Package day03 is our package for the 3rd AoC day
package day03

import (
	"strconv"
	"strings"
)

// Coordinate is our struct for holding a point in our grid
type Coordinate struct {
	Y       int
	X       int
	Counter int
	Value   int
	Char    string
}

func checkAdjacent(grid [][]string, y int, x int) bool {
	startY := y - 1
	endY := y + 1
	startX := x - 1
	endX := x + 1
	if startY < 0 {
		startY = 0
	}
	if endY >= len(grid) {
		endY = len(grid) - 1
	}
	if startX < 0 {
		startX = 0
	}
	if endX >= len(grid[0]) {
		endX = len(grid[0]) - 1
	}
	for checkY := startY; checkY <= endY; checkY++ {
		for checkX := startX; checkX <= endX; checkX++ {
			switch grid[checkY][checkX] {
			case ".":
				continue
			case "0":
				continue
			case "1":
				continue
			case "2":
				continue
			case "3":
				continue
			case "4":
				continue
			case "5":
				continue
			case "6":
				continue
			case "7":
				continue
			case "8":
				continue
			case "9":
				continue
			default:
				return true
			}
		}
	}
	return false
}

func isNumber(input string) bool {
	switch input {
	case "0":
		return true
	case "1":
		return true
	case "2":
		return true
	case "3":
		return true
	case "4":
		return true
	case "5":
		return true
	case "6":
		return true
	case "7":
		return true
	case "8":
		return true
	case "9":
		return true
	}
	return false
}

func intExists(a []int, i int) bool {
	for _, v := range a {
		if v == i {
			return true
		}
	}
	return false
}

// RunPart1 is for the first star of the day
func RunPart1(input []string) int {
	var grid [][]string
	sum := 0
	for _, line := range input {
		row := strings.Split(line, "")
		grid = append(grid, row)
	}
	var number strings.Builder
	for y := 0; y < len(grid); y++ {
		number.Reset()
		isAdjacent := false
		numberFound := false
		for x := 0; x < len(grid[y]); x++ {
			_, numCheck := strconv.Atoi(grid[y][x])
			if numCheck == nil {
				numberFound = true
				number.WriteString(grid[y][x])
				if checkAdjacent(grid, y, x) {
					isAdjacent = true
				}
			} else {
				if numberFound {
					if isAdjacent {
						numberInt, _ := strconv.Atoi(number.String())
						sum += numberInt
					}
					number.Reset()
					isAdjacent = false
					numberFound = false
				}
			}
		}
		if numberFound {
			if isAdjacent {
				numberInt, _ := strconv.Atoi(number.String())
				sum += numberInt
			}
		}
	}
	return sum
}

// RunPart2 is for the second star of the day
func RunPart2(input []string) int {
	var grid [][]Coordinate
	sum := 0
	for y, line := range input {
		row := strings.Split(line, "")
		var newY []Coordinate
		for x, r := range row {
			var newX Coordinate
			newX.X = x
			newX.Y = y
			newX.Counter = 0
			newX.Value = 0
			newX.Char = r
			newY = append(newY, newX)
		}
		grid = append(grid, newY)
	}
	counter := 1
	for y, yV := range grid {
		var number strings.Builder
		inNumber := 0
		for x, xV := range yV {
			if isNumber(xV.Char) {
				number.WriteString(xV.Char)
				inNumber++
			} else {
				if inNumber > 0 {
					for i := x - inNumber; i < x; i++ {
						numberInt, _ := strconv.Atoi(number.String())
						grid[y][i].Counter = counter
						grid[y][i].Value = numberInt
					}
					counter++
				}
				number.Reset()
				inNumber = 0
			}
		}
		if inNumber > 0 {
			for i := len(grid[y]) - inNumber; i < len(grid[y]); i++ {
				numberInt, _ := strconv.Atoi(number.String())
				grid[y][i].Counter = counter
				grid[y][i].Value = numberInt
			}
			counter++
		}
	}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x].Char == "*" {
				startY := y - 1
				endY := y + 1
				startX := x - 1
				endX := x + 1
				var adjacent []Coordinate
				var uniqCounterStore []int
				for checkY := startY; checkY <= endY; checkY++ {
					for checkX := startX; checkX <= endX; checkX++ {
						if grid[checkY][checkX].Value > 0 {
							if !intExists(uniqCounterStore, grid[checkY][checkX].Counter) {
								uniqCounterStore = append(uniqCounterStore, grid[checkY][checkX].Counter)
								adjacent = append(adjacent, grid[checkY][checkX])
							}
						}
					}
				}
				if len(adjacent) == 2 {
					sum += adjacent[0].Value * adjacent[1].Value
				}
			}
		}
	}
	return sum
}
