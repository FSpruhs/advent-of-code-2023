package day1

import (
	"fmt"
	"github.com/FSpruhs/advent-of-code-2023/util"
)

func Solve(filePath string) {

	x := util.ReadFile(filePath)
	for _, s := range x {
		fmt.Println(s)
	}

}
