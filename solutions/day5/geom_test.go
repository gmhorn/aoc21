package day5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLine(t *testing.T) {
	l, err := NewLine("1,2 -> 30,40")
	assert.NoError(t, err)
	assert.Equal(t, 1, l.X1)
	assert.Equal(t, 2, l.Y1)
	assert.Equal(t, 30, l.X2)
	assert.Equal(t, 40, l.Y2)
}

func TestBounds(t *testing.T) {
	actual := Bounds([]Line{{3, 15, 10, 11}})
	assert.Equal(t, Point{10, 15}, actual)
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

	assert.Equal(t, 5, CountOverlaps(lines, 2))
}
