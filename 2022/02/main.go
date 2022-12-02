package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type game struct {
	opponent int
	self     int
}

var pointMap = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,

	"X": 1,
	"Y": 2,
	"Z": 3,
}

func main() {
	result1 := solveChallengeOne()
	fmt.Printf("Result-1: %d\n", result1)

	result2 := solveChallengeTwo()
	fmt.Printf("Result-2: %d\n", result2)
}

func solveChallengeOne() int {
	totalPoints := 0
	games := readGames()

	for _, g := range games {
		points := calcScore(g)
		totalPoints += points
	}

	return totalPoints
}

func solveChallengeTwo() int {
	totalPoints := 0
	games := readGames()

	for _, g := range games {
		points := calcNewScore(g)
		totalPoints += points
	}

	return totalPoints
}

func readGames() []game {
	// open file
	file, err := os.Open("./input_01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	// access file buffer
	scanner := bufio.NewScanner(file)
	var games []game

	lineNumber := 0

	for scanner.Scan() {
		lineNumber += 1
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")
		if len(lineSplit) != 2 {
			log.Fatalf("Invalid input at line %d", lineNumber)
		} else {
			currGame := game{
				opponent: pointMap[lineSplit[0]],
				self:     pointMap[lineSplit[1]],
			}

			games = append(games, currGame)
		}
	}

	return games
}

func calcScore(g game) int {
	points := g.self
	diff := g.opponent - g.self

	if diff < 0 {
		diff += 3
	}

	// diff == 1 => lose game, dont give points
	switch diff {
	case 0:
		points += 3
		break
	case 2:
		points += 6
		break
	}

	return points
}

func calcNewScore(g game) int {
	var newChoice int

	switch g.self {
	case 1:
		newChoice = g.opponent - 1
		break
	case 2:
		newChoice = g.opponent
		break
	case 3:
		newChoice = g.opponent - 2
		break
	}

	if newChoice < 1 {
		newChoice += 3
	}

	g.self = newChoice

	return calcScore(g)
}
