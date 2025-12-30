package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/randy115/advent-of-code-2025/internal/input"
)

func minMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func getLargestTile(tiles []string) {
	length := len(tiles)
	re := regexp.MustCompile(`(\d+),(\d+)`)
	largestArea := 0
	for i := 0; i < length; i++ {
		t1 := re.FindStringSubmatch(tiles[i])
		x1, _ := strconv.Atoi(t1[2])
		y1, _ := strconv.Atoi(t1[1])
		for j := i + 1; j < length; j++ {
			fmt.Println("-----------------------------------")
			t2 := re.FindStringSubmatch(tiles[j])
			x2, _ := strconv.Atoi(t2[2])
			y2, _ := strconv.Atoi(t2[1])
			fmt.Printf("first tile (%d,%d)\n", y1, x1)
			fmt.Printf("second tile (%d,%d)\n", y2, x2)
			if x1 == x2 && y1 == y2 {
				continue
			}
			xmin, xmax := minMax(x1, x2)
			ymin, ymax := minMax(y1, y2)
			row := (ymax - ymin) + 1
			col := (xmax - xmin) + 1
			fmt.Printf("row : %d - %d = %d\n", ymax, ymin, row)
			fmt.Printf("col : %d - %d = %d\n", xmax, xmin, col)
			fmt.Printf("area : %d\n", row*col)
			largestArea = max(largestArea, row*col)
			fmt.Println("-----------------------------------")
		}
	}
	fmt.Println(largestArea)
}

func main() {
	file, err := input.ReadFile("day09.txt")
	if err != nil {
		log.Fatal(err)
	}
	getLargestTile(file)
}
