package day9

import (
	"sort"
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
	sort.Slice(actual, func(i, j int) bool {
		return actual[i] < actual[j]
	})

	assert.Equal(t, expected, actual)
}
