package day8

import (
	"errors"
	"fmt"
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

// Contains determines if the given Signal contains the provided rune.
func (s Signal) Contains(r rune) bool {
	idx := sort.Search(len(s), func(i int) bool {
		return s[i] >= r
	})
	return idx < len(s) && s[idx] == r
}

// Intersect constructs a new Signal that is the rune-wise intersection of this
// and the other Signal.
func (s Signal) Intersect(other Signal) Signal {
	res := make([]rune, 0)

	for _, r := range s {
		if other.Contains(r) {
			res = append(res, r)
		}
	}

	return res
}

// Intersection returns a Signal that is the set intersection of the constituent
// runes of the input Signals.
func Intersection(input []Signal) Signal {
	res := input[0]
	for _, sig := range input[1:] {
		res = res.Intersect(sig)
	}
	return res
}

// Occurrences counts the number of occurences of the given run in the input
// signals.
func Occurrences(r rune, input []Signal) int {
	occ := 0
	for _, sig := range input {
		if sig.Contains(r) {
			occ++
		}
	}
	return occ
}

func Frequencies(input []Signal) map[rune]int {
	freq := make(map[rune]int)
	for _, sig := range input {
		for _, r := range sig {
			freq[r]++
		}
	}
	return freq
}

func process(line string) map[rune]Signal {
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

	// Build a list of candidates to map unscrambled to scrambed runes
	candidates := make(map[rune]Signal)
	for _, unscrambled := range alphabetUnscrambled {
		candidates[unscrambled] = alphabetScrambled
	}

	// Pattern 0 is for '1', so its runes are candidates for C and F.
	// candidates['C'] = patterns[0]
	// candidates['F'] = patterns[0]
	// Pattern 1 is for '7' and pattern 0 is for '1'.
	// Exclude '7' - '1' gives the value for A.
	candidates['A'] = patterns[1].Exclude(patterns[0])
	// Pattern 0 is for '1', and contains candidates for 'C' and 'F'.
	// In a 10-element pattern, 'C' will occur 8 times and 'F' 9 times.
	// Count occurrences to get which is which
	occ := Occurrences(patterns[0][0], patterns)
	switch occ {
	case 8:
		candidates['C'] = Signal([]rune{patterns[0][0]})
		candidates['F'] = Signal([]rune{patterns[0][1]})
	case 9:
		candidates['C'] = Signal([]rune{patterns[0][1]})
		candidates['F'] = Signal([]rune{patterns[0][0]})
	default:
		// ERROR
		panic(fmt.Sprintf("Got %d occurences", occ))
	}

	// We can now exclude Pattern 1 from all other candidates
	for _, r := range []rune("BDEG") {
		candidates[r] = candidates[r].Exclude(patterns[1])
	}
	// Pattern 2 is for '4'. Exclude pattern for '1' to get candidates for B, D
	candidates['B'] = patterns[2].Exclude(patterns[0])
	candidates['D'] = patterns[2].Exclude(patterns[0])
	// Exclude candidates for B/D for E and G
	candidates['E'] = candidates['E'].Exclude(candidates['B'])
	candidates['G'] = candidates['G'].Exclude(candidates['B'])
	// Intersect 2, 3, 5 to get horizontal runes.
	// Then remove value for A, which we already know from earlier.
	// These are the candidates for D and G
	horiz := Intersection(patterns[3:5])
	horiz = horiz.Exclude(candidates['A'])
	candidates['D'] = horiz
	candidates['G'] = horiz

	return candidates
}
