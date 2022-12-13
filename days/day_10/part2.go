package main

import (
	"context"
	"fmt"
	"math"
	"strings"
)

// Part2 solves the second part of the day's puzzle
func Part2(ctx context.Context, input string) (interface{}, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	var program []Instruction

	for _, line := range lines {
		i, err := ParseInstruction(line)
		if err != nil {
			return nil, err
		}
		program = append(program, i)
	}

	var crt string
	c := NewCPU(program, WithHook(StatusExecuting, func(c *CPU) error {
		cycle := int(math.Mod(float64(c.CurrentCycle()-1), 40))

		if cycle == 0 {
			crt += "\n"
		}

		if cycle == c.RegisterX() || cycle == c.RegisterX()-1 || cycle == c.RegisterX()+1 {
			crt += "#"
		} else {
			crt += "."
		}
		return nil
	}))

	for c.HasMoreInstructions() {
		err := c.Tick()
		if err != nil {
			return nil, err
		}
	}

	// Need to print the answer here so we can read the letters
	fmt.Println(crt)

	return crt, nil
}
