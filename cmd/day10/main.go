package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/randy115/advent-of-code-2025/internal/input"
)

func configureLights(schemes []string) {
	for _, scheme := range schemes {
		line := strings.Fields(scheme)
		light := []byte(strings.Trim(line[0], "[]"))
		fmt.Printf("%s\n", light)
		var buttons [][]int
		for i := 1; i < len(light); i++ {
			button := []int{}
			if strings.HasPrefix(line[i], "(") {
				nums := strings.Trim(line[i], "()")
				for num := range strings.SplitSeq(nums, ",") {
					res, _ := strconv.Atoi(num)
					button = append(button, res)
				}
			}
			buttons = append(buttons, button)
		}
		fmt.Printf("buttons %v\n", buttons)

	}
}

func main() {
	file, err := input.ReadFile("day10e.txt")
	if err != nil {
		log.Fatal(err)
	}
	configureLights(file)
}
