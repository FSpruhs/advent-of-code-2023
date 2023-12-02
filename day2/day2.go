package day2

import (
	"fmt"
	"github.com/FSpruhs/advent-of-code-2023/util"
	"regexp"
	"strings"
)

var bag = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

var numericAndColorRegex = regexp.MustCompile(`[^\d\sredbluegreen]+`)

func clearString(str *string) string {
	return numericAndColorRegex.ReplaceAllString(*str, "")
}

func findHigherColorCount(rowArray []string) bool {
	for i := 1; i < len(rowArray); i += 2 {
		if bag[rowArray[i+1]] < util.ConvertToInt(rowArray[i]) {
			return true
		}
	}
	return false
}

func Solve(filePath string) {
	resultPartOne := 0
	resultPartTwo := 0
	for _, row := range *util.ReadFile(filePath) {

		rowArray := strings.Fields(clearString(&row))[1:]
		if !findHigherColorCount(rowArray) {
			resultPartOne += util.ConvertToInt(rowArray[0])
		}
		resultPartTwo += valueFrom(findHighestColorCount(rowArray))
	}
	fmt.Printf("Soloution of day 2 part 1: %d\n", resultPartOne)
	fmt.Printf("Soloution of day 2 part 2: %d\n", resultPartTwo)

}

func valueFrom(highestColor map[string]int) int {
	result := 1
	for _, count := range highestColor {
		result *= count
	}
	return result
}

func findHighestColorCount(rowArray []string) map[string]int {
	highestColorCount := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for i := 1; i < len(rowArray); i += 2 {
		if highestColorCount[rowArray[i+1]] < util.ConvertToInt(rowArray[i]) {
			highestColorCount[rowArray[i+1]] = util.ConvertToInt(rowArray[i])
		}
	}
	return highestColorCount
}
