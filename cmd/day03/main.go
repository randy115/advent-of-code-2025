package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/randy115/advent-of-code-2025/internal/input"
)

func maxJoltage(jolts []string) {
	var totalJolt int
	for _, jolt := range jolts {
		var bankMax int
		maxRight, _ := strconv.Atoi(string(jolt[len(jolt)-1]))
		for i := len(jolt) - 2; i >= 0; i-- {
			leftJoltString := string(jolt[i])
			rightJoltString := strconv.Itoa(maxRight)
			currMaxJolt, _ := strconv.Atoi(leftJoltString + rightJoltString)
			bankMax = max(bankMax, currMaxJolt)
			leftJoltInt, _ := strconv.Atoi(leftJoltString)
			maxRight = max(leftJoltInt, maxRight)
		}
		totalJolt += bankMax
	}
	fmt.Println(totalJolt)
}

func main() {
	file, err := input.ReadFile("day03.txt")
	if err != nil {
		log.Fatal(err)
	}
	maxJoltage(file)
}
