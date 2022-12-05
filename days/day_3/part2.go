package main

import (
	"context"
	"strings"

	"golang.org/x/text/runes"
	"golang.org/x/text/unicode/rangetable"
)

// Part2 solves the second part of the day's puzzle
func Part2(ctx context.Context, input string) (interface{}, error) {
	input = strings.TrimSpace(input)
	var total int

	lines := strings.Split(input, "\n")

lines:
	for i := 0; i+3 <= len(lines); i += 3 {
		r1, r2, r3 := []rune(lines[i]), []rune(lines[i+1]), []rune(lines[i+2])

		r1Set := runes.In(rangetable.New(r1...))
		r2Set := runes.In(rangetable.New(r2...))

		for _, r := range r3 {
			if r1Set.Contains(r) && r2Set.Contains(r) {
				p, err := FromRune(r)
				if err != nil {
					return nil, err
				}
				total += int(p)
				continue lines
			}
		}
	}

	return total, nil
}
