package day9

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadGrid(t *testing.T) {
	g, err := LoadGrid("testdata.txt")
	assert.NoError(t, err)
	assert.Equal(t, 4, g.YMax)
	assert.Equal(t, 9, g.XMax)
}

func TestMinima(t *testing.T) {
	g, err := LoadGrid("testdata.txt")
	assert.NoError(t, err)

	expected := []int8{0, 1, 5, 5}
	actual := g.Minima()

	assert.Equal(t, expected, actual)
}

func TestBasins(t *testing.T) {
	g, err := LoadGrid("testdata.txt")
	assert.NoError(t, err)

	expected := []int{14, 9, 9, 3}
	actual := g.Basins()

	assert.Equal(t, expected, actual)
}
