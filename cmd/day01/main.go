package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/randy115/advent-of-code-2025/internal/input"
)

func calulateRotations(start *int, dir string, num int) (beg, result int) {
	begg := *start
	if dir == "L" {
		*start = *start - num
		if *start < 0 {
			*start = *start + 100
		}
	}
	if dir == "R" {
		*start = *start + num
		if *start > 99 {
			*start = *start - 100
		}
	}
	end := *start
	return begg, end
}

func partOne(combination []string) {
	start := 50
	count := 0
	for i, combination := range combination {
		direction := string(combination[0])
		num, _ := strconv.Atoi(combination[1:])
		begg, end := calulateRotations(&start, direction, num)
		if i < 200 {
			fmt.Printf("start %d | direction %s | turn %d | end %d\n", begg, direction, num, end)
		}
		if start == 0 {
			count++
		}
	}
	fmt.Println(count)
}

func main() {
	file, err := input.ReadFile("day01.txt")
	if err != nil {
		log.Fatal(err)
	}
	partOne(file)
}
