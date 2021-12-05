package main

import (
	"testing"
)

func TestReadBingoInput(t *testing.T) {
	CalledNumbers := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
	BingoCard1Nums := [5][5]int{{22, 13, 17, 11, 0}, {8, 2, 23, 4, 24}, {21, 9, 14, 16, 7}, {6, 10, 3, 18, 5}, {1, 12, 20, 15, 19}}
	BingoCard1Mark := [5][5]bool{{false, false, false, false, false}, {false, false, false, false, false}, {false, false, false, false, false}, {false, false, false, false, false}, {false, false, false, false, false}}
	BingoCard2Nums := [5][5]int{{3, 15, 0, 2, 22}, {9, 18, 13, 17, 5}, {19, 8, 7, 25, 23}, {20, 11, 10, 24, 4}, {14, 21, 16, 12, 6}}
	BingoCard2Mark := [5][5]bool{{false, false, false, false, false}, {false, false, false, false, false}, {false, false, false, false, false}, {false, false, false, false, false}, {false, false, false, false, false}}
	BingoCard3Nums := [5][5]int{{14, 21, 17, 24, 4}, {10, 16, 15, 9, 19}, {18, 8, 23, 26, 20}, {22, 11, 13, 6, 5}, {2, 0, 12, 3, 7}}
	BingoCard3Mark := [5][5]bool{{false, false, false, false, false}, {false, false, false, false, false}, {false, false, false, false, false}, {false, false, false, false, false}, {false, false, false, false, false}}

	Bingo1 := BingoCard{Numbers: BingoCard1Nums, MarkedNumbers: BingoCard1Mark, Won: false}
	Bingo2 := BingoCard{Numbers: BingoCard2Nums, MarkedNumbers: BingoCard2Mark, Won: false}
	Bingo3 := BingoCard{Numbers: BingoCard3Nums, MarkedNumbers: BingoCard3Mark, Won: false}

	allCards := []BingoCard{Bingo1, Bingo2, Bingo3}

	returnedNumbers, ReturnedStructs, err := ReadBingoInput("testInput")
	if err != nil {
		t.Fatalf("ReadBingoInput error'd: %v", err)
	}

	if !compareCalledNumbersArray(returnedNumbers, CalledNumbers) {
		t.Errorf("Called Numbers array does not match.\nExpected: %v\n     Got: %v", CalledNumbers, returnedNumbers)
	}

	if !compareBingoCardArray(allCards, ReturnedStructs) {
		t.Errorf("Bingo Cards array does not match.\nExpected: %v\n     Got: %v", allCards, ReturnedStructs)
	}
}

func compareCalledNumbersArray(one, two []int) bool {
	if one == nil || two == nil {
		return false
	}
	if len(one) != len(two) {
		return false
	}

	for i, num1 := range one {
		if two[i] != num1 {
			return false
		}
	}
	return true
}

func compareBingoCardArray(one, two []BingoCard) bool {
	if one == nil || two == nil {
		return false
	}
	for i, bingo := range two {
		bingo2 := one[i]

		// Compare Nums and amrked off
		for r := 0; r < 5; r++ {
			for c := 0; c < 5; c++ {
				if bingo.MarkedNumbers[r][c] != bingo2.MarkedNumbers[r][c] {
					return false
				}

				if bingo.Numbers[r][c] != bingo2.Numbers[r][c] {
					return false
				}
			}
		}
	}
	return true
}

func TestCalculateSolution(t *testing.T) {
	CalledNumbers := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
	BingoCard1Nums := [5][5]int{{22, 13, 17, 11, 0}, {8, 2, 23, 4, 24}, {21, 9, 14, 16, 7}, {6, 10, 3, 18, 5}, {1, 12, 20, 15, 19}}
	BingoCard1Mark := [5][5]bool{{false, false, false, false, false}, {false, false, false, false, false}, {false, false, false, false, false}, {false, false, false, false, false}, {false, false, false, false, false}}
	BingoCard2Nums := [5][5]int{{3, 15, 0, 2, 22}, {9, 18, 13, 17, 5}, {19, 8, 7, 25, 23}, {20, 11, 10, 24, 4}, {14, 21, 16, 12, 6}}
	BingoCard2Mark := [5][5]bool{{false, false, false, false, false}, {false, false, false, false, false}, {false, false, false, false, false}, {false, false, false, false, false}, {false, false, false, false, false}}
	BingoCard3Nums := [5][5]int{{14, 21, 17, 24, 4}, {10, 16, 15, 9, 19}, {18, 8, 23, 26, 20}, {22, 11, 13, 6, 5}, {2, 0, 12, 3, 7}}
	BingoCard3Mark := [5][5]bool{{false, false, false, false, false}, {false, false, false, false, false}, {false, false, false, false, false}, {false, false, false, false, false}, {false, false, false, false, false}}

	Bingo1 := BingoCard{Numbers: BingoCard1Nums, MarkedNumbers: BingoCard1Mark, Won: false}
	Bingo2 := BingoCard{Numbers: BingoCard2Nums, MarkedNumbers: BingoCard2Mark, Won: false}
	Bingo3 := BingoCard{Numbers: BingoCard3Nums, MarkedNumbers: BingoCard3Mark, Won: false}

	allCards := []BingoCard{Bingo1, Bingo2, Bingo3}

	expectedAnswer := 1924

	answer, err := calcSolution(CalledNumbers, allCards)
	if err != nil {
		t.Fatalf("calcSolution error'd: %v", err)
	}

	if answer != expectedAnswer {
		t.Errorf("calcSolution answer does not match.\nExpected: %v\n     Got: %v", expectedAnswer, answer)
	}
}
