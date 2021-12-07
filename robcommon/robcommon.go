package robcommon

import (
	"bufio"
	"io"
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

func ReadInputAsBinaryArray(filename string) ([][]bool, error) {
	var output [][]bool
	fileH, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fileH.Close()
	log.Printf("Opened Input file\n")

	fileReader := bufio.NewReader(fileH)
	for err == nil {
		var line string
		var lineoutput []bool
		line, err = fileReader.ReadString('\n')
		if err != nil {
			log.Printf("Failed to read line %v\n", err)
			break
		}
		line = strings.Trim(line, "\n\r ")

		// Now parse character by character building an array
		for _, c := range line {
			if c == '0' {
				lineoutput = append(lineoutput, false)
			} else if c == '1' {
				lineoutput = append(lineoutput, true)
			} else {
				log.Printf("Unknown character '%v'\n", c)
				break
			}
		}

		// add on to the end
		output = append(output, lineoutput)
	}

	return output, nil
}

func BitArrayToNumber(input []bool) int {
	starting := 1
	solution := 0
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] {
			solution += starting
		}

		starting *= 2
	}

	return solution
}

func CompareBitArraysEqual(first, second [][]bool) bool {
	if len(first) != len(second) {
		return false
	}

	for i := 0; i < len(first); i++ {
		if len(first[i]) != len(second[i]) {
			return false
		}
		for d := 0; d < len(first[i]); d++ {
			if first[i][d] != second[i][d] {
				return false
			}
		}
	}

	return true
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
