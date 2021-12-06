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

type Coords struct {
	StartX int
	StartY int
	EndX   int
	EndY   int
}

func main() {
	log.Println(">Grabbing input")
	numbers, err := ReadCSIntList("input")
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

func calcSolution(numbers []int, iterations int) (int, error) {
	if numbers == nil {
		return 0, fmt.Errorf("numbers array was nil")
	}
	if iterations <= 0 {
		return 0, fmt.Errorf("iterations must be above 0")
	}

	clonedNumbers := numbers

	// oh god
	for iter := 0; iter < iterations; iter++ {
		var replaceNumbers []int
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

func ReadCSIntList(filename string) ([]int, error) {
	var finalNumbers []int
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

		numbers := strings.Split(line, ",")

		for _, numS := range numbers {
			thisNum, err := strconv.Atoi(numS)
			if err != nil {
				log.Printf("Failed to convert int %v - %v\n", numS, err)
				break
			}
			finalNumbers = append(finalNumbers, thisNum)
		}
	}

	return finalNumbers, nil
}
