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

func (p Point) DistanceToUsingNoDiagonals(p2 Point) int {
	return Abs(p.X-p2.X) + Abs(p.Y-p2.Y)
}

type Direction Point

func (d Direction) TurnRight() Direction {
	return Direction{-d.Y, d.X}
}

func (d Direction) TurnLeft() Direction {
	return Direction{d.Y, -d.X}
}

func (d Direction) Opposite() Direction {
	return Direction{-d.X, -d.Y}
}

func (d Direction) Rune() rune {
	switch d {
	case Up:
		return '^'
	case Left:
		return '<'
	case Right:
		return '>'
	case Down:
		return 'v'
	case UpLeft, DownRight:
		return '\\'
	case UpRight, DownLeft:
		return '/'
	default:
		return '?'
	}
}

func (d Direction) String() string {
	return string(d.Rune())
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

func NewGridWithDefaultValue(width, height int, val rune) *Grid {
	data := make(map[Point]rune)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			data[Point{x, y}] = val
		}
	}

	return &Grid{
		Width:  width,
		Height: height,
		Data:   data,
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

func (g *Grid) MustFindFirstPointWithValue(value rune) Point {
	for cell := range g.Iter() {
		if cell.Value == value {
			return cell.Point
		}
	}
	panic("not found in grid: " + string(value))
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
