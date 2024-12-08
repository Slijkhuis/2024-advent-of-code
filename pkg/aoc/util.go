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
func Debug(v ...any) {
	if debug {
		fmt.Println(v...)
	}
}

func Error(v ...any) {
	v = append([]any{"ERROR:"}, v...)
	fmt.Println(v...)
	os.Exit(1)
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

func StringFromFile(path string) string {
	b, err := os.ReadFile(path)
	fail(err)
	if b[len(b)-1] == '\n' {
		b = b[:len(b)-1]
	}

	return string(b)
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

func Atoi(str string) int64 {
	n, err := strconv.ParseInt(str, 10, 64)
	fail(err)
	return n
}

func Map[T, U any](collection []T, fn func(T) U) []U {
	result := make([]U, len(collection))
	for i := range collection {
		result[i] = fn(collection[i])
	}
	return result
}

func In[T comparable](collection []T, value T) bool {
	for i := range collection {
		if collection[i] == value {
			return true
		}
	}
	return false
}

func Abs[T ~int64](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

func ReverseString(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}
