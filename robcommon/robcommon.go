package robcommon

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadInputAsIntArray(filename string) ([]int, error) {
	var output []int
	fileH, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fileH.Close()
	log.Printf("Opened Input file\n")

	fileReader := bufio.NewReader(fileH)
	for err == nil {
		var line string
		var number int
		line, err = fileReader.ReadString('\n')
		if err != nil {
			log.Printf("Failed to read line %v\n", err)
			break
		}
		line = strings.Trim(line, "\n\r ")
		number, err = strconv.Atoi(line)
		if err != nil {
			log.Printf("Failed to convert int %v\n", err)
			break
		}
		output = append(output, number)
	}

	return output, nil
}

type DIRECTION int

const (
	FORWARD DIRECTION = iota
	BACKWARD
	DOWN
	UP
)

type DirDist struct {
	Dir  DIRECTION
	Dist int
}

func ReadInputAsDirectionDistance(filename string) ([]DirDist, error) {
	var output []DirDist
	fileH, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fileH.Close()
	log.Printf("Opened Input file\n")

	fileReader := bufio.NewReader(fileH)
	for err == nil {
		var line string
		var di DIRECTION
		var number int

		// read a line
		line, err = fileReader.ReadString('\n')
		if err != nil {
			log.Printf("Failed to read line %v\n", err)
			break
		}
		line = strings.Trim(line, "\n\r ")

		// parse a line
		splitted := strings.Split(line, " ")
		if len(splitted) != 2 {
			log.Printf("Split looked weird: %v\n", splitted)
			break
		}

		// Work out which direction
		if strings.EqualFold(splitted[0], "forward") {
			di = FORWARD
		} else if strings.EqualFold(splitted[0], "backward") {
			di = BACKWARD
		} else if strings.EqualFold(splitted[0], "up") {
			di = UP
		} else if strings.EqualFold(splitted[0], "down") {
			di = DOWN
		} else {
			log.Fatalf("Unknown direction: '%s'", splitted[0])
			break
		}

		// and distance
		number, err = strconv.Atoi(splitted[1])
		if err != nil {
			log.Printf("Failed to convert int %v\n", err)
			break
		}
		toAppend := DirDist{Dir: di, Dist: number}
		output = append(output, toAppend)
	}

	return output, nil
}
