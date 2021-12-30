package day09

import (
	"testing"

	"github.com/gmhorn/aoc21/lib"
	"github.com/stretchr/testify/assert"
)

func TestMinima(t *testing.T) {
	g, err := lib.LoadGrid("testdata.txt")
	assert.NoError(t, err)

	expect := []uint8{0, 1, 5, 5}
	actual := minima(g)

	assert.Equal(t, expect, actual)
}

func TestBasins(t *testing.T) {
	g, err := lib.LoadGrid("testdata.txt")
	assert.NoError(t, err)

	expected := []int{14, 9, 9, 3}
	actual := basins(g, 9)

	assert.Equal(t, expected, actual)
}
