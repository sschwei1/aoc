package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type section struct {
	from int
	to   int
}

type pair struct {
	e1 section
	e2 section
}

func main() {
	result1 := solveChallengeOne()
	fmt.Printf("Result-1: %d\n", result1)

	result2 := solveChallengeTwo()
	fmt.Printf("Result-2: %d\n", result2)
}

func solveChallengeOne() int {
	pairs := readSectionPairs()
	containedCount := 0

	for _, pair := range pairs {
		if checkPairContained(pair) {
			containedCount += 1
		}
	}

	return containedCount
}

func solveChallengeTwo() int {
	return 0
}

func readSectionPairs() []pair {
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

	lineNumber := 0

	var pairs []pair

	for scanner.Scan() {
		lineNumber += 1
		line := scanner.Text()

		newPair := parseInput(line, lineNumber)
		pairs = append(pairs, newPair)
	}

	return pairs
}

func parseInput(line string, lineNumber int) pair {
	var inpPair pair

	lineSplit := strings.Split(line, ",")

	if len(lineSplit) != 2 {
		log.Fatalf("Invalid input at line %d", lineNumber)
	} else {
		inpPair.e1 = parseSection(lineSplit[0], lineNumber)
		inpPair.e2 = parseSection(lineSplit[1], lineNumber)
	}

	return inpPair
}

func parseSection(sectionStr string, lineNumber int) section {
	var newSection section

	lineSplit := strings.Split(sectionStr, "-")

	if len(lineSplit) != 2 {
		log.Fatalf("Invalid input at line %d", lineNumber)
	} else {
		var err error

		newSection.from, err = strconv.Atoi(lineSplit[0])
		if err != nil {
			log.Fatal(err)
		}

		newSection.to, err = strconv.Atoi(lineSplit[1])
		if err != nil {
			log.Fatal(err)
		}
	}

	return newSection
}

func checkPairContained(pair pair) bool {
	return isSectionContained(pair.e1, pair.e2) || isSectionContained(pair.e2, pair.e1)
}

func isSectionContained(s1 section, s2 section) bool {
	return s1.from <= s2.from && s1.to >= s2.to
}
