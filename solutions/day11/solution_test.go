package day11

import (
	"io/ioutil"
	"testing"

	"github.com/gmhorn/aoc21/lib"
	"github.com/stretchr/testify/assert"
)

func TestAdvance(t *testing.T) {
	grid, err := lib.LoadGrid("testdata/init.txt")
	assert.NoError(t, err)

	advance(grid, 9)
	expected, err := ioutil.ReadFile("testdata/step1.txt")
	assert.NoError(t, err)

	assert.Equal(t, string(expected), grid.String(), "step 1 failed")

	advance(grid, 9)
	expected, err = ioutil.ReadFile("testdata/step2.txt")
	assert.NoError(t, err)

	assert.Equal(t, string(expected), grid.String(), "step 2 failed")

	advance(grid, 9)
	expected, err = ioutil.ReadFile("testdata/step3.txt")
	assert.NoError(t, err)

	assert.Equal(t, string(expected), grid.String(), "step 3 failed")
}
