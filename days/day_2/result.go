package main

import "fmt"

type Result int

const (
	Win  Result = 6
	Lose Result = 0
	Draw Result = 3
)

func ParseResult(s string) (Result, error) {
	switch s {
	case "X":
		return Lose, nil
	case "Y":
		return Draw, nil
	case "Z":
		return Win, nil
	default:
		return Win, fmt.Errorf("unknown result %s", s)
	}
}
