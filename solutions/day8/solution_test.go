package day8

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	output := process("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf")
	assert.NotNil(t, output)
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
func TestIntersection(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{{
		input:    []string{"acd", "abc", "abcd"},
		expected: "ac",
	}}

	for idx, tt := range tests {
		t.Run(fmt.Sprintf("case %d", idx), func(t *testing.T) {
			sigs := toSignals(tt.input)
			assert.Equal(t, NewSignal(tt.expected), Intersection(sigs))
		})
	}
}

func TestOccurrences(t *testing.T) {
	tests := []struct {
		input    []string
		r        rune
		expected int
	}{{
		input:    []string{"abcd", "acd", "cd", "bd", "ad"},
		r:        'a',
		expected: 3,
	}, {
		input:    []string{"abcd", "acd", "cd", "bd", "ad"},
		r:        'z',
		expected: 0,
	}}

	for idx, tt := range tests {
		t.Run(fmt.Sprintf("case %d", idx), func(t *testing.T) {
			sigs := toSignals(tt.input)
			assert.Equal(t, tt.expected, Occurrences(tt.r, sigs))
		})
	}
}

func toSignals(input []string) []Signal {
	out := make([]Signal, 0, len(input))
	for _, in := range input {
		out = append(out, NewSignal(in))
	}
	return out
}
