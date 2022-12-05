package main

import (
	"context"
)

// Part1 solves the first part of the day's puzzle
func Part1(ctx context.Context, input string) (interface{}, error) {
	stacks, instructions := ParseInput(input)

	// Process all instructions
	for _, instruction := range instructions {
		for i := 0; i < instruction.Count; i++ {
			toIdx, fromIdx := instruction.To-1, instruction.From-1
			stacks[toIdx] = append(stacks[toIdx], stacks[fromIdx][len(stacks[fromIdx])-1])
			stacks[fromIdx] = stacks[fromIdx][:len(stacks[fromIdx])-1]
		}
	}

	var s string
	for _, stack := range stacks {
		if len(stack) <= 0 {
			continue
		}
		s = s + stack[len(stack)-1]
	}

	return s, nil
}
