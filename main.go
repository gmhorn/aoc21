package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := loadArray("input/day1.txt")
	part1 := Reduce(data, Increase)
	part2 := Reduce(WindowSum(data, 3), Increase)
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

func WindowSum(data []int, width int) []int {
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

func Reduce(data []int, fn func(x, y int) int) int {
	total := 0
	a := data[0]

	for _, b := range data[1:] {
		total += fn(a, b)
		a = b
	}

	return total
}

func Increase(x, y int) int {
	if y > x {
		return 1
	}
	return 0
}

func loadArray(path string) []int {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(1)
	}

	data := make([]int, 0)
	for i, line := range strings.Split(string(b), "\n") {
		val, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("Error parsing line %d in %s: %v\n", i, path, err)
			os.Exit(1)
		}
		data = append(data, val)
	}

	return data
}
