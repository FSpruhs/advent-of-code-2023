package day6

import (
	"fmt"
	"github.com/FSpruhs/advent-of-code-2023/util"
	"log"
	"regexp"
	"strings"
)

var whiteSpaceRegex = regexp.MustCompile(`\s{2,}`)

func clearString(str *string) string {
	return whiteSpaceRegex.ReplaceAllString(*str, " ")
}

var numericRegex = regexp.MustCompile(`[^0-9]+`)

func clearStringAll(str *string) string {
	return numericRegex.ReplaceAllString(*str, "")
}

func Solve(filepath string) {
	input := util.ReadFile(filepath)
	time := util.StringSliceToIntSlice(strings.Split(clearString(&(*input)[0]), " ")[1:])
	distance := util.StringSliceToIntSlice(strings.Split(clearString(&(*input)[1]), " ")[1:])
	solutionPartOne := 1
	for i := 0; i < len(distance); i++ {
		counter := 0
		for j := 1; j < time[i]; j++ {
			speed := j
			drivingTime := time[i] - j
			dist := speed * drivingTime
			if dist > distance[i] {
				counter++
			}
		}
		solutionPartOne *= counter
	}
	timeTwo := util.StringSliceToIntSlice(strings.Split(clearStringAll(&(*input)[0]), " "))
	distanceTwo := util.StringSliceToIntSlice(strings.Split(clearStringAll(&(*input)[1]), " "))
	log.Println(timeTwo, distanceTwo)
	solutionPartTwo := 0
	for j := 1; j < timeTwo[0]; j++ {
		speed := j
		drivingTime := timeTwo[0] - j
		dist := speed * drivingTime
		if dist > distanceTwo[0] {
			solutionPartTwo++
		}
	}

	fmt.Printf("Solution of day 6 part 1: %d\n", solutionPartOne)
	fmt.Printf("Solution of day 6 part 2: %d\n", solutionPartTwo)
}
