package main

import (
	"fmt"
	"log"
	"regexp"
	"slices"
	"strconv"

	"github.com/randy115/advent-of-code-2025/internal/input"
)

func validateId(file []string) {
	re := regexp.MustCompile("([0-9]+)-([0-9]+)")
	var idRange [][]int64
	var ids []int64
	var validId int
	for _, ingredientIds := range file {
		match := re.FindStringSubmatch(ingredientIds)
		if match != nil {
			start, _ := strconv.ParseInt(match[1], 10, 64)
			end, _ := strconv.ParseInt(match[2], 10, 64)
			idRange = append(idRange, []int64{start, end})
		} else if ingredientIds != "" {
			id, _ := strconv.ParseInt(ingredientIds, 10, 64)
			ids = append(ids, id)
		}
	}

	slices.SortFunc(idRange, func(a, b []int64) int {
		if a[0] < b[0] {
			return -1
		} else if a[0] > b[0] {
			return 1
		}
		return 0
	})

	for _, id := range ids {
		for _, ingredientId := range idRange {
			if id >= ingredientId[0] && id <= ingredientId[1] {
				validId += 1
				fmt.Printf("%d IS IN RANGES %d - %d\n", id, ingredientId[0], ingredientId[1])
				break
			}
		}
	}
	fmt.Println(validId)
}

func main() {
	file, err := input.ReadFile("day05.txt")
	if err != nil {
		log.Fatal(err)
	}
	validateId(file)
}
