package one

import (
	"aoc21/util"
	"bufio"
	"os"
	"strconv"
)

func readInput() []int {
	// Sorting the given slice
	f, err := os.Open("day1/input.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var fileLines []int

	for scanner.Scan() {
		line, err := strconv.Atoi(scanner.Text())

		if err != nil {
			panic(err)
		}
		fileLines = append(fileLines, line)
	}

	return fileLines
}

func partOne(input []int) int {
	var previousValue int
	counter := 0
	for index, value := range input {
		if index == 0 {
			previousValue = value
		}
		if previousValue < value {
			counter += 1
		}
		previousValue = value
	}
	return counter

}

func partTwo(input []int) int {
	solutionCounter := 0
	for i := 0; i < len(input)-3; i++ {
		sumOne := input[i] + input[i+1] + input[i+2]
		sumTwo := input[i+1] + input[i+2] + input[i+3]

		if sumOne < sumTwo {
			solutionCounter += 1
		}
	}

	return solutionCounter
}

func Run() {
	var resultPartOne = strconv.Itoa(partOne(readInput()))
	var resultPartTwo = strconv.Itoa(partTwo(readInput()))
	util.PrintHeader(1, resultPartOne, resultPartTwo)
}
