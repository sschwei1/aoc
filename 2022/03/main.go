package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

func main() {
    result1 := solveChallengeOne()
    fmt.Printf("Result-1: %d\n", result1)

    result2 := solveChallengeTwo()
    fmt.Printf("Result-2: %d\n", result2)
}

func solveChallengeOne() int {
    rucksacks := readRucksacks()
    totalPriority := 0

    for _, currRucksack := range rucksacks {
        duplicateItem := getDuplicatedItem(currRucksack)
        totalPriority += getItemPriority(duplicateItem)
    }

    return totalPriority
}

func solveChallengeTwo() int {
    rucksacks := readRucksacks()
    badgePriority := 0

    for i := 0; i < len(rucksacks); i += 3 {
        currGroup := rucksacks[i:i+3]
        badge := getCommonItem(currGroup)
        badgePriority += getItemPriority(badge)
    }

    return badgePriority
}

func readRucksacks() []string {
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

    var lines []string
    for scanner.Scan() {
        newLine := scanner.Text()
        lines = append(lines, newLine)
    }

    return lines
}

func getDuplicatedItem(rucksack string) int {
    compartmentSize := len(rucksack) / 2

    for i := 0; i < compartmentSize; i++ {
        currItem := rucksack[i]

        for j := compartmentSize; j < 2 * compartmentSize; j++ {
            itemToCheck := rucksack[j]

            if(currItem == itemToCheck) {
                return int(currItem)
            }
        }
    }

    // should never occur due to given challenge
    return 0;
}

func getCommonItem(rucksacks []string) int {
    sRucksack, sRucksackIndex := getSmallestRucksack(rucksacks)

    for i := 0; i < len(sRucksack); i++ {
        currChar := string(sRucksack[i])
        contained := true

        for j := 0; contained && j < len(rucksacks); j++ {
            if(j == sRucksackIndex) {
                continue
            }

            rucksack := rucksacks[j]

            if(!strings.Contains(rucksack, currChar)) {
                contained = false
            }
        }

        if contained {
            return int(currChar[0])
        }
    }

    // should never occur due to given challenge
    return 0
}

func getSmallestRucksack(rucksacks []string) (string, int) {
    if(len(rucksacks) == 0) {
        return "", -1
    }

    smallestRucksack := rucksacks[0]
    smallestRucksackIndex := 0

    for i, rucksack := range rucksacks {
        if(len(rucksack) < len(smallestRucksack)) {
            smallestRucksack = rucksack
            smallestRucksackIndex = i
        }
    }

    return smallestRucksack, smallestRucksackIndex
}

func getItemPriority(item int) int {
    var priority int

    if(item < 96) {
        priority = item - 64 + 26
    } else {
        priority = item - 96
    }

    return priority
}