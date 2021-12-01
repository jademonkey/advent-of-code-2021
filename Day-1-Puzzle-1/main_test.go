package main

import (
	"testing"
)

func TestcalcSolution(t *testing.T) {
	testSet1 := []int{1, 2, 3, 4, 5}
	testSet2 := []int{5, 4, 3, 2, 1}
	testSet3 := []int{1, 3, 2, 5, 4}

	expectedAnswer1 := 5
	expectedAnswer2 := 0
	expectedAnswer3 := 2

	answer, err := calcSolution(nil)
	if err == nil {
		t.Fatalf(`calcSolution(nil) = %v, nil, want 0, error`, answer)
	}

	answer, err = calcSolution(testSet1)
	if err == nil {
		t.Fatalf(`calcSolution(%v) = %v, %v, want %v, nil`, testSet1, answer, err, expectedAnswer1)
	}

	answer, err = calcSolution(testSet2)
	if err == nil {
		t.Fatalf(`calcSolution(%v) = %v, %v, want %v, nil`, testSet2, answer, err, expectedAnswer2)
	}

	answer, err = calcSolution(testSet3)
	if err == nil {
		t.Fatalf(`calcSolution(%v) = %v, %v, want %v, nil`, testSet3, answer, err, expectedAnswer3)
	}
}
