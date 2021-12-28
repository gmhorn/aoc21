package day9

import (
	"fmt"
	"strconv"

	"github.com/gmhorn/aoc21/lib"
)

type Grid struct {
	Vals       [][]int8
	XMax, YMax int
}

func LoadGrid(path string) (*Grid, error) {
	lines, err := lib.ReadLines(path)
	if err != nil {
		return nil, err
	}

	grid := &Grid{
		Vals: make([][]int8, 0),
		XMax: len(lines[0]) - 1,
		YMax: len(lines) - 1,
	}

	for idxL, line := range lines {
		arr := make([]int8, 0, len(line))
		for idxC := range line {
			val, err := strconv.ParseInt(line[idxC:idxC+1], 10, 8)
			if err != nil {
				return grid, fmt.Errorf("could not parse char at %d:%d: %v", idxL, idxC, err)
			}
			arr = append(arr, int8(val))
		}
		grid.Vals = append(grid.Vals, arr)
	}

	return grid, nil
}

func (g *Grid) Minima() []int8 {
	minima := make([]int8, 0)
	for x := 0; x <= g.XMax; x++ {
		for y := 0; y <= g.YMax; y++ {
			val := g.Vals[y][x]
			if x < g.XMax && g.Vals[y][x+1] <= val {
				continue
			}
			if x > 0 && g.Vals[y][x-1] <= val {
				continue
			}
			if y < g.YMax && g.Vals[y+1][x] <= val {
				continue
			}
			if y > 0 && g.Vals[y-1][x] <= val {
				continue
			}
			minima = append(minima, val)
		}
	}

	return minima
}
