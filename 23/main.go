package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/Slijkhuis/2024-advent-of-code/pkg/aoc"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <part> <input-file>")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "1":
		part1()
	case "2":
		part2()
	default:
		fmt.Println("Invalid part")
		os.Exit(1)
	}
}

func part1() {
	t := time.Now()
	g := aoc.NewGraph[string, struct{}]()

	for line := range aoc.LinesFromFile(os.Args[2]) {
		computers := strings.Split(line, "-")
		n1 := g.NewOrExistingNode(computers[0], struct{}{})
		n2 := g.NewOrExistingNode(computers[1], struct{}{})
		g.AddEdge(n1, n2)
		g.AddEdge(n2, n1)
	}

	aoc.Println(t, "built graph")

	groups := make(map[string]struct{})
	for key, node := range g.Nodes {
		if key[0] != 't' {
			continue
		}

		neighs := g.Neighbors(node)
		for _, neigh := range neighs {
			neights2 := g.Neighbors(neigh)
			for _, neigh2 := range neights2 {
				if aoc.In(neighs, neigh2) { // node, neigh, and neigh2 are all connected
					groupKey := []string{key, neigh.Key, neigh2.Key}
					sort.Strings(groupKey)
					groups[strings.Join(groupKey, "-")] = struct{}{}
				}
			}
		}
	}

	aoc.Debug(groups)

	aoc.Println(t, len(groups))
}

func part2() {
	for line := range aoc.LinesFromFile(os.Args[2]) {
		fmt.Println(line)
	}
}
