package main

import (
	"flag"
	"github.com/FSpruhs/advent-of-code-2023/day1"
	"github.com/FSpruhs/advent-of-code-2023/day10"
	"github.com/FSpruhs/advent-of-code-2023/day2"
	"github.com/FSpruhs/advent-of-code-2023/day3"
	"github.com/FSpruhs/advent-of-code-2023/day4"
	"github.com/FSpruhs/advent-of-code-2023/day5"
	"github.com/FSpruhs/advent-of-code-2023/day6"
	"github.com/FSpruhs/advent-of-code-2023/day7"
	"github.com/FSpruhs/advent-of-code-2023/day8"
	"github.com/FSpruhs/advent-of-code-2023/day9"
	"github.com/FSpruhs/advent-of-code-2023/util"
	"log"
	"os"
)

var (
	flagDay     = flag.Int("d", 0, "sets the day to solve")
	flagExample = flag.Bool("e", false, "use example data when set")
)

func main() {
	flag.Parse()

	filePath := util.BuildFilePath(*flagDay, *flagExample)

	switch *flagDay {
	case 1:
		day1.Solve(filePath)
	case 2:
		day2.Solve(filePath)
	case 3:
		day3.Solve(filePath)
	case 4:
		day4.Solve(filePath)
	case 5:
		day5.Solve(filePath)
	case 6:
		day6.Solve(filePath)
	case 7:
		day7.Solve(filePath)
	case 8:
		day8.Solve(filePath)
	case 9:
		day9.Solve(filePath)
	case 10:
		day10.Solve(filePath)

	default:
		log.Printf("can not find puzzle for day %d", *flagDay)
		os.Exit(1)
	}

}
