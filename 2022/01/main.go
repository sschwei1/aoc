package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	result1 := solveChallengeOne()
	fmt.Printf("Result-1: %d\n", result1)

	result2 := solveChallengeTwo()
    fmt.Printf("Result-2: %d\n", result2)
}

func solveChallengeOne() int {
	elves := readElves()
	maxVal, _ := getMaxValue(elves)
	return maxVal
}

func solveChallengeTwo() int {
	elves := readElves()

	topThree := 0
	for i := 0; i < 3; i++ {
		maxVal, index := getMaxValue(elves)
		topThree += maxVal
		elves = removeElementAtIndex(elves, index)
	}

	return topThree
}

func readElves() []int {
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

	// init variables
	var elves []int
	counter := 0

	// setup start
	addElv(&elves)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			counter = counter + 1
			addElv(&elves)
		} else {
			calories, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			elves[counter] += calories
		}
	}

	return elves
}

func addElv(elves *[]int) {
	*elves = append(*elves, 0)
}

func getMaxValue(s []int) (int, int) {
	if len(s) == 0 {
		return 0, -1
	}

	max := s[0]
	index := 0

	for i, value := range s {
		if value > max {
			max = value
			index = i
		}
	}

	return max, index
}

func removeElementAtIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
