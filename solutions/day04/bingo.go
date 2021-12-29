package day4

import (
	"fmt"
	"strconv"
	"strings"
)

const size = 5

// Bingo is a struct representing a bingo board.
type Bingo struct {
	vals   [5][5]int
	marked [5][5]bool
}

func NewBingo(lines []string) (*Bingo, error) {
	b := new(Bingo)
	// MISSING: all kinds of bounds checking
	for row := 0; row < size; row++ {
		vals := strings.Fields(lines[row])
		for col := 0; col < size; col++ {
			val, err := strconv.Atoi(vals[col])
			if err != nil {
				return nil, fmt.Errorf("error converting %s (%d:%d): %v", vals[col], row, col, err)
			}
			b.vals[row][col] = val
		}
	}
	return b, nil
}

// Mark marks a value on the board. It returns true if, as a result of marking
// the new value, the board is now a winner.
func (b *Bingo) Mark(val int) bool {
	var row, col int
	var found bool

search:
	for row = 0; row < 5; row++ {
		for col = 0; col < 5; col++ {
			if b.vals[row][col] == val {
				b.marked[row][col] = true
				found = true
				break search
			}
		}
	}

	if !found {
		return false
	}

	// search for winners along same row as newly-marked value
	return b.winningRow(row) || b.winningCol(col)
}

func (b *Bingo) SumOfUnmarked() int {
	sum := 0
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			if !b.marked[row][col] {
				sum += b.vals[row][col]
			}
		}
	}
	return sum
}

func (b *Bingo) winningRow(row int) bool {
	for _, val := range b.marked[row] {
		if !val {
			return false
		}
	}
	return true
}

func (b *Bingo) winningCol(col int) bool {
	for i := 0; i < size; i++ {
		if !b.marked[i][col] {
			return false
		}
	}
	return true
}
