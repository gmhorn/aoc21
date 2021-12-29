package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gmhorn/aoc21/lib"
)

type Solution struct{}

func (sln Solution) Part1(input string) (int, error) {
	s, err := computeState(input, simple)
	if err != nil {
		return -1, err
	}

	return s.x * s.y, nil
}

func (sln Solution) Part2(input string) (int, error) {
	s, err := computeState(input, complex)
	if err != nil {
		return -1, err
	}

	return s.x * s.y, nil
}

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

var complex = behavior{
	down: func(amt int, old state) state {
		return state{old.x, old.y, old.aim + amt}
	},
	up: func(amt int, old state) state {
		return state{old.x, old.y, old.aim - amt}
	},
	forward: func(amt int, old state) state {
		return state{old.x + amt, old.y + old.aim*amt, old.aim}
	},
}

func computeState(input string, behavior behavior) (state, error) {
	s := state{}

	lines, err := lib.ReadLines(input)
	if err != nil {
		return s, err
	}

	for idx, line := range lines {
		parts := strings.Split(line, " ")
		action := parts[0]
		amt, err := strconv.Atoi(parts[1])
		if err != nil {
			return s, fmt.Errorf("could not parse line %d: %v", idx, err)
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

	return s, nil
}
