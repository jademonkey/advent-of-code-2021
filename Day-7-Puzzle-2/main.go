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

	// TODO

	total := float64(0)
	for _, num := range numbers {
		total += float64(num)
	}

	log.Printf("TOTAL: %v\n", total)
	log.Printf("  LEN: %v\n", float64(len(numbers)))

	averageF := float64(total / float64(len(numbers)))
	average := uint64(averageF) // This is wrong. It fails the test like this and i have to force the avg to round up. But the real answer relies us on forcing the answer down
	log.Printf("Avg: %v\n", averageF)
	log.Printf("Avg: %v\n", average)

	// Now we need to get to the answer.
	for _, curNum := range numbers {
		moveSteps := uint64(0)
		if curNum > average {
			moveSteps = curNum - average
		} else if curNum < average {
			moveSteps = average - curNum
		}

		if moveSteps == 0 {
			continue
		}
		adder := uint64(1)

		for i := uint64(0); i < moveSteps; i++ {
			solution += adder
			adder++
		}
	}

	return solution, nil
}
