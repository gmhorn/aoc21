package day7

import (
	"math"
	"sort"

	"github.com/gmhorn/aoc21/lib"
)

type Solution struct{}

func (sln Solution) Part1(input string) (int, error) {
	return solve(input, linear{})
}

func (sln Solution) Part2(input string) (int, error) {
	return solve(input, newGeometric())
}

func solve(input string, model costModel) (int, error) {
	lines, err := lib.ReadLines(input)
	if err != nil {
		return -1, err
	}
	positions, err := lib.ParseCSVInts(lines[0])
	if err != nil {
		return -1, nil
	}

	sort.Ints(positions)
	start, end := positions[0], positions[len(positions)-1]

	minCost := math.MaxInt
	for target := start; target <= end; target++ {
		cost := totalCost(positions, target, model)
		if cost < minCost {
			minCost = cost
		}
	}

	return minCost, nil
}

func totalCost(positions []int, target int, model costModel) int {
	total := 0
	for _, pos := range positions {
		total += model.cost(abs(pos - target))
	}
	return total
}

type costModel interface {
	cost(dist int) int
}

type linear struct{}

func (l linear) cost(dist int) int {
	return dist
}

type geometric struct {
	memo map[int]int
}

func newGeometric() *geometric {
	return &geometric{memo: make(map[int]int)}
}

func (g *geometric) cost(dist int) int {
	if dist <= 1 {
		return dist
	}
	if cost, found := g.memo[dist]; found {
		return cost
	}
	cost := dist + g.cost(dist-1)
	g.memo[dist] = cost
	return cost
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
