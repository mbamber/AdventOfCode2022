package main

import (
	"context"
	"strconv"
	"strings"
)

// Part1 solves the first part of the day's puzzle
func Part1(ctx context.Context, input string) (interface{}, error) {
	head, tail := &Point{0, 0}, &Point{0, 0}

	locations := Locations{tail.Copy()}

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		fields := strings.Fields(line)

		dir := fields[0]
		n, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, err
		}

		for i := 0; i < n; i++ {
			switch dir {
			case "U":
				head.Y++
			case "D":
				head.Y--
			case "L":
				head.X--
			case "R":
				head.X++
			}

			if !head.IsAdjacent(tail) {
				tail.MoveTowards(head)
			}

			locations = append(locations, tail.Copy())
		}
	}

	return locations.CountUnique(), nil
}
