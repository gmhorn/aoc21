package day08

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
	lines, err := lib.ReadLines(input)
	if err != nil {
		return -1, err
	}

	ans := 0
	for idx, line := range lines {
		v, err := valueOf(line)
		if err != nil {
			return -1, fmt.Errorf("error processing line %d: %v", idx, err)
		}
		ans += v
	}
	return ans, nil
}

// digits is an array where each (unscrambled) Signal's index represents that
// Signal's digit value.
var digits = []Signal{
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

func valueOf(line string) (int, error) {
	halves := strings.Split(line, "|")
	if len(halves) != 2 {
		return -1, errors.New("could not split line into patterns and output halves")
	}

	patterns := make([]Signal, 0)
	outputs := make([]Signal, 0)
	for _, pattern := range strings.Split(strings.TrimSpace(halves[0]), " ") {
		patterns = append(patterns, NewSignal(pattern))
	}
	for _, output := range strings.Split(strings.TrimSpace(halves[1]), " ") {
		outputs = append(outputs, NewSignal(output))
	}

	decoder, err := decode(patterns)
	if err != nil {
		return -1, fmt.Errorf("could not decode line: %v", err)
	}

	ans := 0
	pow := 1
	for i := len(outputs) - 1; i >= 0; i-- {
		d, ok := digitValue(outputs[i], decoder)
		if !ok {
			return -1, fmt.Errorf("could not decode %dth output", i)
		}
		ans = ans + (pow * d)
		pow = pow * 10
	}
	return ans, nil
}

func digitValue(scrambled Signal, decoder map[rune]rune) (int, bool) {
	decoded := make([]rune, 0)
	for _, r := range scrambled {
		decoded = append(decoded, decoder[r])
	}
	unscrambled := NewSignal(string(decoded))

	var digit int
	var sig Signal
	for digit, sig = range digits {
		if unscrambled.Equals(sig) {
			break
		}
	}

	return digit, digit < len(digits) && unscrambled.Equals(digits[digit])
}

// decode takes 10 pattern Signals representing the digits 0-9 in the
// "scrambled" alphabet [a,g] and returns an unscrambling map [a,g] -> [A,G].
func decode(patterns []Signal) (map[rune]rune, error) {
	if len(patterns) != 10 {
		return nil, fmt.Errorf("Need 10 patterns, got %d", len(patterns))
	}

	// Sort patterns by length
	sort.Slice(patterns, func(i, j int) bool {
		return len(patterns[i]) < len(patterns[j])
	})

	// Now we have the following:
	//
	// pattern Idx | 0 | 1 | 2 | 3 4 5 | 6 7 8 | 9
	// ------------|---|---|---|-------|-------|---
	// Display Num | 1 | 7 | 4 | 2,3,5 | 0,6,9 | 8

	// Candidates is the map of unscrambled runes to the list of scrambled runes
	// which may match it
	candidates := make(map[rune]Signal)

	// Start with frequency analysis to build our list of candidates
	// In a 10-element list, we have the following frequencies
	//
	// Element   | A | B | C | D | E | F | G
	// ----------|---|---|---|---|---|---|---
	// Frequency | 8 | 6 | 8 | 7 | 4 | 9 | 7
	freqs := Frequencies(patterns)

	// We can immediately decode 'B', 'E' and 'F' because they have unique
	// frequencies
	candidates['B'] = freqs.SignalFor(6)
	candidates['E'] = freqs.SignalFor(4)
	candidates['F'] = freqs.SignalFor(9)

	// To get 'C' exclude the signal for 'F' from the pattern for '1'
	candidates['C'] = patterns[0].Exclude(candidates['F'])
	// Similarly, to get 'A' exclude the pattern for '1' from the pattern for '7'
	candidates['A'] = patterns[1].Exclude(patterns[0])

	// All that remains is candidates for 'D' and 'G' (both freq 7)
	// To get 'D', intersect the remaining candidates with the pattern for '4'
	candidates['D'] = patterns[2].Intersect(freqs.SignalFor(7))
	// Then to get 'G' exclude 'D' from remaining candidates
	candidates['G'] = freqs.SignalFor(7).Exclude(candidates['D'])

	ans := make(map[rune]rune)
	for unscrambled, scrambled := range candidates {
		if len(scrambled) != 1 {
			return nil, fmt.Errorf("failed to decode '%c', '%d candidates", unscrambled, len(scrambled))
		}
		ans[scrambled[0]] = unscrambled
	}

	return ans, nil
}
