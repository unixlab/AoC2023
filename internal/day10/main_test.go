package day10

import (
	"testing"

	"github.com/unixlab/AoC2023/internal/aoeinput"
)

func TestRunPart1Example(t *testing.T) {
	expectedResult := 8
	actualResult := RunPart1(aoeinput.Read("../../", "day10", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart1(t *testing.T) {
	expectedResult := 6690
	actualResult := RunPart1(aoeinput.Read("../../", "day10", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2Example(t *testing.T) {
	expectedResult := 0
	actualResult := RunPart2(aoeinput.Read("../../", "day10", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2(t *testing.T) {
	expectedResult := 0
	actualResult := RunPart2(aoeinput.Read("../../", "day10", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestLeftPossible(t *testing.T) {
	actualResult := leftPossible("S")
	if !actualResult {
		t.Fatalf("RunPart2() = expected %v, got %v", true, actualResult)
	}
	actualResult = leftPossible("x")
	if actualResult {
		t.Fatalf("RunPart2() = expected %v, got %v", false, actualResult)
	}
}

func TestBottomPossible(t *testing.T) {
	actualResult := bottomPossible("S")
	if !actualResult {
		t.Fatalf("RunPart2() = expected %v, got %v", true, actualResult)
	}
}

func TestFindS(t *testing.T) {
	var grid Grid
	expectedResult := Cord{Y: -1, X: -1}
	actualResult := grid.FindS()
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestFindNextBestPath(t *testing.T) {
	var grid Grid
	grid = append(grid, []string{"", ""})
	grid = append(grid, []string{"", ""})
	expectedResult := Cord{Y: -1, X: -1}
	actualResult := grid.FindNextBestPath(Cord{Y: 0, X: 0})
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
	grid[0][0] = "S"
	grid[0][1] = "S"
	expectedResult = Cord{Y: 0, X: 1}
	actualResult = grid.FindNextBestPath(Cord{Y: 0, X: 0})
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
	expectedResult = Cord{Y: 0, X: 0}
	actualResult = grid.FindNextBestPath(Cord{Y: 0, X: 1})
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
	grid[0][0] = "S"
	grid[0][1] = ""
	grid[1][0] = "S"
	expectedResult = Cord{Y: 1, X: 0}
	actualResult = grid.FindNextBestPath(Cord{Y: 0, X: 0})
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}
