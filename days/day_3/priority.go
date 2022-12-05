package main

import "errors"

var ErrNotAPriority = errors.New("not a rune")

type Priority int

func FromRune(r rune) (Priority, error) {
	if r >= 'A' && r <= 'Z' {
		return Priority(r - 38), nil
	}

	if r >= 'a' && r <= 'z' {
		return Priority(r - 96), nil
	}

	return 0, ErrNotAPriority
}
