package day3

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gmhorn/aoc21/lib"
)

const (
	high = 1
	low  = -1

	highChar = '1'
	lowChar  = '0'
)

type Solution struct{}

func (sln Solution) Part1(input string) (int, error) {
	lines, err := lib.ReadLines(input)
	if err != nil {
		return -1, err
	}

	acc, err := accumulate(lines)
	if err != nil {
		return -1, err
	}

	g, err := gamma(acc)
	if err != nil {
		return -1, err
	}
	e, err := epsilon(acc)
	if err != nil {
		return -1, err
	}

	return int(g * e), nil
}

func (sln Solution) Part2(input string) (int, error) {
	return -1, errors.New("not implemented")
}

func accumulate(lines []string) ([]int, error) {
	// MISSING: check that lines not nil and has at least 1 element
	// MISSING: check that all lines are equal length
	acc := make([]int, len(lines[0]))

	for idxL, line := range lines {
		for idxC, c := range line {
			switch c {
			case highChar:
				acc[idxC] += high
			case lowChar:
				acc[idxC] += low
			default:
				return acc, fmt.Errorf("unknown char '%c' at %d:%d", c, idxL, idxC)
			}
		}
	}

	return acc, nil
}

func gamma(accumulator []int) (int64, error) {
	s := ""
	for idx, v := range accumulator {
		switch {
		case v >= high:
			s += string(highChar)
		case v <= low:
			s += string(lowChar)
		default:
			return -1, fmt.Errorf("ambiguous gamma digit at %d position", idx)
		}
	}
	return strconv.ParseInt(s, 2, 64)
}

func epsilon(accumulator []int) (int64, error) {
	s := ""
	for idx, v := range accumulator {
		switch {
		case v >= high:
			s += string(lowChar)
		case v <= low:
			s += string(highChar)
		default:
			return -1, fmt.Errorf("ambiguous epsilon digit at %d position", idx)
		}
	}
	return strconv.ParseInt(s, 2, 64)
}
