package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type position struct {
	x int
	y int
}

func main() {
	result1 := solveChallengeOne()
	fmt.Printf("Result-1: %d\n", result1)

	result2 := solveChallengeTwo()
	fmt.Printf("Result-2: %d\n", result2)
}

func solveChallengeOne() int {
	lines := readInput()

	visitedPos := make(map[position]bool)
	headPos := position{x: 0, y: 0}
	tailPos := position{x: 0, y: 0}

	for _, line := range lines {
		handleMove(line, &headPos, &tailPos, &visitedPos)
	}

	return len(visitedPos)
}

func solveChallengeTwo() int {
	return 0
}

func handleMove(move string, headPos *position, tailPos *position, visitedPos *map[position]bool) {
	moveSplit := strings.Split(move, " ")

	moveDir := moveSplit[0]
	moveAmount, err := strconv.Atoi(moveSplit[1])
	handleError(err)

	for i := 0; i < moveAmount; i++ {
		if moveDir == "U" {
			headPos.y += 1
		} else if moveDir == "D" {
			headPos.y += -1
		} else if moveDir == "L" {
			headPos.x += -1
		} else if moveDir == "R" {
			headPos.x += 1
		}

		moveTail(*headPos, tailPos)
		(*visitedPos)[*tailPos] = true
	}
}

func moveTail(headPos position, tailPos *position) {
	xDiff := headPos.x - tailPos.x
	yDiff := headPos.y - tailPos.y

	if getAbs(xDiff) <= 1 && getAbs(yDiff) <= 1 {
		return
	}

	if getAbs(xDiff) > 0 {
		xDir := xDiff / getAbs(xDiff)
		tailPos.x += xDir
	}

	if getAbs(yDiff) > 0 {
		yDir := yDiff / getAbs(yDiff)
		tailPos.y += yDir
	}
}

func getAbs(num int) int {
	if num >= 0 {
		return num
	}
	return num * -1
}

func readInput() []string {
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
	var lines []string

	for scanner.Scan() {
		newLine := scanner.Text()
		lines = append(lines, newLine)
	}

	return lines
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
