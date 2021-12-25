package day3

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gmhorn/aoc21/lib"
)

const (
	lo = '0'
	hi = '1'

	loVal int8 = -1
	hiVal int8 = 1
)

type Report [][]int8

func NewReport(input string) (Report, error) {
	lines, err := lib.ReadLines(input)
	if err != nil {
		return nil, err
	}

	report := make([][]int8, 0, len(lines))
	for idxL, line := range lines {
		row := make([]int8, 0)
		for idxC, char := range line {
			switch char {
			case lo:
				row = append(row, loVal)
			case hi:
				row = append(row, hiVal)
			default:
				return nil, fmt.Errorf("unknown char '%c' at %d:%d", char, idxL, idxC)
			}
		}
		report = append(report, row)
	}

	return report, nil
}

func (r Report) Gamma() (int64, error) {
	s := ""
	for col := 0; col < r.width(); col++ {
		sum := r.sumColumn(col)
		switch {
		case sum >= hiVal:
			s += string(hi)
		case sum <= loVal:
			s += string(lo)
		default:
			return -1, fmt.Errorf("ambiguous gamma value at %d digit", col)
		}
	}
	return strconv.ParseInt(s, 2, 64)
}

func (r Report) Epsilon() (int64, error) {
	s := ""
	for col := 0; col < r.width(); col++ {
		sum := r.sumColumn(col)
		switch {
		case sum >= hiVal:
			s += string(lo)
		case sum <= loVal:
			s += string(hi)
		default:
			return -1, fmt.Errorf("ambiguous epsilon value at %d digit", col)
		}
	}
	return strconv.ParseInt(s, 2, 64)
}

func (r Report) Oxygen() (int64, error) {
	return -1, errors.New("not implemented")
}

func (r Report) C02() (int64, error) {
	return -1, errors.New("not implemented")
}

func (r Report) sumColumn(digit int) int8 {
	acc := int8(0)
	for line := 0; line < len(r); line++ {
		acc += r[line][digit]
	}

	return acc
}

func (r Report) width() int {
	return len(r[0])
}
