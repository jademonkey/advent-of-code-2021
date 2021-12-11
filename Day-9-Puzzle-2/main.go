package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type coords struct {
	Row   int
	Col   int
	Value int
}

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
	var lowPoints []coords
	var totalSizes []int

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
			lowPoints = append(lowPoints, coords{Row: r, Col: c, Value: thisNum})
		}
	}

	// Now iterate through the low points creating the basin
	for _, firstLow := range lowPoints {
		thisBasin := []coords{firstLow}
		thisBasin = buildBasin(thisBasin, firstLow.Row, firstLow.Col, numbers)
		// Now add the size
		totalSizes = append(totalSizes, len(thisBasin))
	}

	sort.Slice(totalSizes, func(i, j int) bool {
		return totalSizes[i] > totalSizes[j]
	})

	solution = totalSizes[0] * totalSizes[1] * totalSizes[2]

	return solution, nil
}

func buildBasin(currentArray []coords, ThisRow, ThisColumn int, set [][]int) []coords {
	// Check up
	if ThisRow != 0 && set[(ThisRow - 1)][ThisColumn] != 9 {
		if !containsRowCol(ThisRow-1, ThisColumn, currentArray) {
			currentArray = append(currentArray, coords{Row: ThisRow - 1, Col: ThisColumn, Value: set[(ThisRow - 1)][ThisColumn]})
			currentArray = buildBasin(currentArray, ThisRow-1, ThisColumn, set)
		}
	}

	// check down
	if ThisRow != (len(set)-1) && set[(ThisRow + 1)][ThisColumn] != 9 {
		if !containsRowCol(ThisRow+1, ThisColumn, currentArray) {
			currentArray = append(currentArray, coords{Row: ThisRow + 1, Col: ThisColumn, Value: set[(ThisRow + 1)][ThisColumn]})
			currentArray = buildBasin(currentArray, ThisRow+1, ThisColumn, set)
		}
	}

	// Check left
	if ThisColumn != 0 && set[ThisRow][ThisColumn-1] != 9 {
		if !containsRowCol(ThisRow, ThisColumn-1, currentArray) {
			currentArray = append(currentArray, coords{Row: ThisRow, Col: ThisColumn - 1, Value: set[ThisRow][ThisColumn-1]})
			currentArray = buildBasin(currentArray, ThisRow, ThisColumn-1, set)
		}
	}

	// Check right
	if ThisColumn != (len(set[ThisRow])-1) && set[ThisRow][ThisColumn+1] != 9 {
		if !containsRowCol(ThisRow, ThisColumn+1, currentArray) {
			currentArray = append(currentArray, coords{Row: ThisRow, Col: ThisColumn + 1, Value: set[ThisRow][ThisColumn+1]})
			currentArray = buildBasin(currentArray, ThisRow, ThisColumn+1, set)
		}
	}

	return currentArray
}

func containsRowCol(row, col int, thisArray []coords) bool {
	for _, co := range thisArray {
		if co.Col == col && co.Row == row {
			return true
		}
	}
	return false
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
