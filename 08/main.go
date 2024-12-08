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
	aoc.Debug(g)

	frequencies := map[rune][]aoc.Point{}
	for cell := range g.Iter() {
		if cell.Value == '.' {
			continue
		}
		frequencies[cell.Value] = append(frequencies[cell.Value], cell.Point)
	}

	antinodes := map[aoc.Point]struct{}{}
	for frequency, positions := range frequencies {
		aoc.Debug(string(frequency), positions)
		for _, p1 := range positions {
			for _, p2 := range positions {
				if p1 == p2 {
					continue
				}

				v1 := aoc.Point{X: p2.X - p1.X, Y: p2.Y - p1.Y}
				v2 := aoc.Point{X: p1.X - p2.X, Y: p1.Y - p2.Y}

				antinode1 := aoc.Point{X: p1.X - v1.X, Y: p1.Y - v1.Y}
				antinode2 := aoc.Point{X: p2.X + v2.X, Y: p2.Y + v2.Y}

				if g.Data[antinode1] == '.' {
					g.Data[antinode1] = '#'
				}

				if g.Data[antinode2] == '.' {
					g.Data[antinode2] = '#'
				}

				if g.InBounds(antinode1) && !aoc.In(positions, antinode1) {
					antinodes[antinode1] = struct{}{}
				}

				if g.InBounds(antinode2) && !aoc.In(positions, antinode2) {
					antinodes[antinode2] = struct{}{}
				}
			}
		}
	}

	aoc.Debug(g)

	fmt.Println(len(antinodes))
}

func part2() {
	g := aoc.BuildGridFromFile(os.Args[2])
	aoc.Debug(g)

	frequencies := map[rune][]aoc.Point{}
	for cell := range g.Iter() {
		if cell.Value == '.' {
			continue
		}
		frequencies[cell.Value] = append(frequencies[cell.Value], cell.Point)
	}

	antinodes := map[aoc.Point]struct{}{}
	for frequency, positions := range frequencies {
		aoc.Debug(string(frequency), positions)

		if len(positions) == 1 {
			continue
		}

		for _, p1 := range positions {
			for _, p2 := range positions {
				if p1 == p2 {
					continue
				}

				v1 := aoc.Point{X: p2.X - p1.X, Y: p2.Y - p1.Y}
				v2 := aoc.Point{X: p1.X - p2.X, Y: p1.Y - p2.Y}

				antinode1 := p1
				for g.InBounds(antinode1) {
					antinodes[antinode1] = struct{}{}
					antinode1 = aoc.Point{X: antinode1.X - v1.X, Y: antinode1.Y - v1.Y}
				}

				antinode2 := p2
				for g.InBounds(antinode2) {
					antinodes[antinode2] = struct{}{}
					antinode2 = aoc.Point{X: antinode2.X + v2.X, Y: antinode2.Y + v2.Y}
				}
			}
		}
	}

	aoc.Debug(g)

	fmt.Println(len(antinodes))
}
