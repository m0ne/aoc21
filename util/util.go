package util

func PrintHeader(day int, solutionPartOne string, solutionPartTwo string) {
	const longDecoration = "------------------------------"
	println(longDecoration)
	println("Day: ", day)
	println("Solution PartOne: ", solutionPartOne)
	println("Solution PartTwo: ", solutionPartTwo)
	println(longDecoration)
}
