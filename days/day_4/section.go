package main

import (
	"errors"
	"strconv"
	"strings"
)

type Section struct {
	Start, End int
}

func ParseSection(s string) (Section, error) {
	sec := Section{}

	parts := strings.Split(s, "-")
	if len(parts) != 2 {
		return sec, errors.New("unable to parse section")
	}

	var err error
	sec.Start, err = strconv.Atoi(parts[0])
	if err != nil {
		return sec, err
	}
	sec.End, err = strconv.Atoi(parts[1])
	if err != nil {
		return sec, err
	}

	return sec, nil

}
