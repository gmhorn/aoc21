package day8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValueOf(t *testing.T) {
	val, err := valueOf("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf")
	assert.NoError(t, err)
	assert.Equal(t, 5353, val)
}

// func TestIntersection(t *testing.T) {
// 	tests := []struct {
// 		input    []string
// 		expected string
// 	}{{
// 		input:    []string{"acd", "abc", "abcd"},
// 		expected: "ac",
// 	}}

// 	for idx, tt := range tests {
// 		t.Run(fmt.Sprintf("case %d", idx), func(t *testing.T) {
// 			sigs := toSignals(tt.input)
// 			assert.Equal(t, NewSignal(tt.expected), Intersection(sigs))
// 		})
// 	}
// }

// func TestOccurrences(t *testing.T) {
// 	tests := []struct {
// 		input    []string
// 		r        rune
// 		expected int
// 	}{{
// 		input:    []string{"abcd", "acd", "cd", "bd", "ad"},
// 		r:        'a',
// 		expected: 3,
// 	}, {
// 		input:    []string{"abcd", "acd", "cd", "bd", "ad"},
// 		r:        'z',
// 		expected: 0,
// 	}}

// 	for idx, tt := range tests {
// 		t.Run(fmt.Sprintf("case %d", idx), func(t *testing.T) {
// 			sigs := toSignals(tt.input)
// 			assert.Equal(t, tt.expected, Occurrences(tt.r, sigs))
// 		})
// 	}
// }

func toSignals(input []string) []Signal {
	out := make([]Signal, 0, len(input))
	for _, in := range input {
		out = append(out, NewSignal(in))
	}
	return out
}
