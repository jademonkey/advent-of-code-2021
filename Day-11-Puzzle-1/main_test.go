package main

import (
	"testing"

	util "github.com/jademonkey/advent-of-code-2021/robcommon"
)

func TestCalculateSolution(t *testing.T) {
	expectedAnswer := 1656

	Input, err := util.ReadHeightMap("testInput")
	if err != nil {
		t.Fatalf("ReadChunkMap error'd: %v", err)
	}

	answer, err := calcSolution(Input, 100)
	if err != nil {
		t.Fatalf("calcSolution error'd: %v", err)
	}

	if expectedAnswer != answer {
		t.Errorf("Answer wrong\nExpected: %v\n     Got: %v", expectedAnswer, answer)
	}
}

func TestCalculateSolutionShort(t *testing.T) {
	expectedAnswer := 9

	Input, err := util.ReadHeightMap("testInputShort")
	if err != nil {
		t.Fatalf("ReadChunkMap error'd: %v", err)
	}

	answer, err := calcSolution(Input, 2)
	if err != nil {
		t.Fatalf("calcSolution error'd: %v", err)
	}

	if expectedAnswer != answer {
		t.Errorf("Answer wrong\nExpected: %v\n     Got: %v", expectedAnswer, answer)
	}
}

func TestReadHeighMap(t *testing.T) {

	expectedAnswer := [][]int{{5, 4, 8, 3, 1, 4, 3, 2, 2, 3}, {2, 7, 4, 5, 8, 5, 4, 7, 1, 1}, {5, 2, 6, 4, 5, 5, 6, 1, 7, 3}, {6, 1, 4, 1, 3, 3, 6, 1, 4, 6}, {6, 3, 5, 7, 3, 8, 5, 4, 7, 8}, {4, 1, 6, 7, 5, 2, 4, 6, 4, 5}, {2, 1, 7, 6, 8, 4, 1, 7, 2, 1}, {6, 8, 8, 2, 8, 8, 1, 1, 3, 4}, {4, 8, 4, 6, 8, 4, 8, 5, 5, 4}, {5, 2, 8, 3, 7, 5, 1, 5, 2, 6}}

	answer, err := util.ReadHeightMap("testInput")
	if err != nil {
		t.Fatalf("calcSolution error'd: %v", err)
	}

	if !util.CompareHeightMap(expectedAnswer, answer) {
		t.Errorf("Read wrong\nExpected: %v\n     Got: %v", expectedAnswer, answer)
	}
}
