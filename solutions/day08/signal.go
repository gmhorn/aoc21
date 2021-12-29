package day8

import "sort"

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

// Equals compares Signals for equality.
func (s Signal) Equals(other Signal) bool {
	if len(s) != len(other) {
		return false
	}

	for i := 0; i < len(s); i++ {
		if s[i] != other[i] {
			return false
		}
	}

	return true
}

// Frequency is a counter for the occurrencies of each rune.
type Frequency map[rune]int

// Frequencies constructs a new Frequency from the list of input Signals.
func Frequencies(input []Signal) Frequency {
	freq := make(map[rune]int)
	for _, sig := range input {
		for _, r := range sig {
			freq[r]++
		}
	}
	return freq
}

// SignalFor constructs a signal consisting only of runes that have the given
// frequency.
func (f Frequency) SignalFor(count int) Signal {
	runes := make([]rune, 0)
	for r, cnt := range f {
		if cnt == count {
			runes = append(runes, r)
		}
	}
	// can't just return runes because no guarantee that `range` iterates
	// in any particular order, and Signals must be sorted.
	// So go []rune -> string -> Signal, and NewSignal sorts for us.
	return NewSignal(string(runes))
}
