package main_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	main "github.com/mbamber/aoc22/days/day_2"
	"github.com/mbamber/aoc22/inputs"
)

func TestPart1Answer(t *testing.T) {
	expected := 10624

	input := inputs.Load(2)
	out, err := main.Part1(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}

func TestPart1Example1(t *testing.T) {
	expected := 15

	input := inputs.LoadExample(2, 1)
	out, err := main.Part1(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}
