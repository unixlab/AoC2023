// Package day09 is our package for the 9th AoC day
package day09

import (
	"strconv"
	"strings"
)

// Reading is our own int array struct to build functions on
type Reading []int

// AllZero returns true if all numbers in the int array are 0
// otherwise returns false
func (r Reading) AllZero() bool {
	for _, v := range r {
		if v != 0 {
			return false
		}
	}
	return true
}

// RunPart1 is for the first star of the day
func RunPart1(input []string) int {
	sum := 0
	for _, line := range input {
		var readings []Reading
		var initialReading Reading
		for _, v := range strings.Fields(line) {
			n, _ := strconv.Atoi(v)
			initialReading = append(initialReading, n)
		}
		readings = append(readings, initialReading)
		for i := 0; i < len(readings); i++ {
			var newReading Reading
			for y := 0; y < len(readings[i])-1; y++ {
				newReading = append(newReading, readings[i][y+1]-readings[i][y])
			}
			if !newReading.AllZero() {
				readings = append(readings, newReading)
			}
		}
		for _, r := range readings {
			sum += r[len(r)-1]
		}
	}
	return sum
}

// RunPart2 is for the second star of the day
func RunPart2(input []string) int {
	sum := 0
	for _, line := range input {
		var readings []Reading
		var initialReading Reading
		for _, v := range strings.Fields(line) {
			n, _ := strconv.Atoi(v)
			initialReading = append(initialReading, n)
		}
		readings = append(readings, initialReading)
		for i := 0; i < len(readings); i++ {
			var newReading Reading
			for y := 0; y < len(readings[i])-1; y++ {
				newReading = append(newReading, readings[i][y+1]-readings[i][y])
			}
			if !newReading.AllZero() {
				readings = append(readings, newReading)
			}
		}
		last := 0
		for i := len(readings) - 1; i >= 0; i-- {
			last = readings[i][0] - last
		}
		sum += last
	}
	return sum
}
