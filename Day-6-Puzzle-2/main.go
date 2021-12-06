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

func ReadCSIntList(filename string) ([]uint64, error) {
	var finalNumbers []uint64
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
			finalNumbers = append(finalNumbers, uint64(thisNum))
		}
	}

	return finalNumbers, nil
}
