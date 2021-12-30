package lib

import (
	"fmt"
	"strconv"
	"strings"
)

// GridCoord is a simple struct to how a (Row, Column) Grid coordinate.
// Typically only useful for storing in larger data structures such as slices
// or maps.
type GridCoord struct {
	Row, Col int
}

// Grid is a 2D grid of uint8 values.
//
// Internally data is stored row-major, so Vals[0] gives the first row, and
// Vals[1][2] gives the 2nd row's 3rd element.
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

type AdjacencyMode int

const (
	AdjacencyMode4Way AdjacencyMode = 1
	AdjacencyMode8Way AdjacencyMode = 2
)

type VisitFunc func(row, col int, val uint8) bool

func (g *Grid) ForEach(fn VisitFunc) {
	for r, row := range g.Vals {
		for c, val := range row {
			if !fn(r, c, val) {
				return
			}
		}
	}
}

func (g *Grid) ForEachNeighbor(row, col int, mode AdjacencyMode, fn VisitFunc) {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			// Get row and column we're visiting
			rowVisit, colVisit := row+dx, col+dy
			// Skip if out of bounds
			if rowVisit < 0 || rowVisit >= g.Rows {
				continue
			}
			if colVisit < 0 || colVisit >= g.Cols {
				continue
			}
			// Calculate "distance" and exclude dist 0 (origin itself) and
			// distances greater than the mode. So for mode 4 we only visit
			// cardinal directions (dist==1) and for mode 8 we visit cardinal
			// and diagonal (dist==2).
			dist := dx*dx + dy*dy
			if dist == 0 || dist > int(mode) {
				continue
			}

			// Now actual visit the node
			// If fn returns false, stop visiting and return
			if !fn(rowVisit, colVisit, g.Vals[rowVisit][colVisit]) {
				return
			}
		}
	}
}

func (g *Grid) VisitNeighbors(row, col int, mode AdjacencyMode, fn VisitFunc) {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			// Get row and column we're visiting
			rowVisit, colVisit := row+dx, col+dy
			// Skip if out of bounds
			if rowVisit < 0 || rowVisit >= g.Rows {
				continue
			}
			if colVisit < 0 || colVisit >= g.Cols {
				continue
			}
			// Calculate "distance" and exclude dist 0 (origin itself) and
			// distances greater than the mode. So for mode 4 we only visit
			// cardinal directions (dist==1) and for mode 8 we visit cardinal
			// and diagonal (dist==2).
			dist := dx*dx + dy*dy
			if dist == 0 || dist > int(mode) {
				continue
			}

			// Now actual visit the node
			// If fn returns false, stop visiting and return
			if !fn(rowVisit, colVisit, g.Vals[rowVisit][colVisit]) {
				return
			}
		}
	}
}

func (g *Grid) Neighbors(origin GridCoord, mode AdjacencyMode) []GridCoord {
	nbrs := make([]GridCoord, 0)
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			// First calculate the "distance" and exclude dist 0 (origin itself)
			// and distances greater than the mode. This has the effect where
			// if mode is 4way, only distance 1 works (cardinal directions) but
			// for 8 way we get distance 2 (diagonals) included as well
			dist := IntAbs(dx) + IntAbs(dy)
			if dist == 0 || dist > int(mode) {
				continue
			}

			pt := GridCoord{origin.Row + dx, origin.Col + dy}
			if pt.Row < 0 || pt.Col < 0 || pt.Row >= g.Rows || pt.Col >= g.Cols {
				continue
			}

			nbrs = append(nbrs, pt)
		}
	}
	return nbrs
}
