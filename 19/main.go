package main

import (
	"fmt"
	"os"
	"strings"
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

	dataSections := aoc.SectionsFromFileAsString(os.Args[2])
	towels := strings.Split(dataSections[0], ", ")
	patterns := strings.Split(dataSections[1], "\n")

	cache := map[string]int{}

	var numberOfPatternOptions func(pattern string) int
	numberOfPatternOptions = func(pattern string) int {
		if numberOfOptions, ok := cache[pattern]; ok {
			return numberOfOptions
		}

		numberOfOptions := 0

		if aoc.In(towels, pattern) {
			numberOfOptions++
		}

		for length := range len(pattern) {
			if length > 10 { // no larger towels in input
				break
			}

			potentialTowel := pattern[:length]
			restOfThePattern := pattern[length:]
			if aoc.In(towels, potentialTowel) {
				numberOfOptions += numberOfPatternOptions(restOfThePattern)
			}
		}

		cache[pattern] = numberOfOptions

		return numberOfOptions
	}

	possible := 0
	for _, pattern := range patterns {
		if numberOfPatternOptions(pattern) > 0 {
			possible++
		}
	}

	aoc.Println(t, possible)
}

func part2() {
	t := time.Now()

	dataSections := aoc.SectionsFromFileAsString(os.Args[2])
	towels := strings.Split(dataSections[0], ", ")
	patterns := strings.Split(dataSections[1], "\n")

	cache := map[string]int{}

	var numberOfPatternOptions func(pattern string) int
	numberOfPatternOptions = func(pattern string) int {
		if numberOfOptions, ok := cache[pattern]; ok {
			return numberOfOptions
		}

		numberOfOptions := 0

		if aoc.In(towels, pattern) {
			numberOfOptions++
		}

		for length := range len(pattern) {
			if length > 10 { // no larger towels in input
				break
			}

			potentialTowel := pattern[:length]
			restOfThePattern := pattern[length:]
			if aoc.In(towels, potentialTowel) {
				numberOfOptions += numberOfPatternOptions(restOfThePattern)
			}
		}

		cache[pattern] = numberOfOptions

		return numberOfOptions
	}

	numberOfAllPossibleArrangements := 0
	for _, pattern := range patterns {
		numberOfAllPossibleArrangements += numberOfPatternOptions(pattern)
	}

	aoc.Println(t, numberOfAllPossibleArrangements)
}
