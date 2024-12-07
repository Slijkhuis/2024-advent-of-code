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
	sum := int64(0)
	for line := range aoc.LinesFromFile(os.Args[2]) {
		data := strings.Split(line, ": ")
		testResult := aoc.Atoi(data[0])
		numbers := aoc.Map(strings.Split(data[1], " "), aoc.Atoi)
		startingNode := &aoc.TreeNode[int64]{Value: numbers[0]}

		result := getNumberOfPossibilities(startingNode, numbers[1:], testResult)
		if result > 0 {
			sum += testResult
		}
	}

	fmt.Println(sum)
}

func getNumberOfPossibilities(
	node *aoc.TreeNode[int64], numbers []int64, testResult int64,
) int {
	if len(numbers) == 0 {
		if node.Value == testResult {
			return 1
		} else {
			return 0
		}
	}

	node1 := &aoc.TreeNode[int64]{Value: node.Value * numbers[0]}
	node2 := &aoc.TreeNode[int64]{Value: node.Value + numbers[0]}

	return getNumberOfPossibilities(node1, numbers[1:], testResult) +
		getNumberOfPossibilities(node2, numbers[1:], testResult)
}

func part2() {
	sum := int64(0)
	for line := range aoc.LinesFromFile(os.Args[2]) {
		data := strings.Split(line, ": ")
		testResult := aoc.Atoi(data[0])
		numbers := aoc.Map(strings.Split(data[1], " "), aoc.Atoi)
		startingNode := &aoc.TreeNode[int64]{Value: numbers[0]}

		result := getNumberOfPossibilities2(startingNode, numbers[1:], testResult)
		if result > 0 {
			sum += testResult
		}
	}

	fmt.Println(sum)
}
func getNumberOfPossibilities2(
	node *aoc.TreeNode[int64], numbers []int64, testResult int64,
) int {
	if len(numbers) == 0 {
		if node.Value == testResult {
			return 1
		} else {
			return 0
		}
	}

	node1 := &aoc.TreeNode[int64]{Value: node.Value * numbers[0]}
	node2 := &aoc.TreeNode[int64]{Value: node.Value + numbers[0]}
	node3 := &aoc.TreeNode[int64]{Value: aoc.Atoi(fmt.Sprintf("%d%d", node.Value, numbers[0]))}

	return getNumberOfPossibilities2(node1, numbers[1:], testResult) +
		getNumberOfPossibilities2(node2, numbers[1:], testResult) +
		getNumberOfPossibilities2(node3, numbers[1:], testResult)
}
