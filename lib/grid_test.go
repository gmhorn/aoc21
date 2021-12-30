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

func TestForEachNeighbor(t *testing.T) {
	grid := NewGrid(3, 3)

	type coord struct {
		row, col int
	}

	tests := []struct {
		name   string
		origin coord
		mode   AdjacencyMode
		in     []coord
		out    []coord
	}{{
		name:   "middle, 4 way",
		origin: coord{1, 1},
		mode:   AdjacencyMode4Way,
		in:     []coord{{0, 1}, {2, 1}, {1, 0}, {1, 2}},
		out:    []coord{{1, 1}, {0, 0}, {2, 2}, {2, 0}, {0, 2}},
	}, {
		name:   "middle, 8 way",
		origin: coord{1, 1},
		mode:   AdjacencyMode8Way,
		in:     []coord{{0, 1}, {2, 1}, {1, 0}, {1, 2}, {0, 0}, {2, 2}, {2, 0}, {0, 2}},
		out:    []coord{{1, 1}},
	}, {
		name:   "corner, 4 way",
		origin: coord{0, 0},
		mode:   AdjacencyMode4Way,
		in:     []coord{{0, 1}, {1, 0}},
	}, {
		name:   "corner, 8 way",
		origin: coord{0, 0},
		mode:   AdjacencyMode8Way,
		in:     []coord{{0, 1}, {1, 0}, {1, 1}},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nbrs := make(map[coord]bool)
			fn := VisitFunc(func(row, col int, val uint8) bool {
				nbrs[coord{row, col}] = true
				return true
			})

			grid.ForEachNeighbor(tt.origin.row, tt.origin.col, tt.mode, fn)

			for _, expectIn := range tt.in {
				assert.True(t, nbrs[expectIn])
			}
			for _, expectOut := range tt.out {
				assert.False(t, nbrs[expectOut])
			}
		})
	}
}
