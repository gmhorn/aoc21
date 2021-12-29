package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBingo(t *testing.T) {
	lines := []string{
		"22 13 17 11  0",
		" 8  2 23  4 24",
		"21  9 14 16  7",
		" 6 10  3 18  5",
		" 1 12 20 15 19",
	}
	b, err := NewBingo(lines)
	assert.NoError(t, err)
	assert.NotNil(t, b)
}
