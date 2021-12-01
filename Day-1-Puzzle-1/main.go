package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(">Grabbing input")
	inputList, err := readInput()
	if err != nil {
		fmt.Printf("Got error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(">Running solution")
	answer, err := calcSolution(inputList)
	if err != nil {
		fmt.Printf("Got error: %v\n", err)
		os.Exit(2)
	}
	fmt.Printf("Answer: %v\n", answer)
}

func readInput() ([]int, error) {
	var output []int
	fileH, err := os.Open("input")
	if err != nil {
		return nil, err
	}
	defer fileH.Close()
	fmt.Printf("Opened Input file\n")

	fileReader := bufio.NewReader(fileH)
	for err == nil {
		var line string
		var number int
		line, err = fileReader.ReadString('\n')
		if err != nil {
			fmt.Printf("Failed to read line %v\n", err)
			break
		}
		line = strings.Trim(line, "\n\r ")
		number, err = strconv.Atoi(line)
		if err != nil {
			fmt.Printf("Failed to convert int %v\n", err)
			break
		}
		output = append(output, number)
	}

	return output, nil
}

func calcSolution(inputList []int) (int, error) {
	if inputList == nil {
		return 0, fmt.Errorf("input list array was nil")
	}
	curNum := inputList[0]
	solution := 0
	for i := 1; i < len(inputList); i++ {
		if inputList[i] > curNum {
			solution++
		}
		curNum = inputList[i]
	}
	return solution, nil
}
