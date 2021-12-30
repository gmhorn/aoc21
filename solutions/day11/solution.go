package day11

import (
	"errors"

	"github.com/gmhorn/aoc21/lib"
)

type Solution struct{}

func (sln Solution) Part1(input string) (int, error) {
	steps := 100
	threshold := 9

	grid, err := lib.LoadGrid(input)
	if err != nil {
		return -1, err
	}

	flashes := 0
	for i := 0; i < steps; i++ {
		flashes += advance(grid, threshold)
	}
	return flashes, nil
}

func (sln Solution) Part2(input string) (int, error) {
	return -1, errors.New("not implemented")
}

// advance advances the Grid according to the rules of the problem statement. It
// returns the number of elements that "flashed" above the given threshold.
func advance(grid *lib.Grid, threshold uint8) int {
	toFlash := make([]lib.GridCoord, 0)
	flashed := make(map[lib.GridCoord]bool)

	// do initial increment
	for r := 0; r < grid.Rows; r++ {
		for c := 0; c < grid.Cols; c++ {
			grid.Vals[r][c]++
			// if we're above threshold, add that point to our list of coords
			// that need to be flashed
			if grid.Vals[r][c] > threshold {
				toFlash = append(toFlash, lib.GridCoord{r, c})
			}
		}
	}

	// Process flashes until we quiesce
	for len(toFlash) != 0 {
		// node := toFlash[0]
		// toFlash = toFlash[1:]
	}

	// Reset all flashed coords to 0 while counting total
	total := 0
	for node, _ := range flashed {
		grid.Vals[node.Row][node.Col] = 0
		total++
	}
	return total
}

func surroundingPoints(grid *lib.Grid, origin )