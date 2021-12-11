package main

import (
	"fmt"
	"log"
	"os"

	util "github.com/jademonkey/advent-of-code-2021/robcommon"
)

func main() {
	log.Println(">Grabbing input")
	numbers, err := util.ReadHeightMap("input")
	if err != nil {
		log.Printf("Got error: %v\n", err)
		os.Exit(1)
	}

	log.Println(">Running solution")
	answer, err := calcSolution(numbers, 100)
	if err != nil {
		log.Printf("Got error: %v\n", err)
		os.Exit(2)
	}
	log.Printf("Answer: %v\n", answer)
}

func calcSolution(lines [][]int, iterations int) (int, error) {
	if lines == nil {
		return 0, fmt.Errorf("lines array was nil")
	}
	var solution int
	var totalOcts int

	for r := 0; r < len(lines); r++ {
		for c := 0; c < len(lines[0]); c++ {
			totalOcts++
		}
	}

	for steps := 0; ; steps++ {
		var flashesThisStep int

		// First increment everything and flash as necessary
		for r := 0; r < len(lines); r++ {
			for c := 0; c < len(lines[0]); c++ {
				lines[r][c]++
				if lines[r][c] == 10 {
					// FLASH TIME recursion
					lines, _ = doFlash(lines, r, c)
				}
			}
		}

		// Reset all above 9 to 0
		for r := 0; r < len(lines); r++ {
			for c := 0; c < len(lines[0]); c++ {
				if lines[r][c] > 9 {
					lines[r][c] = 0
					flashesThisStep++
				}
			}
		}
		// Check if ALL the buggers flashed at once and add to the solution list
		log.Printf("Step: %v Flashes: %v/%v", steps, flashesThisStep, totalOcts)
		if flashesThisStep == totalOcts {
			solution = (steps + 1)
			break
		}
		// Go to the next step if not
	}

	return solution, nil
}

func doFlash(lines [][]int, r, c int) ([][]int, int) {
	flashes := 0

	for newRow := r - 1; newRow < (r + 2); newRow++ {
		if newRow < 0 || newRow >= len(lines) {
			// out of bounds
			continue
		}
		for newCol := c - 1; newCol < (c + 2); newCol++ {
			if newCol < 0 || newCol >= len(lines[0]) {
				// out of bounds
				continue
			}
			lines[newRow][newCol]++
			if lines[newRow][newCol] == 10 {
				var extraFlashes int
				flashes++
				// we're flashing again!
				lines, extraFlashes = doFlash(lines, newRow, newCol)
				flashes += extraFlashes
			}
		}
	}

	return lines, flashes
}
