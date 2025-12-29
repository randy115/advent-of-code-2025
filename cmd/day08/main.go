package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"sort"
	"strconv"

	"github.com/randy115/advent-of-code-2025/internal/input"
)

type Point struct {
	x, y, z float64
}

type Edge struct {
	a, b int
	dist float64
}

func euclideanDistance(x1, y1, z1, x2, y2, z2 float64) float64 {
	dx := x2 - x1
	dy := y2 - y1
	dz := z2 - z1
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func find(dsu []int, x int) int {
	if dsu[x] != x {
		dsu[x] = find(dsu, dsu[x])
	}
	return dsu[x]
}

func union(dsu []int, a, b int) {
	roota := find(dsu, a)
	rootb := find(dsu, b)

	if roota != rootb {
		dsu[rootb] = roota
	}
}

func connectJunctionBoxes(boxes []string) {
	re := regexp.MustCompile(`(\d+),(\d+),(\d+)`)
	var points []Point
	var edges []Edge

	for _, box := range boxes {
		matches := re.FindStringSubmatch(box)
		x, _ := strconv.ParseFloat(matches[1], 64)
		y, _ := strconv.ParseFloat(matches[2], 64)
		z, _ := strconv.ParseFloat(matches[3], 64)
		points = append(points, Point{x, y, z})
	}
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			a := points[i]
			b := points[j]
			distance := euclideanDistance(a.x, a.y, a.z, b.x, b.y, b.z)
			edges = append(edges, Edge{i, j, distance})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dist < edges[j].dist
	})
	dsu := make([]int, len(points))
	for i := range dsu {
		dsu[i] = i
	}

	for i := 0; i < 10; i++ {
		union(dsu, edges[i].a, edges[i].b)
	}
	counts := make(map[int]int)

	for i := 0; i < len(points); i++ {
		root := find(dsu, i)
		counts[root]++
	}
	for i := range dsu {
		fmt.Println(i, find(dsu, i))
	}
}

func main() {
	file, err := input.ReadFile("day08e.txt")
	if err != nil {
		log.Fatal(err)
	}
	connectJunctionBoxes(file)
}
