package main

import (
	"fmt"
	"os"
	"strconv"
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
	data := aoc.StringFromFile(os.Args[2])
	parts := strings.Split(data, "\n\n")
	registers := strings.Split(parts[0], "\n")
	registerA := aoc.IntsFromString(registers[0])[0]
	registerB := aoc.IntsFromString(registers[1])[0]
	registerC := aoc.IntsFromString(registers[2])[0]
	program := aoc.IntsFromString(parts[1])

	aoc.Debugf("Registers: %d %d %d", registerA, registerB, registerC)
	aoc.Debugf("Program: %v", program)

	result, _, _, _ := executeProgram(program, registerA, registerB, registerC)

	fmt.Println(strings.Join(aoc.Map(result, strconv.Itoa), ","))
}

func executeProgram(program []int, a, b, c int) ([]int, int, int, int) {
	combo := func(operand int) int {
		if operand <= 3 {
			return operand
		}

		switch operand {
		case 4:
			return a
		case 5:
			return b
		case 6:
			return c
		default:
			panic("Invalid operand")
		}
	}

	dv := func(operand int) int {
		return (a / (1 << combo(operand)))
	}

	var result []int

	for i := 0; i < len(program); i += 2 {
		instruction := program[i]
		operand := program[i+1]

		aoc.Debugf("Instruction: %d, Operand: %d\n", instruction, operand)

		switch instruction {
		case 0: // adv
			a = dv(operand)
		case 1: // bxl
			b = b ^ operand
		case 2: // bst
			b = combo(operand) % 8
		case 3: // jnz
			if a == 0 {
				continue
			}
			i = operand - 2
		case 4: // bxc
			b = b ^ c
		case 5: // out
			result = append(result, combo(operand)%8)
		case 6: // bdv
			b = dv(operand)
		case 7: // cdv
			c = dv(operand)
		default:
			panic("Invalid instruction")
		}
	}

	return result, a, b, c
}

func part2() {
	for line := range aoc.LinesFromFile(os.Args[2]) {
		fmt.Println(line)
	}
}
