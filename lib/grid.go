package lib

import (
	"fmt"
	"strconv"
	"strings"
)

// Grid is a 2D grid of uint8 values.
type Grid struct {
	Vals       [][]uint8
	Rows, Cols int
}

// NewGrid creates a new Grid with the given number of rows and columns and all
// values initialized to 0.
func NewGrid(rows, cols int) *Grid {
	vals := make([][]uint8, rows)
	for r := range vals {
		vals[r] = make([]uint8, cols)
	}

	return &Grid{
		Vals: vals,
		Rows: rows,
		Cols: cols,
	}
}

func LoadGrid(path string) (*Grid, error) {
	lines, err := ReadLines(path)
	if err != nil {
		return nil, err
	}

	grid := NewGrid(len(lines), len(lines[0]))

	for r, line := range lines {
		for c := range line {
			val, err := strconv.ParseUint(line[c:c+1], 10, 8)
			if err != nil {
				return nil, fmt.Errorf("could not parse %d:%d: %v", r, c, err)
			}
			grid.Vals[r][c] = uint8(val)
		}
	}

	return grid, nil
}

func (g *Grid) String() string {
	lines := make([]string, g.Rows)
	for r, row := range g.Vals {
		for _, val := range row {
			lines[r] += strconv.FormatUint(uint64(val), 10)
		}
	}
	return strings.Join(lines, "\n")
}
