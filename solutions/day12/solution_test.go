package day12

import (
	"io/ioutil"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExplore(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		rule     rule
		expected string
	}{{
		name:     "visit once",
		input:    "testdata/input.txt",
		rule:     visitOnce,
		expected: "testdata/visitOnce.txt",
	}, {
		name:     "visit twice",
		input:    "testdata/input.txt",
		rule:     visitTwice,
		expected: "testdata/visitTwice.txt",
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, err := loadGraph(tt.input)
			assert.NoError(t, err)

			paths := make([]string, 0)
			for _, p := range explore(g, tt.rule) {
				paths = append(paths, p.String())
			}

			sort.Slice(paths, func(i, j int) bool {
				return paths[i] < paths[j]
			})

			actual := strings.Join(paths, "\n")

			expected, err := ioutil.ReadFile(tt.expected)
			assert.NoError(t, err)

			assert.Equal(t, string(expected), actual)
		})
	}
}
