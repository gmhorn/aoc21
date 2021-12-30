package day11

import (
	"errors"

	"github.com/gmhorn/aoc21/lib"
)

type Solution struct{}

func (sln Solution) Part1(input string) (int, error) {
	var steps = 100
	var threshold uint8 = 9

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
		// Pop node off list
		node := toFlash[0]
		toFlash = toFlash[1:]

		// Check if we've already flashed this node. If so, don't process a
		// second time
		if flashed[node] {
			continue
		}

		// At this point, execute the node's flash:
		// 1. Mark it as flashed
		flashed[node] = true
		// 2. Increment all its neighbors
		for _, nbr := range grid.Neighbors(node, lib.AdjacencyMode8Way) {
			grid.Vals[nbr.Row][nbr.Col]++
			// 3. If neighbor is above flash threshold add it to the list
			if grid.Vals[nbr.Row][nbr.Col] > threshold {
				toFlash = append(toFlash, nbr)
			}
		}
	}

	// Reset all flashed coords to 0 while counting total
	total := 0
	for node := range flashed {
		grid.Vals[node.Row][node.Col] = 0
		total++
	}
	return total
}
