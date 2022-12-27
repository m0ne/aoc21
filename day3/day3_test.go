package three

import (
	"testing"
)

var puzzleInput = []string{"00100",
"11110",
"10110",
"10111",
"10101",
"01111",
"00111",
"11100",
"10000",
"11001",
"00010",
"01010"}

const solutionPartOne int = 198

const solutionPartTwo int = 230

func TestPartOne(t *testing.T) {
	answer1 := partOne(puzzleInput)
	if answer1 != solutionPartOne {
		t.Fatalf("Answer 1: %v is not equal to correct answer %v", answer1, solutionPartOne)
	}
}

func TestPartTw(t *testing.T) {
	answer2 := partTwo(puzzleInput)
	if answer2 != solutionPartTwo {
		t.Fatalf("Answer 2: %v is not equal to correct answer %v", answer2, solutionPartTwo)
	}
} 
