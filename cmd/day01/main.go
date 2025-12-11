package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/randy115/advent-of-code-2025/internal/input"
)

func calulateRotations(start *int, dir string, num int) (beg, result, quotient int) {
	begg := *start
	var q int
	if dir == "L" {
		q = num / 100
		if *start > 0 && num%100 >= *start {
			q += 1
		}
		*start = ((*start-num)%100 + 100) % 100
	}
	if dir == "R" {
		q = (*start + num) / 100
		*start = (*start + num) % 100
	}

	end := *start
	return begg, end, q
}

func rotate(combination []string) {
	start := 50
	partOne := 0
	partTwo := 0
	for i, combination := range combination {
		direction := string(combination[0])
		num, _ := strconv.Atoi(combination[1:])
		begg, end, q := calulateRotations(&start, direction, num)
		if i < 20 {
			fmt.Printf("start %d | direction %s | turn %d | end %d\n", begg, direction, num, end)
			fmt.Println(q)
		}
		if start == 0 {
			partOne++
		}
		partTwo += q
	}
	fmt.Printf("Part One Solution: %d | Part Two Solution: %d\n", partOne, partTwo)
}

func main() {
	file, err := input.ReadFile("day01.txt")
	if err != nil {
		log.Fatal(err)
	}
	rotate(file)
}
