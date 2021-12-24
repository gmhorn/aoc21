package solutions

import "github.com/gmhorn/aoc21/lib"

type Day1 struct{}

func (d Day1) Part1() (int, error) {
	data, err := lib.ReadLinesToInts("input/day1.txt")
	if err != nil {
		return -1, err
	}

	return lib.Reduce(data, lib.Increase), nil
}

func (d Day1) Part2() (int, error) {
	data, err := lib.ReadLinesToInts("input/day1.txt")
	if err != nil {
		return -1, err
	}

	return lib.Reduce(lib.WindowSum(data, 3), lib.Increase), nil
}
