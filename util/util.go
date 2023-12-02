package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func BuildFilePath(day int, example bool) string {
	if example {
		return fmt.Sprintf("./day%d/example.txt", day)
	}
	return fmt.Sprintf("./day%d/puzzle.txt", day)
}

func ReadFile(path string) *[]string {

	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("error %s while reading file %s", err, path)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		zeile := scanner.Text()
		lines = append(lines, zeile)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("error %s while scanning file %s", err, path)
		os.Exit(1)
	}

	return &lines
}

func ConvertToInt(str string) int {
	number, err := strconv.Atoi(str)
	if err != nil {
		log.Printf("could not transform %s to an int", str)
		os.Exit(1)
	}
	return number
}
