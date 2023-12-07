package day07

import (
	"testing"

	"github.com/unixlab/AoC2023/internal/aoeinput"
)

func TestRunPart1Example(t *testing.T) {
	expectedResult := 6440
	actualResult := RunPart1(aoeinput.Read("../../", "day07", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart1(t *testing.T) {
	expectedResult := 246912307
	actualResult := RunPart1(aoeinput.Read("../../", "day07", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2Example(t *testing.T) {
	expectedResult := 5905
	actualResult := RunPart2(aoeinput.Read("../../", "day07", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2(t *testing.T) {
	expectedResult := 246894760
	actualResult := RunPart2(aoeinput.Read("../../", "day07", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestGetCardValue(t *testing.T) {
	expectedResult := -1
	actualResult := getCardValue("F")
	if actualResult != expectedResult {
		t.Fatalf("getCardValue() = expected %v, got %v", expectedResult, actualResult)
	}
}
