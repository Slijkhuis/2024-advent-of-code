package main

import (
	"fmt"
	"os"
	"strings"

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
	data := aoc.StringFromFile(os.Args[2])
	parts := strings.Split(data, "\n\n")
	grid := aoc.BuildGridFromString(parts[0])
	var moves []aoc.Direction

	for _, r := range parts[1] {
		switch r {
		case '^':
			moves = append(moves, aoc.Up)
		case 'v':
			moves = append(moves, aoc.Down)
		case '<':
			moves = append(moves, aoc.Left)
		case '>':
			moves = append(moves, aoc.Right)
		case '\n':
			continue
		default:
			panic("invalid move: " + string(r))
		}
	}

	robotCell, ok := grid.FindFirstCellWithValue('@')
	if !ok {
		panic("no robot found")
	}
	robotPos := robotCell.Point
	grid.Data[robotPos] = '.'

	for _, move := range moves {
		pos := robotPos.Add(aoc.Point(move))
		val := grid.Data[pos]

		if val == '#' {
			continue
		}

		if val == '.' {
			robotPos = pos
			continue
		}

		if val != 'O' {
			panic("invalid state")
		}

		// - find the last box in line
		// - check if that one can be moved
		// - if so, teleport the current one there and the robot to `pos`

		lastBoxPos := pos
		numberOfBoxes := 1
		for {
			tmpPos := lastBoxPos.Add(aoc.Point(move))
			if grid.Data[tmpPos] == 'O' {
				numberOfBoxes++
				lastBoxPos = tmpPos
			} else {
				break
			}
		}

		nextToLastBoxPos := lastBoxPos.Add(aoc.Point(move))
		if grid.Data[nextToLastBoxPos] == '#' {
			continue
		}
		if grid.Data[nextToLastBoxPos] != '.' {
			panic("invalid state")
		}
		grid.Data[nextToLastBoxPos] = 'O'
		grid.Data[pos] = '.'
		robotPos = pos
	}

	grid.Data[robotPos] = '@'
	aoc.Debug(grid)

	sumOfGPSCoordinates := 0
	for cell := range grid.Iter() {
		if cell.Value == 'O' {
			sumOfGPSCoordinates += cell.Point.X + 100*cell.Point.Y
		}
	}

	fmt.Println(sumOfGPSCoordinates)
}

func part2() {
	aoc.Debug("todo")
}
