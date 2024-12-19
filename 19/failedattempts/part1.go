package failedattempts

import (
	"fmt"
	"os"
	"strings"

	"github.com/Slijkhuis/2024-advent-of-code/pkg/aoc"
)

func Part1WorksForTestButIsTooSlowForRealInput() {
	dataSections := aoc.SectionsFromFileAsString(os.Args[2])
	towels := strings.Split(dataSections[0], ", ")
	patterns := strings.Split(dataSections[1], "\n")

	possible := 0
	for i, pattern := range patterns {
		fmt.Println(pattern, i+1, "out of", len(patterns))

		options := map[string]struct{}{}

		for position := range len(pattern) {
			aoc.Debug(position, options)

			if position == 0 {
				for _, towel := range towels {
					if towel[0] == pattern[0] {
						options[towel] = struct{}{}
					}
				}
				continue
			}

			color := pattern[position]
			for option := range options {
				optionTowels := strings.Split(option, ",")
				optionPattern := strings.Join(optionTowels, "")

				if len(optionPattern) == position {
					for _, towel := range towels {
						if towel[0] == pattern[position] {
							delete(options, option)

							tmp := optionPattern + towel
							if len(tmp) <= len(pattern) && pattern[:len(tmp)] == tmp {
								options[option+","+towel] = struct{}{}
							}
						}
					}
				} else {
					if optionPattern[position] != color {
						delete(options, option)
					}
				}
			}
		}
		aoc.Debug(options)
		for option := range options {
			optionTowels := strings.Split(option, ",")
			optionPattern := strings.Join(optionTowels, "")
			if optionPattern == pattern {
				possible++
				break
			}
		}
	}

	fmt.Println(possible)
}
