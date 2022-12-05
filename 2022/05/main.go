package main

import (
	"05/h"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type move struct {
	amount int
	from   int
	to     int
}

func main() {
	result1 := solveChallengeOne()
	fmt.Printf("Result-1: %s\n", result1)

	result2 := solveChallengeTwo()
	fmt.Printf("Result-2: %s\n", result2)
}

func solveChallengeOne() string {
	towers, moves := readInput()
	handleMoves9000(towers, moves)

	topChars := ""

	for _, t := range towers {
		topChars += t.Peek()
	}

	return topChars
}

func solveChallengeTwo() string {
	towers, moves := readInput()
	handleMoves9001(towers, moves)

	topChars := ""

	for _, t := range towers {
		topChars += t.Peek()
	}

	return topChars
}

func readInput() ([]h.Stack, []move) {
	// open file
	file, err := os.Open("./input_01.txt")
	handleError(err)

	defer func(file *os.File) {
		err := file.Close()
		handleError(err)
	}(file)

	// access file buffer
	scanner := bufio.NewScanner(file)

	// init variables
	var towers []h.Stack
	var moves []move

	for scanner.Scan() {
		line := scanner.Text()

		// check if line is part of tower or a move
		if strings.Contains(line, "[") {
			parseLineToTower(&towers, line)
		} else if len(line) > 0 && string(line[0]) == "m" {
			parseLineToMove(&moves, line)
		}
	}

	return towers, moves
}

func parseLineToTower(t *[]h.Stack, line string) {
	cnt := 0

	for i := 0; i < len(line); i += 4 {
		currChar := string(line[i+1])

		if cnt >= len(*t) {
			var newTower h.Stack
			*t = append(*t, newTower)
		}

		if currChar != " " {
			(*t)[cnt].PushFront(currChar)
		}

		cnt += 1
	}
}

func parseLineToMove(m *[]move, line string) {
	lineSplit := strings.Split(line, " ")

	amount, err := strconv.Atoi(lineSplit[1])
	handleError(err)

	from, err := strconv.Atoi(lineSplit[3])
	handleError(err)

	to, err := strconv.Atoi(lineSplit[5])
	handleError(err)

	// from/to needs to be -1, so it matches index of tower
	newMove := move{
		amount: amount,
		from:   from - 1,
		to:     to - 1,
	}

	*m = append(*m, newMove)
}

func handleMoves9000(towers []h.Stack, moves []move) {
	for _, m := range moves {
		for i := 0; i < m.amount; i++ {
			val, success := towers[m.from].Pop()

			if success {
				towers[m.to].Push(val)
			}
		}
	}
}

func handleMoves9001(towers []h.Stack, moves []move) {
	var tmpStack h.Stack

	for _, m := range moves {
		for i := 0; i < m.amount; i++ {
			val, success := towers[m.from].Pop()

			if success {
				tmpStack.Push(val)
			}
		}

		for !tmpStack.IsEmpty() {
			val, success := tmpStack.Pop()

			if success {
				towers[m.to].Push(val)
			}
		}
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
