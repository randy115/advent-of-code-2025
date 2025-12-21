package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/randy115/advent-of-code-2025/internal/input"
)

func computeMathProblems(file []string) {
	var problems [][]string
	for _, num := range file {
		problems = append(problems, strings.Fields(num))
	}
	row := len(problems)
	col := len(problems[0])
	var total int
	for c := 0; c < col; c++ {
		curr := 0
		currSym := problems[row-1][c]
		for r := row - 2; r >= 0; r-- {
			num, _ := strconv.Atoi(problems[r][c])
			if currSym == "+" {
				curr += num
			} else if curr == 0 {
				curr = 1
				curr *= num
			} else {
				curr *= num
			}
			fmt.Printf("matrix[%d][%d] = %s\n", r, c, problems[r][c])
		}
		total += curr
	}
	fmt.Println(total)
}

func computeRightMostMathProblems(file []string) {
	var problems [][]string
	for _, num := range file {
		problems = append(problems, strings.Fields(num))
	}
	// row := len(problems)
	col := len(problems[0])
	// var total int
	for c := 0; c < col; c++ {
		// fmt.Printf("matrix[%d][%d] = %s\n", r, c, problems[r][c])
	}
	for _, problem := range problems {
		fmt.Println(problem)
	}
}

func main() {
	file, err := input.ReadFile("day06e.txt")
	if err != nil {
		log.Fatal(err)
	}
	computeRightMostMathProblems(file)
}
