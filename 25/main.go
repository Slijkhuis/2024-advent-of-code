package main

import (
	"fmt"
	"os"
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
	sections := aoc.SectionsFromFileAsString(os.Args[2])
	var locks []*aoc.Grid
	var keys []*aoc.Grid
	for _, section := range sections {
		grid := aoc.BuildGridFromString(section)
		if grid.Data[aoc.Point{X: 0, Y: 0}] == '#' {
			keys = append(keys, grid)
		} else {
			locks = append(locks, grid)
		}
	}

	aoc.Println(t, "parsed input")

	var numFittingPairs int

	for _, lock := range locks {
		for _, key := range keys {
			overlaps := false
			for col := range lock.Width {
				for row := range lock.Height {
					if lock.Data[aoc.Point{X: col, Y: row}] == '#' && key.Data[aoc.Point{X: col, Y: row}] == '#' {
						overlaps = true
						break
					}
				}
				if overlaps {
					break
				}
			}
			if !overlaps {
				numFittingPairs++
			}
		}
	}

	aoc.Println(t, numFittingPairs)
}

func part2() {
	for line := range aoc.LinesFromFile(os.Args[2]) {
		fmt.Println(line)
	}
}
