package day12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountPaths(t *testing.T) {
	g, err := loadGraph("testdata.txt")
	assert.NoError(t, err)

	actual := countPaths(g, cave("start"), cave("end"))
	expected := 10

	assert.Equal(t, expected, actual)
}
