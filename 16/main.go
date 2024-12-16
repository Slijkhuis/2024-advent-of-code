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
	fmt.Println("built graph")

	startNode := aoc.NewNode(aoc.PointWithDirection{Point: start.Point, Direction: aoc.Right}, start.Value)
	g.AddNode(startNode)
	buildGraph(startNode, aoc.Right)

	endNode1 := aoc.NewNode(aoc.PointWithDirection{Point: end.Point, Direction: aoc.Right}, end.Value)
	endNode2 := aoc.NewNode(aoc.PointWithDirection{Point: end.Point, Direction: aoc.Up}, end.Value)

	path1, cost1 := g.FindShortestPath(startNode, endNode1)
	grid1 := grid.Copy()
	for _, node := range path1 {
		grid1.Data[node.Key.Point] = 'P'
	}
	aoc.Debug(grid1.String())
	aoc.Debug(cost1)
	fmt.Println("shortest path1")

	path2, cost2 := g.FindShortestPath(startNode, endNode2)
	grid2 := grid.Copy()
	for _, node := range path2 {
		grid2.Data[node.Key.Point] = 'P'
	}
	aoc.Debug(grid2.String())
	aoc.Debug(cost2)
	fmt.Println("shortest path2")

	if cost1 < cost2 {
		fmt.Println(cost1)
	} else {
		fmt.Println(cost2)
	}
}

func part2() {
	aoc.Debug("todo")
}
