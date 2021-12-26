package day5

import (
	"fmt"
	"regexp"
	"strconv"
)

var validLine = regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)$`)

var emptyLine = Line{0, 0, 0, 0}

type Point struct {
	X, Y int
}

type Line struct {
	X1, Y1, X2, Y2 int
}

func (l Line) IsRect() bool {
	return l.X1 == l.X2 || l.Y1 == l.Y2
}

func CountOverlaps(lines []Line, threshold int) int {
	counts := make(map[Point]int)
	for _, line := range lines {
		for _, point := range line.Points() {
			counts[point]++
		}
	}

	total := 0
	for _, count := range counts {
		if count >= threshold {
			total++
		}
	}
	return total
}

func (l Line) Points() []Point {
	pts := make([]Point, 0)

	dx := sign(l.X2 - l.X1)
	dy := sign(l.Y2 - l.Y1)
	for t := 0; ; t++ {
		x := l.X1 + t*dx
		y := l.Y1 + t*dy
		pts = append(pts, Point{x, y})
		if x == l.X2 && y == l.Y2 {
			break
		}
	}
	return pts
}

// func (l Line) Points() []Point {
// 	pts := make([]Point, 0)
// 	if l.X1 == l.X2 {
// 		for y := min(l.Y1, l.Y2); y <= max(l.Y1, l.Y2); y++ {
// 			pts = append(pts, Point{l.X1, y})
// 		}
// 	}
// 	if l.Y1 == l.Y2 {
// 		for x := min(l.X1, l.X2); x <= max(l.X1, l.X2); x++ {
// 			pts = append(pts, Point{x, l.Y1})
// 		}
// 	}

// 	return pts
// }

func ParseLine(s string) (Line, error) {
	parts := validLine.FindStringSubmatch(s)
	if len(parts) != 5 {
		return emptyLine, fmt.Errorf("could not parse '%s'", s)
	}
	parts = parts[1:]

	vals := make([]int, 0)
	for _, part := range parts {
		val, err := strconv.Atoi(part)
		if err != nil {
			return emptyLine, fmt.Errorf("could not parse '%s' in '%s' as number: %v", part, s, err)
		}
		vals = append(vals, val)
	}

	return Line{vals[0], vals[1], vals[2], vals[3]}, nil
}

func max(vals ...int) int {
	ret := vals[0]
	for _, v := range vals[1:] {
		if v > ret {
			ret = v
		}
	}
	return ret
}

func min(vals ...int) int {
	ret := vals[0]
	for _, v := range vals[1:] {
		if v < ret {
			ret = v
		}
	}
	return ret
}

func sign(a int) int {
	switch {
	case a < 0:
		return -1
	case a > 0:
		return 1
	default:
		return 0
	}
}
