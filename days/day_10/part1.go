package main

import (
	"context"
	"math"
	"strings"
)

// Part1 solves the first part of the day's puzzle
func Part1(ctx context.Context, input string) (interface{}, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	var program []Instruction

	for _, line := range lines {
		i, err := ParseInstruction(line)
		if err != nil {
			return nil, err
		}
		program = append(program, i)
	}

	var signalStrength int
	c := NewCPU(program, WithHook(StatusExecuting, func(c *CPU) error {
		if math.Mod(float64(c.CurrentCycle())+20, 40) == 0 || c.CurrentCycle() == 20 {
			signalStrength += (c.CurrentCycle() * c.RegisterX())
		}
		return nil
	}))

	for c.HasMoreInstructions() {
		err := c.Tick()
		if err != nil {
			return nil, err
		}
	}

	return signalStrength, nil
}
