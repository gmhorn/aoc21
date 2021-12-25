package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	s := Solution{}
	actual, err := s.Part1("testdata.txt")
	assert.NoError(t, err)
	assert.Equal(t, 198, actual)
}
