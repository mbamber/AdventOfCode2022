package main

import (
	"context"
	"strconv"
	"strings"
)

// Part2 solves the second part of the day's puzzle
func Part2(ctx context.Context, input string) (interface{}, error) {
	rope := []*Point{}
	for i := 0; i < 10; i++ {
		rope = append(rope, &Point{0, 0})
	}

	locations := Locations{rope[9].Copy()}

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
				rope[0].Y++
			case "D":
				rope[0].Y--
			case "L":
				rope[0].X--
			case "R":
				rope[0].X++
			}

			for i := 1; i < 10; i++ {
				if !rope[i-1].IsAdjacent(rope[i]) {
					rope[i].MoveTowards(rope[i-1])
				}
			}

			locations = append(locations, rope[9].Copy())
		}
	}

	return locations.CountUnique(), nil
}
