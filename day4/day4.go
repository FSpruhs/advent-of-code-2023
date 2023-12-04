package day4

import (
	"fmt"
	"github.com/FSpruhs/advent-of-code-2023/util"
	"log"
	"math"
	"os"
	"regexp"
	"strings"
)

var numericRegex = regexp.MustCompile(`[^0-9 ]+`)

func clearString(str *string) string {
	return numericRegex.ReplaceAllString(*str, "")
}

func Solve(filePath string) {
	input := util.ReadFile(filePath)
	resultPartOne := 0
	cardPile := initCarPile(len(*input))
	for cardNumber, line := range *input {
		card, givenNumbers := readLine(line)
		winingNumbersInGivenNumbers := countWiningNumbersInGivenNumbers(card, givenNumbers)
		resultPartOne += calculateValue(winingNumbersInGivenNumbers)
		for i := 1; i <= winingNumbersInGivenNumbers; i++ {
			if cardNumber+1 <= len(*input) {
				cardPile[cardNumber+1+i] += cardPile[cardNumber+1]
			}

		}
	}
	fmt.Printf("Solution of day 4 part 1 is %d\n", resultPartOne)
	fmt.Printf("Solution of day 4 part 2 is %d\n", calculateCardPileValue(cardPile))
}

func initCarPile(length int) map[int]int {
	cardPile := map[int]int{}
	for i := 0; i < length; i++ {
		cardPile[i+1] = 1
	}
	return cardPile
}

func calculateCardPileValue(pile map[int]int) int {
	counter := 0
	for _, numberOfCards := range pile {
		counter += numberOfCards
	}
	return counter
}

func calculateValue(count int) int {
	if count == 0 {
		return 0
	}
	return int(math.Pow(2, float64(count-1)))
}

func countWiningNumbersInGivenNumbers(card map[string]bool, numbers map[string]bool) int {
	counter := 0
	for key := range card {
		if numbers[key] {
			counter++
		}
	}
	return counter
}

func splitLine(line string) []string {
	splittedInput := strings.Split(strings.Split(line, ":")[1], "|")
	if len(splittedInput) != 2 {
		log.Printf("splitted inut has more than 2 parts %s", splittedInput)
		os.Exit(1)
	}
	return splittedInput
}

func readLine(line string) (map[string]bool, map[string]bool) {
	splittedLine := splitLine(line)
	cardSet := map[string]bool{}
	card := readCard(strings.TrimSpace(splittedLine[0]))
	for _, cardItem := range card {
		if cardItem != "" {
			cardSet[cardItem] = true
		}
	}
	givenNumbers := strings.Split(strings.TrimSpace(splittedLine[1]), " ")
	givenNumbersSet := map[string]bool{}
	for _, givenNumberItem := range givenNumbers {
		if givenNumberItem != "" {
			givenNumbersSet[givenNumberItem] = true
		}
	}
	return cardSet, givenNumbersSet
}

func readCard(cardValue string) []string {
	return strings.Split(clearString(&cardValue), " ")
}
