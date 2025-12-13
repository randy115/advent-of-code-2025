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
			currentId := strconv.Itoa(i)
			fmt.Printf("Inside the range - current id: %s\n", currentId)
			for j := 0; j < len(currentId); j++ {
				currSubId := currentId[:j+1]
				fmt.Printf("Sub Id: %s\n", currSubId)
			}
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
