package two

import (
	"testing"
)

var puzzleInput = [][]string{{"forward", "5"},
	{"down", "5"},
	{"forward", "8"},
	{"up", "3"},
	{"down", "8"},
	{"forward", "2"}}

const solutionPartOne int = 150

//const solutionPartTwo int = 5

func TestPartOne(t *testing.T) {
	answer1 := partOne(puzzleInput)
	if answer1 != solutionPartOne {
		t.Fatalf("Answer 1: %v is not equal to correct answer %v", answer1, solutionPartOne)
	}
}

/*func TestPartTw(t *testing.T) {
	answer2 := partTwo(puzzleInput)
	if answer2 != solutionPartTwo {
		t.Fatalf("Answer 2: %v is not equal to correct answer %v", answer2, solutionPartTwo)
	}
} */
