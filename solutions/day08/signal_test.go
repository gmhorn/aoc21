package day08

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestContains(t *testing.T) {
	tests := []struct {
		input    string
		search   rune
		expected bool
	}{{
		input:    "abcd",
		search:   'b',
		expected: true,
	}, {
		input:    "abd",
		search:   'c',
		expected: false,
	}}

	for idx, tt := range tests {
		t.Run(fmt.Sprintf("case %d", idx), func(t *testing.T) {
			sig := NewSignal(tt.input)
			assert.Equal(t, tt.expected, sig.Contains(tt.search))
		})
	}

}

func TestIntersect(t *testing.T) {
	tests := []struct {
		a, b, expected string
	}{{
		a:        "abd",
		b:        "acd",
		expected: "ad",
	}}

	for idx, tt := range tests {
		t.Run(fmt.Sprintf("case %d", idx), func(t *testing.T) {
			sig := NewSignal(tt.a)
			other := NewSignal(tt.b)

			assert.Equal(t, NewSignal(tt.expected), sig.Intersect(other))
		})
	}
}
