package solutions

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gmhorn/aoc21/lib"
)

type Day2 struct{}

func (d Day2) Part1() (int, error) {
	lines, err := lib.ReadLines("input/day2.txt")
	if err != nil {
		return -1, err
	}

	x, y := 0, 0
	for idx, line := range lines {
		parts := strings.Split(line, " ")
		action := parts[0]
		amt, err := strconv.Atoi(parts[1])
		if err != nil {
			return -1, fmt.Errorf("could not parse line %d: %v", idx, err)
		}
		switch action {
		case "forward":
			x += amt
		case "down":
			y += amt
		case "up":
			y -= amt
		}
	}
	return x * y, nil
}

func (d Day2) Part2() (int, error) {
	return -4, errors.New("not implemented")
}
