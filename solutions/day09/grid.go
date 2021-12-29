package day09

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/gmhorn/aoc21/lib"
)

// Grid is a structure representing a heightmap of values
//
//      0 ------>xMax
//    0 2199943210
//    | 3987894921
//    | 9856789892
//    V 8767896789
// yMax 9899965678
type Grid struct {
	Vals       [][]int8
	XMax, YMax int
}

// Point represents a x,y point on a Grid.
type Point struct {
	X, Y int
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

func (g *Grid) lowPoints() []Point {
	minima := make([]Point, 0)
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
			minima = append(minima, Point{x, y})
		}
	}

	return minima
}

// Minima returns the list of minima values, sorted ascending
func (g *Grid) Minima() []int8 {
	minima := make([]int8, 0)
	for _, pt := range g.lowPoints() {
		minima = append(minima, g.Vals[pt.Y][pt.X])
	}

	sort.Slice(minima, func(i, j int) bool {
		return minima[i] < minima[j]
	})

	return minima
}

// Basins returns the list of basin sizes, sorted descending
func (g *Grid) Basins() []int {
	sizes := make([]int, 0)

	for _, seed := range g.lowPoints() {
		sizes = append(sizes, g.basinSize(seed))
	}

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	return sizes
}

// Calculates basin size using flood fill algo (with explicit stack).
// See https://en.wikipedia.org/wiki/Flood_fill
func (g *Grid) basinSize(point Point) int {
	toSearch := []Point{point}
	visited := make(map[Point]bool)
	acc := 0

	for len(toSearch) != 0 {
		node := toSearch[0]
		toSearch = toSearch[1:]

		if g.Vals[node.Y][node.X] < 9 && !visited[node] {
			acc++
			visited[node] = true

			if node.X < g.XMax {
				toSearch = append(toSearch, Point{node.X + 1, node.Y})
			}
			if node.X > 0 {
				toSearch = append(toSearch, Point{node.X - 1, node.Y})
			}
			if node.Y < g.YMax {
				toSearch = append(toSearch, Point{node.X, node.Y + 1})
			}
			if node.Y > 0 {
				toSearch = append(toSearch, Point{node.X, node.Y - 1})
			}
		}
	}

	return acc
}
