package main

import (
	"testing"

	"github.com/jademonkey/advent-of-code-2021/robcommon"
)

func TestCalculateSolution(t *testing.T) {
	expectedAnswer := 15

	Input, err := robcommon.ReadHeightMap("testInput")
	if err != nil {
		t.Fatalf("ReadSegmentDisplays error'd: %v", err)
	}

	answer, err := calcSolution(Input)
	if err != nil {
		t.Fatalf("calcSolution error'd: %v", err)
	}

	if expectedAnswer != answer {
		t.Errorf("Answer wrong\nExpected: %v\n     Got: %v", expectedAnswer, answer)
	}

}

func TestReadHeightMap(t *testing.T) {

	expectedAnswer := [][]int{{2, 1, 9, 9, 9, 4, 3, 2, 1, 0}, {3, 9, 8, 7, 8, 9, 4, 9, 2, 1}, {9, 8, 5, 6, 7, 8, 9, 8, 9, 2}, {8, 7, 6, 7, 8, 9, 6, 7, 8, 9}, {9, 8, 9, 9, 9, 6, 5, 6, 7, 8}}

	answer, err := robcommon.ReadHeightMap("testInput")
	if err != nil {
		t.Fatalf("ReadHeightMap error'd: %v", err)
	}

	if !robcommon.CompareHeightMap(expectedAnswer, answer) {
		t.Errorf("Read wrong\nExpected: %v\n     Got: %v", expectedAnswer, answer)
	}
}

func TestReadHeightMap2(t *testing.T) {
	expectedAnswer := [][]int{
		{6, 5, 4, 6, 7, 9, 8, 7, 8, 9, 1, 2, 3, 5, 6, 7, 9, 6, 5, 5, 6, 7, 8, 9, 5, 3, 2, 3, 9, 4, 3, 2, 1, 2, 3, 4, 5, 8, 9, 2, 1, 2, 9, 7, 6, 3, 2, 3, 5, 9, 9, 9, 4, 3, 2, 3, 4, 5, 6, 7, 8, 9, 7, 6, 4, 3, 2, 3, 4, 5, 6, 7, 9, 9, 8, 7, 6, 5, 4, 5, 2, 4, 5, 6, 7, 9, 9, 8, 7, 6, 5, 3, 2, 5, 6, 8, 9, 2, 3, 6},
		{5, 4, 3, 5, 7, 9, 9, 6, 5, 4, 0, 1, 2, 4, 5, 9, 9, 4, 3, 4, 5, 6, 7, 8, 9, 5, 1, 9, 8, 9, 4, 3, 0, 1, 2, 3, 4, 6, 7, 9, 0, 9, 8, 6, 5, 4, 5, 6, 9, 8, 9, 8, 9, 2, 1, 2, 3, 4, 5, 8, 9, 7, 6, 5, 3, 2, 1, 2, 3, 4, 5, 8, 9, 8, 7, 8, 5, 4, 3, 2, 1, 3, 5, 6, 7, 8, 9, 9, 8, 8, 6, 4, 3, 4, 9, 9, 0, 1, 3, 5},
		{6, 6, 5, 6, 9, 8, 7, 6, 4, 3, 2, 3, 4, 7, 6, 7, 8, 9, 2, 3, 4, 5, 9, 9, 5, 4, 3, 9, 7, 8, 9, 2, 1, 3, 3, 4, 5, 8, 9, 9, 9, 9, 8, 7, 6, 5, 6, 9, 9, 7, 8, 7, 8, 9, 2, 3, 4, 5, 6, 9, 9, 8, 7, 6, 5, 4, 2, 3, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 1, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 5, 6, 7, 8, 9, 2, 9, 6}}

	answer, err := robcommon.ReadHeightMap("testInput2")
	if err != nil {
		t.Fatalf("ReadHeightMap error'd: %v", err)
	}

	if !robcommon.CompareHeightMap(expectedAnswer, answer) {
		t.Errorf("Read wrong\nExpected: %v\n     Got: %v", expectedAnswer, answer)
	}
}

func TestCalculateSolution2(t *testing.T) {
	input := [][]int{
		{6, 5, 4, 6, 7, 9, 8, 7, 8, 9, 1, 2, 3, 5, 6, 7, 9, 6, 5, 5, 6, 7, 8, 9, 5, 3, 2, 3, 9, 4, 3, 2, 1, 2, 3, 4, 5, 8, 9, 2, 1, 2, 9, 7, 6, 3, 2, 3, 5, 9, 9, 9, 4, 3, 2, 3, 4, 5, 6, 7, 8, 9, 7, 6, 4, 3, 2, 3, 4, 5, 6, 7, 9, 9, 8, 7, 6, 5, 4, 5, 2, 4, 5, 6, 7, 9, 9, 8, 7, 6, 5, 3, 2, 5, 6, 8, 9, 2, 3, 6},
		{5, 4, 3, 5, 7, 9, 9, 6, 5, 4, 0, 1, 2, 4, 5, 9, 9, 4, 3, 4, 5, 6, 7, 8, 9, 5, 1, 9, 8, 9, 4, 3, 0, 1, 2, 3, 4, 6, 7, 9, 0, 9, 8, 6, 5, 4, 5, 6, 9, 8, 9, 8, 9, 2, 1, 2, 3, 4, 5, 8, 9, 7, 6, 5, 3, 2, 1, 2, 3, 4, 5, 8, 9, 8, 7, 8, 5, 4, 3, 2, 1, 3, 5, 6, 7, 8, 9, 9, 8, 8, 6, 4, 3, 4, 9, 9, 0, 1, 3, 5},
		{6, 6, 5, 6, 9, 8, 7, 6, 4, 3, 2, 3, 4, 7, 6, 7, 8, 9, 2, 3, 4, 5, 9, 9, 5, 4, 3, 9, 7, 8, 9, 2, 1, 3, 3, 4, 5, 8, 9, 9, 9, 9, 8, 7, 6, 5, 6, 9, 9, 7, 8, 7, 8, 9, 2, 3, 4, 5, 6, 9, 9, 8, 7, 6, 5, 4, 2, 3, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 1, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 5, 6, 7, 8, 9, 2, 9, 6}}

	expectedAnswer := 3 + 3 + 4 + 1 + 2 + 1 + 1 + 2 + 2 + 1 + 3 + 8 + 8 + 8 + 1

	answer, err := calcSolution(input)
	if err != nil {
		t.Fatalf("calcSolution error'd: %v", err)
	}

	if expectedAnswer != answer {
		t.Errorf("Answer wrong\nExpected: %v\n     Got: %v", expectedAnswer, answer)
	}
}
