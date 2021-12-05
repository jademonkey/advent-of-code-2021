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
	allCoords, maxX, maxY, err := ReadLineCoords("input")
	if err != nil {
		log.Printf("Got error: %v\n", err)
		os.Exit(1)
	}
	log.Println(">Running solution")
	answer, err := calcSolution(allCoords, maxX, maxY)
	if err != nil {
		log.Printf("Got error: %v\n", err)
		os.Exit(2)
	}
	log.Printf("Answer: %v\n", answer)
}

func calcSolution(allCoords []Coords, maxX, maxY int) (int, error) {
	if allCoords == nil {
		return 0, fmt.Errorf("allCoords array was nil")
	}

	var Board [][]int
	var solution int

	for i := 0; i < maxY+1; i++ {
		// array using make method
		arrayWithMake := make([]int, (maxX + 1))
		Board = append(Board, arrayWithMake)
	}

	for _, cor := range allCoords {
		//	log.Printf("Looking at: %v", cor)
		var myStart, myEnd int
		if cor.StartX == cor.EndX {
			// Column
			if cor.StartY < cor.EndY {
				myStart = cor.StartY
				myEnd = cor.EndY
			} else if cor.EndY < cor.StartY {
				myStart = cor.EndY
				myEnd = cor.StartY
			} else {
				myStart = cor.StartY
				myEnd = cor.StartY
			}
			for c := myStart; c < (myEnd + 1); c++ {
				Board[cor.StartX][c]++
			}
		} else if cor.StartY == cor.EndY {
			// Row
			if cor.StartX < cor.EndX {
				myStart = cor.StartX
				myEnd = cor.EndX
			} else if cor.EndX < cor.StartX {
				myStart = cor.EndX
				myEnd = cor.StartX
			} else {
				myStart = cor.StartX
				myEnd = cor.StartX
			}
			for r := myStart; r < (myEnd + 1); r++ {
				Board[r][cor.StartY]++
			}
		}

	}
	// Now we iterate through
	for r := 0; r < maxX+1; r++ {
		for c := 0; c < maxY+1; c++ {
			if Board[r][c] > 1 {
				solution++
			}
		}
	}
	return solution, nil
}

func ReadLineCoords(filename string) ([]Coords, int, int, error) {
	var finalCoords []Coords
	var err error
	var largestX, largestY int
	fileH, err := os.Open(filename)
	if err != nil {
		return nil, 0, 0, nil
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
		line = strings.ReplaceAll(line, " -> ", ",")
		if len(line) == 0 {
			// skip blank lines
			continue
		}

		numbers := strings.Split(line, ",")
		if len(numbers) != 4 {
			err = fmt.Errorf("array of numbers was wrong")
			log.Printf("%v\n", err)
			break
		}

		startX, err := strconv.Atoi(numbers[0])
		if err != nil {
			log.Printf("Failed to convert int %v - %v\n", numbers[0], err)
			break
		}

		if startX > largestX {
			largestX = startX
		}

		startY, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Printf("Failed to convert int %v - %v\n", numbers[1], err)
			break
		}

		if startY > largestY {
			largestY = startY
		}

		endX, err := strconv.Atoi(numbers[2])
		if err != nil {
			log.Printf("Failed to convert int %v - %v\n", numbers[2], err)
			break
		}

		if endX > largestX {
			largestX = endX
		}

		endY, err := strconv.Atoi(numbers[3])
		if err != nil {
			log.Printf("Failed to convert int %v - %v\n", numbers[3], err)
			break
		}

		if endY > largestY {
			largestY = endY
		}

		coord := Coords{startX, startY, endX, endY}
		finalCoords = append(finalCoords, coord)
	}

	return finalCoords, largestX, largestY, nil
}
