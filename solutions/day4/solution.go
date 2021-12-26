package day4

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gmhorn/aoc21/lib"
)

type Solution struct{}

func (sln Solution) Part1(input string) (int, error) {
	vals, boards, err := load(input)
	if err != nil {
		return -1, err
	}

	for _, val := range vals {
		for _, board := range boards {
			if board.Mark(val) {
				return val * board.SumOfUnmarked(), nil
			}
		}
	}
	return -1, errors.New("no winners found")
}

func (sln Solution) Part2(input string) (int, error) {
	return -1, errors.New("not implemented")
}

func load(input string) ([]int, []*Bingo, error) {
	rawLines, err := lib.ReadLines(input)
	if err != nil {
		return nil, nil, err
	}

	// Time blank lines
	lines := make([]string, 0)
	for _, line := range rawLines {
		if line != "" {
			lines = append(lines, line)
		}
	}

	// Parse first line
	vals := make([]int, 0)
	for _, sval := range strings.Split(lines[0], ",") {
		val, err := strconv.Atoi(sval)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing '%s' in values line: %v", sval, err)
		}
		vals = append(vals, val)
	}
	// Shift slice forward 1
	lines = lines[1:]

	// Parse through rest of lines 5 at a time, making boards
	boards := make([]*Bingo, 0)
	for len(lines) > 0 {
		board, err := NewBingo(lines[:5])
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing bingo board: %v", err)
		}
		boards = append(boards, board)
		lines = lines[5:]
	}

	return vals, boards, nil
}
