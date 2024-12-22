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
	t := time.Now()

	var initialNumbers []int
	for line := range aoc.LinesFromFile(os.Args[2]) {
		initialNumbers = append(initialNumbers, aoc.Atoi(line))
	}

	numbers := make([][]int, len(initialNumbers))
	lastDigits := make([][]int, len(initialNumbers))
	changes := make([][]int, len(initialNumbers))

	var allKeys [][4]int
	pricePer4Changes := make([]map[[4]int]int, len(initialNumbers))

	for i, n := range initialNumbers {
		numbers[i] = make([]int, 2000)
		lastDigits[i] = make([]int, 2000)
		pricePer4Changes[i] = make(map[[4]int]int)

		for j := range 2000 {
			n = generateSecretNumber(n)
			numbers[i][j] = n
			lastDigits[i][j] = n % 10
			if j > 0 {
				changes[i] = append(changes[i], lastDigits[i][j]-lastDigits[i][j-1])
			}
			cs := changes[i]
			lcs := len(cs)
			if len(cs) >= 4 {
				key := [4]int{cs[lcs-4], cs[lcs-3], cs[lcs-2], cs[lcs-1]}

				// TODO: not sure if I have to overwrite the value if the newer occurrence has a higher price? Reading the
				// puzzle text I don't think so, but I need to double check.
				if _, ok := pricePer4Changes[i][key]; !ok {
					pricePer4Changes[i][key] = lastDigits[i][j]
				}

				allKeys = append(allKeys, key)
			}
		}
	}

	aoc.Println(t, "processed all changes")

	allUniqueKeys := aoc.Unique(allKeys)
	aoc.Println(t, "processed unique keys, went from", len(allKeys), "to", len(allUniqueKeys))

	// Find the best key.
	var highestKey [4]int
	var highestValue int
	for i, key := range allUniqueKeys {
		if (i+1)%10_000 == 0 {
			aoc.Println(t, "processing", i+1, "out of", len(allUniqueKeys))
		}

		var totalPrice int
		for i := range len(pricePer4Changes) {
			totalPrice += pricePer4Changes[i][key]
		}

		if totalPrice > highestValue {
			highestKey = key
			highestValue = totalPrice
		}
	}

	aoc.Debug(highestKey)
	aoc.Println(t, highestValue)
}
