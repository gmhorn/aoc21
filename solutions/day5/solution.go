package day5

import (
	"fmt"

	"github.com/gmhorn/aoc21/lib"
)

type Solution struct{}

func (sln Solution) Part1(input string) (int, error) {
	lines, err := readInput(input, true)
	if err != nil {
		return -1, err
	}
	return CountOverlaps(lines, 2), nil
}

func (sln Solution) Part2(input string) (int, error) {
	lines, err := readInput(input, false)
	if err != nil {
		return -1, err
	}
	return CountOverlaps(lines, 2), nil
}

func readInput(input string, rectOnly bool) ([]Line, error) {
	strs, err := lib.ReadLines(input)
	if err != nil {
		return nil, err
	}

	lines := make([]Line, 0)
	for idx, str := range strs {
		line, err := ParseLine(str)
		if err != nil {
			return nil, fmt.Errorf("could not parse line %d: %v", idx, err)
		}

		if rectOnly && !line.IsRect() {
			continue
		}

		lines = append(lines, line)
	}

	return lines, nil
}
