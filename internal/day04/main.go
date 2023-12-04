// Package day04 is our package for the 4th AoC day
package day04

import (
	"math"
	"strings"
)

func evaluateCard(input string) int {
	card := strings.Split(input, ":")
	sections := strings.Split(card[1], "|")
	winNumbers := make(map[string]bool, 100)
	cardSum := 0
	for _, w := range strings.Split(sections[0], " ") {
		if w == "" {
			continue
		}
		winNumbers[w] = true
	}
	for _, n := range strings.Split(sections[1], " ") {
		if n == "" {
			continue
		}
		_, found := winNumbers[n]
		if found {
			cardSum++
		}
	}
	return cardSum
}

func recursiveSum(input []int) int {
	sum := 0
	for i := 1; i <= input[0]; i++ {
		sum += recursiveSum(input[i:])
	}
	return 1 + sum
}

// RunPart1 is for the first star of the day
func RunPart1(input []string) int {
	sum := 0
	for _, line := range input {
		sum += int(math.Pow(2, float64(evaluateCard(line)-1)))
	}
	return sum
}

// RunPart2 is for the second star of the day
func RunPart2(input []string) int {
	var cards []int
	for _, line := range input {
		cards = append(cards, evaluateCard(line))
	}
	sum := 0
	for k := range cards {
		sum += recursiveSum(cards[k:])
	}
	return sum
}
