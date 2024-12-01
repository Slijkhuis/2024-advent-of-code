package aoc

import (
	"bufio"
	"fmt"
	"iter"
	"os"
	"strconv"
)

func fail(err error) {
	if err != nil {
		panic(err)
	}
}

var debug = false

func init() {
	debug = os.Getenv("DEBUG") == "1"
}

// Debug input if DEBUG=1 is set.
func Debug(v any) {
	if debug {
		fmt.Println(v)
	}
}

func LinesFromFile(path string) iter.Seq[string] {
	f, err := os.Open(path)
	fail(err)

	scanner := bufio.NewScanner(f)

	return func(yield func(string) bool) {
		for scanner.Scan() {
			if !yield(scanner.Text()) {
				return
			}
		}

		fail(scanner.Err())
		fail(f.Close())
	}
}

func Count[T comparable](collection []T, value T) (count int64) {
	for i := range collection {
		if collection[i] == value {
			count++
		}
	}

	return count
}

func Sum[T ~float32 | ~float64 | ~int | ~int64](collection []T) T {
	sum := T(0)
	for i := range collection {
		sum += collection[i]
	}
	return sum
}
