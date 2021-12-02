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
	answer, err := calcSolution2(inputList)
	if err != nil {
		fmt.Printf("Got error: %v\n", err)
		os.Exit(2)
	}
	fmt.Printf("Answer: %v\n", answer)
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
