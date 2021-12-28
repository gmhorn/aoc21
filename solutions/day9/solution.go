package day9

import "errors"

type Solution struct{}

func (sln Solution) Part1(input string) (int, error) {
	grid, err := LoadGrid(input)
	if err != nil {
		return -1, err
	}

	risk := 0
	for _, minimum := range grid.Minima() {
		risk += int(minimum) + 1
	}
	return risk, nil
}

func (sln Solution) Part2(input string) (int, error) {
	return -1, errors.New("not implemented")
}
