package day8

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIt(t *testing.T) {
	output := process("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf")
	assert.NotNil(t, output)
}

func TestExclude(t *testing.T) {
	tests := []struct {
		base, exclude, expected string
	}{{
		base:     "abcd",
		exclude:  "ab",
		expected: "cd",
	}, {
		base:     "abcd",
		exclude:  "bc",
		expected: "ad",
	}, {
		base:     "abcd",
		exclude:  "ac",
		expected: "bd",
	}, {
		base:     "abcd",
		exclude:  "e",
		expected: "abcd",
	}}

	for idx, tt := range tests {
		t.Run(fmt.Sprintf("case %d", idx), func(t *testing.T) {
			base := NewSignal(tt.base)
			exclude := NewSignal(tt.exclude)
			expected := NewSignal(tt.expected)

			actual := base.Exclude(exclude)

			assert.Equal(t, expected, actual)
		})
	}
}
