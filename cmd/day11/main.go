package main

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/randy115/advent-of-code-2025/internal/input"
)

type (
	Graph   map[string][]string
	Visited map[Node]int
)

type Node struct {
	Server  string
	seenDac bool
	seenFft bool
}

func containsOut(source string, devices Graph) bool {
	g := devices[source]
	if slices.Contains(g, "out") {
		return true
	} else {
		return false
	}
}

func containsPath(path []string) bool {
	if slices.Contains(path, "dac") && slices.Contains(path, "fft") && slices.Contains(path, "out") {
		return true
	} else {
		return false
	}
}

func createMap(devices []string) Graph {
	graph := make(Graph)
	for _, d := range devices {
		device := strings.Fields(d)
		source := strings.Trim(device[0], ":")
		var outputs []string
		for _, output := range device[1:] {
			outputs = append(outputs, output)
		}
		graph[source] = outputs
	}
	return graph
}

func dfs(source string, graph Graph) int {
	if containsOut(source, graph) {
		return 1
	}

	var res int
	outputs := graph[source]
	for _, output := range outputs {
		res += dfs(output, graph)
	}

	return res
}

func dfsDacAndFft(source string, graph Graph, visited Visited, dac, fft bool) int {
	// fmt.Printf("current path %s\n", paths)
	// if _, ok := visited[source]; ok {
	// 	return visited[source]
	// }

	key := Node{source, dac, fft}
	if val, ok := visited[key]; ok {
		return val
	}

	seenDac, seenFft := dac, fft

	if source == "dac" {
		seenDac = true
	} else if source == "fft" {
		seenFft = true
	}

	if seenDac && seenFft && source == "out" {
		return 1
	}

	// if containsPath(paths) {
	// 	fmt.Printf("current path valid %s \n", paths)
	// 	return 1
	// }

	var res int
	outputs := graph[source]
	for _, output := range outputs {
		res += dfsDacAndFft(output, graph, visited, seenDac, seenFft)
	}
	visited[key] = res
	return res
}

func main() {
	file, err := input.ReadFile("day11.txt")
	if err != nil {
		log.Fatal(err)
	}
	visited := make(Visited)
	g := createMap(file)
	totalPaths := dfsDacAndFft("svr", g, visited, false, false)
	fmt.Println(totalPaths)
}
