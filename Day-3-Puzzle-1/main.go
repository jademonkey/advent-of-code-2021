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
	var gamma, epsilon, solution int
	var gammabuilder []bool
	if inputList == nil {
		return 0, fmt.Errorf("input list array was nil")
	}

	// Now calculate the epsilon and gamma bits
	log.Printf(">got input array=%v\n", inputList)
	log.Println(">calculating bits")
	for i := 0; i < len(inputList[0]); i++ {
		var tot1 int
		for i2 := 0; i2 < len(inputList); i2++ {
			if inputList[i2][i] {
				tot1++
			}
		}

		// Now add a 0 or 1
		if tot1 > (len(inputList) / 2) {
			gammabuilder = append(gammabuilder, true)
		} else {
			gammabuilder = append(gammabuilder, false)
		}
	}

	// convert to numbers
	log.Printf(">got bit array=%v\n", gammabuilder)
	log.Println(">Converting to number")
	starting := 1
	for i := len(gammabuilder) - 1; i >= 0; i-- {
		if gammabuilder[i] {
			gamma += starting
		} else {
			epsilon += starting
		}

		starting *= 2
	}

	log.Printf(">got numbers gamma=%d epsilon=%d\n", gamma, epsilon)

	// solution
	solution = gamma * epsilon

	return solution, nil
}
