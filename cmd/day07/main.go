package main

import (
	"fmt"
	"log"

	"github.com/randy115/advent-of-code-2025/internal/input"
)

func dfs(row, col int, grid [][]byte, count *int) {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
		*count += 1
		return
	}

	if grid[row][col] == '.' {
		grid[row][col] = '|'
		for _, val := range grid {
			fmt.Println(string(val))
		}
		dfs(row+1, col, grid, count)
	}

	if grid[row][col] == '^' {
		dfs(row, col-1, grid, count)
		dfs(row, col+1, grid, count)
	}
}

func dfsTimeline(row, col int, grid [][]byte, memo [][]int) int {
	if row < 0 || col < 0 || col >= len(grid[0]) {
		return 0
	}

	if row >= len(grid) {
		fmt.Printf("path found\n")
		return 1
	}

	if memo[row][col] != 0 {
		return memo[row][col]
	}
	res := 0
	if grid[row][col] == '.' {
		grid[row][col] = '|'
		for _, val := range grid {
			fmt.Println(string(val))
		}
		fmt.Printf("adding beam\n")
		res = dfsTimeline(row+1, col, grid, memo)
		fmt.Printf("removing beam\n")
		grid[row][col] = '.'
	}

	if grid[row][col] == '^' {
		res = dfsTimeline(row, col-1, grid, memo) + dfsTimeline(row, col+1, grid, memo)
	}

	memo[row][col] = res
	return res
}

func countBeams(beams []string) {
	grid := make([][]byte, len(beams))
	memo := make([][]int, len(grid))
	for i := range beams {
		grid[i] = []byte(beams[i])
	}
	for i := range memo {
		memo[i] = make([]int, len(grid[0]))
	}

	for i, row := range grid {
		for j := range row {
			if grid[i][j] == 'S' {
				num := dfsTimeline(i+1, j, grid, memo)
				fmt.Println(num)
				break
			}
		}
	}
}

func main() {
	file, err := input.ReadFile("day07.txt")
	if err != nil {
		log.Fatal(err)
	}
	countBeams(file)
}
