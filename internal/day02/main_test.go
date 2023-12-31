package day02

import (
	"testing"

	"github.com/unixlab/AoC2023/internal/aoeinput"
)

func TestRunPart1Example(t *testing.T) {
	expectedResult := 8
	actualResult := RunPart1(aoeinput.Read("../../", "day02", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart1(t *testing.T) {
	expectedResult := 2204
	actualResult := RunPart1(aoeinput.Read("../../", "day02", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2Example(t *testing.T) {
	expectedResult := 2286
	actualResult := RunPart2(aoeinput.Read("../../", "day02", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2(t *testing.T) {
	expectedResult := 71036
	actualResult := RunPart2(aoeinput.Read("../../", "day02", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}
