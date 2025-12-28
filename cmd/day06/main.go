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

func computeRightMostMathProblems(problems []string) {
	row := len(problems)
	col := len(problems[0]) - 1
	var total int64
	var stack []int64
	var symbol byte
	for c := col; c >= 0; c-- {
		var builder strings.Builder
		for r := 0; r < row; r++ {
			if problems[r][c] != '+' && problems[r][c] != '*' {
				builder.WriteByte(problems[r][c])
			} else {
				symbol = problems[r][c]
			}
		}
		numString := strings.TrimSpace(builder.String())

		if numString != "" {
			fmt.Printf("appending string %s\n", numString)
			num, _ := strconv.ParseInt(numString, 10, 64)
			stack = append(stack, num)
			fmt.Println(stack)
		} else if len(stack) > 0 {
			var result int64
			if symbol == '+' {
				result = 0
				for _, num := range stack {
					result += num
				}
			}

			if symbol == '*' {
				result = 1
				for _, num := range stack {
					result *= num
				}
			}
			stack = stack[:0]
			total += result
		}
	}
	if len(stack) > 0 {
		var result int64
		if symbol == '+' {
			result = 0
			for _, num := range stack {
				result += num
			}
		}

		if symbol == '*' {
			result = 1
			for _, num := range stack {
				result *= num
			}
		}
		total += result
	}
	fmt.Println(total)
}

func main() {
	file, err := input.ReadFile("day06.txt")
	if err != nil {
		log.Fatal(err)
	}
	computeRightMostMathProblems(file)
}
