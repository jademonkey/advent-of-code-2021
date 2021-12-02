package main

import (
	"fmt"
	"os"

	util "github.com/jademonkey/advent-of-code-2021/robcommon"
)

func main() {
	fmt.Println(">Grabbing input")
	inputList, err := util.ReadInputAsIntArray("input")
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
