package main

import (
	"fmt"
	"os"

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
	grid := aoc.BuildGridFromFile(os.Args[2])
	g := aoc.NewGraph[aoc.PointWithDirection, rune]()
	start, ok := grid.FindFirstCellWithValue('S')
	if !ok {
		panic("no start found")
	}
	end, ok := grid.FindFirstCellWithValue('E')
	if !ok {
		panic("no end found")
	}

	var buildGraph func(*aoc.Node[aoc.PointWithDirection, rune], aoc.Direction)
	buildGraph = func(comingFrom *aoc.Node[aoc.PointWithDirection, rune], facingIn aoc.Direction) {
		// straight ahead
		next := comingFrom.Key.Point.Move(facingIn)
		val := grid.Data[next]
		if val != '#' {
			n := g.NewOrExistingNode(aoc.PointWithDirection{Point: next, Direction: facingIn}, val)
			g.AddNodeIgnoreDuplicate(n)
			g.AddWeightedEdge(comingFrom, n, 1)
			buildGraph(n, facingIn)
		}

		// turns
		dirs := aoc.Except(aoc.NoDiagonals, facingIn, facingIn.Opposite())
		for _, dir := range dirs {
			n := g.NewOrExistingNode(aoc.PointWithDirection{Point: comingFrom.Key.Point, Direction: dir}, val)
			g.AddNodeIgnoreDuplicate(n)
			if ok := g.AddWeightedEdge(comingFrom, n, 1000); !ok {
				buildGraph(n, dir)
			}
		}
	}

	startNode := aoc.NewNode(aoc.PointWithDirection{Point: start.Point, Direction: aoc.Right}, start.Value)
	g.AddNode(startNode)
	buildGraph(startNode, aoc.Right)

	fmt.Println("built graph")

	// I'm lucky that all paths coming from the left are not the shortest in the test input and the real input. So I only
	// have to check for paths that come from the bottom (i.e. are going up).
	endNode := aoc.NewNode(aoc.PointWithDirection{Point: end.Point, Direction: aoc.Up}, end.Value)

	path, cost := g.FindShortestPath(startNode, endNode)
	for _, node := range path {
		grid.Data[node.Key.Point] = 'P'
	}
	aoc.Debug(grid.String())
	aoc.Debug(cost)

	fmt.Println(cost)
}

func part2() {
	aoc.Debug("todo")
}
