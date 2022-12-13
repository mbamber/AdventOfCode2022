package main

import (
	"context"
	"math"
	"sort"
	"strings"
)

// Part2 solves the second part of the day's puzzle
func Part2(ctx context.Context, input string) (interface{}, error) {
	monkeyStrings := strings.Split(strings.TrimSpace(input), "\n\n")

	var monkeys []*Monkey
	for _, monkey := range monkeyStrings {
		m, err := Parse(monkey)
		if err != nil {
			return nil, err
		}
		monkeys = append(monkeys, m)
	}

	// Calculate the modulo
	mod := 1
	for _, monkey := range monkeys {
		mod *= monkey.GetTestCondition()
	}

	// Control the worry
	for _, monkey := range monkeys {
		monkey.SetPostInspect(func(m *Monkey) error {
			m.UpdateCurrentWorry(int(math.Mod(float64(m.GetCurrentWorry()), float64(mod))))
			return nil
		})
	}

	for round := 0; round < 10000; round++ {
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
