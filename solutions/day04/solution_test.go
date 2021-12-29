package day04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	soln := Solution{}
	actual, err := soln.Part1("testdata.txt")
	assert.NoError(t, err)
	assert.Equal(t, 4512, actual)
}

func TestPart2(t *testing.T) {
	soln := Solution{}
	actual, err := soln.Part2("testdata.txt")
	assert.NoError(t, err)
	assert.Equal(t, 1924, actual)
}
