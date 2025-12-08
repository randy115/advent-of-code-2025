package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/randy115/advent-of-code-2025/internal/input"
)

func calulateRotations(start *int, dir string, num int) int {
	*start++
	return *start
}

func partOne(combination []string) int {
	start := 50
	for _, combination := range combination {
		direction := string(combination[0])
		num, _ := strconv.Atoi(combination[1:])
		fmt.Println(calulateRotations(&start, direction, num))
	}
	return 1
}

func main() {
	file, err := input.ReadFile("day01.txt")
	if err != nil {
		log.Fatal(err)
	}
	partOne(file)
}
