package day11

import (
	"testing"

	"github.com/unixlab/AoC2023/internal/aoeinput"
)

func TestRunPart1Example(t *testing.T) {
	expectedResult := 374
	actualResult := Run(aoeinput.Read("../../", "day11", true), 2)
	if actualResult != expectedResult {
		t.Fatalf("Run(2) = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart1(t *testing.T) {
	expectedResult := 9693756
	actualResult := Run(aoeinput.Read("../../", "day11", false), 2)
	if actualResult != expectedResult {
		t.Fatalf("Run(2) = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2Example1(t *testing.T) {
	expectedResult := 1030
	actualResult := Run(aoeinput.Read("../../", "day11", true), 10)
	if actualResult != expectedResult {
		t.Fatalf("Run(10) = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2Example2(t *testing.T) {
	expectedResult := 8410
	actualResult := Run(aoeinput.Read("../../", "day11", true), 100)
	if actualResult != expectedResult {
		t.Fatalf("Run(100) = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2Example3(t *testing.T) {
	expectedResult := 82000210
	actualResult := Run(aoeinput.Read("../../", "day11", true), 1000000)
	if actualResult != expectedResult {
		t.Fatalf("Run(1000000) = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2(t *testing.T) {
	expectedResult := 717878258016
	actualResult := Run(aoeinput.Read("../../", "day11", false), 1000000)
	if actualResult != expectedResult {
		t.Fatalf("Run(1000000) = expected %v, got %v", expectedResult, actualResult)
	}
}
