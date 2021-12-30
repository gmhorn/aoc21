package day11

import (
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
	var threshold uint8 = 9

	grid, err := lib.LoadGrid(input)
	if err != nil {
		return -1, err
	}

	allFlashed := grid.Rows * grid.Cols
	steps := 1
	for ; ; steps++ {
		if advance(grid, threshold) == allFlashed {
			break
		}
	}
	return steps, nil
}

type coord struct {
	row, col int
}

// advance advances the Grid according to the rules of the problem statement. It
// returns the number of elements that "flashed" above the given threshold.
func advance(grid *lib.Grid, threshold uint8) int {
	toFlash := make([]coord, 0)
	flashed := make(map[coord]bool)

	increment := lib.VisitFunc(func(row, col int, _ uint8) bool {
		grid.Vals[row][col]++
		// If above threshold, add to list of nodes to process for flash
		if grid.Vals[row][col] > threshold {
			toFlash = append(toFlash, coord{row, col})
		}
		return true
	})

	// Do initial increment
	grid.ForEach(increment)

	// Process flashes until we quiesce
	for len(toFlash) != 0 {
		// Pop node off list
		node := toFlash[0]
		toFlash = toFlash[1:]

		// Its possible this node could have already been flashed and re-added
		// to the list this round, so check first to avoid processing twice
		if !flashed[node] {
			// Now execute the flash:
			// 1. Mark as flashed
			flashed[node] = true
			// 2. Increment its neighbors
			grid.ForEachNeighbor(node.row, node.col, lib.AdjacencyMode8Way, increment)
		}
	}

	// Reset all flashed coords to 0 while counting total
	total := 0
	for node := range flashed {
		grid.Vals[node.row][node.col] = 0
		total++
	}
	return total
}
