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
	width := 6 + 1
	height := 6 + 1
	fallingBytes := 12
	if !aoc.DebugMode {
		width = 70 + 1
		height = 70 + 1
		fallingBytes = 1024
	}

	grid := aoc.NewGridWithDefaultValue(width, height, '.')

	var byteCoords []aoc.Point
	for line := range aoc.LinesFromFile(os.Args[2]) {
		coords := aoc.IntsFromString(line)
		byteCoords = append(byteCoords, aoc.Point{X: coords[0], Y: coords[1]})
	}

	for i := range fallingBytes {
		grid.Data[byteCoords[i]] = '#'
	}

	aoc.Debug(grid.String())

	g := aoc.NewGraph[aoc.Point, rune]()

	visited := make(map[aoc.Point]struct{})
	var buildGraph func(*aoc.Node[aoc.Point, rune])
	buildGraph = func(comingFrom *aoc.Node[aoc.Point, rune]) {
		if _, ok := visited[comingFrom.Key]; ok {
			return
		}
		visited[comingFrom.Key] = struct{}{}

		for _, dir := range aoc.NoDiagonals {
			next := comingFrom.Key.Move(dir)
			val := grid.Data[next]
			if val == '.' {
				n := g.NewOrExistingNode(next, val)
				g.AddNodeIgnoreDuplicate(n)
				g.AddWeightedEdge(comingFrom, n, 1)
				buildGraph(n)
			}
		}
	}

	startNode := aoc.NewNode(aoc.Point{X: 0, Y: 0}, '.')
	g.AddNode(startNode)
	buildGraph(startNode)

	endNode := g.Nodes[aoc.Point{X: width - 1, Y: height - 1}]

	path, _ := g.FindShortestPath(startNode, endNode)
	aoc.Debug(path)

	for _, node := range path {
		grid.Data[node.Key] = 'P'
	}

	aoc.Debug(grid.String())

	fmt.Println(len(path) - 1)
}

func part2() {
	width := 6 + 1
	height := 6 + 1
	skipCheckUntilByte := 12
	if !aoc.DebugMode {
		width = 70 + 1
		height = 70 + 1
		skipCheckUntilByte = 1024
	}

	grid := aoc.NewGridWithDefaultValue(width, height, '.')

	var byteCoords []aoc.Point
	for line := range aoc.LinesFromFile(os.Args[2]) {
		coords := aoc.IntsFromString(line)
		byteCoords = append(byteCoords, aoc.Point{X: coords[0], Y: coords[1]})
	}

	for i := range skipCheckUntilByte + 1 {
		grid.Data[byteCoords[i]] = '#'
	}

	g := aoc.NewGraph[aoc.Point, rune]()

	visited := make(map[aoc.Point]struct{})
	var buildGraph func(*aoc.Node[aoc.Point, rune])
	buildGraph = func(comingFrom *aoc.Node[aoc.Point, rune]) {
		if _, ok := visited[comingFrom.Key]; ok {
			return
		}
		visited[comingFrom.Key] = struct{}{}

		for _, dir := range aoc.NoDiagonals {
			next := comingFrom.Key.Move(dir)
			val := grid.Data[next]
			if val == '.' {
				n := g.NewOrExistingNode(next, val)
				g.AddNodeIgnoreDuplicate(n)
				g.AddWeightedEdge(comingFrom, n, 1)
				buildGraph(n)
			}
		}
	}

	startNode := aoc.NewNode(aoc.Point{X: 0, Y: 0}, '.')
	g.AddNode(startNode)
	buildGraph(startNode)

	endNode := g.Nodes[aoc.Point{X: width - 1, Y: height - 1}]

	index := skipCheckUntilByte + 1
	for g.AreConnected(startNode, endNode) {
		index++

		if index >= len(byteCoords) {
			panic("something is badly wrong")
		}

		fmt.Println(index, "/", len(byteCoords))

		grid.Data[byteCoords[index]] = '#'
		for edge := range g.Edges {
			if edge.From.Key == byteCoords[index] || edge.To.Key == byteCoords[index] {
				delete(g.Edges, edge)
			}
		}
	}

	aoc.Debug(grid.String())

	// it's very slow... but it does work
	fmt.Println(byteCoords[index])
}
