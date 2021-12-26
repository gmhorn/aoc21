package day7

import (
	"errors"
	"fmt"
	"math"
	"sort"

	"github.com/gmhorn/aoc21/lib"
)

type Solution struct{}

func (sln Solution) Part1(input string) (int, error) {
	lines, err := lib.ReadLines(input)
	if err != nil {
		return -1, err
	}
	positions, err := lib.ParseCSVInts(lines[0])
	if err != nil {
		return -1, nil
	}
	sort.Ints(positions)
	minCost := math.MaxInt
	for target := positions[0]; target <= positions[len(positions)-1]; target++ {
		c := cost(positions, target)
		if c < minCost {
			minCost = c
		}
	}
	return minCost, nil
}

func (sln Solution) Part2(input string) (int, error) {
	return -1, errors.New("not implemented")
}

func cost(positions []int, target int) int {
	total := 0
	for _, pos := range positions {
		total += abs(pos - target)
	}
	return total
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
