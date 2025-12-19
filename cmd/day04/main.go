package main

import (
	"fmt"
	"log"

	"github.com/randy115/advent-of-code-2025/internal/input"
)

func allowedPaperAccess(matrix [][]byte) {
	dirs := [][]int{
		{1, 0},
		{1, -1},
		{1, 1},
		{0, -1},
		{0, 1},
		{-1, 0},
		{-1, -1},
		{-1, 1},
	}

	var access int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '@' {
				fmt.Println("---------------------------------")
				rolls := 0
				// if i == 0 {
				// 	fmt.Printf("current row %d\n", i)
				// 	fmt.Printf("current col %d\n", j)
				// 	fmt.Printf("current val %s\n", string(matrix[i][j]))
				// }
				for _, d := range dirs {
					r := i + d[0]
					c := j + d[1]
					if r >= 0 && c >= 0 && r < len(matrix) && c < len(matrix[0]) {
						// if i == 0 {
						// 	fmt.Println("-----------------------")
						// 	fmt.Printf("new row %d\n", r)
						// 	fmt.Printf("new col %d\n", c)
						// 	fmt.Printf("new val %s\n", string(matrix[r][c]))
						// 	fmt.Println("-----------------------")
						//
						// }
						if matrix[r][c] == '@' || matrix[r][c] == 'x' {
							fmt.Println("FOUND ROLL")
							rolls += 1
						}
					}
				}
				if rolls < 4 {
					access += 1
					matrix[i][j] = 'x'
				}
				fmt.Println("FINISHED")
				fmt.Println("---------------------------------")

			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 'x' {
				matrix[i][j] = '.'
			}
		}
	}
	fmt.Println(access)
	for _, val := range matrix {
		fmt.Println(string(val))
	}
}

func main() {
	file, err := input.Create2DMatrix("day04e.txt")
	if err != nil {
		log.Fatal(err)
	}

	allowedPaperAccess(file)
}
