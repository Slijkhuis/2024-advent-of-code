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
	regions := make([][]aoc.Point, 0)
	partOfRegion := make(map[aoc.Point]bool)
	for cell := range g.Iter() {
		if _, ok := partOfRegion[cell.Point]; ok {
			continue
		}

		region := []aoc.Point{cell.Point}

		var expandRegion func(p aoc.Point)
		expandRegion = func(p aoc.Point) {
			if _, ok := partOfRegion[p]; ok {
				return
			}
			partOfRegion[p] = true

			for _, dir := range aoc.NoDiagonals {
				newPoint := p.Add(aoc.Point(dir))
				r, ok := g.Data[newPoint]
				if !ok {
					continue
				}
				if _, ok := partOfRegion[newPoint]; ok {
					continue
				}

				if r == cell.Value {
					region = append(region, newPoint)
					expandRegion(newPoint)
				}
			}
		}

		expandRegion(cell.Point)

		regions = append(regions, region)
	}

	perimiters := make(map[string]int)
	areas := make(map[string]int)
	for i, region := range regions {
		r := g.Data[region[0]]
		key := fmt.Sprintf("%c-%d", r, i)
		areas[key] += len(region)

		for _, p := range region {
			for _, dir := range aoc.NoDiagonals {
				r2 := g.AdjOrNull(p, dir)
				if r != r2 {
					perimiters[key]++
				}
			}
		}
	}

	answer := 0
	for plot, area := range areas {
		aoc.Debug("A region of", string(plot), "plants with price", area, "*", perimiters[plot], "=", area*perimiters[plot])
		answer += area * perimiters[plot]
	}
	fmt.Println(answer)
}

func part2() {
	g := aoc.BuildGridFromFile(os.Args[2])
	regions := make([][]aoc.Point, 0)
	partOfRegion := make(map[aoc.Point]bool)
	answer := 0
	for cell := range g.Iter() {
		if _, ok := partOfRegion[cell.Point]; ok {
			continue
		}

		region := []aoc.Point{cell.Point}

		var expandRegion func(p aoc.Point)
		expandRegion = func(p aoc.Point) {
			if _, ok := partOfRegion[p]; ok {
				return
			}
			partOfRegion[p] = true

			for _, dir := range aoc.NoDiagonals {
				newPoint := p.Add(aoc.Point(dir))
				r, ok := g.Data[newPoint]
				if !ok {
					continue
				}
				if _, ok := partOfRegion[newPoint]; ok {
					continue
				}

				if r == cell.Value {
					region = append(region, newPoint)
					expandRegion(newPoint)
				}
			}
		}

		expandRegion(cell.Point)

		regions = append(regions, region)
		corners := 0
		for _, p := range region {
			for _, dirs := range [][2]aoc.Direction{
				{aoc.Up, aoc.Left},
				{aoc.Up, aoc.Right},
				{aoc.Down, aoc.Left},
				{aoc.Down, aoc.Right},
			} {
				a := g.AdjOrNull(p, dirs[0])
				b := g.AdjOrNull(p, dirs[1])
				c := g.AdjOrNull(p, aoc.Direction{X: dirs[0].X + dirs[1].X, Y: dirs[0].Y + dirs[1].Y})

				if a != cell.Value && b != cell.Value {
					corners++
				} else if a == cell.Value && b == cell.Value && c != cell.Value {
					corners++
				}
			}
		}

		aoc.Debug(string(cell.Value), region, corners)

		sides := corners
		answer += sides * len(region)
	}

	fmt.Println(answer)
}
