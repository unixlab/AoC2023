// Package day02 is our package for the 2th AoC day
package day02

import (
	"strconv"
	"strings"
)

// RunPart1 is for the first star of the day
func RunPart1(input []string) int {
	gameSum := 0
	maxRed := 12
	maxGreen := 13
	maxBlue := 14
	for _, game := range input {
		gamePossible := true
		// split string so we can avoid regex
		colonSep := strings.Split(game, ":")
		semicolonSep := strings.Split(colonSep[1], ";")
		// parse gameID
		gameID, _ := strconv.Atoi(colonSep[0][5:])
		// loop over sets
		for _, set := range semicolonSep {
			set = strings.TrimSpace(set)
			sumRed := 0
			sumGreen := 0
			sumBlue := 0
			// loop over draws
			for _, draw := range strings.Split(set, ",") {
				draw = strings.TrimSpace(draw)
				drawData := strings.Split(draw, " ")
				drawCount, _ := strconv.Atoi(drawData[0])
				switch drawData[1] {
				case "red":
					sumRed += drawCount
				case "green":
					sumGreen += drawCount
				case "blue":
					sumBlue += drawCount
				}
				if sumRed > maxRed || sumGreen > maxGreen || sumBlue > maxBlue {
					gamePossible = false
				}
			}
		}
		if gamePossible {
			gameSum += gameID
		}
	}
	return gameSum
}

// RunPart2 is for the second star of the day
func RunPart2(input []string) int {
	gameSum := 0
	for _, game := range input {
		minRed := 0
		minGreen := 0
		minBlue := 0
		// split string so we can avoid regex
		colonSep := strings.Split(game, ":")
		semicolonSep := strings.Split(colonSep[1], ";")
		// loop over sets
		for _, set := range semicolonSep {
			set = strings.TrimSpace(set)
			sumRed := 0
			sumGreen := 0
			sumBlue := 0
			// loop over draws
			for _, draw := range strings.Split(set, ",") {
				draw = strings.TrimSpace(draw)
				drawData := strings.Split(draw, " ")
				drawCount, _ := strconv.Atoi(drawData[0])
				switch drawData[1] {
				case "red":
					sumRed += drawCount
				case "green":
					sumGreen += drawCount
				case "blue":
					sumBlue += drawCount
				}
				if sumRed > minRed {
					minRed = sumRed
				}
				if sumGreen > minGreen {
					minGreen = sumGreen
				}
				if sumBlue > minBlue {
					minBlue = sumBlue
				}
			}
		}
		gameSum += minRed * minGreen * minBlue
	}
	return gameSum
}
