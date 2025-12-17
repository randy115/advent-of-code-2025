package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/randy115/advent-of-code-2025/internal/input"
)

func maxTwoBatteries(banks []string) {
	var totalJolt int
	for _, jolt := range banks {
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

func maxTwelveBatteries(banks []string) {
	stack := make([]int, 0, 12)
	totalJolt := 0
	for _, jolt := range banks {
		remove := len(jolt) - 12
		fmt.Printf("CURRENT BATTERY BANK: %s\n", jolt)
		for i := 0; i < len(jolt); i++ {
			fmt.Println("-------------------------------------------------")
			battery, _ := strconv.Atoi(string(jolt[i]))
			fmt.Printf("remove count %d\n", remove)
			fmt.Printf("Current Stack: %d\n", stack)
			fmt.Printf("Current Battery: %d\n", battery)
			for len(stack) > 0 && stack[len(stack)-1] < battery && remove > 0 {
				fmt.Printf("Battery %d < Top Stack Battery %d\n", stack[len(stack)-1], battery)
				stack = stack[:len(stack)-1]
				remove -= 1
			}
			fmt.Printf("Pushing %d\n", battery)
			stack = append(stack, battery)
		}
		fmt.Println("-------------------------------------------------")
		result := 0
		for _, num := range stack[:12] {
			result = result*10 + num
		}
		totalJolt += result
		stack = stack[:0]
		fmt.Println(result)
	}
	fmt.Println(totalJolt)
}

func main() {
	file, err := input.ReadFile("day03.txt")
	if err != nil {
		log.Fatal(err)
	}
	maxTwelveBatteries(file)
}
