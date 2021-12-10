package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	log.Println(">Grabbing input")
	numbers, err := ReadHeightMap("input")
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

func ReadHeightMap(filename string) ([][]int, error) {
	var finalHM [][]int
	var err error
	fileH, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fileH.Close()
	log.Printf("Opened Input file\n")

	fileReader := bufio.NewReader(fileH)
	for {
		var line string
		line, err = fileReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				err = nil
				break
			}
			log.Printf("Failed to read line %v\n", err)
			break
		}
		line = strings.Trim(line, "\n\r ")
		if len(line) == 0 {
			// skip blank lines
			continue
		}

		// Read each number and weeee
		var toAppend []int
		for _, num := range line {
			convNum, err := strconv.Atoi(string(num))
			if err != nil {
				return nil, err
			}
			toAppend = append(toAppend, convNum)
		}

		finalHM = append(finalHM, toAppend)
	}

	return finalHM, nil
}
