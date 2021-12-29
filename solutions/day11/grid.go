package day11

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gmhorn/aoc21/lib"
)

type Grid struct {
	vals       [][]uint8
	xmax, ymax int
}

func LoadGrid(path string) (*Grid, error) {
	lines, err := lib.ReadLines(path)
	if err != nil {
		return nil, err
	}

	grid := &Grid{
		vals: make([][]uint8, 0),
		xmax: len(lines[0]) - 1,
		ymax: len(lines[0]),
	}

	for idxL, line := range lines {
		row := make([]uint8, 0, len(lines))
		for idxC := range line {
			val, err := strconv.ParseUint(line[idxC:idxC+1], 10, 8)
			if err != nil {
				return nil, fmt.Errorf("could not parse %d:%d: %v", idxL, idxC, err)
			}
			row = append(row, uint8(val))
		}
		grid.vals = append(grid.vals, row)
	}

	return grid, nil
}

func (g *Grid) String() string {
	rows := make([]string, 0)
	for _, row := range g.vals {
		s := ""
		for _, val := range row {
			s += strconv.FormatUint(uint64(val), 10)
		}
		rows = append(rows, s)
	}
	return strings.Join(rows, "\n")
}
