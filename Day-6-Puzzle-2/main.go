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
	answer, err := calcSolution(numbers, 256)
	if err != nil {
		log.Printf("Got error: %v\n", err)
		os.Exit(2)
	}
	log.Printf("Answer: %v\n", answer)
}

func calcSolution(numbers []uint64, iterations int) (uint64, error) {
	var sumOfAll uint64
	if numbers == nil {
		return 0, fmt.Errorf("numbers array was nil")
	}
	if iterations <= 0 {
		return 0, fmt.Errorf("iterations must be above 0")
	}

	finalNumbers := [9]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0}

	// First go through the numbers and add them to a array
	for _, num := range numbers {
		finalNumbers[num]++
	}

	// oh god
	for iter := 0; iter < iterations; iter++ {
		var replaceNumbers [9]uint64

		for i := 8; i >= 0; i-- {
			if i == 0 {
				// special case
				replaceNumbers[6] += finalNumbers[0]
				replaceNumbers[8] += finalNumbers[0]
			} else {
				replaceNumbers[i-1] = finalNumbers[i]
			}
		}

		finalNumbers = replaceNumbers
	}

	for i := 0; i < len(finalNumbers); i++ {
		sumOfAll += finalNumbers[i]
	}

	return sumOfAll, nil
}
