package day1

import (
	"fmt"
	"github.com/FSpruhs/advent-of-code-2023/util"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var numericRegex = regexp.MustCompile(`[^0-9 ]+`)

func clearString(str *string) string {
	return numericRegex.ReplaceAllString(*str, "")
}

func convertToInt(str *string) int {
	number, err := strconv.Atoi(readTwoDigitNumber(clearString(str)))
	if err != nil {
		log.Printf("could not transform %s to an int", *str)
		os.Exit(1)
	}
	return number

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
		resultPartOne += convertToInt(&inputLine)
		resultPartTwo += convertToInt(mapTextToNumber(inputLine))
	}

	fmt.Printf("Solution of day 1 part 1 is: %d\n", resultPartOne)
	fmt.Printf("Solution of day 1 part 2 is: %d\n", resultPartTwo)
}

func mapTextToNumber(line string) *string {
	result := strings.ReplaceAll(line, "one", "one1one")
	result = strings.ReplaceAll(result, "two", "two2two")
	result = strings.ReplaceAll(result, "three", "three3three")
	result = strings.ReplaceAll(result, "four", "four4four")
	result = strings.ReplaceAll(result, "five", "five5five")
	result = strings.ReplaceAll(result, "six", "six6six")
	result = strings.ReplaceAll(result, "seven", "seven7seven")
	result = strings.ReplaceAll(result, "eight", "eight8eight")
	result = strings.ReplaceAll(result, "nine", "nine9nine")
	return &result
}
