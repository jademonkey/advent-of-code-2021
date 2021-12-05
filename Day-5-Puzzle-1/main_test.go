package main

import (
	"testing"
)

func TestReadLineCoords(t *testing.T) {
	Input1 := Coords{0, 9, 5, 9}
	Input2 := Coords{8, 0, 0, 8}
	Input3 := Coords{9, 4, 3, 4}
	Input4 := Coords{2, 2, 2, 1}
	Input5 := Coords{7, 0, 7, 4}
	Input6 := Coords{6, 4, 2, 0}
	Input7 := Coords{0, 9, 2, 9}
	Input8 := Coords{3, 4, 1, 4}
	Input9 := Coords{0, 0, 8, 8}
	Input0 := Coords{5, 5, 8, 2}

	CoordsArray := []Coords{Input1, Input2, Input3, Input4, Input5, Input6, Input7, Input8, Input9, Input0}
	MaxX := 9
	MaxY := 9

	MyArray, X, Y, err := ReadLineCoords("testInput")
	if err != nil {
		t.Fatalf("ReadLineCoords error'd: %v", err)
	}

	if MaxX != X {
		t.Errorf("ReadLineCoords MaxX does not match.\nExpected: %v\n     Got: %v", MaxX, X)
	}

	if MaxY != X {
		t.Errorf("ReadLineCoords MaxY does not match.\nExpected: %v\n     Got: %v", MaxY, Y)
	}

	if !CompareCoordsArray(CoordsArray, MyArray) {
		t.Errorf("ReadLineCoords CoordArray does not match.\nExpected: %v\n     Got: %v", CoordsArray, MyArray)
	}

}

func CompareCoordsArray(one, two []Coords) bool {
	if one == nil || two == nil {
		return false
	}
	if len(one) != len(two) {
		return false
	}

	for i := 0; i < len(one); i++ {
		if one[i].StartX != two[i].StartX {
			return false
		}
		if one[i].StartY != two[i].StartY {
			return false
		}
		if one[i].EndX != two[i].EndX {
			return false
		}
		if one[i].EndY != two[i].EndY {
			return false
		}
	}

	return true
}

func TestCalculateSolution(t *testing.T) {
	Input1 := Coords{0, 9, 5, 9}
	Input2 := Coords{8, 0, 0, 8}
	Input3 := Coords{9, 4, 3, 4}
	Input4 := Coords{2, 2, 2, 1}
	Input5 := Coords{7, 0, 7, 4}
	Input6 := Coords{6, 4, 2, 0}
	Input7 := Coords{0, 9, 2, 9}
	Input8 := Coords{3, 4, 1, 4}
	Input9 := Coords{0, 0, 8, 8}
	Input0 := Coords{5, 5, 8, 2}

	CoordsArray := []Coords{Input1, Input2, Input3, Input4, Input5, Input6, Input7, Input8, Input9, Input0}
	MaxX := 9
	MaxY := 9

	expectedAnswer := 5

	answer, err := calcSolution(CoordsArray, MaxX, MaxY)
	if err != nil {
		t.Fatalf("calcSolution error'd: %v", err)
	}

	if answer != expectedAnswer {
		t.Errorf("calcSolution answer does not match.\nExpected: %v\n     Got: %v", expectedAnswer, answer)
	}
}
