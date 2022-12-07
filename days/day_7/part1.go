package main

import (
	"context"
	"strings"
)

// Part1 solves the first part of the day's puzzle
func Part1(ctx context.Context, input string) (interface{}, error) {
	root, err := Parse(strings.Split(strings.TrimSpace(input), "\n"))
	if err != nil {
		return nil, err
	}

	var total int
	err = root.Traverse(func(c *Dir) error {
		size := c.Size()
		if size < 100000 {
			total += size
		}
		return nil
	})
	if err != nil {
		return "", err
	}

	return total, nil
}
