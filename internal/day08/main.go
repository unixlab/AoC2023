// Package day08 is our package for the 8th AoC day
package day08

import (
	"strings"
)

// Node is our struct to hold our nodes
type Node struct {
	Left  string
	Right string
}

func parseInput(input []string) ([]string, map[string]Node) {
	var instructions []string
	nodes := make(map[string]Node, 1000)
	for _, line := range input {
		if strings.Index(line, "=") > 0 {
			node := strings.Split(line, "=")[0]
			node = strings.ReplaceAll(node, " ", "")
			line = strings.Split(line, "=")[1]
			line = strings.ReplaceAll(line, "(", "")
			line = strings.ReplaceAll(line, ")", "")
			line = strings.ReplaceAll(line, " ", "")
			var newNode Node
			newNode.Left = strings.Split(line, ",")[0]
			newNode.Right = strings.Split(line, ",")[1]
			nodes[node] = newNode
		} else if line == "" {
			continue
		} else {
			instructions = strings.Split(line, "")
		}
	}
	return instructions, nodes
}

func greatestCommonDivisor(num1 int, num2 int) int {
	for num2 > 0 {
		store := num2
		num2 = num1 % num2
		num1 = store
	}
	return num1
}

func leastCommonMultiplier(numN ...int) int {
	result := numN[0] * numN[1] / greatestCommonDivisor(numN[0], numN[1])
	if len(numN) > 2 {
		for i := 2; i < len(numN); i++ {
			result = leastCommonMultiplier(result, numN[i])
		}
	}
	return result
}

// RunPart1 is for the first star of the day
func RunPart1(input []string) int {
	instructions, nodes := parseInput(input)
	steps := 0
	currentNode := "AAA"
	instructionIdx := 0
	for currentNode != "ZZZ" {
		if instructions[instructionIdx] == "L" {
			currentNode = nodes[currentNode].Left
		} else {
			currentNode = nodes[currentNode].Right
		}
		steps++
		instructionIdx = steps % len(instructions)
	}
	return steps
}

// RunPart2 is for the second star of the day
func RunPart2(input []string) int {
	instructions, nodes := parseInput(input)
	steps := 0
	var currentNodes []string
	var nodeSteps []int
	for node := range nodes {
		if strings.HasSuffix(node, "A") {
			currentNodes = append(currentNodes, node)
			nodeSteps = append(nodeSteps, 0)
		}
	}
	zCounter := 0
	instructionIdx := 0
	for zCounter != len(currentNodes) {
		for i := range currentNodes {
			if instructions[instructionIdx] == "L" {
				currentNodes[i] = nodes[currentNodes[i]].Left
			} else {
				currentNodes[i] = nodes[currentNodes[i]].Right
			}
		}
		for i, node := range currentNodes {
			if strings.HasSuffix(node, "Z") {
				if nodeSteps[i] == 0 {
					nodeSteps[i] = steps + 1
					zCounter++
				}
			}
		}
		steps++
		instructionIdx = steps % len(instructions)
	}
	return leastCommonMultiplier(nodeSteps...)
}
