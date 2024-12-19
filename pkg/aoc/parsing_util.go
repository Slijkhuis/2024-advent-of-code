package aoc

import (
	"bufio"
	"iter"
	"os"
	"regexp"
	"strings"
)

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

func StringFromFile(path string) string {
	b, err := os.ReadFile(path)
	fail(err)
	if b[len(b)-1] == '\n' {
		b = b[:len(b)-1]
	}

	return string(b)
}

func SectionsFromFileAsString(path string) []string {
	return strings.Split(StringFromFile(path), "\n\n")
}

var intsFromStringRegex = regexp.MustCompile(`-?\d+`)

func IntsFromString(s string) []int {
	return Map(intsFromStringRegex.FindAllString(s, -1), Atoi)
}
