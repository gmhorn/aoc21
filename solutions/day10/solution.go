package day10

import (
	"errors"
	"strings"

	"github.com/gmhorn/aoc21/lib"
)

type Solution struct{}

func (sln Solution) Part1(input string) (int, error) {
	lines, err := lib.ReadLines(input)
	if err != nil {
		return -1, err
	}

	total := 0
	for _, line := range lines {
		total += score(line)
	}
	return total, nil
}

func (sln Solution) Part2(input string) (int, error) {
	return -1, errors.New("not implemented")
}

const openers = "([{<"
const closers = ")]}>"

var openerFor = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

var values = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

// score returns the s
func score(line string) int {
	stack := NewStack()
	for _, char := range line {
		switch {
		case strings.ContainsRune(openers, char):
			stack.Push(char)
		case strings.ContainsRune(closers, char):
			expect := openerFor[char]
			actual, _ := stack.Pop()
			if actual != expect {
				return values[char]
			}
		}
	}
	return 0
}
