package day10

import (
	"fmt"
	"github.com/FSpruhs/advent-of-code-2023/util"
	"log"
	"strings"
)

type direction int

const (
	north direction = iota
	east
	south
	west
)

type position struct {
	row    int
	column int
}

func Solve(filePath string) {
	input := util.ReadFile(filePath)
	grid := [][]string{}
	var start position
	for index, line := range *input {
		grid = append(grid, strings.Split(line, ""))
		if strings.Contains(line, "S") {
			start = position{
				row:    index,
				column: strings.Index(line, "S"),
			}
		}
	}

	comeFrom, currentPosition := findFirstMove(grid, start)

	counter := 1
	for grid[currentPosition.row][currentPosition.column] != "S" {
		log.Println(grid[currentPosition.row][currentPosition.column])
		comeFrom, currentPosition = nextPosition(comeFrom, currentPosition, grid)
		counter++
	}
	fmt.Printf("Solution of day 10 part 1 is: %d\n", counter/2)
}

func nextPosition(from direction, currentPosition position, grid [][]string) (direction, position) {
	switch grid[currentPosition.row][currentPosition.column] {
	case "|":
		if from == south {
			return south, position{
				row:    currentPosition.row - 1,
				column: currentPosition.column,
			}
		}
		return north, position{
			row:    currentPosition.row + 1,
			column: currentPosition.column,
		}
	case "-":
		if from == west {
			return west, position{
				row:    currentPosition.row,
				column: currentPosition.column + 1,
			}
		}
		return east, position{
			row:    currentPosition.row,
			column: currentPosition.column - 1,
		}
	case "L":
		if from == north {
			return west, position{
				row:    currentPosition.row,
				column: currentPosition.column + 1,
			}
		}
		return south, position{
			row:    currentPosition.row - 1,
			column: currentPosition.column,
		}
	case "J":
		if from == north {
			return east, position{
				row:    currentPosition.row,
				column: currentPosition.column - 1,
			}
		}
		return south, position{
			row:    currentPosition.row - 1,
			column: currentPosition.column,
		}
	case "7":
		if from == south {
			return east, position{
				row:    currentPosition.row,
				column: currentPosition.column - 1,
			}
		}
		return north, position{
			row:    currentPosition.row + 1,
			column: currentPosition.column,
		}
	}
	if from == south {
		return west, position{
			row:    currentPosition.row,
			column: currentPosition.column + 1,
		}
	}
	return north, position{
		row:    currentPosition.row + 1,
		column: currentPosition.column,
	}
}

func findFirstMove(grid [][]string, start position) (direction, position) {
	if start.row > 0 {
		if strings.Contains("|7F", grid[start.row-1][start.column]) {
			return south, position{
				row:    start.row - 1,
				column: start.column,
			}
		}
	}
	if start.column < len(grid[0])-2 {
		if strings.Contains("-J7", grid[start.row][start.column+1]) {
			return west, position{
				row:    start.row,
				column: start.column + 1,
			}
		}
	}
	if start.row < len(grid)-2 {
		if strings.Contains("|LJ", grid[start.row+1][start.column]) {
			return north, position{
				row:    start.row + 1,
				column: start.column,
			}
		}
	}

	return east, position{
		row:    start.row,
		column: start.column - 1,
	}
}
