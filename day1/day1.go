package day1

import (
	"fmt"
	"github.com/FSpruhs/advent-of-code-2023/util"
	"log"
	"regexp"
	"strings"
)

var numericRegex = regexp.MustCompile(`[^0-9 ]+`)

func clearString(str *string) string {
	return numericRegex.ReplaceAllString(*str, "")
}

func getChar(str *string, index int) rune {
	return []rune(*str)[index]
}

func readTwoDigitNumber(numberString string) string {
	if len(numberString) == 0 {
		log.Printf("line %s has no numbers", numberString)
		return "0"
	}
	return string(getChar(&numberString, 0)) + string(getChar(&numberString, len(numberString)-1))
}

func Solve(filePath string) {
	resultPartOne := 0
	resultPartTwo := 0
	for _, inputLine := range *util.ReadFile(filePath) {
		resultPartOne += util.ConvertToInt(readTwoDigitNumber(clearString(&inputLine)))
		resultPartTwo += util.ConvertToInt(readTwoDigitNumber(clearString(mapTextToNumber(inputLine))))
	}

	fmt.Printf("Solution of day 1 part 1 is: %d\n", resultPartOne)
	fmt.Printf("Solution of day 1 part 2 is: %d\n", resultPartTwo)
}

func mapTextToNumber(line string) *string {
	result := strings.ReplaceAll(line, "one", "o1e")
	result = strings.ReplaceAll(result, "two", "t2o")
	result = strings.ReplaceAll(result, "three", "t3e")
	result = strings.ReplaceAll(result, "four", "f4r")
	result = strings.ReplaceAll(result, "five", "f5e")
	result = strings.ReplaceAll(result, "six", "s6x")
	result = strings.ReplaceAll(result, "seven", "s7n")
	result = strings.ReplaceAll(result, "eight", "e8t")
	result = strings.ReplaceAll(result, "nine", "n9e")
	return &result
}
