package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Slijkhuis/2024-advent-of-code/pkg/aoc"
)

func main() {
	part1()
	part2()
}

func part1() {
	var safeReports int
	for line := range aoc.LinesFromFile(os.Args[1]) {
		numbers := aoc.Map(strings.Fields(line), aoc.Atoi)

		if checkSafety(numbers) {
			safeReports++
		}
	}

	fmt.Println(safeReports)
}

func part2() {
	var safeReports int
	for line := range aoc.LinesFromFile(os.Args[1]) {
		numbers := aoc.Map(strings.Fields(line), aoc.Atoi)

		aoc.Debug(line)

		// skip first number (no need to check checkSafety(numbers) since that would be safe too, if this is safe)
		safe := checkSafety(numbers[1:])

		if !safe { // skip second number
			aoc.Debug("second check: unsafe")
			safe = checkSafety(append(numbers[:0], numbers[2:]...))
		}
		if !safe { // skip any number after that
			aoc.Debug("third check: unsafe")
			skipIndex := getSkipIndex(numbers)
			if skipIndex >= 0 {
				safe = checkSafety(append(numbers[:skipIndex], numbers[skipIndex+1:]...))
			}
		}

		if safe {
			aoc.Debug("safe!")
			safeReports++
		} else {
			aoc.Debug("last check: unsafe")
		}
	}

	fmt.Println(safeReports)
}

func checkSafety(numbers []int) bool {
	lastNumber := numbers[0]
	increasing := false
	safe := true
	for i, number := range numbers {
		if i == 0 {
			continue
		}

		if i == 1 {
			increasing = number > lastNumber
		}

		if i > 1 && increasing != (number > lastNumber) {
			safe = false
			break
		}

		diff := aoc.Abs(number - lastNumber)
		if diff < 1 || diff > 3 {
			safe = false
			break
		}

		lastNumber = number
	}

	return safe
}
func getSkipIndex(numbers []int) int {
	lastNumber := numbers[0]
	increasing := false
	for i, number := range numbers {
		if i == 0 {
			continue
		}

		if i == 1 {
			increasing = number > lastNumber
		}

		// The increasing/decreasing check should not be skipped since we're checking for numbers[1:] as well.
		if i > 1 && increasing != (number > lastNumber) {
			break
		}

		diff := aoc.Abs(number - lastNumber)
		if diff < 1 || diff > 3 {
			return i
		}

		lastNumber = number
	}

	return -1
}
