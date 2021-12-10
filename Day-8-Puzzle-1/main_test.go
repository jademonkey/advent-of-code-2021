package main

import (
	"testing"
)

func TestCalculateSolution(t *testing.T) {
	expectedAnswer := 26

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
	zero := []string{"b", "e"}
	one := []string{"a", "b", "c", "d", "e", "f", "g"}
	two := []string{"b", "c", "d", "e", "f", "g"}
	three := []string{"a", "c", "d", "e", "f", "g"}
	four := []string{"b", "c", "e", "g"}
	five := []string{"c", "d", "e", "f", "g"}
	six := []string{"a", "b", "d", "e", "f", "g"}
	seven := []string{"b", "c", "d", "e", "f"}
	eight := []string{"a", "b", "c", "d", "f"}
	nine := []string{"b", "d", "e"}
	ten := []string{"a", "b", "c", "d", "e", "f", "g"}
	eleven := []string{"b", "c", "d", "e", "f"}
	twelve := []string{"b", "c", "d", "e", "f", "g"}
	thirteen := []string{"b", "c", "e", "g"}

	AllSegs := [10][]string{zero, one, two, three, four, five, six, seven, eight, nine}
	AllDigs := [4][]string{ten, eleven, twelve, thirteen}

	expectedAnswer := []SegmentDisplay{SegmentDisplay{Seg: AllSegs, Digit: AllDigs}}

	answer, err := ReadSegmentDisplays("testInputshort")
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
