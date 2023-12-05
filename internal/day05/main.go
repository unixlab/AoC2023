// Package day05 is our package for the 5th AoC day
package day05

import (
	"sort"
	"strconv"
	"strings"
)

func parseInputMutations(input []string) [][][]int {
	var mutations [][][]int
	var newMutation [][]int
	for _, line := range input {
		if line == "" {
			continue
		}
		_, isNumber := strconv.Atoi(line[0:1])
		if isNumber == nil {
			mutationData := strings.Split(line, " ")
			var newMut []int
			for _, v := range mutationData {
				number, _ := strconv.Atoi(v)
				newMut = append(newMut, number)
			}
			newMutation = append(newMutation, newMut)
		} else {
			if len(newMutation) > 0 {
				mutations = append(mutations, newMutation)
				newMutation = [][]int{}
			}
		}
	}
	return append(mutations, newMutation)
}

func parseSeeds(input []string) []int {
	var seeds []int
	for _, line := range input {
		if strings.HasPrefix(line, "seeds: ") {
			seedLine := strings.Split(line, ":")
			seedLine = strings.Split(strings.TrimSpace(seedLine[1]), " ")
			for _, v := range seedLine {
				number, _ := strconv.Atoi(v)
				seeds = append(seeds, number)
			}
			continue
		}
	}
	return seeds
}

// RunPart1 is for the first star of the day
func RunPart1(input []string) int {
	minSeed := -1
	seeds := parseSeeds(input)
	mutations := parseInputMutations(input)
	for _, seed := range seeds {
		for _, mutation := range mutations {
			for _, mutationData := range mutation {
				dest := mutationData[0]
				source := mutationData[1]
				length := mutationData[2]
				if seed >= source && seed <= source+length {
					seed += dest - source
					break
				}
			}
		}
		if minSeed == -1 || minSeed > seed {
			minSeed = seed
		}
	}
	return minSeed
}

// RunPart2 is for the second star of the day
func RunPart2(input []string) int {
	minSeed := -1
	seeds := parseSeeds(input)
	mutations := parseInputMutations(input)
	for i := 0; i < len(seeds); i += 2 {
		for seedIter := seeds[i]; seedIter <= seeds[i]+seeds[i+1]; seedIter += 0 {
			seed := seedIter
			var offsets []int
			for _, mutation := range mutations {
				for _, mutationData := range mutation {
					dest := mutationData[0]
					source := mutationData[1]
					length := mutationData[2]
					if seed >= source && seed <= source+length {
						offsetFromSource := seed - source
						offsetUntilEnd := length - offsetFromSource
						offsets = append(offsets, offsetUntilEnd)
						seed += dest - source
						break
					}
				}
			}
			sort.Ints(offsets)
			if offsets[0] > 0 {
				seedIter += offsets[0]
			} else {
				seedIter++
			}
			if minSeed == -1 || minSeed > seed {
				minSeed = seed
			}
		}
	}
	return minSeed
}
