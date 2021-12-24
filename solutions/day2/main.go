package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gmhorn/aoc21/lib"
)

type state struct {
	x, y, aim int
}

type transform func(amt int, old state) state
type behavior struct {
	down, up, forward transform
}

var simple = behavior{
	down: func(amt int, old state) state {
		return state{old.x, old.y + amt, old.aim}
	},
	up: func(amt int, old state) state {
		return state{old.x, old.y - amt, old.aim}
	},
	forward: func(amt int, old state) state {
		return state{old.x + amt, old.y, old.aim}
	},
}

func process(input string, behavior behavior) (int, error) {
	s := state{}

	lines, err := lib.ReadLines(input)
	if err != nil {
		return -1, err
	}

	for idx, line := range lines {
		parts := strings.Split(line, " ")
		action := parts[0]
		amt, err := strconv.Atoi(parts[1])
		if err != nil {
			return -1, fmt.Errorf("could not parse line %d: %v", idx, err)
		}

		switch action {
		case "forward":
			s = behavior.forward(amt, s)
		case "down":
			s = behavior.down(amt, s)
		case "up":
			s = behavior.up(amt, s)
		}
	}

	return s.x * s.y, nil
}

func main() {
	lib.PrintSolution("Day 2, Part 1", func() (int, error) {
		return process("input.txt", simple)
	})
	fmt.Println()

}
