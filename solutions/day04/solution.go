package day04

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
	vals, boards, err := load(input)
	if err != nil {
		return -1, err
	}

	remaining := boards
	for _, val := range vals {
		// Iterate from back to front, since more than 1 board can win each round
		// Assumes final round there's only 1 left...
		// IDK this seems hacky
		for idx := len(remaining) - 1; idx >= 0; idx-- {
			board := remaining[idx]
			if board.Mark(val) {
				if len(remaining) == 1 {
					return val * board.SumOfUnmarked(), nil
				}
				remaining = evict(remaining, idx)
			}
		}
	}

	return -1, errors.New("something went wrong")
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

func evict(boards []*Bingo, index int) []*Bingo {
	return append(boards[:index], boards[index+1:]...)
}
