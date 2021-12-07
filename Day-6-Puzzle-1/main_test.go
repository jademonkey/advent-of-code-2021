package main

import (
	"testing"
)

func TestCalculateSolution(t *testing.T) {
	Input := []uint64{3, 4, 3, 1, 2}
	expectedAnswer18 := 26
	expectedAnswer80 := 5934

	answer18, err := calcSolution(Input, 18)
	if err != nil {
		t.Fatalf("calcSolution(18) error'd: %v", err)
	}

	answer80, err := calcSolution(Input, 80)
	if err != nil {
		t.Fatalf("calcSolution(80) error'd: %v", err)
	}
	if expectedAnswer18 != answer18 {
		t.Errorf("Answer 18 wrong\nExpected: %v\n     Got: %v", expectedAnswer18, answer18)
	}
	if expectedAnswer80 != answer80 {
		t.Errorf("Answer 18 wrong\nExpected: %v\n     Got: %v", expectedAnswer80, answer80)
	}
}
