package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/randy115/advent-of-code-2025/internal/input"
)

func checkForInvalidIds(file []string) {
	re := regexp.MustCompile("([0-9]+)-([0-9]+)")
	for id := range strings.SplitSeq(file[0], ",") {
		matches := re.FindStringSubmatch(id)
		start, _ := strconv.Atoi(matches[1])
		end, _ := strconv.Atoi(matches[2])
		fmt.Printf("Start range: %d    End range: %d\n", start, end)
		for i := start; i <= end; i++ {
			currentId := string(i)
			for j := 0; j <= len(currentId); j++ {
			}
			fmt.Printf("Inside the range - current id: %d\n", i)
		}
	}
}

func main() {
	file, err := input.ReadFile("day02e.txt")
	if err != nil {
		log.Fatal(err)
	}
	checkForInvalidIds(file)
}
