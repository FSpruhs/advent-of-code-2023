package day9

import (
	"fmt"
	"github.com/FSpruhs/advent-of-code-2023/util"
	"strings"
)

func decodeLine(decodedLine [][]int) [][]int {
	if decodedLineIsZeroOnly(decodedLine[len(decodedLine)-1]) {
		return decodedLine
	}
	var newLine []int
	for i := 0; i < len(decodedLine[len(decodedLine)-1])-1; i++ {
		newLine = append(newLine, decodedLine[len(decodedLine)-1][i+1]-decodedLine[len(decodedLine)-1][i])
	}
	decodedLine = append(decodedLine, newLine)
	return decodeLine(decodedLine)

}

func decodedLineIsZeroOnly(line []int) bool {
	for _, number := range line {
		if number != 0 {
			return false
		}
	}
	return true
}

func Solve(filePath string) {
	input := util.ReadFile(filePath)
	solutionPartOne := 0
	for _, line := range *input {
		numberLine := util.StringSliceToIntSlice(strings.Split(line, " "))
		decodedLine := [][]int{numberLine}
		decodedLine = decodeLine(decodedLine)
		for i := len(decodedLine) - 2; i >= 0; i-- {
			decodedLine[i] = append(decodedLine[i], decodedLine[i][len(decodedLine[i])-1]+decodedLine[i+1][len(decodedLine[i+1])-1])
		}
		solutionPartOne += decodedLine[0][len(decodedLine[0])-1]
	}

	fmt.Printf("Solution of day 9 part 1 is: %d\n", solutionPartOne)

}
