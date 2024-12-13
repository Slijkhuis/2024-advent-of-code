package main

import (
	"fmt"
	"os"
	"regexp"
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

type ClawMachine struct {
	A     aoc.Direction
	B     aoc.Direction
	Prize aoc.Point
}

func part1() {
	data := aoc.StringFromFile(os.Args[2])
	clawMachinesStrs := strings.Split(data, "\n\n")
	r := regexp.MustCompile(`\d+`)
	clawMachines := aoc.Map(clawMachinesStrs, func(s string) ClawMachine {
		lines := strings.Split(s, "\n")
		aMatches := r.FindAllString(lines[0], -1)
		bMatches := r.FindAllString(lines[1], -1)
		prizeMatches := r.FindAllString(lines[2], -1)

		return ClawMachine{
			A:     aoc.Direction{X: int(aoc.Atoi(aMatches[0])), Y: int(aoc.Atoi(aMatches[1]))},
			B:     aoc.Direction{X: int(aoc.Atoi(bMatches[0])), Y: int(aoc.Atoi(bMatches[1]))},
			Prize: aoc.Point{X: int(aoc.Atoi(prizeMatches[0])), Y: int(aoc.Atoi(prizeMatches[1]))},
		}
	})
	aoc.Debug(clawMachines)

	tokens := 0
	for _, cm := range clawMachines {
		hasSolution := false
		maxButtonBPresses := aoc.Min((cm.Prize.X/cm.B.X)+1, (cm.Prize.Y/cm.B.Y)+1)
		for buttonBPresses := maxButtonBPresses; buttonBPresses >= 0; buttonBPresses-- {
			maxButtonAPresses := aoc.Min(((cm.Prize.X-buttonBPresses*cm.B.X)/cm.A.X)+1, ((cm.Prize.Y-buttonBPresses*cm.B.Y)/cm.A.Y)+1)
			for buttonAPresses := maxButtonAPresses; buttonAPresses >= 0; buttonAPresses-- {
				x := buttonAPresses*cm.A.X + buttonBPresses*cm.B.X
				y := buttonAPresses*cm.A.Y + buttonBPresses*cm.B.Y
				if x == cm.Prize.X && y == cm.Prize.Y {
					aoc.Debug(cm, "has solution", buttonAPresses, buttonBPresses)
					hasSolution = true
					tokens += buttonAPresses*3 + buttonBPresses*1
					break
				}
			}
			if hasSolution {
				break
			}
		}

		if !hasSolution {
			aoc.Debug(cm, "no solution", maxButtonBPresses)
		}
	}

	fmt.Println(tokens)
}

func part2() {
	data := aoc.StringFromFile(os.Args[2])
	clawMachinesStrs := strings.Split(data, "\n\n")
	r := regexp.MustCompile(`\d+`)
	clawMachines := aoc.Map(clawMachinesStrs, func(s string) ClawMachine {
		lines := strings.Split(s, "\n")
		aMatches := r.FindAllString(lines[0], -1)
		bMatches := r.FindAllString(lines[1], -1)
		prizeMatches := r.FindAllString(lines[2], -1)

		return ClawMachine{
			A:     aoc.Direction{X: int(aoc.Atoi(aMatches[0])), Y: int(aoc.Atoi(aMatches[1]))},
			B:     aoc.Direction{X: int(aoc.Atoi(bMatches[0])), Y: int(aoc.Atoi(bMatches[1]))},
			Prize: aoc.Point{X: int(10_000_000_000_000 + aoc.Atoi(prizeMatches[0])), Y: int(10_000_000_000_000 + aoc.Atoi(prizeMatches[1]))},
		}
	})
	aoc.Debug(clawMachines)

	tokens := 0
	for _, cm := range clawMachines {
		determinant := cm.A.X*cm.B.Y - cm.A.Y*cm.B.X
		if determinant == 0 {
			aoc.Debug(cm, "no solution")
			continue
		}
		numeratorA := cm.Prize.X*cm.B.Y - cm.Prize.Y*cm.B.X
		numeratorB := cm.A.X*cm.Prize.Y - cm.A.Y*cm.Prize.X

		if numeratorA%determinant != 0 || numeratorB%determinant != 0 {
			aoc.Debug(cm, "no solution")
			continue
		}

		a := numeratorA / determinant
		b := numeratorB / determinant

		aoc.Debug(cm, "a", a, "b", b)

		tokens += a*3 + b*1
	}

	fmt.Println(tokens)
}
