package main

import (
	"context"
	"errors"
)

// Part1 solves the first part of the day's puzzle
func Part1(ctx context.Context, input string) (interface{}, error) {
	last4 := make([]rune, 4)
	for i, r := range input {
		last4 = last4[1:]
		last4 = append(last4, r)

		if i < 4 {
			continue // Cant find marker until 4 runes checked
		}

		if !containsDuplicates(last4) {
			return i + 1, nil
		}
	}

	return nil, errors.New("input exhausted")
}
