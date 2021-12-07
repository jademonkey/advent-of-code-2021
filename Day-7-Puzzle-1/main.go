package main

import (
	"fmt"
	"log"
	"os"
	"sort"

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
	answer, err := calcSolution(numbers)
	if err != nil {
		log.Printf("Got error: %v\n", err)
		os.Exit(2)
	}
	log.Printf("Answer: %v\n", answer)
}

func calcSolution(numbers []uint64) (uint64, error) {
	if numbers == nil {
		return 0, fmt.Errorf("numbers array was nil")
	}
	var solution uint64

	// Find the median number
	numClone := numbers
	sort.Slice(numClone, func(i, j int) bool {
		return numClone[i] < numClone[j]
	})
	var median uint64
	if (len(numClone) % 2) == 1 {
		// odd
		median = numClone[(len(numClone)-1)/2]
	} else {
		median = (numClone[len(numClone)/2] + numClone[(len(numClone)/2)-1]) / 2.
	}

	// Now we need to get to the answer
	for _, curNum := range numbers {
		if curNum > median {
			solution += curNum - median
		} else if curNum < median {
			solution += median - curNum
		}
	}

	return solution, nil
}
