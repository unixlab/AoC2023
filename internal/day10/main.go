// Package day10 is our package for the 10th AoC day
package day10

import (
	"errors"
	"fmt"
	"strings"

	"github.com/RyanCarrier/dijkstra/v2"
)

func calculateIndex(length, y, x int) int {
	return y*length + x
}

func calcPart1(grid [][]string) (int, []int) {
	start := 0
	graph := dijkstra.NewGraph()
	vertices := make(map[int]bool)
	for y := 0; y < len(grid); y += 2 {
		xLen := len(grid[y])
		for x := 0; x < len(grid[y]); x += 2 {
			cidx := calculateIndex(xLen, y, x)
			arcs := make(map[int]uint64)
			switch grid[y][x] {
			case ".":
				continue
			case "S":
				start = cidx
				if cidx-1 > 0 {
					arcs[cidx-1] = 1
				}
				arcs[cidx+1] = 1
				if cidx-xLen > 0 {
					arcs[cidx-xLen] = 1
				}
				arcs[cidx+xLen] = 1
			case "-":
				if cidx-1 > 0 {
					arcs[cidx-1] = 1
				}
				arcs[cidx+1] = 1
			case "|":
				if cidx-xLen > 0 {
					arcs[cidx-xLen] = 1
				}
				arcs[cidx+xLen] = 1
			case "F":
				arcs[cidx+1] = 1
				arcs[cidx+xLen] = 1
			case "7":
				if cidx-1 > 0 {
					arcs[cidx-1] = 1
				}
				arcs[cidx+xLen] = 1
			case "J":
				if cidx-1 > 0 {
					arcs[cidx-1] = 1
				}
				if cidx-xLen > 0 {
					arcs[cidx-xLen] = 1
				}
			case "L":
				if cidx-xLen > 0 {
					arcs[cidx-xLen] = 1
				}
				arcs[cidx+1] = 1
			}
			vertices[cidx] = true
			_ = graph.AddEmptyVertex(cidx)
			for arc, distance := range arcs {
				_, err := graph.GetVertexArcs(arc)
				if errors.Is(err, dijkstra.ErrVertexNotFound) {
					_ = graph.AddEmptyVertex(arc)
				}
				_ = graph.AddArc(cidx, arc, distance)
				_ = graph.AddArc(arc, cidx, distance)
			}
		}
	}
	var maxDistanceVertex uint64
	var bestPath []int
	maxDistanceVertex = 0
	for vertex := range vertices {
		arcs, _ := graph.GetVertexArcs(vertex)
		if len(arcs) == 2 {
			path, _ := graph.Shortest(start, vertex)
			if path.Distance > maxDistanceVertex {
				maxDistanceVertex = path.Distance
				bestPath = path.Path
			}
			for i := 1; i < len(path.Path); i++ {
				delete(vertices, path.Path[i])
			}
		} else {
			delete(vertices, vertex)
		}
	}
	for i := 0; i < len(bestPath); i += 2 {
		fmt.Printf("%d ", bestPath[i]/2)
	}
	fmt.Println()
	return int(maxDistanceVertex) / 2, bestPath
}

// RunPart1 is for the first star of the day
func RunPart1(input []string) int {
	var grid [][]string
	var result int
	for _, line := range input {
		if line == "" {
			result, _ = calcPart1(grid)
			grid = [][]string{}
			continue
		}
		var row []string
		var emptyRow []string
		for _, char := range strings.Split(line, "") {
			row = append(row, char, "#")
			emptyRow = append(emptyRow, "#", "#")
		}
		grid = append(grid, row, emptyRow)
	}
	return result
}

// RunPart2 is for the second star of the day
func RunPart2(input []string) int {
	for _, line := range input {
		fmt.Println(line)
	}
	return 0
}
