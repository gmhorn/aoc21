package day09

import (
	"sort"

	"github.com/gmhorn/aoc21/lib"
)

type Solution struct{}

func (sln Solution) Part1(input string) (int, error) {
	grid, err := lib.LoadGrid(input)
	if err != nil {
		return -1, err
	}

	risk := 0
	for _, minimum := range minima(grid) {
		risk += int(minimum) + 1
	}
	return risk, nil
}

func (sln Solution) Part2(input string) (int, error) {
	grid, err := lib.LoadGrid(input)
	if err != nil {
		return -1, err
	}

	acc := 1
	for _, basin := range basins(grid, 9)[:3] {
		acc *= basin
	}
	return acc, nil
}

type coord struct {
	row, col int
}

// lowPoints returns the list of coords corresponding to the low points of the
// grid.
func lowPoints(grid *lib.Grid) []coord {
	minima := make([]coord, 0)
	grid.ForEach(func(row, col int, val uint8) bool {
		isMin := true
		grid.ForEachNeighbor(row, col, lib.AdjacencyMode4Way, func(_, _ int, valOther uint8) bool {
			if valOther <= val {
				isMin = false
			}
			return isMin
		})

		if isMin {
			minima = append(minima, coord{row, col})
		}
		return true
	})
	return minima
}

// basinSize calculates basin size using flood fill algo (with explicit stack.)
// See https://en.wikipedia.org/wiki/Flood_fill
func basinSize(grid *lib.Grid, start coord, threshold uint8) int {
	toSearch := []coord{start}
	visited := make(map[coord]bool)
	acc := 0

	addToSearch := lib.VisitFunc(func(row, col int, val uint8) bool {
		toSearch = append(toSearch, coord{row: row, col: col})
		return true
	})

	for len(toSearch) != 0 {
		// pop off next coord to start searching at
		node := toSearch[0]
		toSearch = toSearch[1:]

		// if node is not a boundary node, and we haven't visited it yet
		if grid.Vals[node.row][node.col] < threshold && !visited[node] {
			// mark as visited and increment our accumulator
			visited[node] = true
			acc++

			// add its neighbors to our list of nodes to search
			grid.ForEachNeighbor(node.row, node.col, lib.AdjacencyMode4Way, addToSearch)
		}
	}
	return acc
}

// minima returns the list of minima values, sorted ascending.
func minima(grid *lib.Grid) []uint8 {
	minima := make([]uint8, 0)
	for _, pt := range lowPoints(grid) {
		minima = append(minima, grid.Vals[pt.row][pt.col])
	}

	sort.Slice(minima, func(i, j int) bool {
		return minima[i] < minima[j]
	})

	return minima
}

// basins returns the list of basin sizes, sorted descending.
func basins(grid *lib.Grid, threshold uint8) []int {
	sizes := make([]int, 0)

	for _, seed := range lowPoints(grid) {
		sizes = append(sizes, basinSize(grid, seed, threshold))
	}

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	return sizes
}
