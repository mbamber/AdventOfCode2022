package main

import (
	"context"
	"log"
	"strings"
)

// Part2 solves the second part of the day's puzzle
func Part2(ctx context.Context, input string) (interface{}, error) {
	games := strings.Split(input, "\n")

	var total int

	for _, game := range games {
		choices := strings.Fields(game)
		opp, err := ParseAction(choices[0])
		if err != nil {
			return nil, err
		}
		res, err := ParseResult(choices[1])
		if err != nil {
			return nil, err
		}

		me := opp.ToGetResult(res)
		log.Println(res)
		log.Println(me)
		total += int(me.Vs(opp)) + int(me)
	}

	return total, nil
}
