package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

/*
  0:
 aaaa
b    c
b    c
 dddd
e    f
e    f
 gggg

a: 0000001 = 1
b: 0000010 = 2
c: 0000100 = 4
d: 0001000 = 8
e: 0010000 = 16
f: 0100000 = 32
g: 1000000 = 64

0 = 1110111 = 119
1 = 0100100 = 36
2 = 1011101 = 93
3 = 1101101 = 109
4 = 0101110 = 46
5 = 1101011 = 107
6 = 1111011 = 123
7 = 0100101 = 37
8 = 1111111 = 127
9 = 1101111 = 111

*/

type SegmentDisplay struct {
	Seg   [10][]string
	Digit [4][]string
}

func main() {
	log.Println(">Grabbing input")
	numbers, err := ReadSegmentDisplays("input")
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

func calcSolution(numbers []SegmentDisplay) (int, error) {
	if numbers == nil {
		return 0, fmt.Errorf("numbers array was nil")
	}
	var solution int

	// TODO

	return solution, nil
}

func ReadSegmentDisplays(filename string) ([]SegmentDisplay, error) {
	var finalSegs []SegmentDisplay
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
		line = strings.ReplaceAll(line, "| ", "")
		if len(line) == 0 {
			// skip blank lines
			continue
		}

		// Read segments and output digits
		LineSplitted := strings.Split(line, " ")
		if len(LineSplitted) != 14 {
			return nil, fmt.Errorf("length of string array wrong: %v - %v", len(LineSplitted), LineSplitted)
		}

		var thisSegDisSegs [10][]string
		var thisSegDisDigs [4][]string
		for i := 0; i < len(LineSplitted); i++ {
			var toAppend []string
			for _, c := range LineSplitted[i] {
				toAppend = append(toAppend, string(c))
			}

			if i <= 9 {
				thisSegDisSegs[i] = toAppend
			} else {
				thisSegDisDigs[i-10] = toAppend
			}
		}

		// add to output
		thisSegDis := SegmentDisplay{Digit: thisSegDisDigs, Seg: thisSegDisSegs}
		finalSegs = append(finalSegs, thisSegDis)
	}

	return finalSegs, nil
}
