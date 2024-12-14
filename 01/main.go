package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/Slijkhuis/2024-advent-of-code/pkg/aoc"
)

func main() {
	aoc01a()
	aoc01b()
}

func aoc01b() {
	list1 := make([]int, 0)
	list2 := make([]int, 0)

	for line := range aoc.LinesFromFile(os.Args[1]) {
		cols := strings.Fields(line)
		if len(cols) != 2 {
			continue
		}
		i1, err := strconv.ParseInt(cols[0], 10, 64)
		fail(err)
		list1 = append(list1, int(i1))

		i2, err := strconv.ParseInt(cols[1], 10, 64)
		fail(err)
		list2 = append(list2, int(i2))
	}

	var result int
	for i := 0; i < len(list1); i++ {
		number := list1[i]
		result += number * aoc.Count(list2, number)
	}

	fmt.Println(result)
}

func aoc01a() {
	list1 := make([]int64, 0)
	list2 := make([]int64, 0)

	for line := range aoc.LinesFromFile(os.Args[1]) {
		cols := strings.Fields(line)
		if len(cols) != 2 {
			continue
		}
		i1, err := strconv.ParseInt(cols[0], 10, 64)
		fail(err)
		list1 = append(list1, i1)

		i2, err := strconv.ParseInt(cols[1], 10, 64)
		fail(err)
		list2 = append(list2, i2)
	}

	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})
	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})

	pointer2 := 0

	results := make([]int64, 0)

	for i := 0; i < len(list1); i++ {
		var diff int64
		if list1[i]-list2[pointer2] >= 0 {
			diff = list1[i] - list2[pointer2]
		} else {
			diff = list2[pointer2] - list1[i]
		}

		results = append(results, diff)

		if pointer2 < len(list2)-1 {
			pointer2++
		}
	}

	fmt.Println(aoc.Sum(results))
}

func fail(err error) {
	if err != nil {
		panic(err)
	}
}
