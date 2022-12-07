package main

import (
	"2022/h"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	result1 := solveChallengeOne()
	fmt.Printf("Result-1: %d\n", result1)

	result2 := solveChallengeTwo()
	fmt.Printf("Result-2: %d\n", result2)
}

func solveChallengeOne() int {
	lines := readInput()

	if len(lines) != 1 {
		return -1
	}

	return getMarkerPosition(lines[0], 4)
}

func solveChallengeTwo() int {
	lines := readInput()

	if len(lines) != 1 {
		return -1
	}

	return getMarkerPosition(lines[0], 14)
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

func getMarkerPosition(line string, mLen int) int {
	var marker h.Stack

	for i := 0; i < mLen; i++ {
		marker.PushFront(string(line[i]))
	}

	for i := mLen; i < len(line); i++ {
		if marker.CollectionUnique() {
			return i
		}

		marker.Pop()
		marker.PushFront(string(line[i]))
	}

	return -1
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
