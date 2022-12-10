package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type extraCycleFnc func([]string, *int, *extraCycleFnc)

var commands = map[string]extraCycleFnc{
	"noop": nil,
	"addx": handleAddX,
}

func main() {
	result1 := solveChallengeOne()
	fmt.Printf("Result-1: %d\n", result1)

	result2 := solveChallengeTwo()
	fmt.Printf("Result-2:\n%s", result2)
}

func solveChallengeOne() int {
	lines := readInput()

	cycle := 0
	register := 1

	var lineSplit []string
	var execNext extraCycleFnc
	execNext = nil

	regCountTotal := 0

	for i := 0; i < len(lines); i++ {
		cycle += 1

		if (cycle-20)%40 == 0 {
			regCountTotal += register * cycle
		}

		if execNext != nil {
			i -= 1
			execNext(lineSplit, &register, &execNext)
		} else {
			lineSplit = strings.Split(lines[i], " ")

			if val, ok := commands[lineSplit[0]]; ok {
				execNext = val
			}
		}
	}

	return regCountTotal
}

func solveChallengeTwo() string {
	lines := readInput()

	cycle := 0
	register := 1

	result := ""

	var lineSplit []string
	var execNext extraCycleFnc
	execNext = nil

	for i := 0; i < len(lines); i++ {
		cycle += 1

		regCycleDiff := register - (cycle%40 - 1)

		if getAbs(regCycleDiff) > 1 {
			result += " "
		} else {
			result += "#"
		}

		if (cycle)%40 == 0 {
			result += "\n"
		}

		if execNext != nil {
			i -= 1
			execNext(lineSplit, &register, &execNext)
		} else {
			lineSplit = strings.Split(lines[i], " ")

			if val, ok := commands[lineSplit[0]]; ok {
				execNext = val
			}
		}
	}

	return result
}

func handleAddX(command []string, register *int, execNext *extraCycleFnc) {
	val, err := strconv.Atoi(command[1])

	if err != nil {
		handleError(err)
	}

	*register += val
	*execNext = nil
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
