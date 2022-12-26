package two

import (
	"aoc21/util"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func readInput(path string) [][]string {
	// Sorting the given slice
	f, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var fileLines [][]string

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		fileLines = append(fileLines, line)
	}

	return fileLines
}

func partOne(input [][]string) int {
	depth, forward := 0, 0
	for _, value := range input {
		units, error := strconv.Atoi(value[1])

		if error != nil {
			panic(error)
		}

		switch value[0] {
		case "forward":
			forward += units
		case "down":
			depth += units
		case "up":
			depth -= units
		}
	}

	return depth * forward
}

func partTwo(input []string) int {
	return 2
}

func Run() {
	const input = "day2/input.txt"
	serializedInput := readInput(input)

	var resultPartOne = strconv.Itoa(partOne(serializedInput))
	//var resultPartTwo = strconv.Itoa(partTwo(serializedInput))
	util.PrintHeader(2, resultPartOne, "none")
}
