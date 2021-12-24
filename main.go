package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/gmhorn/aoc21/solutions"
)

type solution interface {
	Part1() (int, error)
	Part2() (int, error)
}

var solns = []solution{
	solutions.Day1{},
}

func main() {
	var day, part int
	flag.IntVar(&day, "day", day, "Day to run")
	flag.IntVar(&part, "part", part, "part to run")
	flag.Parse()

	// Skip a BOATLOAD of input checking. Just don't use nonsensical stuff like
	// --day 73 or --day -12, or you'll get a panic.

	// Days are 1-indexed, arrays are 0-indexed
	soln := solns[day-1]

	switch part {
	case 1:
		runSolution(day, part, soln.Part1)
	case 2:
		runSolution(day, part, soln.Part2)
	default:
		runSolution(day, 1, soln.Part1)
		fmt.Println()
		runSolution(day, 2, soln.Part2)
	}
}

func runSolution(day, part int, fn func() (int, error)) {
	start := time.Now()
	defer func() {
		fmt.Printf("(Took %s)\n", time.Since(start))
	}()

	fmt.Printf("Day %d, Part %d:\n", day, part)

	ans, err := fn()
	if err != nil {
		fmt.Printf("Error ocurred: %v\n", err)
		return
	}
	fmt.Println(ans)
}
