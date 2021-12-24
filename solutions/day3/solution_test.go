package day3

import (
	"testing"

	"github.com/gmhorn/aoc21/lib"
	"github.com/stretchr/testify/assert"
)

func TestGreeks(t *testing.T) {
	tests := []struct {
		name     string
		fn       func([]int) (int64, error)
		expected int
	}{{
		name:     "epsilon",
		fn:       epsilon,
		expected: 9,
	}, {
		name:     "gamma",
		fn:       gamma,
		expected: 22,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := lib.ReadLines("testdata.txt")
			assert.NoError(t, err)

			acc, err := accumulate(lines)
			assert.NoError(t, err)

			actual, err := tt.fn(acc)
			assert.NoError(t, err)

			assert.Equal(t, tt.expected, int(actual))
		})
	}
}

func TestPart1(t *testing.T) {
	s := Solution{}
	actual, err := s.Part1("testdata.txt")
	assert.NoError(t, err)
	assert.Equal(t, 198, actual)
}
