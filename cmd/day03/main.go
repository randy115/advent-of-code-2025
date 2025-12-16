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
	for _, jolt := range banks {
		remove := len(jolt) - 12
		fmt.Printf("CURRENT BATTERY BANK: %s\n", jolt)
		for i := 0; i < len(jolt); i++ {
			fmt.Println("-------------------------------------------------")
			battery, _ := strconv.Atoi(string(jolt[i]))
			fmt.Printf("remove count %d\n", remove)
			fmt.Printf("Current Stack: %d\nCurrent Battery: %d\n", stack, battery)
			fmt.Println("Current length:", len(stack))
			if len(stack) == 0 {
				fmt.Printf("Stack is empty appending battery %d\n", battery)
				stack = append(stack, battery)
			} else if stack[len(stack)-1] < battery && remove > 0 {
				remainingString := len(jolt) - i
				for len(stack) > 0 && stack[len(stack)-1] < battery && remove > 0 && remainingString > 0 {
					fmt.Printf("Battery %d < Top Stack Battery %d\nPushing %d\n", stack[len(stack)-1], battery, battery)
					stack = stack[:len(stack)-1]
					remove -= 1
					remainingString -= 1
				}
				stack = append(stack, battery)
			} else if stack[len(stack)-1] < battery && remove == 0 && len(stack) < 12 {
				stack = append(stack, battery)
			} else if stack[len(stack)-1] >= battery && len(stack) < 12 {
				fmt.Printf("Battery %d > Top Stack Battery %d\nPushing %d\n", stack[len(stack)-1], battery, battery)
				stack = append(stack, battery)
			} else if len(stack) == 12 {
				fmt.Println("last else statement")
				stack[len(stack)-1] = battery
			}
		}
		fmt.Println(len(stack))
		fmt.Println(stack)
		fmt.Println("len:", len(stack))
		fmt.Println("cap:", cap(stack))
		stack = stack[:0]
		fmt.Println("-------------------------------------------------")
	}
}

func main() {
	file, err := input.ReadFile("day03e.txt")
	if err != nil {
		log.Fatal(err)
	}
	maxTwelveBatteries(file)
}
