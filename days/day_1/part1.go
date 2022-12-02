package main

import (
	"context"
	"sort"
	"strconv"
	"strings"
)

// Part1 solves the first part of the day's puzzle
func Part1(ctx context.Context, input string) (interface{}, error) {
	totals := []int{0}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			totals = append(totals, 0)
			continue
		}

		i, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		totals[len(totals)-1] += i
	}

	sort.Ints(totals)
	return totals[len(totals)-1], nil
}
