package main

import (
	"context"
	"strings"
)

// Part1 solves the first part of the day's puzzle
func Part1(ctx context.Context, input string) (interface{}, error) {
	games := strings.Split(input, "\n")

	var total int

	for _, game := range games {
		choices := strings.Fields(game)
		opp, err := ParseAction(choices[0])
		if err != nil {
			return nil, err
		}
		me, err := ParseAction(choices[1])
		if err != nil {
			return nil, err
		}

		total += int(me.Vs(opp)) + int(me)
	}

	return total, nil
}
