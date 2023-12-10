package day8

import (
	"fmt"
	"github.com/FSpruhs/advent-of-code-2023/util"
	"log"
)

type node struct {
	name  string
	left  string
	right string
}

func Solve(filePath string) {
	input := util.ReadFile(filePath)
	instructions := (*input)[0]
	nodes := map[string]node{}
	for _, line := range (*input)[2:] {
		nodes[line[0:3]] = node{
			name:  line[0:3],
			left:  line[7:10],
			right: line[12:15],
		}
	}
	actualNode := nodes["AAA"]
	solutionPartOne := 0
	for actualNode.name != "ZZZ" {
		log.Println("find ZZZ")
		for index := 0; index < len(instructions); index++ {
			switch fmt.Sprintf("%c", instructions[index]) {
			case "L":
				actualNode = nodes[actualNode.left]
			case "R":
				actualNode = nodes[actualNode.right]
			}
			solutionPartOne++
			if actualNode.name == "ZZZ" {
				break
			}
		}
	}

	fmt.Printf("Solution of day 8 part 1 is: %d\n", solutionPartOne)
}
