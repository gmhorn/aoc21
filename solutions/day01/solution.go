package day01

import "github.com/gmhorn/aoc21/lib"

type Solution struct{}

func (s Solution) Part1(input string) (int, error) {
	data, err := lib.ReadLinesToInts(input)
	if err != nil {
		return -1, err
	}

	return reduce(data, increase), nil
}

func (s Solution) Part2(input string) (int, error) {
	data, err := lib.ReadLinesToInts(input)
	if err != nil {
		return -1, err
	}

	return reduce(windowSum(data, 3), increase), nil
}

func windowSum(data []int, width int) []int {
	// TODO: check width < len(data)
	sums := make([]int, 0)

	sum := 0
	for i := 0; i < width; i++ {
		sum += data[i]
	}
	sums = append(sums, sum)

	for i := width; i < len(data); i++ {
		sum = sum + data[i] - data[i-width]
		sums = append(sums, sum)
	}

	return sums
}

func reduce(data []int, fn func(x, y int) int) int {
	total := 0
	a := data[0]

	for _, b := range data[1:] {
		total += fn(a, b)
		a = b
	}

	return total
}

func increase(x, y int) int {
	if y > x {
		return 1
	}
	return 0
}
