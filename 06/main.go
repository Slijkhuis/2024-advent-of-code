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
	g := aoc.BuildGridFromFile(os.Args[2])
	start, ok := g.FindFirstCellWithValue('^')
	if !ok {
		aoc.Error("No start found")
	}

	p := start.Point
	dir := aoc.Up
	visited := map[aoc.Point]bool{
		p: true,
	}
	for {
		newPoint := p.Add(aoc.Point(dir))
		r, ok := g.Data[newPoint]
		if !ok {
			break
		}

		if r == '#' {
			dir = dir.TurnRight()
			continue
		}

		g.Data[newPoint] = 'X'
		visited[newPoint] = true
		p = newPoint

		if len(visited)%5 == 0 {
			aoc.Debug(g)
		}
	}

	fmt.Println(len(visited))
}

func part2() {
	g := aoc.BuildGridFromFile(os.Args[2])
	start, ok := g.FindFirstCellWithValue('^')
	if !ok {
		aoc.Error("No start found")
	}

	p := start.Point
	dir := aoc.Up
	visited := map[aoc.Point]bool{
		p: true,
	}
	for {
		newPoint := p.Add(aoc.Point(dir))
		r, ok := g.Data[newPoint]
		if !ok {
			break
		}

		if r == '#' {
			dir = dir.TurnRight()
			continue
		}

		visited[newPoint] = true
		p = newPoint
	}

	// Add obstacle to any of the visited points, except the starting point, and check if it loops.
	positions := 0
	count := 0
	for p := range visited {
		if p == start.Point {
			continue
		}

		aoc.Debug("Checking", p, " ", count, " of ", len(visited))
		count++

		g2 := g.Copy()
		g2.Data[p] = '#'
		if doesItLoop(g2, start.Point, aoc.Up) {
			positions++
		}
	}

	fmt.Println(positions)
}

func doesItLoop(g *aoc.Grid, p aoc.Point, dir aoc.Direction) bool {
	start := p
	visited := map[aoc.PointWithDirection]bool{
		{Point: start, Direction: dir}: true,
	}
	for {
		if p != start {
			if _, ok := visited[aoc.PointWithDirection{Point: p, Direction: dir}]; ok {
				return true
			}
		}

		newPoint := p.Add(aoc.Point(dir))
		r, ok := g.Data[newPoint]
		if !ok {
			break
		}

		if r == '#' {
			aoc.Debug("Turning right at", p)

			dir = dir.TurnRight()
			continue
		}

		visited[aoc.PointWithDirection{Point: p, Direction: dir}] = true
		p = newPoint
	}

	return false
}
