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

func TestNeighbors(t *testing.T) {
	grid := NewGrid(3, 3)

	tests := []struct {
		name   string
		origin GridCoord
		mode   AdjacencyMode
		in     []GridCoord
		out    []GridCoord
	}{{
		name:   "middle, 4 way",
		origin: GridCoord{1, 1},
		mode:   AdjacencyMode4Way,
		in:     []GridCoord{{0, 1}, {2, 1}, {1, 0}, {1, 2}},
		out:    []GridCoord{{1, 1}, {0, 0}, {2, 2}, {2, 0}, {0, 2}},
	}, {
		name:   "middle, 8 way",
		origin: GridCoord{1, 1},
		mode:   AdjacencyMode8Way,
		in:     []GridCoord{{0, 1}, {2, 1}, {1, 0}, {1, 2}, {0, 0}, {2, 2}, {2, 0}, {0, 2}},
		out:    []GridCoord{{1, 1}},
	}, {
		name:   "corner, 4 way",
		origin: GridCoord{0, 0},
		mode:   AdjacencyMode4Way,
		in:     []GridCoord{{0, 1}, {1, 0}},
	}, {
		name:   "corner, 8 way",
		origin: GridCoord{0, 0},
		mode:   AdjacencyMode8Way,
		in:     []GridCoord{{0, 1}, {1, 0}, {1, 1}},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nbrs := make(map[GridCoord]bool)
			for _, pt := range grid.Neighbors(tt.origin, tt.mode) {
				nbrs[pt] = true
			}

			for _, expectIn := range tt.in {
				assert.True(t, nbrs[expectIn])
			}
			for _, expectOut := range tt.out {
				assert.False(t, nbrs[expectOut])
			}
		})
	}
}
