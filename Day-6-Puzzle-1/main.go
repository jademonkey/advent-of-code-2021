package main

import (
	"fmt"
	"log"
	"os"

	util "github.com/jademonkey/advent-of-code-2021/robcommon"
)

func main() {
	log.Println(">Grabbing input")
	numbers, err := util.ReadCSIntList("input")
	if err != nil {
		log.Printf("Got error: %v\n", err)
		os.Exit(1)
	}
	log.Println(">Running solution")
	answer, err := calcSolution(numbers, 80)
	if err != nil {
		log.Printf("Got error: %v\n", err)
		os.Exit(2)
	}
	log.Printf("Answer: %v\n", answer)
}

func calcSolution(numbers []uint64, iterations int) (int, error) {
	if numbers == nil {
		return 0, fmt.Errorf("numbers array was nil")
	}
	if iterations <= 0 {
		return 0, fmt.Errorf("iterations must be above 0")
	}

	clonedNumbers := numbers

	// oh god
	for iter := 0; iter < iterations; iter++ {
		var replaceNumbers []uint64
		for i := 0; i < len(clonedNumbers); i++ {
			if clonedNumbers[i] == 0 {
				replaceNumbers = append(replaceNumbers, 6)
				replaceNumbers = append(replaceNumbers, 8)
			} else {
				replaceNumbers = append(replaceNumbers, (clonedNumbers[i] - 1))
			}
		}
		clonedNumbers = replaceNumbers
	}

	return len(clonedNumbers), nil
}
