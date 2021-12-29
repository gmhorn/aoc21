package day05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLine(t *testing.T) {
	actual, err := ParseLine("1,2 -> 30,40")
	assert.NoError(t, err)
	assert.Equal(t, Line{1, 2, 30, 40}, actual)
}

func TestCountOverlaps(t *testing.T) {
	lines := []Line{
		{0, 9, 5, 9},
		{8, 0, 0, 8},
		{9, 4, 3, 4},
		{2, 2, 2, 1},
		{7, 0, 7, 4},
		{6, 4, 2, 0},
		{0, 9, 2, 9},
		{3, 4, 1, 4},
		{0, 0, 8, 8},
		{5, 5, 8, 2},
	}

	assert.Equal(t, 12, CountOverlaps(lines, 2))
}
