package main

import (
	"fmt"
)

type Action int

func (a Action) Vs(v Action) Result {
	if a == v {
		return Draw
	}

	if a == Rock && v == Scissors {
		return Win
	}

	if a == Paper && v == Rock {
		return Win
	}

	if a == Scissors && v == Paper {
		return Win
	}

	return Lose
}

func (a Action) ToGetResult(r Result) Action {
	switch r {
	case Win:
		switch a {
		case Rock:
			return Paper
		case Paper:
			return Scissors
		case Scissors:
			return Rock
		}
	case Lose:
		switch a {
		case Rock:
			return Scissors
		case Paper:
			return Rock
		case Scissors:
			return Paper
		}
	case Draw:
		return a
	}

	panic("something went wrong...")
}

const (
	Rock Action = iota + 1
	Paper
	Scissors
)

func ParseAction(s string) (Action, error) {
	switch s {
	case "A", "X":
		return Rock, nil
	case "B", "Y":
		return Paper, nil
	case "C", "Z":
		return Scissors, nil
	default:
		return Rock, fmt.Errorf("unknown action %s", s)
	}
}
