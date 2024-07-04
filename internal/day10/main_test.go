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
