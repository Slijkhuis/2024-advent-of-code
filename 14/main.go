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

type Robot struct {
	Pos      aoc.Point
	Velocity aoc.Direction
}

func part1() {
	var robots []Robot

	width := 101
	height := 103
	if aoc.DebugMode {
		width = 11
		height = 7
	}

	for line := range aoc.LinesFromFile(os.Args[2]) {
		ints := aoc.IntsFromString(line)
		robots = append(robots, Robot{
			Pos:      aoc.Point{X: ints[0], Y: ints[1]},
			Velocity: aoc.Direction{X: ints[2], Y: ints[3]},
		})
	}

	for t := range 100 {
		aoc.Debug(t, robots)
		for i := range robots {
			robots[i].Pos.X += robots[i].Velocity.X
			if robots[i].Pos.X < 0 {
				robots[i].Pos.X = width + robots[i].Pos.X
			}
			if robots[i].Pos.X >= width {
				robots[i].Pos.X = robots[i].Pos.X - width
			}

			robots[i].Pos.Y += robots[i].Velocity.Y
			if robots[i].Pos.Y < 0 {
				robots[i].Pos.Y = height + robots[i].Pos.Y
			}
			if robots[i].Pos.Y >= height {
				robots[i].Pos.Y = robots[i].Pos.Y - height
			}
		}
	}

	answer := 1
	for _, quadrant := range [][4]int{
		{0, 0, width / 2, height / 2},
		{width/2 + 1, 0, width / 2, height / 2},
		{0, height/2 + 1, width / 2, height / 2},
		{width/2 + 1, height/2 + 1, width / 2, height / 2},
	} {
		quadrantCount := 0
		for _, robot := range robots {
			if robot.Pos.X >= quadrant[0] && robot.Pos.X < quadrant[0]+quadrant[2] && robot.Pos.Y >= quadrant[1] && robot.Pos.Y < quadrant[1]+quadrant[3] {
				quadrantCount++
			}
		}
		aoc.Debug(quadrant, quadrantCount)
		answer *= quadrantCount
	}

	fmt.Println(answer)
}

func part2() {
	var robots []Robot

	width := 101
	height := 103
	if aoc.DebugMode {
		width = 11
		height = 7
	}

	for line := range aoc.LinesFromFile(os.Args[2]) {
		ints := aoc.IntsFromString(line)
		robots = append(robots, Robot{
			Pos:      aoc.Point{X: ints[0], Y: ints[1]},
			Velocity: aoc.Direction{X: ints[2], Y: ints[3]},
		})
	}

	f, err := os.Create("14/data/output.txt")
	aoc.Fail(err)
	defer f.Close()

	printRobots := func(t int, robotCount map[aoc.Point]int) {
		fmt.Fprintln(f, "time", t+1)
		for x := range width {
			for y := range height {
				if robotCount[aoc.Point{X: x, Y: y}] > 0 {
					fmt.Fprint(f, robotCount[aoc.Point{X: x, Y: y}])
				} else {
					fmt.Fprint(f, ".")
				}
			}
			fmt.Fprintln(f)
		}
		fmt.Fprintln(f)
	}

	for t := 0; t < 10000; t++ {
		robotCount := map[aoc.Point]int{}
		robotCountPerRow := map[int]int{}
		for i := range robots {
			robots[i].Pos.X += robots[i].Velocity.X
			if robots[i].Pos.X < 0 {
				robots[i].Pos.X = width + robots[i].Pos.X
			}
			if robots[i].Pos.X >= width {
				robots[i].Pos.X = robots[i].Pos.X - width
			}

			robots[i].Pos.Y += robots[i].Velocity.Y
			if robots[i].Pos.Y < 0 {
				robots[i].Pos.Y = height + robots[i].Pos.Y
			}
			if robots[i].Pos.Y >= height {
				robots[i].Pos.Y = robots[i].Pos.Y - height
			}

			robotCount[robots[i].Pos]++
			robotCountPerRow[robots[i].Pos.X]++
		}

		robotsWithMoreThanThreeNeighbours := 0
		for i := range robots {
			neighbors := 0
			for _, direction := range aoc.Directions {
				if robotCount[aoc.Point{X: robots[i].Pos.X + direction.X, Y: robots[i].Pos.Y + direction.Y}] > 0 {
					neighbors++
				}
			}
			if neighbors > 3 {
				robotsWithMoreThanThreeNeighbours++
			}
		}

		if robotsWithMoreThanThreeNeighbours > 10 {
			printRobots(t, robotCount)
		}
	}
}
