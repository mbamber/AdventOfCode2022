package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	main "github.com/mbamber/aoc22/days/day_3"
)

func TestFromRune(t *testing.T) {
	cases := map[string]struct {
		r        rune
		expected main.Priority
	}{
		"a is 1": {
			r:        'a',
			expected: 1,
		},
		"z is 26": {
			r:        'z',
			expected: 26,
		},
		"A is 27": {
			r:        'A',
			expected: 27,
		},
		"Z is 52": {
			r:        'Z',
			expected: 52,
		},
	}

	for name, data := range cases {
		name, data := name, data
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			p, err := main.FromRune(data.r)
			assert.NoError(t, err)
			assert.Equal(t, data.expected, p)
		})
	}
}
