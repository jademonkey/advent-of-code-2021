package main

import (
	"fmt"
	"log"
	"os"

	util "github.com/jademonkey/advent-of-code-2021/robcommon"
)

func main() {
	log.Println(">Grabbing input")
	inputList, err := util.ReadInputAsDirectionDistance("input")
	if err != nil {
		log.Printf("Got error: %v\n", err)
		os.Exit(1)
	}
	log.Println(">Running solution")
	answer, err := calcSolution1(inputList)
	if err != nil {
		log.Printf("Got error: %v\n", err)
		os.Exit(2)
	}
	log.Printf("Answer: %v\n", answer)
}

func calcSolution1(inputList []util.DirDist) (int, error) {
	var solution, hori, vert int
	if inputList == nil {
		return 0, fmt.Errorf("input list array was nil")
	}
	if len(inputList) < 4 {
		return 0, fmt.Errorf("input list must be larger than 3")
	}
	for i := 0; i < len(inputList); i++ {
		thisone := inputList[i]
		switch thisone.Dir {
		case util.FORWARD:
			hori += thisone.Dist
		case util.BACKWARD:
			hori -= thisone.Dist
		case util.UP:
			vert -= thisone.Dist
		case util.DOWN:
			vert += thisone.Dist
		}
	}

	solution = hori * vert

	return solution, nil
}
