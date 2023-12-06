package day06

import (
	"testing"

	"github.com/unixlab/AoC2023/internal/aoeinput"
)

func TestRunPart1Example(t *testing.T) {
	expectedResult := 288
	actualResult := RunPart1(aoeinput.Read("../../", "day06", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart1(t *testing.T) {
	expectedResult := 170000
	actualResult := RunPart1(aoeinput.Read("../../", "day06", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2Example(t *testing.T) {
	expectedResult := 71503
	actualResult := RunPart2(aoeinput.Read("../../", "day06", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2(t *testing.T) {
	expectedResult := 20537782
	actualResult := RunPart2(aoeinput.Read("../../", "day06", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestCalcLeftCutoff(t *testing.T) {
	expectedResult := -1
	actualResult := calcLeftCutoff(0, 0)
	if actualResult != expectedResult {
		t.Fatalf("calcLeftCutoff() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestCalcRightCutoff(t *testing.T) {
	expectedResult := -1
	actualResult := calcRightCutoff(0, 0)
	if actualResult != expectedResult {
		t.Fatalf("calcRightCutoff() = expected %v, got %v", expectedResult, actualResult)
	}
}
