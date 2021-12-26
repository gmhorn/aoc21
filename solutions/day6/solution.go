package day6

import (
	"errors"

	"github.com/gmhorn/aoc21/lib"
)

type Solution struct{}

func (sln Solution) Part1(input string) (int, error) {
	s, err := loadState(input)
	if err != nil {
		return -1, err
	}

	for i := 0; i < 80; i++ {
		s = nextState(s)
	}

	return sum(s), nil
}

func (sln Solution) Part2(input string) (int, error) {
	return -1, errors.New("not implemented")
}

const (
	tDouble = 7
	tStart  = 2
	tTotal  = tDouble + tStart
)

type state [tTotal]int

func nextState(init state) state {
	var next state

	// Handle all fish not ready to double
	for age, count := range init[1:] {
		// because we're taking a slice [1:] we don't need to decrement
		// age -> e.g. value at idx 1 of init[:] == value at idx 0 of init[1:]
		next[age] += count
	}

	// Handle rollover
	next[tDouble+tStart-1] += init[0]
	next[tDouble-1] += init[0]

	return next
}

func sum(s state) int {
	total := 0
	for _, count := range s {
		total += count
	}
	return total
}

func loadState(input string) (state, error) {
	var init state

	lines, err := lib.ReadLines(input)
	if err != nil {
		return init, err
	}
	vals, err := lib.ParseCSVInts(lines[0])
	if err != nil {
		return init, err
	}

	for _, val := range vals {
		init[val]++
	}
	return init, nil
}
