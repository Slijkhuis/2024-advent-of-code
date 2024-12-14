package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/Slijkhuis/2024-advent-of-code/pkg/aoc"
)

func main() {
	part1()
	part2()
}

func part1() {
	data := aoc.StringFromFile(os.Args[1])
	results := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`).FindAllString(data, -1)
	answer := 0
	for _, result := range results {
		aoc.Debug(result)
		numbers := aoc.Map(strings.Split(result[4:len(result)-1], ","), aoc.Atoi)
		answer += numbers[0] * numbers[1]
		aoc.Debug(numbers)
	}
	fmt.Println(answer)
}

func part2() {
	data := aoc.StringFromFile(os.Args[1])
	results := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\))|(do\(\))|(don't\(\))`).FindAllString(data, -1)
	answer := 0
	do := true
	for _, result := range results {
		if result == "do()" {
			do = true
			continue
		}
		if result == "don't()" {
			do = false
			continue
		}

		if !do {
			continue
		}

		aoc.Debug(result)
		numbers := aoc.Map(strings.Split(result[4:len(result)-1], ","), aoc.Atoi)
		answer += numbers[0] * numbers[1]
		aoc.Debug(numbers)
	}
	fmt.Println(answer)
}

func fail(err error) {
	if err != nil {
		panic(err)
	}
}
