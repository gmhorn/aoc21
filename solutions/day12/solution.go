package day12

import (
	"strings"

	"github.com/gmhorn/aoc21/lib"
)

type cave string
type path []cave
type graph map[cave][]cave
type rule func(path, cave) bool

const (
	start cave = "start"
	end   cave = "end"
)

type Solution struct{}

func (sln Solution) Part1(input string) (int, error) {
	g, err := loadGraph(input)
	if err != nil {
		return -1, err
	}

	return len(explore(g, visitOnce)), nil
}

func (sln Solution) Part2(input string) (int, error) {
	g, err := loadGraph(input)
	if err != nil {
		return -1, err
	}

	return len(explore(g, visitTwice)), nil
}

var visitOnce rule = func(p path, c cave) bool {
	// big caves always get explored
	if c.isBig() {
		return true
	}

	// if c is already in the path, don't visit
	for _, elem := range p {
		if elem == c {
			return false
		}
	}
	return true
}

var visitTwice rule = func(p path, c cave) bool {
	// never return to start
	if len(p) > 0 && c == start {
		return false
	}

	// big caves always get explored
	if c.isBig() {
		return true
	}

	// Assuming we do explore c, count how many times each small cave would be
	// explored. If there are 2 or more greater than 2, don't explore
	counts := make(map[cave]int)
	for _, elem := range p {
		if !elem.isBig() {
			counts[elem]++
		}
	}
	counts[c]++

	cavesExploredTwice := 0
	for _, cnt := range counts {
		// don't explore any small cave more than twice
		if cnt > 2 {
			return false
		}
		// keep track of number of caves explored exactly twice
		if cnt == 2 {
			cavesExploredTwice++
		}
	}
	return cavesExploredTwice < 2
}

func (c cave) isBig() bool {
	return strings.ToUpper(string(c)) == string(c)
}

func (p path) last() cave {
	return p[len(p)-1]
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
	return strings.Join(parts, ",")
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

func explore(g graph, shouldExplore rule) []path {
	paths := make([]path, 0)

	toExplore := make([]path, 0)
	toExplore = append(toExplore, []cave{start})

	for len(toExplore) != 0 {
		// dequeue next path
		p := toExplore[0]
		toExplore = toExplore[1:]

		// if path terminates, increment the total and continue
		if p.last() == end {
			paths = append(paths, p)
			continue
		}

		// otherwise get all caves connected to the last cave in the path
		for _, c := range g[p.last()] {
			// if explore rule determines we should keep exploring, add to our
			// list of paths to keep exploring.
			if shouldExplore(p, c) {
				toExplore = append(toExplore, p.concat(c))
			}
		}
	}

	return paths
}
