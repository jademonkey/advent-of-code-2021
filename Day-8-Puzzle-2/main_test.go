package main

import (
	"testing"
)

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
		t.Errorf("Answer 18 wrong\nExpected: %v\n     Got: %v", expectedAnswer, answer)
	}

}

func TestReadSegmentDisplays(t *testing.T) {

	AllSegs := [10][]string{{"a", "c", "e", "d", "g", "f", "b"}, {"c", "d", "f", "b", "e"}, {"g", "c", "d", "f", "a"}, {"f", "b", "c", "a", "d"}, {"d", "a", "b"}, {"c", "e", "f", "a", "b", "d"}, {"c", "d", "f", "g", "e", "b"}, {"e", "a", "f", "b"}, {"c", "a", "g", "e", "d", "b"}, {"a", "b"}}
	AllDigs := [4][]string{{"c", "d", "f", "e", "b"}, {"f", "c", "a", "d", "b"}, {"c", "d", "f", "e", "b"}, {"c", "d", "b", "a", "f"}}

	expectedAnswer := []SegmentDisplay{SegmentDisplay{Seg: AllSegs, Digit: AllDigs}}

	answer, err := ReadSegmentDisplays("testInput")
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
			if !CompareStringArray(one[i].Digit[i2], two[i].Digit[i2]) {
				return false
			}
		}
		for i2 := 0; i2 < len(one[i].Seg); i2++ {
			if !CompareStringArray(one[i].Seg[i2], two[i].Seg[i2]) {
				return false
			}
		}
	}

	return true
}

func CompareStringArray(one, two []string) bool {
	if one == nil || two == nil {
		return false
	}
	if len(one) != len(two) {
		return false
	}
	for i := 0; i < len(one); i++ {
		if one[i] != two[i] {
			return false
		}
	}
	return true
}
