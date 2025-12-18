package main

import (
	"fmt"
	"log"

	"github.com/randy115/advent-of-code-2025/internal/input"
)

func main() {
	file, err := input.ReadFile("day04e.txt")
	if err != nil {
		log.Fatal(err)
	}
	for _, value := range file {
		for _, char := range value {
			fmt.Print(char)
		}
		fmt.Printf("\n")
	}
}
