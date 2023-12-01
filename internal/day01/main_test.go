package day01

import (
	"testing"

	"github.com/unixlab/AoC2023/internal/aoeinput"
)

func TestRunPart1Example(t *testing.T) {
	expectedResult := 142
	actualResult := RunPart1(aoeinput.Read("../../", "day01p1", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart1(t *testing.T) {
	expectedResult := 54338
	actualResult := RunPart1(aoeinput.Read("../../", "day01", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2Example(t *testing.T) {
	expectedResult := 281
	actualResult := RunPart2(aoeinput.Read("../../", "day01p2", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2(t *testing.T) {
	expectedResult := 53389
	actualResult := RunPart2(aoeinput.Read("../../", "day01", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}
