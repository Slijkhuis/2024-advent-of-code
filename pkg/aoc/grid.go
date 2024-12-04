package aoc

import (
	"fmt"
	"iter"
	"strings"
)

var (
	Up        = Point{0, -1}
	Down      = Point{0, 1}
	Left      = Point{-1, 0}
	Right     = Point{1, 0}
	UpLeft    = Point{-1, -1}
	UpRight   = Point{1, -1}
	DownLeft  = Point{-1, 1}
	DownRight = Point{1, 1}

	Directions = []Point{Up, Down, Left, Right, UpLeft, UpRight, DownLeft, DownRight}
	Diagonals  = []Point{UpLeft, UpRight, DownLeft, DownRight}
)

type Point struct {
	X, Y int
}

func (p Point) Add(other Point) Point {
	return Point{p.X + other.X, p.Y + other.Y}
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

type Cell struct {
	Point
	Value rune
}

type Grid struct {
	Width, Height int
	Data          map[Point]rune
}

func NewGrid(width, height int) *Grid {
	return &Grid{
		Width:  width,
		Height: height,
		Data:   make(map[Point]rune),
	}
}

func BuildGridFromFile(filename string) *Grid {
	data := StringFromFile(filename)
	lines := strings.Split(data, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	grid := NewGrid(len(lines[0]), len(lines))

	for y, line := range lines {
		for x, r := range line {
			grid.Data[Point{x, y}] = r
		}
	}

	return grid
}

func (g *Grid) Iter() iter.Seq[Cell] {
	return func(yield func(Cell) bool) {
		for y := 0; y < g.Height; y++ {
			for x := 0; x < g.Width; x++ {
				if !yield(Cell{Point{x, y}, g.Data[Point{x, y}]}) {
					return
				}
			}
		}
	}
}
func (g *Grid) AdjOrNull(p Point, direction Point) rune {
	value, ok := g.Data[p.Add(direction)]
	if !ok {
		return '\000'
	}
	return value
}

func (g *Grid) AdjString(p Point, direction Point, length int) (string, bool) {
	var result string

	for i := 0; i < length; i++ {
		cell, ok := g.Data[p]
		if !ok {
			return "", false
		}

		result += string(cell)

		p = p.Add(direction)
	}

	return result, true
}

func (g *Grid) String() string {
	var sb strings.Builder
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			sb.WriteRune(g.Data[Point{x, y}])
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}
