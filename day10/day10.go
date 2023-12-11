package day10

import (
	"fmt"
	"github.com/FSpruhs/advent-of-code-2023/util"
	"log"
	"slices"
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
	stretchedGrid := make([][]string, len(*input)*2+1)
	for i := range stretchedGrid {
		stretchedGrid[i] = make([]string, len((*input)[0])*2+1)
	}
	for i := 0; i < len(stretchedGrid); i++ {
		for j := 0; j < len(stretchedGrid[0]); j++ {
			stretchedGrid[i][j] = "."
		}
	}
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
	firstComeFrom := comeFrom
	counter := 1
	for grid[currentPosition.row][currentPosition.column] != "S" {
		stretchedGrid[currentPosition.row*2+1][currentPosition.column*2+1] = grid[currentPosition.row][currentPosition.column]
		comeFrom, currentPosition = nextPosition(comeFrom, currentPosition, grid)
		counter++
	}
	stretchedGrid[start.row*2+1][start.column*2+1] = replaceStartSymbol(firstComeFrom, comeFrom)

	for i := 2; i < len(stretchedGrid); i += 2 {
		for j := 0; j < len(stretchedGrid[0]); j++ {
			if strings.Contains("|F7", stretchedGrid[i-1][j]) {
				stretchedGrid[i][j] = "|"
			}
		}
	}

	for i := 2; i < len(stretchedGrid[0]); i += 2 {
		for j := 0; j < len(stretchedGrid); j++ {
			if strings.Contains("-LF", stretchedGrid[j][i-1]) {
				stretchedGrid[j][i] = "-"
			}
		}
	}

	stretchedGrid = fill(stretchedGrid)
	solutionPartTwo := 0
	for i := 1; i < len(stretchedGrid); i += 2 {
		for j := 1; j < len(stretchedGrid[0]); j += 2 {
			if stretchedGrid[i][j] == "." {
				solutionPartTwo++
			}
		}
	}

	for _, line := range stretchedGrid {
		log.Println(line)
	}
	fmt.Printf("Solution of day 10 part 1 is: %d\n", counter/2)
	fmt.Printf("Solution of day 10 part 2 is: %d\n", solutionPartTwo)
}

func replaceStartSymbol(firstComeFrom direction, comeFrom direction) string {
	if firstComeFrom == south && comeFrom == south {
		return "|"
	}
	if firstComeFrom == north && comeFrom == north {
		return "|"
	}
	if firstComeFrom == east && comeFrom == east {
		return "-"
	}
	if firstComeFrom == west && comeFrom == west {
		return "-"
	}
	if firstComeFrom == south && comeFrom == east {
		return "L"
	}
	if firstComeFrom == west && comeFrom == north {
		return "L"
	}
	if firstComeFrom == north && comeFrom == west {
		return "7"
	}
	if firstComeFrom == east && comeFrom == south {
		return "7"
	}
	if firstComeFrom == north && comeFrom == east {
		return "F"
	}
	if firstComeFrom == west && comeFrom == south {
		return "F"
	}
	if firstComeFrom == south && comeFrom == west {
		return "J"
	}
	if firstComeFrom == east && comeFrom == north {
		return "J"
	}
	return "."
}

func fill(grid [][]string) [][]string {
	toVisit := []position{{
		row:    0,
		column: 0,
	}}

	for len(toVisit) > 0 {
		mark := toVisit[0]
		grid[mark.row][mark.column] = "X"
		if mark.row > 0 {
			if grid[mark.row-1][mark.column] == "." {
				if !slices.Contains(toVisit, position{
					row:    mark.row - 1,
					column: mark.column,
				}) {
					toVisit = append(toVisit, position{
						row:    mark.row - 1,
						column: mark.column,
					})
				}
			}

		}
		if mark.column < len(grid[0])-1 {
			if grid[mark.row][mark.column+1] == "." {
				if !slices.Contains(toVisit, position{
					row:    mark.row,
					column: mark.column + 1,
				}) {
					toVisit = append(toVisit, position{
						row:    mark.row,
						column: mark.column + 1,
					})
				}
			}
		}
		if mark.row < len(grid)-1 {
			if grid[mark.row+1][mark.column] == "." {
				if !slices.Contains(toVisit, position{
					row:    mark.row + 1,
					column: mark.column,
				}) {
					toVisit = append(toVisit, position{
						row:    mark.row + 1,
						column: mark.column,
					})
				}
			}
		}
		if mark.column > 0 {
			if grid[mark.row][mark.column-1] == "." {
				if !slices.Contains(toVisit, position{
					row:    mark.row,
					column: mark.column - 1,
				}) {
					toVisit = append(toVisit, position{
						row:    mark.row,
						column: mark.column - 1,
					})
				}
			}
		}

		toVisit = toVisit[1:]
	}

	return grid
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
