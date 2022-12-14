package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	visibleTrees := make(map[position]bool)

	checkVisibilityTop(lines, &visibleTrees)
	checkVisibilityLeft(lines, &visibleTrees)
	checkVisibilityRight(lines, &visibleTrees)
	checkVisibilityBottom(lines, &visibleTrees)

	return len(visibleTrees)
}

func solveChallengeTwo() int {
	trees := readInput()

	maxScore := 0

	for x := 0; x < len(trees[0]); x++ {
		for y := 0; y < len(trees); y++ {
			pos := position{x, y}
			newScore := calcScenicScore(trees, pos)

			if newScore > maxScore {
				maxScore = newScore
			}
		}
	}

	return maxScore
}

func checkVisibilityTop(trees []string, visTrees *map[position]bool) {
	for x := 0; x < len(trees[0]); x++ {
		maxTreeSize := uint8(0)
		for y := 0; y < len(trees); y++ {
			treeSize := trees[y][x]
			if treeSize <= maxTreeSize && y != 0 {
				continue
			}

			pos := position{x: x, y: y}
			(*visTrees)[pos] = true

			maxTreeSize = treeSize
		}
	}
}

func checkVisibilityLeft(trees []string, visTrees *map[position]bool) {
	for y := 0; y < len(trees); y++ {
		maxTreeSize := uint8(0)
		for x := 0; x < len(trees[0]); x++ {
			treeSize := trees[y][x]
			if treeSize <= maxTreeSize && x != 0 {
				continue
			}

			pos := position{x: x, y: y}
			(*visTrees)[pos] = true

			maxTreeSize = treeSize
		}
	}
}

func checkVisibilityRight(trees []string, visTrees *map[position]bool) {
	for y := 0; y < len(trees); y++ {
		maxTreeSize := uint8(0)
		for x := len(trees[0]) - 1; x >= 0; x-- {
			treeSize := trees[y][x]
			if treeSize <= maxTreeSize && x != 0 {
				continue
			}

			pos := position{x: x, y: y}
			(*visTrees)[pos] = true

			maxTreeSize = treeSize
		}
	}
}

func checkVisibilityBottom(trees []string, visTrees *map[position]bool) {
	for x := 0; x < len(trees[0]); x++ {
		maxTreeSize := uint8(0)
		for y := len(trees) - 1; y >= 0; y-- {
			treeSize := trees[y][x]
			if treeSize <= maxTreeSize && y != 0 {
				continue
			}

			pos := position{x: x, y: y}
			(*visTrees)[pos] = true

			maxTreeSize = treeSize
		}
	}
}

func calcScenicScore(trees []string, pos position) int {
	top := getTreesTop(trees, pos)
	left := getTreesLeft(trees, pos)
	right := getTreesRight(trees, pos)
	bottom := getTreesBottom(trees, pos)

	return top * left * right * bottom
}

func getTreesTop(trees []string, pos position) int {
	treeCnt := 0

	treeSize := trees[pos.y][pos.x]

	for y := pos.y - 1; y >= 0; y-- {
		currTreeSize := trees[y][pos.x]
		treeCnt++

		if treeSize <= currTreeSize {
			break
		}
	}

	return treeCnt
}

func getTreesLeft(trees []string, pos position) int {
	treeCnt := 0

	treeSize := trees[pos.y][pos.x]

	for x := pos.x - 1; x >= 0; x-- {
		currTreeSize := trees[pos.y][x]
		treeCnt++

		if treeSize <= currTreeSize {
			break
		}
	}

	return treeCnt
}

func getTreesRight(trees []string, pos position) int {
	treeCnt := 0

	treeSize := trees[pos.y][pos.x]

	for x := pos.x + 1; x < len(trees[0]); x++ {
		currTreeSize := trees[pos.y][x]
		treeCnt++

		if treeSize <= currTreeSize {
			break
		}
	}

	return treeCnt
}

func getTreesBottom(trees []string, pos position) int {
	treeCnt := 0

	treeSize := trees[pos.y][pos.x]

	for y := pos.y + 1; y < len(trees); y++ {
		currTreeSize := trees[y][pos.x]
		treeCnt++

		if treeSize <= currTreeSize {
			break
		}
	}

	return treeCnt
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
