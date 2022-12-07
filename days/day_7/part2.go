package main

import (
	"context"
	"strings"
)

// Part2 solves the second part of the day's puzzle
func Part2(ctx context.Context, input string) (interface{}, error) {
	root, err := Parse(strings.Split(strings.TrimSpace(input), "\n"))
	if err != nil {
		return nil, err
	}

	const diskSize = 70000000
	const requiredSize = 30000000
	totalSize := root.Size()

	requiredSpace := requiredSize - (diskSize - totalSize)

	smallestSuitableDirSize := totalSize // Starting value

	root.Traverse(func(c *Dir) error {
		size := c.Size()
		if size >= requiredSpace && size < smallestSuitableDirSize {
			smallestSuitableDirSize = size
		}
		return nil
	})

	return smallestSuitableDirSize, nil
}
