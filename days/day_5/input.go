package main

import (
	"regexp"
	"strconv"
	"strings"
)

type Instruction struct {
	Count, From, To int
}

func ParseInput(input string) ([][]string, []Instruction) {
	lines := strings.Split(input, "\n")

	// Only works for up to 9 stacks
	stacks := make([][]string, 9)
	var instructions []Instruction

	// Sanity check the max stacks
	stackIndexesRe := regexp.MustCompile(`^(?:\s+(\d+)\s*)+$`)
	for _, line := range lines {
		matches := stackIndexesRe.FindStringSubmatch(line)
		if len(matches) != 0 {
			stacks, err := strconv.Atoi(matches[len(matches)-1])
			if err != nil {
				panic(err)
			}

			if stacks > 9 {
				panic("more than 9 max stacks")
			}
		}
	}

	instructionRegex := regexp.MustCompile(`^move (\d+) from (\d) to (\d)$`)
	for _, line := range lines {
		stackMatches := parseAsStackLine(line)
		if len(stackMatches) != 0 {
			// We're looking at a starting arrangement
			for i, match := range stackMatches {
				if strings.TrimSpace(match) == "" {
					continue
				}

				// Need to prepend here
				stacks[i] = append([]string{match}, stacks[i]...)
			}

			continue
		}

		instructionMatches := instructionRegex.FindStringSubmatch(line)
		if len(instructionMatches) != 0 {
			// We're looking at an instruction
			var err error

			ins := Instruction{}
			for i, match := range instructionMatches {
				switch i {
				case 1:
					ins.Count, err = strconv.Atoi(match)
					if err != nil {
						panic(err)
					}
				case 2:
					ins.From, err = strconv.Atoi(match)
					if err != nil {
						panic(err)
					}
				case 3:
					ins.To, err = strconv.Atoi(match)
					if err != nil {
						panic(err)
					}
				}
			}
			instructions = append(instructions, ins)
			continue
		}

		// We matched something else (possibly the index list)
	}

	return stacks, instructions
}

func parseAsStackLine(l string) []string {
	var i int
	var items []string

	if !strings.ContainsRune(l, '[') {
		return items
	}

	for {
		if i+1 >= len(l) {
			return items
		}
		items = append(items, string(l[i+1]))

		i += 4
	}
}
