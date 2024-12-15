package main

import (
	"fmt"
	"os"

	"github.com/Slijkhuis/2024-advent-of-code/pkg/aoc"
)

func main() {
	part1()
	part2()
}

func part1() {
	searchFor := "XMAS"
	grid := aoc.BuildGridFromFile(os.Args[1])
	count := 0
	for cell := range grid.Iter() {
		if cell.Value == 'X' {
			for _, dir := range aoc.Directions {
				aoc.Debug("Checking", cell.Point, "in direction", dir)

				result, ok := grid.AdjString(cell.Point, dir, len(searchFor))
				if !ok {
					aoc.Debug("Out of bounds")
					continue
				}
				aoc.Debug("Found", result, "at", cell.Point, "in direction", dir)
				if result == searchFor {
					// aoc.Debug("Found", searchFor, "at", cell.Point, "in direction", dir)
					count++
				}
			}
		}
	}

	fmt.Println("Part 1 result:")
	fmt.Println(count)
}

func part2() {
	searchFor := "MAS"
	grid := aoc.BuildGridFromFile(os.Args[1])
	count := 0
	for cell := range grid.Iter() {
		if cell.Value == 'A' {
			diag1 := string(grid.AdjOrNull(cell.Point, aoc.UpLeft)) + "A" + string(grid.AdjOrNull(cell.Point, aoc.DownRight))
			diag2 := string(grid.AdjOrNull(cell.Point, aoc.UpRight)) + "A" + string(grid.AdjOrNull(cell.Point, aoc.DownLeft))

			diag1Ok := diag1 == searchFor || aoc.ReverseString(diag1) == searchFor
			diag2Ok := diag2 == searchFor || aoc.ReverseString(diag2) == searchFor

			if diag1Ok && diag2Ok {
				count++
			}
		}
	}

	fmt.Println("Part 2 result:")
	fmt.Println(count)
}
