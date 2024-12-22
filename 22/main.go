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

	var initialNumbers []int
	for line := range aoc.LinesFromFile(os.Args[2]) {
		initialNumbers = append(initialNumbers, aoc.Atoi(line))
	}

	sum := 0
	for _, n := range initialNumbers {
		for range 2000 {
			n = generateSecretNumber(n)
		}
		aoc.Debug(n)
		sum += n
	}

	aoc.Println(t, sum)
}

func generateSecretNumber(n int) int {
	n = (n ^ (n * 64)) % 16777216
	n = (n ^ (n / 32)) % 16777216
	n = (n ^ (n * 2048)) % 16777216
	return n
}

func part2() {
	for line := range aoc.LinesFromFile(os.Args[2]) {
		fmt.Println(line)
	}
}
