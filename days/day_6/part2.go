package main

import (
	"context"
	"errors"
)

// Part2 solves the second part of the day's puzzle
func Part2(ctx context.Context, input string) (interface{}, error) {
	last14 := make([]rune, 14)
	for i, r := range input {
		last14 = last14[1:]
		last14 = append(last14, r)

		if i < 4 {
			continue // Cant find marker until 4 runes checked
		}

		if !containsDuplicates(last14) {
			return i + 1, nil
		}
	}

	return nil, errors.New("input exhausted")
}
