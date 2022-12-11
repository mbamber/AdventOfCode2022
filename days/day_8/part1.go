package main

import (
	"context"
	"strconv"
	"strings"
)

// Part1 solves the first part of the day's puzzle
func Part1(ctx context.Context, input string) (interface{}, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	var count int
	for y, line := range lines {
		for x, c := range line {
			cint, err := strconv.Atoi(string(c))
			if err != nil {
				return nil, err
			}

			if x == 0 || y == 0 || x == len(line)-1 || y == len(lines)-1 {
				count++
				continue
			}

			ok := true

			// Check the row
			for x1 := 0; x1 < x; x1++ {
				xint, err := strconv.Atoi(string(lines[y][x1]))
				if err != nil {
					return nil, err
				}
				if xint >= cint {
					ok = false
					break
				}
			}
			if ok {
				count++
				continue
			}

			ok = true
			for x2 := x + 1; x2 < len(line); x2++ {
				xint, err := strconv.Atoi(string(lines[y][x2]))
				if err != nil {
					return nil, err
				}
				if xint >= cint {
					ok = false
					break
				}
			}
			if ok {
				count++
				continue
			}

			// Check the column
			ok = true
			for y1 := 0; y1 < y; y1++ {
				yint, err := strconv.Atoi(string(lines[y1][x]))
				if err != nil {
					return nil, err
				}
				if yint >= cint {
					ok = false
					break
				}
			}
			if ok {
				count++
				continue
			}

			ok = true
			for y2 := y + 1; y2 < len(lines); y2++ {
				yint, err := strconv.Atoi(string(lines[y2][x]))
				if err != nil {
					return nil, err
				}
				if yint >= cint {
					ok = false
					break
				}
			}
			if ok {
				count++
				continue
			}
		}
	}

	return count, nil
}
