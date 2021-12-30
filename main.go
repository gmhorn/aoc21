package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/gmhorn/aoc21/solutions/day01"
	"github.com/gmhorn/aoc21/solutions/day02"
	"github.com/gmhorn/aoc21/solutions/day03"
	"github.com/gmhorn/aoc21/solutions/day04"
	"github.com/gmhorn/aoc21/solutions/day05"
	"github.com/gmhorn/aoc21/solutions/day06"
	"github.com/gmhorn/aoc21/solutions/day07"
	"github.com/gmhorn/aoc21/solutions/day08"
	"github.com/gmhorn/aoc21/solutions/day09"
	"github.com/gmhorn/aoc21/solutions/day10"
	"github.com/gmhorn/aoc21/solutions/day11"
)

type solution interface {
	Part1(input string) (int, error)
	Part2(input string) (int, error)
}

var solns = []solution{
	day01.Solution{},
	day02.Solution{},
	day03.Solution{},
	day04.Solution{},
	day05.Solution{},
	day06.Solution{},
	day07.Solution{},
	day08.Solution{},
	day09.Solution{},
	day10.Solution{},
	day11.Solution{},
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
	input := fmt.Sprintf("input/day%02d.txt", day)

	switch part {
	case 1:
		runSolution(fmt.Sprintf("Day %d, Part 1:", day), input, soln.Part1)
	case 2:
		runSolution(fmt.Sprintf("Day %d, Part 2:", day), input, soln.Part2)
	default:
		runSolution(fmt.Sprintf("Day %d, Part 1:", day), input, soln.Part1)
		fmt.Println()
		runSolution(fmt.Sprintf("Day %d, Part 2:", day), input, soln.Part2)
	}
}

func runSolution(title, input string, fn func(string) (int, error)) {
	start := time.Now()
	defer func() {
		fmt.Printf("(Took %s)\n", time.Since(start))
	}()

	fmt.Println(title)
	ans, err := fn(input)
	if err != nil {
		fmt.Printf("Error ocurred: %v\n", err)
		return
	}
	fmt.Println(ans)
}
