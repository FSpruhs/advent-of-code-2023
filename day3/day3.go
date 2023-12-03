package day3

import (
	"fmt"
	"github.com/FSpruhs/advent-of-code-2023/util"
	"unicode"
)

const gearSymbol = 42
const pointSymbol = 46

type number struct {
	begin     cord
	end       cord
	neighbors bool
	value     string
}

type gear struct {
	coordinate     cord
	adjacentNumber int
	value          int
}

func (g *gear) incrementAdjacentNumber() {
	g.adjacentNumber += 1
}

func (g *gear) setValue(value int) {
	if g.value == 0 {
		g.value = value
	} else {
		g.value *= value
	}

}

func (n *number) isNeighbors(neighbor bool) {
	n.neighbors = neighbor
}

func (g *gear) isAdjacent(input number) bool {
	for i := input.begin.yCord; i <= input.end.yCord; i++ {
		if g.coordinate.xCord-input.begin.xCord > 1 || g.coordinate.xCord-input.begin.xCord < -1 {
			return false
		}
		switch {
		case g.coordinate.xCord+1 == input.begin.xCord && g.coordinate.yCord+1 == i:
			return true
		case g.coordinate.xCord+1 == input.begin.xCord && g.coordinate.yCord == i:
			return true
		case g.coordinate.xCord+1 == input.begin.xCord && g.coordinate.yCord-1 == i:
			return true
		case g.coordinate.xCord == input.begin.xCord && g.coordinate.yCord+1 == i:
			return true
		case g.coordinate.xCord == input.begin.xCord && g.coordinate.yCord-1 == i:
			return true
		case g.coordinate.xCord-1 == input.begin.xCord && g.coordinate.yCord+1 == i:
			return true
		case g.coordinate.xCord-1 == input.begin.xCord && g.coordinate.yCord == i:
			return true
		case g.coordinate.xCord-1 == input.begin.xCord && g.coordinate.yCord-1 == i:
			return true
		}
	}
	return false
}

type cord struct {
	xCord int
	yCord int
}

func Solve(filePath string) {
	input := util.ReadFile(filePath)
	numbers := findNumbers(input)
	gears := findGears(input)
	findNeighbors(&numbers, input)
	findGearsNeighbors(&gears, &numbers)
	fmt.Printf("Solution of day 3 part 1: %d\n", sumNumberValuesWithNeighbors(&numbers))
	fmt.Printf("Solution of day 3 part 2: %d\n", sumGearValuesWithTwoNeighbors(&gears))
}

func sumGearValuesWithTwoNeighbors(gears *[]gear) int {
	result := 0
	for _, gear := range *gears {
		if gear.adjacentNumber == 2 {
			result += gear.value
		}
	}
	return result
}

func findGearsNeighbors(gears *[]gear, input *[]number) {
	for i, gear := range *gears {
		for _, n := range *input {
			if gear.isAdjacent(n) {
				(*gears)[i].incrementAdjacentNumber()
				(*gears)[i].setValue(util.ConvertToInt(n.value))
			}
		}
	}
}

func findGears(input *[]string) []gear {
	var gears []gear
	for x, line := range *input {
		for y, symbol := range line {
			if symbol == gearSymbol {
				gears = append(gears, gear{coordinate: cord{
					xCord: x,
					yCord: y,
				}})
			}

		}
	}
	return gears
}

func sumNumberValuesWithNeighbors(numbers *[]number) int {
	result := 0
	for index := range *numbers {
		if (*numbers)[index].neighbors == true {
			result += getNumberValue((*numbers)[index])
		}
	}
	return result
}

func getNumberValue(n number) int {
	return util.ConvertToInt(n.value)
}

func findNeighbors(numbers *[]number, input *[]string) {
	for index := range *numbers {
		findNeighbor(&(*numbers)[index], input)
	}
}

func calculateBeginEndIndex(value *number, input *[]string) (int, int, int, int) {
	beginX := calculateBeginX(value)
	endX := calculateEndX(value, input)
	beginY := calculateBeginY(value)
	endY := calculateEndY(value, input)
	return beginX, endX, beginY, endY
}

func calculateEndY(value *number, input *[]string) int {
	if value.end.yCord >= len((*input)[0])-1 {
		return value.end.yCord
	}
	return value.end.yCord + 1
}

func calculateBeginY(value *number) int {
	if value.begin.yCord == 0 {
		return 0
	}
	return value.begin.yCord - 1
}

func calculateEndX(value *number, input *[]string) int {
	if value.end.xCord >= len(*input)-1 {
		return value.end.xCord
	}
	return value.end.xCord + 1
}

func calculateBeginX(value *number) int {
	if value.begin.xCord == 0 {
		return 0
	}
	return value.begin.xCord - 1
}

func findNeighbor(value *number, input *[]string) {
	beginX, endX, beginY, endY := calculateBeginEndIndex(value, input)
	for i := beginX; i <= endX; i++ {
		for j := beginY; j <= endY; j++ {
			if !unicode.IsDigit(rune((*input)[i][j])) && (*input)[i][j] != pointSymbol {
				value.isNeighbors(true)
			}
		}
	}
}

func findNumbers(input *[]string) []number {
	var numbers []number
	for x, line := range *input {
		numbers = findNumbersInLine(line, numbers, x)
	}
	return numbers
}

func findNumbersInLine(line string, numbers []number, x int) []number {
	begin := 0
	digitFlag := false
	for y, symbol := range line {
		if unicode.IsDigit(symbol) {
			digitFlag, begin = handleDigit(digitFlag, begin, y)
		} else if digitFlag {
			digitFlag = false
			numbers = addNumber(numbers, x, begin, y, line)
		}
		numbers = findLastNumberInLine(line, numbers, x, symbol, y, begin)
	}
	return numbers
}

func findLastNumberInLine(line string, numbers []number, x int, symbol int32, y int, begin int) []number {
	if unicode.IsDigit(symbol) && y == len(line)-1 {
		numbers = addNumber(numbers, x, begin, y+1, line)
	}
	return numbers
}

func addNumber(numbers []number, x int, begin int, y int, line string) []number {
	return append(numbers, number{
		begin: cord{xCord: x, yCord: begin},
		end: cord{
			xCord: x,
			yCord: y - 1,
		},
		neighbors: false,
		value:     line[begin:y],
	})
}

func handleDigit(digitFlag bool, begin int, y int) (bool, int) {
	if !digitFlag {
		begin = y
		digitFlag = true
	}
	return digitFlag, begin
}
