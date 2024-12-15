package main

import (
	"fmt"
	"os"

	"github.com/Slijkhuis/2024-advent-of-code/pkg/aoc"
)

/*
Thoughts: My solution is very slow. I think it may be more efficient to build a cache of all possible paths from every 8
to every 9, from every 7 to every 8, etcetera.
*/

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
	grid := aoc.BuildGridFromFile(os.Args[2])
	var trailheads []aoc.Point

	for cell := range grid.Iter() {
		if cell.Value == '0' {
			trailheads = append(trailheads, cell.Point)
		}
	}

	g := aoc.NewGraph[aoc.Point, int]()
	for _, trailhead := range trailheads {
		n := aoc.NewNode(trailhead, 0)
		g.AddNode(n)
		addAdjecentNodes(g, grid, n)
	}

	var trailheadNodes []*aoc.Node[aoc.Point, int]
	var endNodes []*aoc.Node[aoc.Point, int]
	for _, n := range g.Nodes {
		if n.Value == 0 {
			trailheadNodes = append(trailheadNodes, n)
		} else if n.Value == 9 {
			endNodes = append(endNodes, n)
		}
	}

	scoreSum := 0
	for i, n := range trailheadNodes {
		fmt.Println("checking trailhead", i+1, "out of", len(trailheadNodes))

		score := 0
		for _, end := range endNodes {
			if g.AreConnected(n, end) {
				score++
			}
		}

		scoreSum += score
	}

	fmt.Println(scoreSum)
}

func part2() {
	grid := aoc.BuildGridFromFile(os.Args[2])
	var trailheads []aoc.Point

	for cell := range grid.Iter() {
		if cell.Value == '0' {
			trailheads = append(trailheads, cell.Point)
		}
	}

	g := aoc.NewGraph[aoc.Point, int]()
	for _, trailhead := range trailheads {
		n := aoc.NewNode(trailhead, 0)
		g.AddNode(n)
		addAdjecentNodes(g, grid, n)
	}

	var trailheadNodes []*aoc.Node[aoc.Point, int]
	var endNodes []*aoc.Node[aoc.Point, int]
	for _, n := range g.Nodes {
		if n.Value == 0 {
			trailheadNodes = append(trailheadNodes, n)
		} else if n.Value == 9 {
			endNodes = append(endNodes, n)
		}
	}

	ratingSum := 0
	for i, n := range trailheadNodes {
		fmt.Println("checking trailhead", i+1, "out of", len(trailheadNodes))

		rating := 0
		for _, end := range endNodes {
			paths := g.FindAllPaths(n, end)
			rating += len(paths)
		}

		ratingSum += rating
	}

	fmt.Println(ratingSum)
}

func addAdjecentNodes(g *aoc.Graph[aoc.Point, int], grid *aoc.Grid, n *aoc.Node[aoc.Point, int]) {
	for _, dir := range aoc.NoDiagonals {
		p := n.Key.Move(dir)
		r, ok := grid.Data[p]
		if !ok {
			continue
		}
		v2 := int(aoc.Atoi(string(r)))

		if v2 == n.Value+1 {
			n2 := g.NewOrExistingNode(p, v2)
			g.AddNode(n2)
			g.AddEdge(n, n2)

			addAdjecentNodes(g, grid, n2)
		}
	}
}
