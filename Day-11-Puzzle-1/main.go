package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var ErrorValue = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func main() {
	log.Println(">Grabbing input")
	numbers, err := ReadChunkMap("input")
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

func calcSolution(lines [][]rune) (int, error) {
	if lines == nil {
		return 0, fmt.Errorf("lines array was nil")
	}
	var solution int

	for l, line := range lines {
		var stack []rune
		// foundError := false
		for cl, c := range line {
			//Need a FILO
			if c == '(' || c == '[' || c == '{' || c == '<' {
				// add to the stack
				stack = append(stack, c)
			} else {
				// Closing bracke
				// Pop out the last one and compare
				popped := stack[len(stack)-1]
				stack = stack[:(len(stack) - 1)]

				if (c == '}' && popped != '{') ||
					(c == ']' && popped != '[') ||
					(c == ')' && popped != '(') ||
					(c == '>' && popped != '<') {
					// Not the right open!
					log.Printf("Found error. Wrong close: Line: %v Column:%v Char:%v Opener:%v", l, cl, string(c), string(popped))
					solution += ErrorValue[c]
					//	foundError = true
					break
				}
			}
		}
	}

	return solution, nil
}

func ReadChunkMap(filename string) ([][]rune, error) {
	var finalHM [][]rune
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
		var toAppend []rune
		for _, num := range line {
			if num != '(' && num != ')' &&
				num != '[' && num != ']' &&
				num != '{' && num != '}' &&
				num != '<' && num != '>' {
				// skip as not valid
				continue
			}
			toAppend = append(toAppend, num)
		}

		finalHM = append(finalHM, toAppend)
	}

	return finalHM, nil
}
