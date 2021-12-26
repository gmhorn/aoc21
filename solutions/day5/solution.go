package day5

import (
	"errors"
	"fmt"

	"github.com/gmhorn/aoc21/lib"
)

type Solution struct{}

func (sln Solution) Part1(input string) (int, error) {
	strs, err := lib.ReadLines(input)
	if err != nil {
		return -1, err
	}

	lines := make([]Line, 0)
	for idx, str := range strs {
		line, err := ParseLine(str)
		if err != nil {
			return -1, fmt.Errorf("could not parse line %d: %v", idx, err)
		}
		lines = append(lines, line)
	}

	return CountOverlaps(lines, 2), nil
}

func (sln Solution) Part2(input string) (int, error) {
	return -1, errors.New("not implemented")
}
