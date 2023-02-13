package main

import (
	"2022/h"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type directory struct {
	subDirs map[string]directory
	files   map[string]int
}

func main() {
	result1 := solveChallengeOne()
	fmt.Printf("Result-1: %d\n", result1)

	result2 := solveChallengeTwo()
	fmt.Printf("Result-2: %d\n", result2)
}

func solveChallengeOne() int {
	lines := readInput()

	srcDir := calcDirectory(lines)
	dirDeleteCnt := 0

	searchDirLessThan100k(srcDir, &dirDeleteCnt)

	return dirDeleteCnt
}

func solveChallengeTwo() int {
	lines := readInput()

	srcDir := calcDirectory(lines)

	// 40k is the max space allowed for all dirs
	spaceToFreeUp := calcDirSize(srcDir) - 40000000
	sizes := getDirSizesToFreeUp(srcDir, spaceToFreeUp)

	return getMinElement(sizes)
}

func calcDirectory(lines []string) directory {
	var currDir h.Stack

	srcDir := directory{
		subDirs: make(map[string]directory),
		files:   make(map[string]int),
	}

	for i, line := range lines {
		if i == 0 {
			continue
		}

		lineSplit := strings.Split(line, " ")

		if lineSplit[0] == "$" {
			if lineSplit[1] == "cd" {
				if lineSplit[2] == ".." {
					currDir.Pop()
				} else {
					currDir.Push(lineSplit[2])
				}
			}
		} else {
			if lineSplit[0] != "dir" {
				fSize, err := strconv.Atoi(lineSplit[0])
				handleError(err)
				srcDir = addFileToDir(srcDir, currDir, lineSplit[1], fSize)
			}
		}
	}

	return srcDir
}

func addFileToDir(srcDir directory, path h.Stack, fileName string, size int) directory {
	currDir := &srcDir

	for _, p := range path {
		_, ok := (*currDir).subDirs[p]

		if !ok {
			(*currDir).subDirs[p] = directory{
				subDirs: make(map[string]directory),
				files:   make(map[string]int),
			}
		}

		newDir := (*currDir).subDirs[p]
		currDir = &newDir
	}

	currDir.files[fileName] = size

	return srcDir
}

func searchDirLessThan100k(dir directory, totalSize *int) {
	for _, subDir := range dir.subDirs {
		searchDirLessThan100k(subDir, totalSize)
	}

	dirSize := calcDirSize(dir)
	if dirSize <= 100000 {
		*totalSize += dirSize
	}
}

func calcDirSize(dir directory) int {
	totalSize := 0

	for _, subDir := range dir.subDirs {
		totalSize += calcDirSize(subDir)
	}

	for _, fileSize := range dir.files {
		totalSize += fileSize
	}

	return totalSize
}

func getDirSizesToFreeUp(dir directory, minSize int) []int {
	var sizes []int

	for _, subDir := range dir.subDirs {
		subDirSizes := getDirSizesToFreeUp(subDir, minSize)
		sizes = append(sizes, subDirSizes...)
	}

	dirSize := calcDirSize(dir)
	if dirSize >= minSize {
		sizes = append(sizes, dirSize)
	}

	return sizes
}

func getMinElement(sizes []int) int {
	if len(sizes) == 0 {
		return -1
	}

	minVal := sizes[0]

	for _, val := range sizes {
		if val < minVal {
			minVal = val
		}
	}

	return minVal
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
