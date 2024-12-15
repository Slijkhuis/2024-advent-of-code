package aoc

import (
	"fmt"
	"iter"
	"strings"
)

var (
	Up        = Direction{0, -1}
	Down      = Direction{0, 1}
	Left      = Direction{-1, 0}
	Right     = Direction{1, 0}
	UpLeft    = Direction{-1, -1}
	UpRight   = Direction{1, -1}
	DownLeft  = Direction{-1, 1}
	DownRight = Direction{1, 1}

	Directions  = []Direction{Up, Down, Left, Right, UpLeft, UpRight, DownLeft, DownRight}
	Diagonals   = []Direction{UpLeft, UpRight, DownLeft, DownRight}
	NoDiagonals = []Direction{Up, Down, Left, Right}
)

type Point struct {
	X, Y int
}

func (p Point) Move(inDirection Direction) Point {
	return Point{p.X + inDirection.X, p.Y + inDirection.Y}
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

type Direction Point

func (d Direction) TurnRight() Direction {
	return Direction{-d.Y, d.X}
}

type PointWithDirection struct {
	Point
	Direction
}

const NullRune rune = '\000'

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
	return BuildGridFromString(data)
}

func BuildGridFromString(data string) *Grid {
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

func (g *Grid) Copy() *Grid {
	newGrid := NewGrid(g.Width, g.Height)
	for point, value := range g.Data {
		newGrid.Data[point] = value
	}
	return newGrid
}

func (g *Grid) FindFirstCellWithValue(value rune) (Cell, bool) {
	for cell := range g.Iter() {
		if cell.Value == value {
			return cell, true
		}
	}
	return Cell{}, false
}

func (g *Grid) AdjOrNull(p Point, direction Direction) rune {
	value, ok := g.Data[p.Move(direction)]
	if !ok {
		return NullRune
	}
	return value
}

func (g *Grid) AdjString(p Point, direction Direction, length int) (string, bool) {
	var result string

	for i := 0; i < length; i++ {
		cell, ok := g.Data[p]
		if !ok {
			return "", false
		}

		result += string(cell)

		p = p.Move(direction)
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

func (g *Grid) InBounds(p Point) bool {
	return p.X >= 0 && p.X < g.Width && p.Y >= 0 && p.Y < g.Height
}
