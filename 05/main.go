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
	sections := strings.Split(data, "\n\n")
	ruleLines := strings.Split(sections[0], "\n")
	orders := strings.Split(sections[1], "\n")

	rulesByNumber := map[int][][2]int{}
	for _, rule := range ruleLines {
		numbers := aoc.Map(strings.Split(rule, "|"), aoc.Atoi)
		rulesByNumber[numbers[0]] = append(rulesByNumber[numbers[0]], [2]int{numbers[0], numbers[1]})
		rulesByNumber[numbers[1]] = append(rulesByNumber[numbers[1]], [2]int{numbers[0], numbers[1]})
	}

	aoc.Debug(ruleLines)

	var result int
	for _, order := range orders {
		numbers := aoc.Map(strings.Split(order, ","), aoc.Atoi)
		aoc.Debug(numbers)

		if findInvalidNumberIndex(numbers, rulesByNumber) == -1 {
			middle := numbers[len(numbers)/2]
			aoc.Debug(middle)
			result += middle
		}
	}

	fmt.Println(result)
}

func findInvalidNumberIndex(numbers []int, rulesByNumber map[int][][2]int) int {
	processed := map[int]bool{}
	for i, number := range numbers {
		rules, ok := rulesByNumber[number]
		if !ok {
			processed[number] = true
			continue
		}

		passedRules := true
		for _, rule := range rules {
			if rule[1] == number {
				continue
			}

			if processed[rule[1]] {
				passedRules = false
				break
			}
		}

		processed[number] = true
		if !passedRules {
			return i
		}
	}

	return -1
}

func part2() {
	data := aoc.StringFromFile(os.Args[2])
	sections := strings.Split(data, "\n\n")
	ruleLines := strings.Split(sections[0], "\n")
	orders := strings.Split(sections[1], "\n")

	rulesByNumber := map[int][][2]int{}
	for _, rule := range ruleLines {
		numbers := aoc.Map(strings.Split(rule, "|"), aoc.Atoi)
		rulesByNumber[numbers[0]] = append(rulesByNumber[numbers[0]], [2]int{numbers[0], numbers[1]})
		rulesByNumber[numbers[1]] = append(rulesByNumber[numbers[1]], [2]int{numbers[0], numbers[1]})
	}

	aoc.Debug(ruleLines)

	var result int
	for _, order := range orders {
		numbers := aoc.Map(strings.Split(order, ","), aoc.Atoi)
		aoc.Debug(numbers)

		invalidNumberIndex := findInvalidNumberIndex(numbers, rulesByNumber)

		if invalidNumberIndex == -1 {
			continue
		}

		loop := 0
		maxLoops := len(numbers) * 10 // rough estimate
		for invalidNumberIndex != -1 {
			numbers[invalidNumberIndex-1], numbers[invalidNumberIndex] = numbers[invalidNumberIndex], numbers[invalidNumberIndex-1]
			invalidNumberIndex = findInvalidNumberIndex(numbers, rulesByNumber)

			loop += 1
			if loop > maxLoops {
				fmt.Println("Max loops reached for", numbers)
				break
			}
		}

		if invalidNumberIndex == -1 {
			middle := numbers[len(numbers)/2]
			aoc.Debug(middle)
			result += middle
		}
	}

	fmt.Println(result)
}
