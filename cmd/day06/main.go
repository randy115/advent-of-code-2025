package main

import (
	"fmt"
	"log"
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

	for c := 0; c < col; c++ {
		for r := row - 2; r >= 0; r-- {
			fmt.Printf("matrix[%d][%d] = %s\n", r, c, problems[r][c])
		}
	}
}

func main() {
	file, err := input.ReadFile("day06e.txt")
	if err != nil {
		log.Fatal(err)
	}
	computeMathProblems(file)
}
