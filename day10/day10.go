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
	log.Println(firstComeFrom, comeFrom)
	if firstComeFrom == south && comeFrom == south {
		stretchedGrid[start.row*2+1][start.column*2+1] = "|"
	}
	if firstComeFrom == north && comeFrom == north {
		stretchedGrid[start.row*2+1][start.column*2+1] = "|"
	}
	if firstComeFrom == east && comeFrom == east {
		stretchedGrid[start.row*2+1][start.column*2+1] = "-"
	}
	if firstComeFrom == west && comeFrom == west {
		stretchedGrid[start.row*2+1][start.column*2+1] = "-"
	}
	if firstComeFrom == south && comeFrom == east {
		stretchedGrid[start.row*2+1][start.column*2+1] = "L"
	}
	if firstComeFrom == west && comeFrom == north {
		stretchedGrid[start.row*2+1][start.column*2+1] = "L"
	}
	if firstComeFrom == north && comeFrom == west {
		stretchedGrid[start.row*2+1][start.column*2+1] = "7"
	}
	if firstComeFrom == east && comeFrom == south {
		stretchedGrid[start.row*2+1][start.column*2+1] = "7"
	}
	if firstComeFrom == north && comeFrom == east {
		stretchedGrid[start.row*2+1][start.column*2+1] = "F"
	}
	if firstComeFrom == west && comeFrom == south {
		stretchedGrid[start.row*2+1][start.column*2+1] = "F"
	}
	if firstComeFrom == south && comeFrom == west {
		stretchedGrid[start.row*2+1][start.column*2+1] = "J"
	}
	if firstComeFrom == east && comeFrom == north {
		stretchedGrid[start.row*2+1][start.column*2+1] = "J"
	}

	//stretchedGrid[start.row*2+1][start.column*2+1] = replaceStart(stretchedGrid, start)

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

func replaceStart(grid [][]string, startInput position, from direction) string {
	start := position{
		row:    startInput.row*2 + 1,
		column: startInput.column*2 + 1,
	}
	if from == south {

	}
	if start.row-2 >= 0 && start.column-2 >= 0 {
		if grid[start.row-2][start.column] != "." && grid[start.row][start.column-2] != "." {
			return "L"
		}
	}
	if start.row-2 >= 0 && start.column+2 <= len(grid[0]) {
		if grid[start.row-2][start.column] != "." && grid[start.row][start.column+2] != "." {
			return "J"
		}
	}
	if start.row+2 <= len(grid) && start.column-2 >= 0 {
		if grid[start.row+2][start.column] != "." && grid[start.row][start.column-2] != "." {
			return "7"
		}
	}
	if start.row+2 <= len(grid) && start.column+2 <= len(grid[0]) {
		if grid[start.row+2][start.column] != "." && grid[start.row][start.column+2] != "." {
			return "F"
		}
	}
	if start.row+2 <= len(grid) && start.row-2 >= 0 {
		if grid[start.row+2][start.column] != "." && grid[start.row-2][start.column] != "." {
			return "|"
		}
	}
	if start.column+2 <= len(grid[0]) && start.column-2 >= 0 {
		if grid[start.row][start.column+2] != "." && grid[start.row][start.column-2] != "." {
			return "-"
		}
	}
	return "."
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
