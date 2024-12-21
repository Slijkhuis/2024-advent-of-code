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

/*
---+---+---+
| 7 | 8 | 9 |
+---+---+---+
| 4 | 5 | 6 |
+---+---+---+
| 1 | 2 | 3 |
+---+---+---+
    | 0 | A |
    +---+---+

    +---+---+
    | ^ | A |
+---+---+---+
| < | v | > |
+---+---+---+
*/

func part1() {
	t := time.Now()

	var codes []Code
	for line := range aoc.LinesFromFile(os.Args[2]) {
		numericPart := aoc.IntsFromString(line)[0]
		buttons := []rune(line)
		codes = append(codes, Code{Numeric: numericPart, Buttons: buttons})
	}
	aoc.Debug(codes)

	numberPad := aoc.NewGrid(3, 4)
	numberPad.Data[aoc.Point{X: 0, Y: 0}] = '7'
	numberPad.Data[aoc.Point{X: 1, Y: 0}] = '8'
	numberPad.Data[aoc.Point{X: 2, Y: 0}] = '9'
	numberPad.Data[aoc.Point{X: 0, Y: 1}] = '4'
	numberPad.Data[aoc.Point{X: 1, Y: 1}] = '5'
	numberPad.Data[aoc.Point{X: 2, Y: 1}] = '6'
	numberPad.Data[aoc.Point{X: 0, Y: 2}] = '1'
	numberPad.Data[aoc.Point{X: 1, Y: 2}] = '2'
	numberPad.Data[aoc.Point{X: 2, Y: 2}] = '3'
	numberPad.Data[aoc.Point{X: 1, Y: 3}] = '0'
	numberPad.Data[aoc.Point{X: 2, Y: 3}] = 'A'

	numberGraph := aoc.NewStandardGraph()
	for cell := range numberPad.Iter() {
		n := numberGraph.NewOrExistingNode(cell.Point, cell.Value)

		for _, dir := range aoc.NoDiagonals {
			neighPos := cell.Point.Move(dir)
			neighVal, ok := numberPad.Data[neighPos]
			if !ok || neighVal == aoc.NullRune {
				continue
			}
			neighNode := numberGraph.NewOrExistingNode(neighPos, neighVal)
			numberGraph.AddWeightedEdge(n, neighNode, 1)
		}
	}

	arrowPad := aoc.NewGrid(3, 2)
	arrowPad.Data[aoc.Point{X: 1, Y: 0}] = '^'
	arrowPad.Data[aoc.Point{X: 2, Y: 0}] = 'A'
	arrowPad.Data[aoc.Point{X: 0, Y: 1}] = '<'
	arrowPad.Data[aoc.Point{X: 1, Y: 1}] = 'v'
	arrowPad.Data[aoc.Point{X: 2, Y: 1}] = '>'

	arrowGraph := aoc.NewStandardGraph()
	for cell := range arrowPad.Iter() {
		n := arrowGraph.NewOrExistingNode(cell.Point, cell.Value)

		for _, dir := range aoc.NoDiagonals {
			neighPos := cell.Point.Move(dir)
			neighVal, ok := arrowPad.Data[neighPos]
			if !ok || neighVal == aoc.NullRune {
				continue
			}
			neighNode := arrowGraph.NewOrExistingNode(neighPos, neighVal)
			arrowGraph.AddWeightedEdge(n, neighNode, 1)
		}
	}

	aoc.Println(t, "parsed input")

	// The plan is to produce a single best sequence every step, instead of finding all shortest paths, using some
	// assumptions (see buttonSequenceForInputAndPad func).
	answer := 0
	for _, code := range codes {
		aoc.Debug(aoc.JoinRunes(code.Buttons, ""))

		seq1 := buttonSequenceForInputAndPad(code.Buttons, numberPad)
		aoc.Debug(aoc.JoinRunes(seq1, ""))
		seq2 := buttonSequenceForInputAndPad(seq1, arrowPad)
		aoc.Debug(aoc.JoinRunes(seq2, ""))
		seq3 := buttonSequenceForInputAndPad(seq2, arrowPad)
		aoc.Debug(aoc.JoinRunes(seq3, ""))

		aoc.Debugf("%d * %d = %d", len(seq3), code.Numeric, code.Numeric*len(seq3))
		answer += code.Numeric * len(seq3)
	}

	aoc.Println(t, answer)
}

func buttonSequenceForInputAndPad(input []rune, pad *aoc.Grid) []rune {
	emptySpot := pad.MustFindFirstPointWithValue(0)

	pos := pad.MustFindFirstPointWithValue('A')
	result := []rune{}

	for _, button := range input {
		next := pad.MustFindFirstPointWithValue(button)

		dx := next.X - pos.X
		dy := next.Y - pos.Y

		// After trial and error, it seems certain order of directions matter:
		// <v is better than v<
		// v> is better than >v
		// <^ is better than ^<
		// not sure about: ^> vs >^
		// However, most important is to always move in the same direction first, if possible.
		// And of course the empty spot needs to be avoided.
		// The below is a result of A LOT of trial and error..
		firstXThanY := false
		if pos.Y == emptySpot.Y && next.X == emptySpot.X {
			firstXThanY = false
		} else if pos.X == emptySpot.X && next.Y == emptySpot.Y {
			firstXThanY = true
		} else if dx < 0 {
			firstXThanY = true
		}

		xdir := '<'
		if dx > 0 {
			xdir = '>'
		}
		if firstXThanY {
			for range aoc.Abs(dx) {
				result = append(result, xdir)
			}
		}

		ydir := '^'
		if dy > 0 {
			ydir = 'v'
		}
		for range aoc.Abs(dy) {
			result = append(result, ydir)
		}

		if !firstXThanY {
			for range aoc.Abs(dx) {
				result = append(result, xdir)
			}
		}

		result = append(result, 'A')
		pos = next
	}

	return result
}

func part2() {
	for line := range aoc.LinesFromFile(os.Args[2]) {
		fmt.Println(line)
	}
}

type Code struct {
	Numeric int
	Buttons []rune
}
