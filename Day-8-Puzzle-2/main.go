package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
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
	Seg   [10]string
	Digit [4]string
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
	var solutions []int
	var solution int

	for _, lines := range numbers {
		CalculatedSegments := make(map[string]int)
		// First the easy ones
		CalculatedSegments[lines.Seg[0]] = 1
		CalculatedSegments[lines.Seg[1]] = 7
		CalculatedSegments[lines.Seg[2]] = 4
		CalculatedSegments[lines.Seg[9]] = 8

		// Now the ones we can work out some
		// Top segment
		topSegAr, _ := missingAndCommonLetters(lines.Seg[1], lines.Seg[0])
		if len(topSegAr) != 1 {
			// what?
			log.Panicf("Top seg wrong")
		}
		topSegL := topSegAr[0]

		// number 9
		Most9 := lines.Seg[2] + topSegL
		var Number9 string
		var botSeg string

		for i := 6; i < 9; i++ {
			if containsAllExact(Most9, lines.Seg[i]) {
				// This is 9
				Number9 = lines.Seg[i]
				CalculatedSegments[lines.Seg[i]] = 9
				botSegAr, _ := missingAndCommonLetters(lines.Seg[i], Most9)
				if len(botSegAr) != 1 {
					// what?
					log.Panicf("Bot seg wrong")
				}
				botSeg = botSegAr[0]
				break
			}
		}

		if len(Number9) != 6 {
			log.Panicf("Didn't get number 9")
		}

		// Number 3
		Most3 := lines.Seg[1] + botSeg
		var Number3 string
		var midSeg string

		for i := 3; i < 6; i++ {
			if containsAllExact(Most3, lines.Seg[i]) {
				// This is 9
				CalculatedSegments[lines.Seg[i]] = 3
				Number3 = lines.Seg[i]
				midSegAr, _ := missingAndCommonLetters(lines.Seg[i], Most3)
				if len(midSegAr) != 1 {
					// what?
					log.Panicf("Mid seg wrong")
				}
				midSeg = midSegAr[0]
				break
			}
		}

		if len(Number3) != 5 {
			log.Panicf("Didn't get number 3")
		}

		// Number 0
		Number0 := RemoveLetter(lines.Seg[9], midSeg)
		CalculatedSegments[Number0] = 0

		if len(Number0) != 6 {
			log.Panicf("Didn't get number 0")
		}

		// Number 6
		var Number6 string
		for i := 6; i < 9; i++ {
			if lines.Seg[i] != Number0 && lines.Seg[i] != Number9 {
				// Number 6
				Number6 = lines.Seg[i]
				CalculatedSegments[lines.Seg[i]] = 6
				break
			}
		}
		if len(Number6) != 6 {
			log.Panicf("Didn't get number 6")
		}

		// Number 2 via bot left
		var Number2 string
		botLeftSegs, _ := missingAndCommonLetters(lines.Seg[9], Number9)
		if len(botLeftSegs) != 1 {
			// what?
			log.Panicf("Bot Left seg wrong")
		}
		botLeft := botLeftSegs[0]

		for i := 3; i < 6; i++ {
			if strings.Contains(lines.Seg[i], botLeft) {
				// Number 2!
				Number2 = lines.Seg[i]
				CalculatedSegments[Number2] = 2
				break
			}
		}

		if len(Number2) != 5 {
			log.Panicf("Didn't get number 2")
		}

		// Finally, Number 5
		var Number5 string
		for i := 3; i < 6; i++ {
			if lines.Seg[i] != Number2 && lines.Seg[i] != Number3 {
				// Number 5!
				Number5 = lines.Seg[i]
				CalculatedSegments[lines.Seg[i]] = 5
				break
			}
		}

		if len(Number5) != 5 {
			log.Panicf("Didn't get number 5")
		}

		// Now finally at this point we can calculate the digits
		solutions = append(solutions, calculateDigitNumber(lines.Digit, CalculatedSegments))
	}

	// add it all
	for _, num := range solutions {
		solution += num
	}

	return solution, nil
}

func calculateDigitNumber(digits [4]string, mappedNumbers map[string]int) int {

	firstNumber := mappedNumbers[digits[0]]
	secondNumber := mappedNumbers[digits[1]]
	thirdNumber := mappedNumbers[digits[2]]
	fourthNumber := mappedNumbers[digits[3]]

	log.Printf("Got digits: %v%v%v%v", firstNumber, secondNumber, thirdNumber, fourthNumber)

	return (firstNumber * 1000) + (secondNumber * 100) + (thirdNumber * 10) + (fourthNumber)
}

func RemoveLetter(word, letters string) string {
	var toReturn string
	for _, c := range word {
		found := false
		for _, c2 := range letters {
			if c == c2 {
				found = true
				break
			}
		}
		if !found {
			toReturn += string(c)
		}
	}
	return toReturn
}

func containsAllExact(first, second string) bool {
	for _, c := range first {
		found := false
		for _, c2 := range second {
			if c == c2 {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func missingAndCommonLetters(first, second string) ([]string, []string) {
	var toReturnMissing []string
	var toReturnCommon []string
	for _, c := range first {
		found := false
		for _, c2 := range second {
			if c == c2 {
				found = true
				toReturnCommon = append(toReturnCommon, string(c2))
				break
			}
		}
		if !found {
			toReturnMissing = append(toReturnMissing, string(c))
		}
	}

	return toReturnMissing, toReturnCommon
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

		var thisSegDisSegs [10]string
		var thisSegDisDigs [4]string
		for i := 0; i < len(LineSplitted); i++ {
			var toAppend []string
			for _, c := range LineSplitted[i] {
				toAppend = append(toAppend, string(c))
			}
			sort.Slice(toAppend, func(i2, j int) bool {
				return rune(toAppend[i2][0]) < rune(toAppend[j][0])
			})

			if i <= 9 {
				thisSegDisSegs[i] = strings.Join(toAppend, "")
			} else {
				thisSegDisDigs[i-10] = strings.Join(toAppend, "")
			}
		}
		thisSegDisSegs = sortSegments(thisSegDisSegs)

		// add to output
		thisSegDis := SegmentDisplay{Digit: thisSegDisDigs, Seg: thisSegDisSegs}
		finalSegs = append(finalSegs, thisSegDis)
	}

	return finalSegs, nil
}

func sortSegments(Seg [10]string) [10]string {
	currentI := 0
	var orderedSeg [10]string
	for i := 2; i < 8; i++ {
		for _, str := range Seg {
			if len(str) == i {
				orderedSeg[currentI] = str
				currentI++
			}
		}
	}
	return orderedSeg
}
