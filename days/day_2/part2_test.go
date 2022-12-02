package main_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	main "github.com/mbamber/aoc22/days/day_2"
	"github.com/mbamber/aoc22/inputs"
)

func TestPart2Answer(t *testing.T) {
	expected := 14060

	input := inputs.Load(2)
	out, err := main.Part2(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}

func TestPart2Example1(t *testing.T) {
	expected := 12

	input := inputs.LoadExample(2, 1)
	out, err := main.Part2(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}
