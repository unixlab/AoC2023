// Package day01 is our package for the 1st AoC day
package day01

import (
	"strconv"
	"strings"
)

type Numbers []Number

type Number struct {
	Identifier  string
	Replacement string
}

func (n Numbers) findAny(input string) bool {
	for _, num := range n {
		idx := strings.Index(input, num.Identifier)
		if idx >= 0 {
			return true
		}
	}
	return false
}

func (n Numbers) replace(input string) string {
	for _, num := range n {
		input = strings.Replace(input, num.Identifier, num.Replacement, 1)
	}
	return input
}

// RunPart1 is for the first star of the day
func RunPart1(input []string) int {
	sum := 0
	for _, line := range input {
		firstNumber := 0
		for i := 0; i < len(line); i++ {
			number, err := strconv.Atoi(line[i : i+1])
			if err == nil && number > 0 {
				firstNumber = number
				break
			}
		}
		secondNumber := 0
		for j := len(line); j > 0; j-- {
			number, err := strconv.Atoi(line[j-1 : j])
			if err == nil && number > 0 {
				secondNumber = number
				break
			}
		}
		sum += firstNumber*10 + secondNumber
	}
	return sum
}

// RunPart2 is for the second star of the day
func RunPart2(input []string) int {
	numbers := Numbers{
		{"1", "1"},
		{"2", "2"},
		{"3", "3"},
		{"4", "4"},
		{"5", "5"},
		{"6", "6"},
		{"7", "7"},
		{"8", "8"},
		{"9", "9"},
		{"one", "1"},
		{"two", "2"},
		{"three", "3"},
		{"four", "4"},
		{"five", "5"},
		{"six", "6"},
		{"seven", "7"},
		{"eight", "8"},
		{"nine", "9"},
	}
	sum := 0
	for _, line := range input {
		extract := ""
		firstNumber := 0
		for i := 0; i <= len(line); i++ {
			extract = line[0:i]
			if numbers.findAny(extract) {
				break
			}
		}
		extract = numbers.replace(extract)
		for i := 0; i < len(extract); i++ {
			number, err := strconv.Atoi(extract[i : i+1])
			if err == nil && number > 0 {
				firstNumber = number
				break
			}
		}
		secondNumber := 0
		for j := len(line) - 1; j >= 0; j-- {
			extract = line[j:]
			if numbers.findAny(extract) {
				break
			}
		}
		extract = numbers.replace(extract)
		for j := len(extract) - 1; j >= 0; j-- {
			number, err := strconv.Atoi(extract[j : j+1])
			if err == nil && number > 0 {
				secondNumber = number
				break
			}
		}
		sum += firstNumber*10 + secondNumber
	}
	return sum
}
