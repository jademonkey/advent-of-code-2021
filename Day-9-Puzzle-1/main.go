package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jademonkey/advent-of-code-2021/robcommon"
)

func main() {
	log.Println(">Grabbing input")
	numbers, err := robcommon.ReadHeightMap("input")
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

func calcSolution(numbers [][]int) (int, error) {
	if numbers == nil {
		return 0, fmt.Errorf("numbers array was nil")
	}
	var solution int
	var lowPoints []int

	for r := 0; r < len(numbers); r++ {
		for c := 0; c < len(numbers[0]); c++ {
			thisNum := numbers[r][c]
			// Check up
			if r != 0 && numbers[(r - 1)][c] <= thisNum {
				continue
			}
			// check down
			if r != (len(numbers)-1) && numbers[(r + 1)][c] <= thisNum {
				continue
			}
			// check left
			if c != 0 && numbers[r][(c-1)] <= thisNum {
				continue
			}

			//check right
			if c != (len(numbers[0])-1) && numbers[r][(c+1)] <= thisNum {
				continue
			}

			// all higher!
			lowPoints = append(lowPoints, thisNum)
		}
	}

	log.Printf("All low points! %v", lowPoints)
	// Now total everything
	for _, num := range lowPoints {
		solution += (num + 1)
	}

	return solution, nil
}
