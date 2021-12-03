package main

import (
	"fmt"
	"log"
	"os"

	util "github.com/jademonkey/advent-of-code-2021/robcommon"
)

func main() {
	log.Println(">Grabbing input")
	inputList, err := util.ReadInputAsBinaryArray("input")
	if err != nil {
		log.Printf("Got error: %v\n", err)
		os.Exit(1)
	}
	log.Println(">Running solution")
	answer, err := calcSolution(inputList)
	if err != nil {
		log.Printf("Got error: %v\n", err)
		os.Exit(2)
	}
	log.Printf("Answer: %v\n", answer)
}

func calcSolution(inputList [][]bool) (int, error) {
	var oxygen, carbon, solution int
	var builtOx, builtCar []bool

	if inputList == nil {
		return 0, fmt.Errorf("input list array was nil")
	}

	currentOxygen := inputList

	// Calculate the oxygen
	i := 0
	for {
		ones, zeros := TrimToOnlyBit(currentOxygen, i)
		if len(ones) > len(zeros) {
			currentOxygen = ones
		} else if len(ones) < len(zeros) {
			currentOxygen = zeros
		} else {
			// Equal
			builtOx = ones[0]
			break
		}
		log.Printf("CurOx: %v\n", currentOxygen)
		i++
	}

	currentCarbon := inputList

	// Calculate the carbon
	for i := 0; true; i++ {
		ones, zeros := TrimToOnlyBit(currentCarbon, i)
		if len(ones) > len(zeros) {
			currentCarbon = zeros
		} else if len(ones) < len(zeros) {
			currentCarbon = ones
		} else {
			// Equal
			builtCar = zeros[0]
			break
		}
	}

	// convert to numbers
	log.Println(">Converting to number")
	oxygen = util.BitArrayToNumber(builtOx)

	log.Printf(">got co2 bit array=%v\n", builtCar)
	log.Println(">Converting to number")
	carbon = util.BitArrayToNumber(builtCar)

	log.Printf(">got numbers oxygen=%d carbon=%d\n", oxygen, carbon)

	// solution
	solution = oxygen * carbon

	return solution, nil
}

func TrimToOnlyBit(inputList [][]bool, pos int) ([][]bool, [][]bool) {
	if inputList == nil {
		return nil, nil
	}

	var OneList, ZeroList [][]bool

	for _, cur := range inputList {
		if cur[pos] {
			OneList = append(OneList, cur)
		} else {
			ZeroList = append(ZeroList, cur)
		}
	}

	return OneList, ZeroList
}
