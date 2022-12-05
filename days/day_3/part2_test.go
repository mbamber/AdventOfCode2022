package main_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	main "github.com/mbamber/aoc22/days/day_3"
	"github.com/mbamber/aoc22/inputs"
)

func TestPart2Answer(t *testing.T) {
	expected := 2587

	input := inputs.Load(3)
	out, err := main.Part2(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}

func TestPart2Example1(t *testing.T) {
	expected := 70

	input := inputs.LoadExample(3, 1)
	out, err := main.Part2(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}
