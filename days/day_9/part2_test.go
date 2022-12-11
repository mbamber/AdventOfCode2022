package main_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	main "github.com/mbamber/aoc22/days/day_9"
	"github.com/mbamber/aoc22/inputs"
)

func TestPart2Answer(t *testing.T) {
	expected := 2485

	input := inputs.Load(9)
	out, err := main.Part2(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}

func TestPart2Example1(t *testing.T) {
	expected := 1

	input := inputs.LoadExample(9, 1)
	out, err := main.Part2(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}

func TestPart2Example2(t *testing.T) {
	expected := 36

	input := inputs.LoadExample(9, 2)
	out, err := main.Part2(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}
