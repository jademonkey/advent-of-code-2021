package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type BingoCard struct {
	Numbers       [5][5]int
	MarkedNumbers [5][5]bool
}

func main() {
	log.Println(">Grabbing input")
	allNums, cards, err := ReadBingoInput("input")
	if err != nil {
		log.Printf("Got error: %v\n", err)
		os.Exit(1)
	}
	log.Println(">Running solution")
	answer, err := calcSolution(allNums, cards)
	if err != nil {
		log.Printf("Got error: %v\n", err)
		os.Exit(2)
	}
	log.Printf("Answer: %v\n", answer)
}

func calcSolution(calledNumbers []int, allCards []BingoCard) (int, error) {
	if calledNumbers == nil {
		return 0, fmt.Errorf("calledNumbers array was nil")
	}
	if allCards == nil {
		return 0, fmt.Errorf("allCards array was nil")
	}
	winningCard := -1
	winningNumber := -1

	// Iterate through each called number
	for _, curNum := range calledNumbers {
		// Mark it all off

		for cardIndex, _ := range allCards {
			for r := 0; r < 5; r++ {
				for c := 0; c < 5; c++ {
					if allCards[cardIndex].Numbers[r][c] == curNum {
						allCards[cardIndex].MarkedNumbers[r][c] = true
					}
				}
			}
		}

		// Check each card if we have a row win
		for cardIndex, card := range allCards {
			for r := 0; r < 5; r++ {
				if card.MarkedNumbers[r][0] && card.MarkedNumbers[r][1] && card.MarkedNumbers[r][2] && card.MarkedNumbers[r][3] && card.MarkedNumbers[r][4] {
					// WINNER WINNER CHICKEN DINNER on a row
					winningCard = cardIndex
					winningNumber = curNum
					break
				}
			}
			if winningCard != -1 {
				break
			}
		}

		if winningCard != -1 {
			break
		}

		// Check each card if we have a column win instead
		for cardIndex, card := range allCards {
			for c := 0; c < 5; c++ {
				if card.MarkedNumbers[0][c] && card.MarkedNumbers[1][c] && card.MarkedNumbers[2][c] && card.MarkedNumbers[3][c] && card.MarkedNumbers[4][c] {
					// WINNER WINNER CHICKEN DINNER on a row
					winningCard = cardIndex
					winningNumber = curNum
					break
				}
			}
			if winningCard != -1 {
				break
			}
		}

		if winningCard != -1 {
			break
		}

		// No winners, keep going
	}

	// We won! somewhere.. Let's now calculate it
	if winningCard == -1 {
		// no one won?!?
		return 0, fmt.Errorf("winningCard ended up being no-one...?")
	}

	WinningBoard := allCards[winningCard]
	sum := 0
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if !WinningBoard.MarkedNumbers[r][c] {
				sum += WinningBoard.Numbers[r][c]
			}
		}
	}

	sum *= winningNumber

	return sum, nil
}

func ReadBingoInput(filename string) ([]int, []BingoCard, error) {
	var calledNumbersOut []int
	var finalCards []BingoCard
	var err error
	fileH, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer fileH.Close()
	log.Printf("Opened Input file\n")

	fileReader := bufio.NewReader(fileH)
	firstLine := true
	curRow := 0
	var currentCard [5][5]int
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

		if firstLine {
			// This is the called numbers string
			numbersSplitted := strings.Split(line, ",")
			for _, num := range numbersSplitted {
				number, err := strconv.Atoi(num)
				if err != nil {
					log.Printf("Failed to convert int %v - %v\n", num, err)
					break
				}
				calledNumbersOut = append(calledNumbersOut, number)
			}
			firstLine = false
			continue
		}

		// After this point we're onto a BingCard array line
		numbersSplitted := strings.Split(line, " ")
		i := 0
		for _, num := range numbersSplitted {
			if num == "" {
				continue
			}
			number, err := strconv.Atoi(num)
			if err != nil {
				log.Printf("Failed to convert int %v - %v\n", num, err)
				break
			}
			if i == 5 {
				log.Printf("Too many numbers? Expected 5 got %v - %v\n", len(numbersSplitted), numbersSplitted)
				break
			}

			currentCard[curRow][i] = number
			i++
		}
		curRow++

		// Check if that was the 5th row!
		if curRow == 5 {
			// It was!
			ToAddCard := BingoCard{Numbers: currentCard, MarkedNumbers: [5][5]bool{{false, false, false, false, false}, {false, false, false, false, false}, {false, false, false, false, false}, {false, false, false, false, false}, {false, false, false, false, false}}}
			finalCards = append(finalCards, ToAddCard)
			curRow = 0
		}
	}

	return calledNumbersOut, finalCards, err
}
