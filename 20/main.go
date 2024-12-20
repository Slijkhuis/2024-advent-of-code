package main

import (
	"fmt"
	"os"
	"sort"
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
	grid := aoc.BuildGridFromFile(os.Args[2])
	_, _, path := buildPath(grid)

	pointIndices := map[aoc.Point]int{}
	for i, p := range path {
		pointIndices[p] = i
	}

	var cheatPaths [][]aoc.Point
	for i, p := range path {
		// pass through walls
		for _, dir := range aoc.NoDiagonals {
			next := p.Move(dir)
			if _, ok := pointIndices[next]; ok {
				continue
			}
			v, ok := grid.Data[next]
			if !ok || v != '#' {
				continue
			}

			nextnext := next.Move(dir)
			index, ok := pointIndices[nextnext]
			if !ok {
				continue
			}

			if index <= i { // not a shortcut
				continue
			}

			pathClone := append(path[:0:0], path...)
			cheatPath := append(pathClone[:i+1], next)
			cheatPath = append(cheatPath, pathClone[index:]...)
			cheatPaths = append(cheatPaths, cheatPath)
		}
	}

	fastestTimeWithoutCheats := len(path) - 1
	cheatsPerSavings := map[int]int{}
	allSavings := []int{}
	cheatsThatSaveAtLeast100ps := 0
	for _, cheatPath := range cheatPaths {
		trackTime := len(cheatPath) - 1
		savings := fastestTimeWithoutCheats - trackTime
		if savings <= 0 {
			continue
		}

		cheatsPerSavings[savings]++
		allSavings = append(allSavings, savings)

		if savings >= 100 {
			cheatsThatSaveAtLeast100ps++
		}
	}

	sort.IntSlice(allSavings).Sort()
	for _, savings := range aoc.Unique(allSavings) {
		aoc.Debug("There are", cheatsPerSavings[savings], "cheats that save", savings, "picoseconds.")
	}

	aoc.Println(t, "cheats that save at least 100 picoseconds:", cheatsThatSaveAtLeast100ps)
}

func buildPath(grid *aoc.Grid) (aoc.Point, aoc.Point, []aoc.Point) {
	start, ok := grid.FindFirstCellWithValue('S')
	if !ok {
		aoc.Error("No start found")
	}
	end, ok := grid.FindFirstCellWithValue('E')
	if !ok {
		aoc.Error("No end found")
	}
	visited := map[aoc.Point]bool{start.Point: true}
	current := start.Point
	path := []aoc.Point{start.Point}
	for {
		for _, dir := range aoc.NoDiagonals {
			next := current.Move(dir)

			_, ok := visited[next]
			if ok {
				continue
			}
			visited[next] = true

			v, ok := grid.Data[next]
			if !ok {
				continue
			}
			if v == '#' {
				continue
			}
			path = append(path, next)
			current = next
		}

		if current == end.Point {
			break
		}
	}
	aoc.Debug(len(path), path)

	return start.Point, end.Point, path
}

func part2() {
	t := time.Now()
	grid := aoc.BuildGridFromFile(os.Args[2])
	start, end, path := buildPath(grid)
	aoc.Debug(start, end)

	var cheatSavings []int
	for i1, p1 := range path {
		for i2, p2 := range path {
			if i1 > i2 {
				continue
			}
			if p1 == p2 {
				continue
			}

			d := p1.DistanceToUsingNoDiagonals(p2)
			if d <= 1 || d > 20 {
				continue
			}

			nonCheatingDistanceToEnd := len(path) - i1
			cheatingDistanceToEnd := len(path) - i2

			savings := nonCheatingDistanceToEnd - cheatingDistanceToEnd - d

			cheatSavings = append(cheatSavings, savings)
		}
	}

	var cheatsThatSaveAtLeast100ps int
	cheatsPerSavings := map[int]int{}
	for _, saving := range cheatSavings {
		cheatsPerSavings[saving]++
		if saving >= 100 {
			cheatsThatSaveAtLeast100ps++
		}
	}

	sort.IntSlice(cheatSavings).Sort()
	for _, savings := range aoc.Unique(cheatSavings) {
		if savings < 50 {
			continue
		}
		aoc.Debug("There are", cheatsPerSavings[savings], "cheats that save", savings, "picoseconds.")
	}

	aoc.Println(t, "cheats that save at least 100 picoseconds:", cheatsThatSaveAtLeast100ps)
}
