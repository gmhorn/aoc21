package lib

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGridLoad(t *testing.T) {
	raw, err := ioutil.ReadFile("testdata/grid.txt")
	assert.NoError(t, err)

	grid, err := LoadGrid("testdata/grid.txt")
	assert.NoError(t, err)

	assert.Equal(t, string(raw), grid.String())
}
