package main

import (
	"log"
	"testing"

	util "github.com/jademonkey/advent-of-code-2021/robcommon"
)

func TestTrimToOnlyBit(t *testing.T) {
	inputList, err := util.ReadInputAsBinaryArray("testInput")
	if err != nil {
		t.Fatalf("Got error: %v\n", err)
	}

	expectedResponseOnes := [][]bool{{true, true, true, true, false}, {true, false, true, true, false}, {true, false, true, true, true}, {true, false, true, false, true},
		{true, true, true, false, false}, {true, false, false, false, false}, {true, true, false, false, true}}

	expectedResponseZeros := [][]bool{{false, false, true, false, false}, {false, true, true, true, true}, {false, false, true, true, true},
		{false, false, false, true, false}, {false, true, false, true, false}}

	ones, zeros := TrimToOnlyBit(inputList, 0)

	if !util.CompareBitArraysEqual(expectedResponseOnes, ones) {
		t.Fatalf("ones failed: Expected %v, got %v", expectedResponseOnes, ones)
	}

	if !util.CompareBitArraysEqual(expectedResponseZeros, zeros) {
		t.Fatalf("zeros failed: Expected %v, got %v", expectedResponseZeros, zeros)
	}

	// test the oxygen
	ones, zeros = TrimToOnlyBit(ones, 1)

	log.Printf("ones: %d zeros %d\n", len(ones), len(zeros))

	expectedResponseZerosTwo := [][]bool{{true, false, true, true, false}, {true, false, true, true, true}, {true, false, true, false, true}, {true, false, false, false, false}}
	if !util.CompareBitArraysEqual(expectedResponseZerosTwo, zeros) {
		t.Fatalf("zeros(2) failed: Expected %v, got %v", expectedResponseZerosTwo, zeros)
	}

	ones, zeros = TrimToOnlyBit(zeros, 2)

	log.Printf("ones: %d zeros %d\n", len(ones), len(zeros))

	expectedResponseThree := [][]bool{{true, false, true, true, false}, {true, false, true, true, true}, {true, false, true, false, true}}
	if !util.CompareBitArraysEqual(expectedResponseThree, ones) {
		t.Fatalf("zeros(3) failed: Expected %v, got %v", expectedResponseThree, ones)
	}

	ones, zeros = TrimToOnlyBit(ones, 3)

	log.Printf("ones: %d zeros %d\n", len(ones), len(zeros))

	expectedResponseFour := [][]bool{{true, false, true, true, false}, {true, false, true, true, true}}
	if !util.CompareBitArraysEqual(expectedResponseFour, ones) {
		t.Fatalf("zeros(3) failed: Expected %v, got %v", expectedResponseFour, ones)
	}

	ones, zeros = TrimToOnlyBit(ones, 4)

	log.Printf("ones: %d zeros %d\n", len(ones), len(zeros))

	expectedResponseFive := [][]bool{{true, false, true, true, true}}
	if !util.CompareBitArraysEqual(expectedResponseFive, ones) {
		t.Fatalf("zeros(3) failed: Expected %v, got %v", expectedResponseFive, ones)
	}

	number := util.BitArrayToNumber(ones[0])

	if number != 23 {
		t.Fatalf("number failed: Expected %d, got %d", 23, number)
	}
}
