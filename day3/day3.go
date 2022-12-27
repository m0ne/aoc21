package three

import (
	"aoc21/util"
	"bufio"
	"os"
	"strconv"
)

func readInput(path string) []string {
	f, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var fileLines []string

	for scanner.Scan() {
		line := scanner.Text()

		fileLines = append(fileLines, line)
	}

	return fileLines
}

func partOne(input []string) int {
	mostSignificantBits := ""
	leastSiginificantBits := ""
	countOnesMS, countZeroesMS := 0, 0

	for i := 0; i < len(input[0]); i++ {
		for k := 0; k < len(input); k++ {
			if input[k][i] == '1' {
				countOnesMS += 1
			} else {
				countZeroesMS += 1
			}
		}

		if countOnesMS > countZeroesMS {
			mostSignificantBits += "1"
			leastSiginificantBits += "0"
		} else {
			mostSignificantBits += "0"
			leastSiginificantBits += "1"
		}
		countOnesMS, countZeroesMS = 0, 0
	}

	mSB, error := strconv.ParseInt(mostSignificantBits, 2, 0)
	lSB, error := strconv.ParseInt(leastSiginificantBits, 2, 0)

	if error != nil {
		panic(error)
	}

	return int(mSB * lSB)
}

func countBits(index int, array []string) int {
	counter := 0
	for k := 0; k < len(array); k++ {
		if array[k][index] == '1' {
			counter += 1
		} else {
			counter -= 1
		}
	}
	return counter
}

func selectValidOxygenGeneratorBits(counter int, position int, array []string) []string {
	newOxygenGeneratorBits := []string{}

	if len(array) > 1 {
		for _, bit := range array {
			if counter >= 0 && bit[position] == '1' {
				newOxygenGeneratorBits = append(newOxygenGeneratorBits, bit)
			} else if counter < 0 && bit[position] == '0' {
				newOxygenGeneratorBits = append(newOxygenGeneratorBits, bit)
			}
		}
	} else {
		return array
	}
	return newOxygenGeneratorBits
}

func selectValidCo2ScrubberBits(counter int, position int, array []string) []string {
	newCO2ScrubberBits := []string{}

	if len(array) > 1 {
		for _, bit := range array {
			if counter < 0 && bit[position] == '1' {
				newCO2ScrubberBits = append(newCO2ScrubberBits, bit)
			} else if counter >= 0 && bit[position] == '0' {
				newCO2ScrubberBits = append(newCO2ScrubberBits, bit)
			}
		}
	} else {
		return array
	}
	return newCO2ScrubberBits
}

func parseBits(bits string) int64 {
	bitsAsDecimal, error := strconv.ParseInt(bits, 2, 64)

	if error != nil {
		panic(error)
	}

	return bitsAsDecimal
}

func partTwo(input []string) int {
	binaryBits := input
	oxygenGeneratorBits := binaryBits
	co2ScrubberBits := binaryBits

	oxygenCounter := 0
	co2Counter := 0
	for i := range binaryBits[0] {
		oxygenCounter = countBits(i, oxygenGeneratorBits)
		co2Counter = countBits(i, co2ScrubberBits)

		oxygenGeneratorBits = selectValidOxygenGeneratorBits(oxygenCounter, i, oxygenGeneratorBits)
		co2ScrubberBits = selectValidCo2ScrubberBits(co2Counter, i, co2ScrubberBits)
	}

	co2ScrubberRating := parseBits(co2ScrubberBits[0])
	oxygenGeneratorRating := parseBits(oxygenGeneratorBits[0])

	resultPartTwo := int(co2ScrubberRating) * int(oxygenGeneratorRating)

	return resultPartTwo
}

func Run() {
	const input = "day3/input.txt"
	serializedInput := readInput(input)

	var resultPartOne = strconv.Itoa(partOne(serializedInput))
	var resultPartTwo = strconv.Itoa(partTwo(serializedInput))
	util.PrintHeader(3, resultPartOne, resultPartTwo)
}
