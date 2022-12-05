package main_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	main "github.com/mbamber/aoc22/days/day_5"
	"github.com/mbamber/aoc22/inputs"
)

func TestPart1Answer(t *testing.T) {
	expected := "VQZNJMWTR"

	input := inputs.Load(5)
	out, err := main.Part1(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}

func TestPart1Example1(t *testing.T) {
	expected := "CMZ"

	input := inputs.LoadExample(5, 1)
	out, err := main.Part1(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}
