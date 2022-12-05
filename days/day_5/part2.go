package main

import (
	"context"
)

// Part2 solves the second part of the day's puzzle
func Part2(ctx context.Context, input string) (interface{}, error) {
	stacks, instructions := ParseInput(input)

	// Process all instructions
	for _, instruction := range instructions {
		toIdx, fromIdx := instruction.To-1, instruction.From-1
		stacks[toIdx] = append(stacks[toIdx], stacks[fromIdx][len(stacks[fromIdx])-instruction.Count:]...)
		stacks[fromIdx] = stacks[fromIdx][:len(stacks[fromIdx])-instruction.Count]
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
