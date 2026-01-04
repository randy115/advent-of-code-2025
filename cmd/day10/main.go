package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/randy115/advent-of-code-2025/internal/input"
)

func parseSchematics(schemes []string) (string, [][]int, []int) {
	light := strings.Trim(schemes[0], "[]")
	var buttons [][]int
	var joltageLevels []int
	for i := 1; i < len(schemes); i++ {
		var button []int
		if strings.HasPrefix(schemes[i], "(") {
			nums := strings.Trim(schemes[i], "()")
			for num := range strings.SplitSeq(nums, ",") {
				res, _ := strconv.Atoi(num)
				button = append(button, res)
			}
		} else if strings.HasPrefix(schemes[i], "[") {
			jolts := strings.Trim(schemes[i], "[]")
			for jolt := range strings.SplitSeq(jolts, ",") {
				res, _ := strconv.Atoi(jolt)
				joltageLevels = append(joltageLevels, res)
			}
		}

		if button != nil {
			buttons = append(buttons, button)
		}
	}
	return light, buttons, joltageLevels
}

func bfs(desiredState string, buttons [][]int) int {
	visited := make(map[string]struct{})
	initial := strings.Repeat(".", len(desiredState))
	queue := []string{initial}
	presses := 0
	// fmt.Printf("DESIRED STATE %s\n", desiredState)
	for len(queue) > 0 {
		nextLevel := []string{}
		for _, pattern := range queue {
			// fmt.Printf("CURRENT PATTERN %s\n", pattern)
			// fmt.Printf("START -------------------------------------------\n\n")

			if desiredState == pattern {
				fmt.Printf("desired light pattern found %s \n", desiredState)
				return presses
			}

			if _, ok := visited[pattern]; !ok {
				// fmt.Printf("light pattern not found: %s adding to map\n", pattern)
				visited[pattern] = struct{}{}
				for _, row := range buttons {
					light := []byte(pattern)
					// fmt.Printf("current button set %d\n", row)
					for _, col := range row {
						if light[col] == '.' {
							light[col] = '#'
						} else {
							light[col] = '.'
						}
					}
					// fmt.Printf("new light pattern %s\n", light)
					nextLevel = append(nextLevel, string(light))
				}
			}
			// fmt.Printf("END -------------------------------------------\n\n")
		}
		// fmt.Printf("current queue %v\n", nextLevel)
		// fmt.Printf("current map %v\n", visited)
		queue = nextLevel
		presses += 1
	}
	return 0
}

func bfsJoltage(desiredState string, buttons [][]int, joltageLevels []int) int {
	visited := make(map[string]struct{})
	initial := make([]int, len(joltageLevels))
	queue := [][]int{initial}
	presses := 0
	// fmt.Printf("DESIRED STATE %s\n", desiredState)
	for len(queue) > 0 {
		nextLevel := []string{}
		for _, pattern := range queue {
			// fmt.Printf("CURRENT PATTERN %s\n", pattern)
			// fmt.Printf("START -------------------------------------------\n\n")

			if desiredState == pattern {
				fmt.Printf("desired light pattern found %s \n", desiredState)
				return presses
			}

			if _, ok := visited[pattern]; !ok {
				// fmt.Printf("light pattern not found: %s adding to map\n", pattern)
				visited[pattern] = struct{}{}
				for _, row := range buttons {
					light := []byte(pattern)
					// fmt.Printf("current button set %d\n", row)
					for _, col := range row {
						if light[col] == '.' {
							light[col] = '#'
						} else {
							light[col] = '.'
						}
					}
					// fmt.Printf("new light pattern %s\n", light)
					nextLevel = append(nextLevel, string(light))
				}
			}
			// fmt.Printf("END -------------------------------------------\n\n")
		}
		// fmt.Printf("current queue %v\n", nextLevel)
		// fmt.Printf("current map %v\n", visited)
		queue = nextLevel
		presses += 1
	}
	return 0
}

func configureLights(schemes []string) {
	total := 0
	for _, scheme := range schemes {
		line := strings.Fields(scheme)
		desiredState, buttons, _ := parseSchematics(line)
		total += bfs(desiredState, buttons)
	}
	fmt.Println(total)
}

func main() {
	file, err := input.ReadFile("day10.txt")
	if err != nil {
		log.Fatal(err)
	}
	configureLights(file)
}
