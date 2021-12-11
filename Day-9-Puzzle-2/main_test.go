package main

import (
	"testing"

	"github.com/jademonkey/advent-of-code-2021/robcommon"
)

func TestCalculateSolution(t *testing.T) {
	expectedAnswer := 1134

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

	if !CompareHeightMap(expectedAnswer, answer) {
		t.Errorf("Read wrong\nExpected: %v\n     Got: %v", expectedAnswer, answer)
	}
}

func CompareHeightMap(one, two [][]int) bool {
	if one == nil || two == nil {
		return false
	}
	if len(one) != len(two) {
		return false
	}

	for i := 0; i < len(one); i++ {
		for i2 := 0; i2 < len(one[i]); i2++ {
			if one[i][i2] != two[i][i2] {
				return false
			}
		}
	}

	return true
}
