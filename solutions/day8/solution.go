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

var (
	alphabetScrambled   = NewSignal("abcdefg")
	alphabetUnscrambled = NewSignal("ABCDEFG")

	// unscrambled is an array where each (unscrambled) Signal's index
	// represents that Signal's digit value.
	unscrambled = []Signal{
		NewSignal("ABCEFG"),  // 0
		NewSignal("CF"),      // 1
		NewSignal("ACDEG"),   // 2
		NewSignal("ACDFG"),   // 3
		NewSignal("BCDF"),    // 4
		NewSignal("ABDFG"),   // 5
		NewSignal("ABDEFG"),  // 6
		NewSignal("ACF"),     // 7
		NewSignal("ABCDEFG"), // 8
		NewSignal("ABCDFG"),  // 9
	}
)

// Signal is a 7-digit output signal. It is represented as a (sorted) list of
// runes from a designated alphabet.
type Signal []rune

// NewSignal creates a new Signal from the input string, making sure to sort
// the consitituent runes lexigraphically.
func NewSignal(str string) Signal {
	sig := Signal(str)
	sort.Slice(sig, func(i int, j int) bool {
		return sig[i] < sig[j]
	})
	return sig
}

// Exclude constructs a new Signal that consists of this Signal, minus any runes
// that appear in the other Signal.
//
// This relies on both this Signal and the other Signal being lexigraphically
// sorted and not containing duplicate runes.
func (s Signal) Exclude(other Signal) Signal {
	res := make([]rune, 0)

	idxS, idxO := 0, 0
	for ; idxS < len(s) && idxO < len(other); idxS++ {
		if s[idxS] == other[idxO] {
			idxO++
			continue
		}
		res = append(res, s[idxS])
	}

	res = append(res, s[idxS:]...)

	return res
}

type Pattern []rune

// func (p Pattern) Exclude(other Pattern) Pattern {
// 	idx := 0
// 	for _, cOther := range other {
// 		for idx, cPattern := range p {
// 			if cOther == cPattern {
// 				break
// 			}
// 		}
// 		if idx < len(pattern)
// 	}
// 	remaining := make([]rune, 0)
// 	for _, c := range p {

// 	}

// 	return remaining
// }

func ToPattern(str string) Pattern {
	pat := Pattern(str)
	sort.Slice(pat, func(i, j int) bool {
		return pat[i] < pat[j]
	})
	return pat
}

func process(line string) []Signal {
	// Parse line into list of pattern Signals and list of output Signals
	// All Signals are using the scrambled alphabet
	halves := strings.Split(line, "|")

	patterns := make([]Signal, 0)
	outputs := make([]Signal, 0)
	for _, part := range strings.Split(strings.TrimSpace(halves[0]), " ") {
		patterns = append(patterns, NewSignal(part))
	}
	for _, part := range strings.Split(strings.TrimSpace(halves[1]), " ") {
		outputs = append(outputs, NewSignal(part))
	}

	// Sort the pattern Signals by length
	sort.Slice(patterns, func(i, j int) bool {
		return len(patterns[i]) < len(patterns[j])
	})

	// candidates := map[rune]Pattern{
	// 	'A': ToPattern("abcdefg"),
	// 	'B': ToPattern("abcdefg"),
	// 	'C': ToPattern("abcdefg"),
	// 	'D': ToPattern("abcdefg"),
	// 	'E': ToPattern("abcdefg"),
	// 	'F': ToPattern("abcdefg"),
	// 	'G': ToPattern("abcdefg"),
	// }
	return patterns
}
