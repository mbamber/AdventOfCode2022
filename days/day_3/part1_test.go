package main_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	main "github.com/mbamber/aoc22/days/day_3"
	"github.com/mbamber/aoc22/inputs"
)

func TestPart1Answer(t *testing.T) {
	expected := 8240

	input := inputs.Load(3)
	out, err := main.Part1(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}

func TestPart1Example1(t *testing.T) {
	expected := 157

	input := inputs.LoadExample(3, 1)
	out, err := main.Part1(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}
