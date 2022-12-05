package main

import (
	"context"
	"strings"
)

// Part2 solves the second part of the day's puzzle
func Part2(ctx context.Context, input string) (interface{}, error) {
	input = strings.TrimSpace(input)
	var overlapCount int

	pairs := strings.Split(input, "\n")

	for _, pair := range pairs {
		sections := strings.Split(pair, ",")
		section1, err := ParseSection(sections[0])
		if err != nil {
			return nil, err
		}
		section2, err := ParseSection(sections[1])
		if err != nil {
			return nil, err
		}

		if section1.Start >= section2.Start && section1.Start <= section2.End {
			overlapCount++
			continue
		}

		if section1.End >= section2.Start && section1.End <= section2.End {
			overlapCount++
			continue
		}

		if section2.Start >= section1.Start && section2.Start <= section1.End {
			overlapCount++
			continue
		}

		if section2.End >= section1.Start && section2.End <= section1.End {
			overlapCount++
			continue
		}
	}

	return overlapCount, nil
}
