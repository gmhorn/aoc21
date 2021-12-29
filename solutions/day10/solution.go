package day10

import (
	"sort"

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
		if score, result := parse(line); result == resultCorrupted {
			total += score
		}
	}
	return total, nil
}

func (sln Solution) Part2(input string) (int, error) {
	lines, err := lib.ReadLines(input)
	if err != nil {
		return -1, err
	}

	scores := make([]int, 0)
	for _, line := range lines {
		if score, result := parse(line); result == resultIncomplete {
			scores = append(scores, score)
		}
	}

	sort.Sort(sort.IntSlice(scores))
	idx := int(len(scores) / 2)
	return scores[idx], nil
}

func scoreAll(input string, result result) (int, error) {
	lines, err := lib.ReadLines(input)
	if err != nil {
		return -1, err
	}

	total := 0
	for _, line := range lines {
		if score, res := parse(line); res == result {
			total += score
		}
	}
	return total, nil
}

type result int

const (
	resultIncomplete result = iota
	resultCorrupted
	resultErrOther
	resultOk
)

var openers = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var closers = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

var corruptValue = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var completeValue = map[rune]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

func parse(line string) (int, result) {
	stack := NewStack()
	for _, char := range line {
		// handle case when charater is an opening char
		if _, found := openers[char]; found {
			stack.Push(char)
		}

		// handle case when character is a closing char
		if opener, found := closers[char]; found {
			partner, err := stack.Pop()
			// handle where we've seen more closers than openers.
			// the problem doesn't have a category for this type of error, so
			// just call it "other" and give a score of -1. We filter it out
			// later anyway.
			if err != nil {
				return -1, resultErrOther
			}
			// if what we've popped off doesn't match what it should be, that's
			// a "corruption" error. The score comes from the current character
			if partner != opener {
				return corruptValue[char], resultCorrupted
			}
		}

		// any other character isn't part of our grammer - just let it go
	}

	// if stack is empty, we have a clean parse
	if stack.Empty() {
		return 0, resultOk
	}

	// Otherwise we have an incomplete parse.
	// Compute the score.
	res := resultOk
	score := 0

	opener, err := stack.Pop()
	for ; err == nil; opener, err = stack.Pop() {
		res = resultIncomplete
		score *= 5
		score += completeValue[opener]
	}

	return score, res
}
