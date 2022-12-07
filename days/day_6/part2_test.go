package main_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	main "github.com/mbamber/aoc22/days/day_6"
	"github.com/mbamber/aoc22/inputs"
)

func TestPart2Answer(t *testing.T) {
	expected := 0

	input := inputs.Load(6)
	out, err := main.Part2(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}

func TestPart2Example1(t *testing.T) {
	expected := 19

	input := inputs.LoadExample(6, 1)
	out, err := main.Part2(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}

func TestPart2Example2(t *testing.T) {
	expected := 23

	input := inputs.LoadExample(6, 2)
	out, err := main.Part2(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}

func TestPart2Example3(t *testing.T) {
	expected := 23

	input := inputs.LoadExample(6, 3)
	out, err := main.Part2(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}

func TestPart2Example4(t *testing.T) {
	expected := 29

	input := inputs.LoadExample(6, 4)
	out, err := main.Part2(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}

func TestPart2Example5(t *testing.T) {
	expected := 26

	input := inputs.LoadExample(6, 5)
	out, err := main.Part2(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}
