package main

import (
	"context"
	"fmt"
	"strings"

	"golang.org/x/text/runes"
	"golang.org/x/text/unicode/rangetable"
)

// Part1 solves the first part of the day's puzzle
func Part1(ctx context.Context, input string) (interface{}, error) {
	var total int

line:
	for _, line := range strings.Split(input, "\n") {
		itemsPerRucksack := len(line) / 2
		r1, r2 := []rune(line[:itemsPerRucksack]), []rune(line[itemsPerRucksack:])

		set := runes.In(rangetable.New(r2...))

		for _, r := range r1 {
			if set.Contains(r) {
				p, err := FromRune(r)
				if err != nil {
					return nil, err
				}
				total += int(p)
				continue line
			}
		}
		return nil, fmt.Errorf("no common rune found between %s and %s", string(r1), string(r2))
	}

	return total, nil
}
