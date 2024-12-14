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
	stones := aoc.Map(strings.Split(aoc.StringFromFile(os.Args[2]), " "), aoc.Atoi)
	aoc.Debug(stones)

	for range 25 {
		newStones := make([]int, 0, len(stones))
		for _, s := range stones {
			if s == 0 {
				newStones = append(newStones, 1)
				continue
			}
			str := fmt.Sprintf("%d", s)
			if len(str)%2 == 0 {
				a := str[:len(str)/2]
				b := str[len(str)/2:]
				newStones = append(newStones, aoc.Atoi(a))
				newStones = append(newStones, aoc.Atoi(b))
				continue
			}
			newStones = append(newStones, s*2024)
		}
		stones = newStones
	}

	fmt.Println(len(stones))
}

func part2() {
	stones := aoc.Map(strings.Split(aoc.StringFromFile(os.Args[2]), " "), aoc.Atoi)
	aoc.Debug(stones)

	var totalNumberOfStones int
	for _, s := range stones {
		totalNumberOfStones += amountOfStonesAfterBlinksForSingleStone(s, 75)
	}

	fmt.Println(totalNumberOfStones)
}

type cacheKeyDay11 struct {
	number int
	blinks int
}

var cacheDay11 = make(map[cacheKeyDay11]int)

func cacheAmountOfStones(key cacheKeyDay11, value int) int {
	cacheDay11[key] = value
	return value
}

func amountOfStonesAfterBlinksForSingleStone(numberOnStone int, blinks int) int {
	key := cacheKeyDay11{numberOnStone, blinks}
	if v, ok := cacheDay11[key]; ok {
		return v
	}

	if blinks == 0 {
		return cacheAmountOfStones(key, 1)
	}
	if numberOnStone == 0 {
		return cacheAmountOfStones(key, amountOfStonesAfterBlinksForSingleStone(1, blinks-1))
	}
	str := fmt.Sprintf("%d", numberOnStone)
	if len(str)%2 == 0 {
		a := aoc.Atoi(str[:len(str)/2])
		b := aoc.Atoi(str[len(str)/2:])
		return cacheAmountOfStones(
			key,
			amountOfStonesAfterBlinksForSingleStone(a, blinks-1)+amountOfStonesAfterBlinksForSingleStone(b, blinks-1),
		)
	}
	return cacheAmountOfStones(
		key,
		amountOfStonesAfterBlinksForSingleStone(numberOnStone*2024, blinks-1),
	)
}
