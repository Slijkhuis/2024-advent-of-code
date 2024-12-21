package aoc

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func fail(err error) {
	if err != nil {
		panic(err)
	}
}

var DebugMode = false

func init() {
	DebugMode = os.Getenv("DEBUG") == "1"
}

// Debug input if DEBUG=1 is set.
func Debug(v ...any) {
	if DebugMode {
		fmt.Println(v...)
	}
}

func Debugf(format string, v ...any) {
	if DebugMode {
		if len(format) > 0 && format[len(format)-1] != '\n' {
			format += "\n"
		}
		fmt.Printf(format, v...)
	}
}

func Println(t time.Time, v ...any) {
	fmt.Println(append([]any{time.Since(t).Round(time.Millisecond)}, v...)...)
}

func Error(v ...any) {
	v = append([]any{"ERROR:"}, v...)
	fmt.Println(v...)
	os.Exit(1)
}

func Atoi(str string) int {
	n, err := strconv.ParseInt(str, 10, 64)
	fail(err)
	return int(n)
}

func Map[T, U any](collection []T, fn func(T) U) []U {
	result := make([]U, len(collection))
	for i := range collection {
		result[i] = fn(collection[i])
	}
	return result
}

func Abs[T ~int64 | ~int](n T) T {
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

func Min[T ~int64 | ~int](a, b T) T {
	if b < a {
		return b
	}
	return a
}

func Max[T ~int64 | ~int](a, b T) T {
	if a >= b {
		return a
	}
	return b
}

func Fail(err error) {
	if err != nil {
		panic(err)
	}
}

func CopyMap[K comparable, V any](m map[K]V) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}

func JoinRunes(rs []rune, sep string) string {
	return strings.Join(
		Map(rs, func(r rune) string { return string(r) }),
		sep,
	)
}
