package day12

import (
	"errors"
	"strings"

	"github.com/gmhorn/aoc21/lib"
)

type Solution struct{}

func (sln Solution) Part1(input string) (int, error) {
	g, err := loadGraph(input)
	if err != nil {
		return -1, err
	}

	return countPaths(g, cave("start"), cave("end")), nil
}

func (sln Solution) Part2(input string) (int, error) {
	return -1, errors.New("not implemented")
}

type cave string
type path []cave
type graph map[cave][]cave

func (c cave) isSmall() bool {
	return strings.ToLower(string(c)) == string(c)
}

func (p path) last() cave {
	return p[len(p)-1]
}

func (p path) contains(c cave) bool {
	for _, elem := range p {
		if c == elem {
			return true
		}
	}
	return false
}

func (p path) concat(c cave) path {
	arr := make([]cave, 0, len(p)+1)
	arr = append(arr, p...)
	arr = append(arr, c)
	return path(arr)
}

func (p path) String() string {
	parts := make([]string, 0, len(p))
	for _, c := range p {
		parts = append(parts, string(c))
	}
	return strings.Join(parts, "-")
}

func loadGraph(input string) (graph, error) {
	lines, err := lib.ReadLines(input)
	if err != nil {
		return nil, err
	}

	g := make(graph)

	for _, line := range lines {
		caves := strings.Split(line, "-")
		a, b := cave(caves[0]), cave(caves[1])
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	return g, nil
}

func countPaths(graph graph, start, end cave) int {
	total := 0

	toExplore := make([]path, 0)
	toExplore = append(toExplore, []cave{start})

	for len(toExplore) != 0 {
		// dequeue next path
		p := toExplore[0]
		toExplore = toExplore[1:]

		// if path terminates, increment the total and continue
		if p.last() == end {
			total++
			continue
		}

		// otherwise get all caves connected to the last cave in the path
		for _, c := range graph[p.last()] {
			// if the cave is small and we've already visited on this path, its
			// a dead end
			if c.isSmall() && p.contains(c) {
				continue
			}

			// otherwise, add it to our list of paths to keep exploring
			toExplore = append(toExplore, p.concat(c))
		}
	}

	return total
}
