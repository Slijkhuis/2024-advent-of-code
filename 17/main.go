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
	data := aoc.StringFromFile(os.Args[2])
	parts := strings.Split(data, "\n\n")
	program := aoc.IntsFromString(parts[1])
	result, a, b, c := executeProgramCustom(59590048, 0, 0)
	fmt.Println(result, a, b, c)
	result, a, b, c = executeProgram(program, 59590048, 0, 0)
	fmt.Println(result, a, b, c)

	type Attempt struct {
		a_guess  int
		position int
	}

	attempts := []Attempt{
		{
			a_guess:  0,
			position: 0,
		},
	}

	for len(attempts) > 0 {
		attempt := attempts[0]
		attempts = attempts[1:]

		if attempt.position > len(program) {
			fmt.Println("FINAL ANSWER:")
			fmt.Println(attempt.a_guess)
			return
		}

		for i := 0; i < 8; i++ {
			guess := (attempt.a_guess << 3) | i
			fmt.Printf("%08b\n", guess)

			result, a, b, c = executeProgramCustom(guess, 0, 0)
			fmt.Println(result, a, b, c)

			shouldMatch := program[len(program)-attempt.position:]

			matches := len(shouldMatch) == len(result)
			if matches {
				for i := range shouldMatch {
					if shouldMatch[i] != result[i] {
						matches = false
						break
					}
				}
			}

			if matches {
				fmt.Println("MATCH", attempt.position)
				attempts = append(attempts, Attempt{
					a_guess:  guess,
					position: attempt.position + 1,
				})
			}
		}
	}

	fmt.Println("NO ANSWER :(.")
}

// executeProgramCustom implements my input as a Go function to make the algorithm more readable.
// I manually went through the instructions and translated them to Go code.
func executeProgramCustom(a, b, c int) ([]int, int, int, int) {
	var result []int
	for a != 0 {
		// 2,4
		b = a % 8
		// 1,5
		b = b ^ 5
		// 7,5
		c = a / (1 << b)
		// 0,3
		a = a / 8
		// 1,6
		b = b ^ 6
		// 4,3
		b = b ^ c
		// 5,5
		result = append(result, b%8)
	} // 3,0

	return result, a, b, c
}
