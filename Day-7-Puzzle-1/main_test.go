package main

import (
	"testing"
)

func TestCalculateSolution(t *testing.T) {
	Input := []uint64{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	expectedAnswer := uint64(37)

	answer, err := calcSolution(Input)
	if err != nil {
		t.Fatalf("calcSolution error'd: %v", err)
	}

	if expectedAnswer != answer {
		t.Errorf("Answer 18 wrong\nExpected: %v\n     Got: %v", expectedAnswer, answer)
	}

}
