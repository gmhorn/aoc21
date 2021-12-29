package day11

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	raw, err := ioutil.ReadFile("testdata.txt")
	assert.NoError(t, err)

	grid, err := LoadGrid("testdata.txt")
	assert.NoError(t, err)

	assert.Equal(t, string(raw), grid.String())
}
