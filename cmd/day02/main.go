package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/randy115/advent-of-code-2025/internal/input"
)

func repeatedIdsAtLeastTwice(file []string) {
	re := regexp.MustCompile("([0-9]+)-([0-9]+)")
	for id := range strings.SplitSeq(file[0], ",") {
		matches := re.FindStringSubmatch(id)
		start, _ := strconv.Atoi(matches[1])
		end, _ := strconv.Atoi(matches[2])
		fmt.Printf("Start range: %d    End range: %d\n", start, end)
		for i := start; i <= end; i++ {
			currentId := strconv.Itoa(i)
			// fmt.Printf("Inside the range - current id: %s\n", currentId)
			for j := 0; j < len(currentId)-1; j++ {
				currSubId := fmt.Sprintf(`^(%s)+$`, currentId[:j+1])
				invalidId, _ := regexp.MatchString(currSubId, currentId)
				if invalidId {
					fmt.Printf("INVALID ID ----> %s\n", currentId)
				}
			}
		}
	}
}

func repeatedIdsTwice(file []string) {
	re := regexp.MustCompile("([0-9]+)-([0-9]+)")
	var sum int
	for id := range strings.SplitSeq(file[0], ",") {
		matches := re.FindStringSubmatch(id)
		start, _ := strconv.Atoi(matches[1])
		end, _ := strconv.Atoi(matches[2])
		fmt.Printf("Start range: %d    End range: %d\n", start, end)
		for i := start; i <= end; i++ {
			currentId := strconv.Itoa(i)
			half := len(currentId) / 2
			if len(currentId)%2 == 0 && currentId[:half] == currentId[half:] {
				sum += i
				fmt.Printf("INVALID ID: %s\n", currentId)
			}
		}
	}
	fmt.Println(sum)
}

func main() {
	file, err := input.ReadFile("day02.txt")
	if err != nil {
		log.Fatal(err)
	}
	repeatedIdsTwice(file)
}
