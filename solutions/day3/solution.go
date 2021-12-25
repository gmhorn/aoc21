package day3

import (
	"errors"
)

type Solution struct{}

func (sln Solution) Part1(input string) (int, error) {
	report, err := NewReport(input)
	if err != nil {
		return -1, err
	}

	g, err := report.Gamma()
	if err != nil {
		return -1, err
	}
	e, err := report.Epsilon()
	if err != nil {
		return -1, err
	}

	return int(g * e), nil
}

func (sln Solution) Part2(input string) (int, error) {
	return -1, errors.New("not implemented")
}
