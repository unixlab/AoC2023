// Package day06 is our package for the 6th AoC day
package day06

import (
	"strconv"
	"strings"
)

func parseRaces(input []string, ignoreSpaces bool) ([]int, []int) {
	var time []int
	var distance []int
	for _, line := range input {
		var counter []int
		if ignoreSpaces {
			line = strings.ReplaceAll(line, " ", "")
		}
		counters := strings.Split(line, ":")
		dataPoints := strings.Split(counters[1], " ")
		for _, dataPoint := range dataPoints {
			if dataPoint == "" {
				continue
			}
			data, _ := strconv.Atoi(dataPoint)
			counter = append(counter, data)
		}
		if strings.HasPrefix(line, "Time") {
			time = counter
		}
		if strings.HasPrefix(line, "Distance") {
			distance = counter
		}
	}
	return time, distance
}

func calcLeftCutoff(time int, distance int) int {
	for j := 1; j < time; j++ {
		speed := j
		runTime := time - j
		traveled := speed * runTime
		if traveled > distance {
			return j
		}
	}
	return -1
}

func calcRightCutoff(time int, distance int) int {
	for j := time - 1; j > 0; j-- {
		speed := j
		runTime := time - j
		traveled := speed * runTime
		if traveled > distance {
			return j
		}
	}
	return -1
}

// RunPart1 is for the first star of the day
func RunPart1(input []string) int {
	sum := 1
	time, distance := parseRaces(input, false)
	for i := 0; i < len(time); i++ {
		leftCutoff := calcLeftCutoff(time[i], distance[i])
		rightCutoff := calcRightCutoff(time[i], distance[i])
		sum *= rightCutoff - leftCutoff + 1
	}
	return sum
}

// RunPart2 is for the second star of the day
func RunPart2(input []string) int {
	time, distance := parseRaces(input, true)
	leftCutoff := calcLeftCutoff(time[0], distance[0])
	rightCutoff := calcRightCutoff(time[0], distance[0])
	return rightCutoff - leftCutoff + 1
}
