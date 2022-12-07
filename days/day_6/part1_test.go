package main_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	main "github.com/mbamber/aoc22/days/day_6"
	"github.com/mbamber/aoc22/inputs"
)

func TestPart1Answer(t *testing.T) {
	expected := 1140

	input := inputs.Load(6)
	out, err := main.Part1(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}

func TestPart1Example1(t *testing.T) {
	expected := 7

	input := inputs.LoadExample(6, 1)
	out, err := main.Part1(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}

func TestPart1Example2(t *testing.T) {
	expected := 5

	input := inputs.LoadExample(6, 2)
	out, err := main.Part1(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}

func TestPart1Example3(t *testing.T) {
	expected := 6

	input := inputs.LoadExample(6, 3)
	out, err := main.Part1(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}

func TestPart1Example4(t *testing.T) {
	expected := 10

	input := inputs.LoadExample(6, 4)
	out, err := main.Part1(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}

func TestPart1Example5(t *testing.T) {
	expected := 11

	input := inputs.LoadExample(6, 5)
	out, err := main.Part1(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}
