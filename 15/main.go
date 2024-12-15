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
		pos := robotPos.Move(move)
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
			tmpPos := lastBoxPos.Move(move)
			if grid.Data[tmpPos] == 'O' {
				numberOfBoxes++
				lastBoxPos = tmpPos
			} else {
				break
			}
		}

		nextToLastBoxPos := lastBoxPos.Move(move)
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
	data := aoc.StringFromFile(os.Args[2])
	parts := strings.Split(data, "\n\n")
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

	lines := strings.Split(parts[0], "\n")
	grid := aoc.NewGrid(len(lines[0])*2, len(lines))
	for y, line := range lines {
		for x, r := range line {
			x1 := x * 2
			x2 := x1 + 1

			var v1, v2 rune

			switch r {
			case '@':
				v1 = '@'
				v2 = '.'
			case 'O':
				v1 = '['
				v2 = ']'
			default:
				v1 = r
				v2 = r
			}

			grid.Data[aoc.Point{X: x1, Y: y}] = v1
			grid.Data[aoc.Point{X: x2, Y: y}] = v2
		}
	}

	robotCell, ok := grid.FindFirstCellWithValue('@')
	if !ok {
		panic("no robot found")
	}
	robotPos := robotCell.Point
	grid.Data[robotPos] = '.'

	for _, move := range moves {
		aoc.Debug("move", move)
		grid.Data[robotPos] = '@'
		aoc.Debug(grid.String())
		grid.Data[robotPos] = '.'

		nextPos := robotPos.Move(move)
		nextVal := grid.Data[nextPos]

		if nextVal == '#' {
			continue
		}

		if nextVal == '.' {
			robotPos = nextPos
			continue
		}

		if isBox(nextVal) {
			if move.Y == 0 { // horizontal
				nextBoxPos := nextPos
				boxPositions := map[aoc.Point]rune{nextPos: nextVal}
				canMove := false
				for {
					nextBoxPos = nextBoxPos.Move(move)
					nextBoxVal := grid.Data[nextBoxPos]

					if nextBoxVal == '#' {
						canMove = false
						break
					}
					if nextBoxVal == '.' {
						canMove = true
						break
					}
					if isBox(nextBoxVal) {
						boxPositions[nextBoxPos] = nextBoxVal
					}
				}
				if !canMove {
					continue
				}

				robotPos = nextPos
				for p := range boxPositions {
					grid.Data[p] = '.'
				}
				for p, v := range boxPositions {
					grid.Data[p.Move(move)] = v
				}
				continue
			}

			// vertical
			otherBoxPartPos, otherBoxPartVal := otherBoxPartIfBox(nextPos, nextVal)
			boxPositions := map[aoc.Point]rune{
				nextPos:         nextVal,
				otherBoxPartPos: otherBoxPartVal,
			}

			canMove := true
			var addMoreBoxes func(aoc.Point, aoc.Direction)

			addMoreBoxes = func(pos aoc.Point, move aoc.Direction) {
				p := pos.Move(move)
				v := grid.Data[p]

				if v == '#' {
					canMove = false
					return
				}

				if isBox(v) {
					boxPositions[p] = v
					op, ov := otherBoxPartIfBox(p, v)
					boxPositions[op] = ov

					addMoreBoxes(p, move)
					addMoreBoxes(op, move)
				}
			}

			addMoreBoxes(nextPos, move)
			addMoreBoxes(otherBoxPartPos, move)

			if !canMove {
				continue
			}

			robotPos = nextPos
			for p := range boxPositions {
				grid.Data[p] = '.'
			}
			for p, v := range boxPositions {
				grid.Data[p.Move(move)] = v
			}
		}
	}

	grid.Data[robotPos] = '@'
	aoc.Debug(grid.String())

	sumOfGPSCoordinates := 0
	for cell := range grid.Iter() {
		if cell.Value == '[' {
			sumOfGPSCoordinates += cell.Point.X + 100*cell.Point.Y
		}
	}

	fmt.Println(sumOfGPSCoordinates)
}

func isBox(val rune) bool {
	return val == '[' || val == ']'
}

func otherBoxPartIfBox(pos aoc.Point, val rune) (aoc.Point, rune) {
	if val == '[' {
		return pos.Move(aoc.Right), ']'
	}
	if val == ']' {
		return pos.Move(aoc.Left), '['
	}
	return aoc.Point{}, aoc.NullRune
}
