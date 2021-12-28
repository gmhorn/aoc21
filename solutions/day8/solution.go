package day8

import (
	"errors"
	"sort"
	"strings"

	"github.com/gmhorn/aoc21/lib"
)

type Solution struct{}

func (sln Solution) Part1(input string) (int, error) {
	lines, err := lib.ReadLines(input)
	if err != nil {
		return -1, err
	}

	uniques := 0
	for _, line := range lines {
		output := strings.Split(line, "|")[1]
		digits := strings.Split(strings.TrimSpace(output), " ")
		for _, digit := range digits {
			switch len(digit) {
			case 2:
				uniques++
			case 3:
				uniques++
			case 4:
				uniques++
			case 7:
				uniques++
			}
		}
	}

	return uniques, nil
}

func (sln Solution) Part2(input string) (int, error) {
	return -1, errors.New("not implemented")
}

var unscrambled = []Pattern{
	ToPattern("ABCEFG"),  // 0
	ToPattern("CF"),      // 1
	ToPattern("ACDEG"),   // 2
	ToPattern("ACDFG"),   // 3
	ToPattern("BCDF"),    // 4
	ToPattern("ABDFG"),   // 5
	ToPattern("ABDEFG"),  // 6
	ToPattern("ACF"),     // 7
	ToPattern("ABCDEFG"), // 8
	ToPattern("ABCDFG"),  // 9
}

type Pattern []rune

func (p Pattern) Exclude(other Pattern) Pattern {
	idx := 0
	for _, cOther := range other {
		for idx, cPattern := range p {
			if cOther == cPattern {
				break
			}
		}
		if idx < len(pattern)
	}
	remaining := make([]rune, 0)
	for _, c := range p {

	}

	return remaining
}
func ToPattern(str string) Pattern {
	pat := Pattern(str)
	sort.Slice(pat, func(i, j int) bool {
		return pat[i] < pat[j]
	})
	return pat
}

func process(line string) []Pattern {
	halves := strings.Split(line, "|")

	signals := make([]Pattern, 0)
	outputs := make([]Pattern, 0)
	for _, part := range strings.Split(strings.TrimSpace(halves[0]), " ") {
		signals = append(signals, ToPattern(part))
	}
	for _, part := range strings.Split(strings.TrimSpace(halves[1]), " ") {
		outputs = append(outputs, ToPattern(part))
	}

	sort.Slice(signals, func(i, j int) bool {
		return len(signals[i]) < len(signals[j])
	})

	candidates := map[rune]Pattern{
		'A': ToPattern("abcdefg"),
		'B': ToPattern("abcdefg"),
		'C': ToPattern("abcdefg"),
		'D': ToPattern("abcdefg"),
		'E': ToPattern("abcdefg"),
		'F': ToPattern("abcdefg"),
		'G': ToPattern("abcdefg"),
	}
	return signals
}
