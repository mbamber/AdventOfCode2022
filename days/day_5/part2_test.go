package main_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	main "github.com/mbamber/aoc22/days/day_5"
	"github.com/mbamber/aoc22/inputs"
)

func TestPart2Answer(t *testing.T) {
	expected := "NLCDCLVMQ"

	input := inputs.Load(5)
	out, err := main.Part2(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}

func TestPart2Example1(t *testing.T) {
	expected := "MCD"

	input := inputs.LoadExample(5, 1)
	out, err := main.Part2(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}
