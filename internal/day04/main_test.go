package day04

import (
	"testing"

	"github.com/unixlab/AoC2023/internal/aoeinput"
)

func TestRunPart1Example(t *testing.T) {
	expectedResult := 13
	actualResult := RunPart1(aoeinput.Read("../../", "day04", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart1(t *testing.T) {
	expectedResult := 21959
	actualResult := RunPart1(aoeinput.Read("../../", "day04", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2Example(t *testing.T) {
	expectedResult := 30
	actualResult := RunPart2(aoeinput.Read("../../", "day04", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2(t *testing.T) {
	expectedResult := 5132675
	actualResult := RunPart2(aoeinput.Read("../../", "day04", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}
