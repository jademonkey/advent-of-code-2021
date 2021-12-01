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
	answer, err := calcSolution2(inputList)
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

func calcSolution2(inputList []int) (int, error) {
	if inputList == nil {
		return 0, fmt.Errorf("input list array was nil")
	}
	if len(inputList) < 4 {
		return 0, fmt.Errorf("input list must be larger than 3")
	}

	curNum := inputList[0] + inputList[1] + inputList[2]
	solution := 0
	for i := 2; (i + 1) < len(inputList); i++ {
		nextNum := inputList[i-1] + inputList[i] + inputList[i+1]
		if nextNum > curNum {
			solution++
		}
		curNum = nextNum
	}
	return solution, nil
}
