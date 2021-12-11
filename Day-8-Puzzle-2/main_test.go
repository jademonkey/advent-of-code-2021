package main

import (
	"testing"
)

func TestCalculateSolutionShort(t *testing.T) {
	expectedAnswer := 5353

	Input, err := ReadSegmentDisplays("testInputShort")
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

func TestCalculateSolution(t *testing.T) {
	expectedAnswer := 61229

	Input, err := ReadSegmentDisplays("testInput")
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

func TestReadSegmentDisplaysShort(t *testing.T) {

	AllSegs := [10]string{"ab", "abd", "abef", "bcdef", "acdfg", "abcdf", "abcdef", "bcdefg", "abcdeg", "abcdefg"}
	AllDigs := [4]string{"bcdef", "abcdf", "bcdef", "abcdf"}

	expectedAnswer := []SegmentDisplay{SegmentDisplay{Seg: AllSegs, Digit: AllDigs}}

	answer, err := ReadSegmentDisplays("testInputShort")
	if err != nil {
		t.Fatalf("calcSolution error'd: %v", err)
	}

	if !CompareSegmentDisplay(expectedAnswer, answer) {
		t.Errorf("Read wrong\nExpected: %v\n     Got: %v", expectedAnswer, answer)
	}
}

func CompareSegmentDisplay(one, two []SegmentDisplay) bool {
	if one == nil || two == nil {
		return false
	}
	if len(one) != len(two) {
		return false
	}

	for i := 0; i < len(one); i++ {
		for i2 := 0; i2 < len(one[i].Digit); i2++ {
			if one[i].Digit[i2] != two[i].Digit[i2] {
				return false
			}
		}
		for i2 := 0; i2 < len(one[i].Seg); i2++ {
			if one[i].Seg[i2] != two[i].Seg[i2] {
				return false
			}
		}
	}

	return true
}
