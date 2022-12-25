package one

import (
	"testing"
)

var puzzleInput = []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

const solutionPartOne int = 7
const solutionPartTwo int = 5

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
