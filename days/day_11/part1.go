package main

import (
	"context"
	"sort"
	"strings"
)

// Part1 solves the first part of the day's puzzle
func Part1(ctx context.Context, input string) (interface{}, error) {
	monkeyStrings := strings.Split(strings.TrimSpace(input), "\n\n")

	var monkeys []*Monkey
	for _, monkey := range monkeyStrings {
		m, err := Parse(monkey)
		if err != nil {
			return nil, err
		}
		m.SetPostInspect(func(m *Monkey) error {
			m.UpdateCurrentWorry(m.GetCurrentWorry() / 3)
			return nil
		})
		monkeys = append(monkeys, m)
	}

	for round := 0; round < 20; round++ {
		for _, monkey := range monkeys {
			for monkey.HasItemsToInspect() {
				monkey.Inspect()
				item, toMonkey := monkey.Throw()
				monkeys[toMonkey].Catch(item)
			}
		}
	}

	var activities []int
	for _, monkey := range monkeys {
		activities = append(activities, monkey.Activity())
	}

	sort.Ints(activities)

	return activities[len(activities)-1] * activities[len(activities)-2], nil
}
