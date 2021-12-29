package day05

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
