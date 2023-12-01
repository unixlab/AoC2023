// Package day01 is our package for the 1st AoC day
package day01

import (
	"strconv"
	"strings"
)

// Numbers is an array of our internal struct Number
// We need a struct in order to write a function on the array
type Numbers []Number

// Number is our struct for finding and replacing numbers in the input string
type Number struct {
	Identifier  string
	Replacement string
}

// FindAny number (digit or word) in the input string
func (n Numbers) FindAny(input string) bool {
	for _, num := range n {
		idx := strings.Index(input, num.Identifier)
		if idx >= 0 {
			return true
		}
	}
	return false
}

// Replace word numbers with digits, Replace digits with itself to simply the code
func (n Numbers) Replace(input string) string {
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
		// go from left(start) to right and break at the first digit
		for i := 0; i < len(line); i++ {
			number, err := strconv.Atoi(line[i : i+1])
			if err == nil && number > 0 {
				firstNumber = number
				break
			}
		}
		secondNumber := 0
		// go from right(end) to left and break at the first digit
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
	// setup replacement and find map
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
		// go from left(start) to right and stop at the first digit or spelled out digit
		for i := 0; i <= len(line); i++ {
			extract = line[0:i]
			if numbers.FindAny(extract) {
				break
			}
		}
		// extract the substring from the start until the found digit/spelled out digit
		extract = numbers.Replace(extract)
		// go from left(start) to right, try to parse the found character as digit and stop at the first found digit
		for i := 0; i < len(extract); i++ {
			number, err := strconv.Atoi(extract[i : i+1])
			if err == nil && number > 0 {
				firstNumber = number
				break
			}
		}
		secondNumber := 0
		// go from right(end) to left and stop at the first digit or spelled out digit
		for j := len(line) - 1; j >= 0; j-- {
			extract = line[j:]
			if numbers.FindAny(extract) {
				break
			}
		}
		// extract the substring from the end until the found digit/spelled out digit
		extract = numbers.Replace(extract)
		// go from right(end) to left, try to parse the found character as digit and stop at the first found digit
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
