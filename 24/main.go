package main

import (
	"fmt"
	"os"
	"strings"
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

type gate struct {
	gateType string
	input1   string
	input2   string
	output   string
}

func part1() {
	t := time.Now()
	sections := aoc.SectionsFromFileAsString(os.Args[2])
	section1 := strings.Split(sections[0], "\n")
	currentValues := map[string]bool{}
	for _, line := range section1 {
		parts := strings.Split(line, ": ")
		currentValues[parts[0]] = aoc.IntsFromString(parts[1])[0] == 1
	}
	var gates []gate
	for _, line := range strings.Split(sections[1], "\n") {
		parts := strings.Split(line, " ")
		gates = append(gates, gate{
			gateType: parts[1],
			input1:   parts[0],
			input2:   parts[2],
			output:   parts[4],
		})
	}

	index := 0
	for {
		if index >= len(gates) {
			break
		}

		gate := gates[index]
		v1, ok1 := currentValues[gate.input1]
		v2, ok2 := currentValues[gate.input2]

		index++

		// wait for input
		if !ok1 || !ok2 {
			gates = append(gates, gate)
			continue
		}

		switch gate.gateType {
		case "AND":
			currentValues[gate.output] = v1 && v2
		case "OR":
			currentValues[gate.output] = v1 || v2
		case "XOR":
			currentValues[gate.output] = v1 != v2
		default:
			panic("Unknown gate type")
		}
	}

	aoc.Debug(currentValues)

	result := 0
	index = 0
	for {
		key := fmt.Sprintf("z%02d", index)
		v, ok := currentValues[key]
		if !ok {
			break
		}

		if v {
			result |= 1 << index
		}
		aoc.Debugf("%08b (%s=%v)", result, key, v)

		index++
	}

	aoc.Debug(currentValues)
	aoc.Debug(gates)

	aoc.Println(t, result)
}

func part2() {
	for line := range aoc.LinesFromFile(os.Args[2]) {
		fmt.Println(line)
	}
}
