package main

import (
	"testing"
)

func TestCalculateSolution(t *testing.T) {
	expectedAnswer := 26397

	Input, err := ReadChunkMap("testInput")
	if err != nil {
		t.Fatalf("ReadChunkMap error'd: %v", err)
	}

	answer, err := calcSolution(Input)
	if err != nil {
		t.Fatalf("calcSolution error'd: %v", err)
	}

	if expectedAnswer != answer {
		t.Errorf("Answer wrong\nExpected: %v\n     Got: %v", expectedAnswer, answer)
	}

}

func TestReadChunkMap(t *testing.T) {

	expectedAnswer := [][]rune{{'[', '(', '{', '(', '<', '(', '(', ')', ')', '[', ']', '>', '[', '[', '{', '[', ']', '{', '<', '(', ')', '<', '>', '>'},
		{'[', '(', '(', ')', '[', '<', '>', ']', ')', ']', '(', '{', '[', '<', '{', '<', '<', '[', ']', '>', '>', '('},
		{'{', '(', '[', '(', '<', '{', '}', '[', '<', '>', '[', ']', '}', '>', '{', '[', ']', '{', '[', '(', '<', '(', ')', '>'},
		{'(', '(', '(', '(', '{', '<', '>', '}', '<', '{', '<', '{', '<', '>', '}', '{', '[', ']', '{', '[', ']', '{', '}'},
		{'[', '[', '<', '[', '(', '[', ']', ')', ')', '<', '(', '[', '[', '{', '}', '[', '[', '(', ')', ']', ']', ']'},
		{'[', '{', '[', '{', '(', '{', '}', ']', '{', '}', '}', '(', '[', '{', '[', '{', '{', '{', '}', '}', '(', '[', ']'},
		{'{', '<', '[', '[', ']', ']', '>', '}', '<', '{', '[', '{', '[', '{', '[', ']', '{', '(', ')', '[', '[', '[', ']'},
		{'[', '<', '(', '<', '(', '<', '(', '<', '{', '}', ')', ')', '>', '<', '(', '[', ']', '(', '[', ']', '(', ')'},
		{'<', '{', '(', '[', '(', '[', '[', '(', '<', '>', '(', ')', ')', '{', '}', ']', '>', '(', '<', '<', '{', '{'},
		{'<', '{', '(', '[', '{', '{', '}', '}', '[', '<', '[', '[', '[', '<', '>', '{', '}', ']', ']', ']', '>', '[', ']', ']'}}

	answer, err := ReadChunkMap("testInput")
	if err != nil {
		t.Fatalf("calcSolution error'd: %v", err)
	}

	if !CompareChunkMap(expectedAnswer, answer) {
		t.Errorf("Read wrong\nExpected: %v\n     Got: %v", expectedAnswer, answer)
	}
}

func CompareChunkMap(one, two [][]rune) bool {
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
